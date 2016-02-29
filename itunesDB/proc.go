package itunesDB

import (
	"io/ioutil"
	"os/user"
	plist "github.com/DHowett/go-plist"
)

type Library struct {
	MajorVersion		int			`plist:"Major Version"`
	MinorVersin			int			`plist:"Minor Version"`
	Features			int
	ApplicationVersion	string		`plist:"Application Version"`
	MusicFolder			string		`plist:"Music Folder"`
	ID					string		`plist:"Library Persistent ID"`
	TrackList			map[string]Track `plist:"Tracks"`
}
type Track struct {
	// For the record: There are many fields that are not included here
	TrackID			int
	Name			string
	Album			string
	Artist			string
	Genre			string
	Kind			string
	TotalTime		int
	DateAdded		string
	PlayCount		int
	PlayDate		string
	ReleaseDate		string
	PersistantID	string
	TrackType		string
	Location		string
	SkipCount		int

}

func ReadLibrary() (artists []string, err error) {
	// Get home dir
	usr, err := user.Current()
	if err != nil {
		return
	}

	// Read Library
	var library Library
	data, err := ioutil.ReadFile(usr.HomeDir + "/Music/iTunes/iTunes Library.xml")
	plist.Unmarshal(data, &library)

	// Extract Artist list
	libArtist := make(map[string]int)
	for _, info := range library.TrackList {
		libArtist[info.Artist] += 1
	}

	// Print the list
	for i, _ := range libArtist {
		artists = append(artists, i)
	}

	return
}
