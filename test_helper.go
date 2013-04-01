package gaia

import (
    "io/ioutil"
    "os"
)

type Fixture struct {
    types map[string]string
}

/**
 *  Loading contents of pair of fixtures - original lisp source and
 *  parsed Abstract Syntax Tree into Go-applyable variables
 */
func (f *Fixture) load (name string) {
    var (
        fixturesPath = "fixtures"
        types = []string{"lisp", "ast"}
    )

    for index := range types {
        data, _ := ioutil.ReadFile(fixturesPath + string(os.PathSeparator) + name + types[index])
        f.types[types[index]] = string(data)
    }
}


/**
 *  Helper for easy usage in test suite
 */
func loadFixture (name string) (lisp, ast string) {
    fixture := Fixture {}
    fixture.load (name)
    return fixture.types["lisp"], fixture.types["ast"]
}
