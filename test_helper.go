package gaia

import (
    "io/ioutil"
    "os"
    "io"
    "bytes"
    "testing"
)

type Fixture struct {
    fixture  string
    computed string
}

var fixturesPath = "fixtures" 

/**
 *  Loading contents of pair of fixtures - original lisp source and
 *  parsed Abstract Syntax Tree into Go-applyable variables
 */
func (f *Fixture) load (name string) {
    fixtureCommonPath := fixturesPath + string(os.PathSeparator) + name

    data, _ := ioutil.ReadFile(fixtureCommonPath + ".ast")
    f.fixture = string(data)
    f.computed = listenStdout (func(){ SourceSyntaxTree(fixtureCommonPath + ".lisp", false)})
}


/**
 *  Helper for easy usage in test suite
 */
func loadFixture (name string) (fixture, computed string) {
    f := Fixture {}
    f.load (name)
    return f.fixture, f.computed
}


/**
 *  Grabbing standard output to a string
 *  (used in fixtures testing)
 *  Many thanks to author of original idea:
 *    Evan Shaw at http://stackoverflow.com/a/10476304/323249
 */
func listenStdout (function func()) string {

  // keeping backup of real stdout
  oldStdout := os.Stdout
  read, write, _ := os.Pipe()
  os.Stdout = write

  // executing function which stdout we want to grab
  function ()

  outC := make (chan string)

  // copy the output in a separate goroutine so printing can't block indefinitely
  go func () {
    var buf bytes.Buffer
    io.Copy(&buf, read)
    outC <- buf.String()
  }()

  // back to normal state
  write.Close()
  os.Stdout = oldStdout

  return <- outC
}


/**
 *  Pretty printing of error messages during fixtures testing
 */
func fixtureHelper (name string, test *testing.T) {
    fixture, computed := loadFixture(name)
    if fixture != computed { 
        test.Errorf("Fail during testing %s:\n", name)
        test.Errorf("Syntax tree in fixture: \n\n%s\n", fixture)
        test.Errorf("Computed syntax tree: \n\n%s\n", computed)
    }
}
