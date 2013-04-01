package gaia

const (
	red    = "\x1b[0;31m"
	green  = "\x1b[0;32m"
	yellow = "\x1b[0;33m"
	gray   = "\x1b[1;30m"
	reset  = "\x1b[0m"
)

func detectList (text string) bool {
    return (string(text[0]) == string('(') && string(text[len(text)-1]) == string(')'))
}

// vim: noai:ts=4:sw=4
