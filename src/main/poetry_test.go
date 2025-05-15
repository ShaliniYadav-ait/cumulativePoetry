package main

import (
	c "cumulativepoetry/src/config"
	"testing"
)

func Test_recitePoem(t *testing.T) {
	type args struct {
		echo bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test recite method",
			args: args{
				echo: false,
			},
			want: poemInStringFormat(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recitePoem(tt.args.echo); got != tt.want {
				t.Errorf("recitePoem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_poemToString(t *testing.T) {
	type args struct {
		poem c.Poem
		echo bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "convert poem to string",
			args: args{
				poem: c.GetPoemConfig(),
				echo : false,
			},
			want: poemInStringFormat(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := poemToString(tt.args.poem,tt.args.echo); got != tt.want {
				t.Errorf("poemToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dayDetailsToString(t *testing.T) {
	type args struct {
		dayDetails c.PoemDetails
		echo       bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check for day 3 details",
			args: args{
				dayDetails: getDayDetails(3),
				echo:       false,
			},
			want: "Day 3 - " + "\n" +
				"and sorry I could not travel both" + "\n" +
				"and be one traveler, long I stood" + "\n" +
				"and looked down one as far as I could",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayDetailsToString(tt.args.dayDetails, tt.args.echo); got != tt.want {
				t.Errorf("dayDetailsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getDayDetails(day int) c.PoemDetails {
	poem := c.GetPoemConfig()
	return poem[day-1]
}

func poemInStringFormat() string {
	poemInString := "\n" + "Day 1 - " + "\n" +
		"Two roads diverged in a yellow wood," + "\n" +
		"Day 2 - " + "\n" +
		"and be one traveler, long I stood" + "\n" +
		"and looked down one as far as I could" + "\n" +
		"Day 3 - " + "\n" +
		"and sorry I could not travel both" + "\n" +
		"and be one traveler, long I stood" + "\n" +
		"and looked down one as far as I could" + "\n" +
		"Day 4 - " + "\n" +
		"To where it bent in the undergrowth." + "\n" +
		"then took the other, as just as fair," + "\n" +
		"and having perhaps the better claim," + "\n" +
		"because it was grassy and wanted wear." + "\n" +
		"Day 5 - " + "\n" +
		"Though as for that the passing there" + "\n" +
		"had worn them really about the same," + "\n" +
		"and both that morning equally lay" + "\n" +
		"in leaves no step had trodden black." + "\n" +
		"Oh, I kept the first for another day!" + "\n" +
		"Day 6 - " + "\n" +
		"Yet knowing how way leads on to way," + "\n" +
		"I doubted if I should ever come back." + "\n" +
		"I shall be telling this with a sigh" + "\n" +
		"somewhere ages and ages hence:" + "\n" +
		"Two roads diverged in a wood, and I—" + "\n" +
		"I took the one less traveled by,And that has made all the difference."
	return poemInString
}
