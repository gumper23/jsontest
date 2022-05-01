package main

type Database struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	Schema   string            `json:"schema"`
	Options  map[string]string `json:"options"`
}
