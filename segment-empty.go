package main

import (
	pwl "github.com/justjanne/powerline-go/powerline"
)

func segmentEmpty(p *powerline) []pwl.Segment {
	l := 0
	for _, s := range p.Segments {
		l = l + len(s)
	}
	if l > 0 {
		return []pwl.Segment{{Content: "", Compact: true}}
	}
	return []pwl.Segment{}
}
