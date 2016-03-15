package summary
import "regexp"

type Analysis struct {
	OriginalWordCount int
	SummaryWordCount  int
	Ratio             float32
}

type Result struct {
	Title    string `json:"title"`
	LexRank  string `json:"lexrank"`
	Natural  string `json:"natural"`
	Analysis map[string]*Analysis `json:"analysis`
}

func (r *Result) Analyze(src *Article) {

	a := make(map[string]*Analysis)

	a["lexrank"] = analyze(src.Content, r.LexRank)
	a["natural"] = analyze(src.Content, r.Natural)

	r.Analysis = a
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
