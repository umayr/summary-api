package summary

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