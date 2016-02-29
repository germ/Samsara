package libTransmission

import "github.com/germ/geoip"

// A mess of types to deal with how RPC calls are made
type Request struct {
	Method    string                 `json:"method"`
	Arguments map[string]interface{} `json:"arguments"`
	Tag       float64                `json:"tag"`

	sessionID string
}
type Response struct {
	Result    string               `json:"result"`
	Arguments map[string]interface{} `json:"arguments"`
	Tag       float64              `json:"tag"`
}

type TorrentResponse struct {
	Result    string               `json:"result"`
	Arguments map[string][]Torrent `json:"arguments"`
	Tag       float64              `json:"tag"`
}
type Torrent struct {
	Id    float64       `json:"id"`
	Peers []TorrentPeer `json:"peers"`
	Name  string        `json:"name"`
}
type TorrentPeer struct {
	Address    string `json:"address"`
	ClientName string `json:"clientName"`
	Geo        geoip.GeoData
}

type AddResponse struct {
	Result    string             `json:"result"`
	Arguments map[string]AddInfo `json:"arguments"`
	Tag       float64            `json:"tag"`
}
type AddInfo struct {
	Id         float64 `json:"id"`
	Name       string  `json:"name"`
	hashString string  `json:"hashString"`
}

type StatsResponse struct {
	Result    string             `json:"result"`
	Arguments Stats				`json:"arguments"`
	Tag       float64            `json:"tag"`
}
type Stats	struct {
	AltSpeedDown          float64 `json:"alt-speed-down"`
	AltSpeedEnabled       bool    `json:"alt-speed-enabled"`
	AltSpeedTimeBegin     float64 `json:"alt-speed-time-begin"`
	AltSpeedTimeEnabled   bool    `json:"alt-speed-time-enabled"`
	AltSpeedTimeEnd       float64 `json:"alt-speed-time-end"`
	AltSpeedTimeDay       float64 `json:"alt-speed-time-day"`
	AltSpeedUp            float64 `json:"alt-speed-up"`
	BlocklistEnabled      bool    `json:"blocklist-enabled"`
	BlocklistSize         float64 `json:"blocklist-size"`
	DhtEnabled            bool    `json:"dht-enabled"`
	Encryption            string  `json:"encryption"`
	DownloadDir           string  `json:"download-dir"`
	PeerLimitGlobabl      float64 `json:"peer-limit-global"`
	PeerLimitPerTorrent   float64 `json:"peer-limit-per-torrent"`
	PexEnabled            bool    `json:"pex-enabled"`
	PeerPort              float64 `json:"peer-port"`
	PeerPortRndmOnStart   bool    `json:"peer-port-random-on-start"`
	PortForwardEnabled    bool    `json:"port-forwarding-enabled"`
	RpcVersion            float64 `json:"rpc-version"`
	RpcVersionMin         float64 `json:"rpc-version-minimum"`
	SeedRationLimit       float64 `json:"seedRatioLimit"`
	SeedRationLimited     bool    `json:"seedRatioLimited"`
	SpeedLimitDown        float64 `json:"speed-limit-down"`
	SpeedLimitDownEnabled bool    `json:"speed-limit-down-enabled"`
	SpeedLimitUp          float64 `json:"speed-limit-up"`
	SpeedLimitUpEnabled   bool    `json:"speed-limit-up-enabled"`
	Version               string  `json:"version"`
}

type SessionResponse struct {
	Result    string             `json:"result"`
	Arguments Session			`json:"arguments"`
	Tag       float64            `json:"tag"`
}
type Session struct {
	ActiveTorrentCount	float64	`json:"activeTorrentCount"`
	DownloadSpeed		float64	`json:"downloadSpeed"`
	PausedTorrentCount	float64	`json:"pausedTorrentCount"`
	TorrentCount		float64	`json:"torrentCount"`
	UploadSpeed			float64	`json:"uploadSpeed"`
	CumulativeStats		CumulativeStats	`json:"cumulative-stats"`
	CurrentStats		CurrentStats`json:"current-stats"`
}
type CumulativeStats struct {
	UploadedBytes	float64	`json:"uploadedBytes"`
	DownloadedBytes	float64	`json:"downloadedBytes"`
	FilesAdded		float64	`json:"filesAdded"`
	SessionCount	float64	`json:"sessionCount"`
	SecondsActive	float64	`json:"secondsActive"`
}
type CurrentStats struct {
	UploadedBytes	float64	`json:"uploadedBytes"`
	DownloadedBytes	float64	`json:"downloadedBytes"`
	FilesAdded		float64	`json:"filesAdded"`
	SessionCount	float64	`json:"sessionCount"`
	SecondsActive	float64	`json"secondsActive"`
}
