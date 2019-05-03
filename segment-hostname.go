package main

import (
	"crypto/md5"
	"os"
	"strconv"
	"strings"

	pwl "github.com/justjanne/powerline-go/powerline"
)

func getHostName(fullyQualifiedDomainName string) string {
	return strings.SplitN(fullyQualifiedDomainName, ".", 2)[0]
}

func getMd5(text string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hasher.Sum(nil)
}

func segmentHost(p *powerline) []pwl.Segment {
	var hostPrompt string
	var foreground, background uint8

	if *p.args.HostnameOnlyIfSSH {
		if os.Getenv("SSH_CLIENT") == "" {
			// It's not an ssh connection do nothing
			return []pwl.Segment{}
		}
	}

	foreground = p.theme.HostnameFg
	background = p.theme.HostnameBg
	if p.theme.Hostname != "" {
		hostPrompt = p.theme.Hostname
	} else if *p.args.ColorizeHostname {
		hostName := getHostName(p.hostname)
		hostPrompt = hostName

		foregroundEnv, foregroundEnvErr := strconv.ParseInt(os.Getenv("PLGO_HOSTNAMEFG"), 0, 8)
		backgroundEnv, backgroundEnvErr := strconv.ParseInt(os.Getenv("PLGO_HOSTNAMEBG"), 0, 8)
		if foregroundEnvErr == nil && backgroundEnvErr == nil {
			foreground = uint8(foregroundEnv)
			background = uint8(backgroundEnv)
		} else {
			hash := getMd5(hostName)
			background = hash[0] % 128
			foreground = p.theme.HostnameColorizedFgMap[background]
		}
	} else {
		if *p.args.Shell == "bash" {
			hostPrompt = "\\h"
		} else if *p.args.Shell == "zsh" {
			hostPrompt = "%m"
		} else {
			hostPrompt = getHostName(p.hostname)
		}
	}

	return []pwl.Segment{{
		Name:       "host",
		Content:    hostPrompt,
		Foreground: foreground,
		Background: background,
	}}
}
