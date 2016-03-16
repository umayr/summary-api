package summary

import "regexp"

type Analysis struct {
	OriginalWordCount int
	SummaryWordCount  int
	Ratio             float32
}

func wordCount(s string) int {
	return len(regexp.MustCompile("\\w+").FindAllString(s, -1))
}

func analyze(c string, s string) *Analysis {
	wc := wordCount(c)
	ws := wordCount(s)
	r := float32(100 * ws) / float32(wc)

	return &Analysis{wc, ws, r}
}
