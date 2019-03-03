package main

type Config struct {
	Port		string		`json:"port,omitempty"`
	Url			string		`json:"url,omitempty"`
	Repos		[]Repo		`json:"repos,omitempty"`
}

type Repo struct {
	Name		string		`json:"name,omitempty"`
	Branches	[]Branch	`json:"branches,omitempty"`
}

type Branch struct {
	Name		string		`json:"name,omitempty"`
	Handle		string		`json:"handle,omitempty"`
}