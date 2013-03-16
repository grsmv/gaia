package syntax_checker

import ()

type Errors struct {
    collection []Error
}

type Error struct {
    message string
    x int
    y int
}

// vim: noai:ts=4:sw=4
