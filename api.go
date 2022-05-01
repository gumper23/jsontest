package main

type API struct {
	Port        string `json:"port"`
	Clientcert  string `json:"clientcert"`
	Clientkey   string `json:"clientkey"`
	Servercert  string `json:"servercert"`
	Serverkey   string `json:"serverkey"`
	Releasemode string `json:"releasemode"`
}
