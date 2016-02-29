// Tests are just non-interactive ways to use librarys right?
package libTransmission

import "testing"

func TestAdd(t *testing.T) {
	t.Log(Add("http://torrents.thepiratebay.se/9379495/Minecraft_1.7.2_Cracked_[Full_Installer]_.9379495.TPB.torrent"))
}
func TestToggle(t *testing.T) {
	t.Log(LimitSpeed(true))
}
func TestGetTorrents(t *testing.T) {
	t.Log(GetTorrents())
}
func TestTorrentStats(t *testing.T) {
	t.Log(GetStats())
}
func TestSession(t *testing.T) {
	t.Log(GetSession())
}
