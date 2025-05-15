package main

import (
	c "cumulativepoetry/src/config"
	"strconv"
)

func recitePoem() string {
	poem := c.GetPoemConfig()
	return poemToString(poem)
}

func revealforDay(day int) string {
	poem := c.GetPoemConfig()
	return poemForDay(day, poem)
}

func poemToString(poem c.Poem) string {
	var finalString string
	for _, dayDetails := range poem {
		finalString += "\n" + dayDetailsToString(dayDetails)
	}
	return finalString
}

func dayDetailsToString(dayDetails c.PoemDetails) string {
	var dayDetailsString, poemLines string
	dayDetailsString = "Day " + strconv.Itoa(dayDetails.Day) + " - "
	totalLines := len(dayDetails.Lines)
	for i := 0; i < totalLines; i++ {
		poemLines += "\n" + dayDetails.Lines[i]
	}

	dayDetailsString += poemLines
	return dayDetailsString
}

func poemForDay(day int, poem c.Poem) string {
	return dayDetailsToString(poem[day-1])
}
