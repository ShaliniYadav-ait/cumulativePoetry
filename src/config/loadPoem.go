package config

import "sync"

var once sync.Once

type PoemDetails struct {
	Day   int      `json:"day"`
	Lines []string `json:"lines"`
}

type Poem struct {
	Poem []PoemDetails `json:"poem"`
}

var p Poem

func getPoemConfig() Poem {
	once.Do(func() {
		LoadConfig("poem.json", &p)
	})
	return p
}
