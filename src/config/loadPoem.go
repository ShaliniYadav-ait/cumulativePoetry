package config

import "sync"

var once sync.Once

type Poem struct {
	Day   int      `json:"day"`
	Lines []string `json:"lines"`
}

var p Poem

func getPoemConfig() Poem {
	once.Do(func() {
		LoadConfig("poem.json", &p)
	})
	return p
}
