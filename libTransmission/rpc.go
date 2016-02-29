// There has gotta be away around the nasty copypasta hacks
// TODO: FIX THAT SHIT

package libTransmission

import (
	"net/http"
	"encoding/json"
	"log"
	"bytes"
	"io"
	"errors"
	"github.com/germ/geoip"
)

var (
	ServerURL	= "http://:9090/transmission/rpc"
	ServerUser	= "germ"
	ServerPass	= "hackersgonnahack"
)

func Add(url, dir string) (err error) {
	req := Request {
		Method: "torrent-add",
		Arguments: map[string]interface{} {
			"filename":url,
			"download-dir":dir,
		},
	}

	reader,err := req.getJson()
	defer reader.Close()

	if err != nil {
		log.Println(err)
		return
	}

	var res AddResponse
	err = json.NewDecoder(reader).Decode(&res)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
	if res.Result != "success" {
		return errors.New("Torrent failed to add")
	}

	return

}
func GetTorrents() (peers []Torrent, err error) {
	// Craft Request
	req := Request {
		Method:		"torrent-get",
		Arguments:	map[string]interface{} {
			"fields":[]string{"peers", "id", "name"},
		},
	}

	// Make request
	jsonReader, err := req.getJson()
	defer jsonReader.Close()
	if err != nil {
		log.Println("Error making request: ", err, jsonReader)
		return
	}

	var resp TorrentResponse
	err = json.NewDecoder(jsonReader).Decode(&resp)
	if err != nil {
		log.Println(err)
		return
	}

	if resp.Result != "success" {
		return
	}

	for _, v := range(resp.Arguments["torrents"]) {
		peers = append(peers, v)
	}

	return
}
func LimitSpeed(enable bool) (err error) {
	req := Request {
		Method:"session-set",
		Arguments: map[string]interface{} {
			"alt-speed-enabled":enable,
		},
	}

	read, err := req.getJson()
	defer read.Close()
	if err != nil {
		log.Println(err)
		return
	}

	var s Response
	err = json.NewDecoder(read).Decode(&s)
	if err != nil {
		log.Println(err)
		return
	}

	if s.Result != "success" {
		err = errors.New("Error: Could not enable alts")
		return
	}

	return
}
func GetStats() (s Stats, err error) {
	req := Request {
		Method:		"session-get",
	}

	read, err := req.getJson()
	defer read.Close()
	if err != nil {
		log.Println(err)
		return
	}

	var resp StatsResponse
	err = json.NewDecoder(read).Decode(&resp)
	if err != nil {
		log.Println(err)
		return
	}

	if resp.Result != "success" {
		err = errors.New("Error: Printer on Fire")
		return
	}

	return resp.Arguments, err
}
func GetSession() (s Session, err error) {
	// So here's something funny. Turns out I named my data structures
	// wrong. Stats relate to Session vars and vice versa.
	// TODO: Learn how to use gofmt rewrite

	req := Request {
		Method:		"session-stats",
	}

	read, err := req.getJson()
	defer read.Close()
	if err != nil {
		log.Println(err)
		return
	}

	var resp SessionResponse
	err = json.NewDecoder(read).Decode(&resp)
	if err != nil {
		log.Println(err)
		return
	}

	if resp.Result != "success" {
		err = errors.New("Error: Printer on Fire")
		return
	}

	return resp.Arguments, err
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetPrefix("[ ScanServer ] ");
}
func (r *Request) getJson() (jsonReader io.ReadCloser, err error) {
	// Can the request into JSON
	msg, err := json.Marshal(r)
	if err != nil {
		log.Println("%v : %v", err, msg)
		return
	}

	buf := bytes.NewBuffer(msg)
	log.Println("Message: ", buf)

	// Craft the request
	var client http.Client
	req, err := http.NewRequest("POST", ServerURL, buf)
	if err != nil {
		log.Println("%v : %v", err, req)
		return
	}
	req.Header.Add("X-Transmission-Session-Id", r.sessionID)
	req.SetBasicAuth(ServerUser, ServerPass)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("%v : %v", err, resp)
		return
	}

	// Set Auth header if needed
	if resp.StatusCode == 409 {
		resp.Body.Close()

		id := resp.Header.Get("X-Transmission-Session-Id")
		log.Println("Changing SessionID to: ", id)
		r.sessionID = id

		return r.getJson()
	}

	jsonReader = resp.Body
	return
}
func (p *TorrentPeer) Locate() (err error) {
	location, err := geoip.LookupString(p.Address)
	if err != nil {
		log.Println(err)
		return
	}

	p.Geo = location
	p.Geo.Description = p.Address + " " + p.ClientName
	return
}
