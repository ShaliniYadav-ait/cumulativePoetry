package config

import "sync"

var once sync.Once

type PoemDetails struct {
	Day   int      `json:"day"`
	Lines []string `json:"lines"`
}

type Poem []PoemDetails 
var p Poem

func GetPoemConfig() Poem {
	once.Do(func() {
		LoadConfig("poem.json", &p)
	})
	return p
}
