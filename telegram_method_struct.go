package gobot

import "os"

type GetUpdates struct {
	Offset         int      `json:"offset"`
	Limit          int      `json:"limit"`
	Timeout        int      `json:"timeout"`
	AllowedUpdates []string `json:"allowed_updates"`
}

type SetWebHook struct {
	Url                string   `json:"url"`
	Certificate        os.File  `json:"certificate"`
	IpAdress           string   `json:"ip_address"`
	MaxConnections     int      `json:"max_connections"`
	AllowedUpdates     []string `json:"allowed_updates"`
	DropPendingUpdates bool     `json:"drop_pending_updates"`
	SecretToken        string   `json:"secret_token"`
}
