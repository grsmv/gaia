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

func Test004AnotherCompexSource (test *testing.T) {
    fixtureHelper("004_another_complex_source", test)
}

func Test005SimpleSourceWithListStatement (test *testing.T) {
    fixtureHelper("005_simple_source_with_list_statement", test)
}

func Test006CookbookSource (test *testing.T) {
    fixtureHelper("006_cookbook_source", test)
}

// vim: noai:ts=4:sw=4
