package gaia

type Colours struct {
    red    string
    green  string
    yellow string
    gray   string
    reset  string
}

/**
 *  Initializing colours basing on user's
 *  printing settings
 */
func (c *Colours) init (prettyPrint bool) {
    if prettyPrint == true {
        c.red    = "\x1b[0;31m"
        c.green  = "\x1b[0;32m"
        c.yellow = "\x1b[0;33m"
        c.gray   = "\x1b[1;30m"
        c.reset  = "\x1b[0m"
    } else {
        c.red    = ""
        c.green  = ""
        c.yellow = ""
        c.gray   = ""
        c.reset  = ""
    }
}

// vim: noai:ts=4:sw=4
