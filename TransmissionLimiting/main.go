package TransmissionLimiting

import (
	"net/http"
	"log"
	"time"
	"math/rand"
	"html/template"
	"fmt"
	"github.com/germ/libTransmission"
)

var (
	timeout	= 108 * time.Minute
	fMap = template.FuncMap {
		"humanSpeed":humanSpeed,
		"humanSize":humanSize,
	}
	rootPage = template.Must(template.New("").Funcs(fMap).Parse(rootTemplate))
)

func ServeRoot(w http.ResponseWriter, r *http.Request) {
	var err error
	log.Println("Page Request:")

	// Grab Meta
	stats, err := libTransmission.GetStats()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 418)
		return
	}

	session, err := libTransmission.GetSession()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 418)
		return
	}

	// Disable torrents on form
	r.ParseForm()
	query := r.Form.Get("enable")

	if query != "" {
		if query == "true" {
			err = libTransmission.LimitSpeed(true)
			log.Println("Alt speeds enabled")
		} else {
			err = libTransmission.LimitSpeed(false)
			log.Println("Alt speeds disabled")
		}

		go func() {
			log.Println("Alt speeds disabled")
			time.Sleep(timeout)
			err := libTransmission.LimitSpeed(false)
			if err != nil {
				log.Fatal("FALLBACK DISABLED. BRING IT DOWN")
			}
		}()

		http.Redirect(w, r, "/", 303)
	}


	// Fill template
	t := TemplateFiller{
		Quote:getTurtleQuote(),
		Stats:stats,
		Session:session,
	}
	rootPage.Execute(w, t)
}

type TemplateFiller struct {
	Stats	libTransmission.Stats
	Session	libTransmission.Session
	Quote	string
}
func getTurtleQuote() string {
	quote := []string{
		"Turtles have been on the earth for more than 200 million years. They evolved before mammals, birds, crocodiles, snakes, and even lizards.",
		"The earliest turtles had teeth and could not retract their heads, but other than this, modern turtles are very similar to their original ancestors.",
		"Several species of turtles can live to be over a hundred years of age including the American Box Turtle.",
		"One documented case of longevity involves an adult Indian Ocean Giant Tortoise that, when captured as an adult, was estimated to be fifty years old. It then lived another 152 years in captivity.",
		"Turtles live on every continent except Antarctica.",
		"Turtles will live in almost any climate warm enough to allow them to complete their breeding cycle.",
		"While most turtles do not tolerate the cold well, the Blanding's turtle has been observed swimming under the ice in the Great Lakes region.",
		"Turtles range in size from the 4-inch Bog Turtle to the 1500-pound Leathery Turtle.",
		"North America contains a large variety of turtle species, but Europe contains only two species of turtle and three species of tortoise.",
		"The top domed part of a turtle's shell is called the carapace, and the bottom underlying part is called the plastron.",
		"The shell of a turtle is made up of 60 different bones all connected together.",
		"The bony portion of the shell is covered with plates (scutes) that are derivatives of skin and offer additional strength and protection.",
		"Most land tortoises have high, domed carapaces that offer protection from the snapping jaws of terrestrial predators. Aquatic turtles tend to have flatter, more aerodynamically shaped shells. An exception to the dome-shaped tortoise shell is the Pancake Tortoise of East Africa that will wedge itself between narrow rocks when threatened and then inflates itself with air making extraction nearly impossible.",
		"Most turtle species have five toes on each limb with a few exceptions including the American Box Turtle of the Carolina species that only has four toes and, in some cases, only three.",
		"Turtles have good eyesight and an excellent sense of smell. Hearing and sense of touch are both good and even the shell contains nerve endings.",
		"Some aquatic turtles can absorb oxygen through the skin on their neck and cloacal areas allowing them to remain submerged underwater for extended periods of time and enabling them to hibernate underwater.",
		"Turtles are one of the oldest and most primitive groups of reptiles and have outlived many other species. One can only wonder if their unique shell is responsible for their longevity.",
	}

	return quote[rand.Intn(len(quote))]
}
func humanSpeed(s float64) (human string) {
	postFixes := []string{"bps", "KB/s", "MB/s", "GB/s"}

	var i int
	for i = 0; s > 1024; i++ {
		s /= 1024
	}

	return fmt.Sprintln(int(s), postFixes[i])
}
func humanSize(s float64) (human string) {
	postFixes := []string{"bytes", "KB", "MB", "GB"}

	var i int
	for i = 0; s > 1024; i++ {
		s /= 1024
	}

	return fmt.Sprintln(int(s), postFixes[i])
}
const rootTemplate = `
<!DOCTYPE html>
<html>
<head>
<style>
body {
	align:			center;
	margin-left:	auto;
	margin-right:	auto;
	text-align:		center;
}
img {
	width: 	30%;
	height: 30%;
}
</style>
</head>
<body>
<h2>Current: [ {{ humanSpeed .Session.DownloadSpeed }} | {{ humanSpeed .Session.UploadSpeed }} ]</h1>
<h3>Limits: [ {{ .Stats.AltSpeedUp }} Kb/s | {{ .Stats.AltSpeedDown }} Kb/s ]</h1>

{{ if .Stats.AltSpeedEnabled }}
<a href="/?enable=false"><img src="/static/slow.jpg"/></a>
<p>{{ .Quote }}</p>
{{ else }}
<a href="/?enable=true"><img src="/static/fast.jpg"/></a>
<p>OPERATIONAL STATUS: VTEC</p>
{{ end }}

<p>
	[ <a href="/">Home</a> ] 
	[ <a href="/pocket">Pocket</a> ] 
	[ <a href="/version">Version</a> ] 
</p>
</body>
</html>
`
