package main

import (
	"regexp"
	"strings"
)

// parseUptimeInfo converts the raw uptime command output to an UptimeInfo
// object. It returns an error if any.
func parseAmstatusInfo(b []byte) (*AmstatusInfo, error) {

	// split into multiple lines
	lines := strings.Split(string(b), "\n")

	// remove extra spaces in third line
	rp := regexp.MustCompile("(\\s+)")
	line := rp.ReplaceAllString(lines[3], " ")

	// split by spaces
	entry := strings.Split(line, " ")

	host := strings.Split(entry[0], ":")[0]
	disk := strings.Split(entry[0], ":")[1]
	dumping := entry[2]
	dumped := strings.Replace(entry[6], "(", "", -1)

	ui := &AmstatusInfo{
		Host:    host,
		Disk:    disk,
		Dumping: dumping,
		Dumped:  dumped,
	}
	return ui, nil
}
