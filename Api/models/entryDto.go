package models

type EntryDto struct {
	Id          int    `json:"id"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Definition  string `json:"definition"`
}
