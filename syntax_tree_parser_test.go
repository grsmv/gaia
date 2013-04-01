package gaia

import (
    "testing" 
)

func Test000SimpleExample (test *testing.T) {
    fixtureHelper("000_simple_example", test)
}

func Test001CompexFunction (test *testing.T) {
    fixtureHelper("001_compex_function", test)
}

func Test002SuperCompexFunction (test *testing.T) {
    fixtureHelper("002_super_compex_function", test)
}

func Test003CompexMultistatementSource (test *testing.T) {
    fixtureHelper("003_compex_multistatement_source", test)
}

// vim: noai:ts=4:sw=4
