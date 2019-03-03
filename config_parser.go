package main

type Config struct {
	Port		string							`json:"port,omitempty"`
	Url			string							`json:"url,omitempty"`
	Repos		map[string]map[string]string	`json:"repos,omitempty"`
}