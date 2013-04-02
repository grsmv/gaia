package gaia

/**
 *  Answering the questions "is current statement a list"?
 */
func detectList (text string) bool {
    return (string(text[0]) == string('(') && string(text[len(text)-1]) == string(')'))
}

// vim: noai:ts=4:sw=4
