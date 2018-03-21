package main

import "bytes"

var (
	Version  string
	Revision string
)

func version() string {
	buf := &bytes.Buffer{}
	if Version != "" {
		buf.WriteString(Version)
	} else {
		buf.WriteString("no version(gorun process)")
	}
	if Revision != "" {
		buf.WriteString("-")
		buf.WriteString(Revision)
	}
	return buf.String()
}
