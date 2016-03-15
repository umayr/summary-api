package summary

import (
	"github.com/urandom/text-summary/summarize"
	"github.com/JesusIslam/tldr"
	"strings"
)

const (
	MaxSentences = 4
	IdealWordCount = 75
)

type Article struct {
	Url     string
	Title   string
	Content string
}

func (a *Article) Summary() *Result {
	var (
		cl = make(chan string)
		cs = make(chan string)
	)

	go lexRank(cl, a.Content);
	go natural(cs, a.Title, a.Content)

	r := &Result{}
	r.LexRank = <-cl
	r.Natural = <-cs
	r.Title = a.Title

	r.Analyze(a)
	return r;
}

// LexRank Algorithm
func lexRank(c chan string, content string) {
	t := tldr.New()
	td, _ := t.Summarize(content, MaxSentences)
	c <- td;
}

// Natural Language Processing Algorithm
func natural(c chan string, title string, content string) {
	s := summarize.NewFromString(title, content)
	s.IdealWordCount = IdealWordCount
	c <- strings.Join(s.KeyPoints(), "")
}
