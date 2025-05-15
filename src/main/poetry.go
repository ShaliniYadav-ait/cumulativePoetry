package main

import (
	c "cumulativepoetry/src/config"
	"strconv"
)

func revealforDay(day int, echo bool) string {
	poem := c.GetPoemConfig()
	return dayDetailsToString(poem[day-1], echo)
}

func poemToString(poem c.Poem, echo bool) string {
	var finalString string
	for _, dayDetails := range poem {
		finalString += "\n" + dayDetailsToString(dayDetails, echo)
	}
	return finalString
}

func recitePoem(echo bool) string {
	poem := c.GetPoemConfig()
	return poemToString(poem, echo)
}

func dayDetailsToString(dayDetails c.PoemDetails, echo bool) string {
	var dayDetailsString, poemLines string
	dayDetailsString = "Day " + strconv.Itoa(dayDetails.Day) + " - "
	totalLines := len(dayDetails.Lines)
	poemLines += "\n" + dayDetails.Lines[0]
	if echo {
		poemLines += "\n" + trimToSecondSpace(dayDetails.Lines[0])
	}
	for i := 1; i < totalLines; i++ {
		poemLines += "\n" + dayDetails.Lines[i]
		if echo {
			poemLines += "\n" + dayDetails.Lines[i]
		}
	}

	dayDetailsString += poemLines
	return dayDetailsString
}

func trimToSecondSpace(s string) string {
	spaceCount := 0
	for i, r := range s {
		if r == ' ' {
			spaceCount++
			if spaceCount == 2 {
				return s[i+1:]
			}
		}
	}
	return s
}
