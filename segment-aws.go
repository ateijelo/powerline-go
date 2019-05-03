package main

import (
	"os"

	pwl "github.com/justjanne/powerline-go/powerline"
)

func segmentAWS(p *powerline) []pwl.Segment {
	profile := os.Getenv("AWS_PROFILE")
	region := os.Getenv("AWS_DEFAULT_REGION")
	if profile != "" {
		var r string
		if region != "" {
			r = " (" + region + ")"
		}
		return []pwl.Segment{{
			Name:       "aws",
			Content:    "\uf1b3  " + profile + r,
			Foreground: p.theme.AWSFg,
			Background: p.theme.AWSBg,
		}}
	}
	return []pwl.Segment{}
}
