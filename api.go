package main

type API struct {
	Port        string `json:"port"`
	ClientCert  string `json:"clientcert"`
	ClientKey   string `json:"clientkey"`
	ServerCert  string `json:"servercert"`
	ServerKey   string `json:"serverkey"`
	ReleaseMode string `json:"releasemode"`
}
