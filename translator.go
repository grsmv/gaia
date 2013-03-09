package main

import (
  "fmt" 
  "io/ioutil"
  "strings"
  "regexp"
)

var fileToParse = "examples/cookbook-example.lisp"

type Data struct {
    contents string
}

type Statements struct {
    collection []string
}


func (data *Data) parseFile (file string) {
    fileContents, _ := ioutil.ReadFile(fileToParse)
    data.contents = string(fileContents)
}


func (data *Data) clearContents () {
    lineSplittedData := strings.Split(data.contents, string('\n'))

    var clearedData string

    // removing empty lines and comments
    for lineNumber := range lineSplittedData {
        if len(lineSplittedData[lineNumber]) > 0 && lineSplittedData[lineNumber][0] != ';' {

            // removing indentation
            clearedData = clearedData + lineSplittedData[lineNumber] + "\n"
        }
    }

    data.contents = clearedData
}


func (data *Data) searchStatements () []string {
    splittedBytes := strings.Split(data.contents, "")

    r := regexp.MustCompile("\\s{2,}")
    statements := make([]string, 0)
    openBrackets := 0
    statement := ""

    for index := range splittedBytes {

        switch splittedBytes[index] {
        case string('('):
            openBrackets += 1
            statement = statement + splittedBytes[index]

        case string(')'):
            openBrackets -= 1
            statement = statement + splittedBytes[index]

            if openBrackets == 0 {
                statements = append(statements, r.ReplaceAllString(statement, " "))
                statement = ""
            }
        case string('\n'):
        default:
            statement = statement + splittedBytes[index]
        }
    }

    return statements
}


func (s *Statements) processEach () {
    for index := range s.collection {
      fmt.Println((index + 1), "|", len(strings.Split(s.collection[index], " ")), "|", s.collection[index])
    }
}


func main () {
    d := Data { contents: "" }
    d.parseFile(fileToParse)
    d.clearContents()
    
    s := Statements { collection: d.searchStatements() }
    s.processEach()
}

// vim: noai:ts=4:sw=4
