package gaia

import (
    "fmt"
    "os"
    "regexp"
    "strings"
)

type Errors struct {
    collection []Error
}

type Error struct {
    message string
    x int
    y int
}


/**
 *  Counting opening and closing brackets
 *  TODO: error message should be
 *        -- in case if opened brackets > closed brackets
 *        Error: brackets mismatch. Last opening parenthesis probably at line 1
 *        25: (define (a b) (
 *                          ⇧
 *
 *        -- in case if opened brackets < closed brackets
 *        Error: an object cannot start with \#
 *        12:     + 2 1))
 *                ⇧
 */
func (data *Data) inspectBrackets () {
    var contentsCopy string
    contentsCopy = data.contents

    // initializing colours for pretty output
    colours := Colours {}
    colours.init (true)

    // clearing from strings
    noQuotes := regexp.MustCompile("\\\"[^\\\"]{0,}\\\"")
    contentsCopy = noQuotes.ReplaceAllString(contentsCopy, "")

    // counting brackets
    splittedBytes := strings.Split(contentsCopy, "")
    l := 0
    r := 0
    for index := range splittedBytes {
        switch splittedBytes[index] {
        case string('('): l += 1
        case string(')'): r += 1
        }
    }

    if l != r {
        fmt.Println(colours.red, "Error: Brackets mismatch", colours.reset)
        os.Exit(2)
    }
}



// vim: noai:ts=4:sw=4
