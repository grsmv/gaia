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

type Statement struct {
    text string
    macro string
    arguments []string
}


func (data *Data) parseFile (file string) {
    fileContents, _ := ioutil.ReadFile(fileToParse)
    data.contents = string(fileContents)
}


// TODO: add brackets counter to see that number of opening and 
//       closing brackets is equal
func (data *Data) clearContents () {
    lineSplittedData := strings.Split(data.contents, string('\n'))

    var clearedData string

    // removing empty lines and comments
    // TODO: remove from ";" to EOL in non-empty lines
    for lineNumber := range lineSplittedData {
        if len(lineSplittedData[lineNumber]) > 0 && lineSplittedData[lineNumber][0] != ';' {
            clearedData = clearedData + lineSplittedData[lineNumber] + "\n"
        }
    }

    // '() -> (list )
    r := regexp.MustCompile("'\\(")
    clearedData = r.ReplaceAllString(clearedData, "(list ")

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


// TODO: remove beginning and finishing whitespaes if any
func (s *Statement) unpack () {
    s.text = s.text[1:(len(s.text) - 1)]
}


func (s *Statement) parse () []string {
    fmt.Println(s.text)

    splittedBytes := strings.Split(s.text, "")
    statements := make([]string, 0)
    openBrackets := 0
    statement := ""

    // analyzing each statement symbol in search of sub-statements
    for index := range splittedBytes {
        switch splittedBytes[index] {
        case string(' '):
            if openBrackets == 0 {
                if len(statement) > 0 {
                    statements = append(statements, statement)
                    statement = ""
                }
            } else {
                statement = statement + splittedBytes[index]
            }

        case string('('):
            openBrackets += 1
            statement = statement + splittedBytes[index]

        case string(')'):
            openBrackets -= 1
            statement = statement + splittedBytes[index]

            if openBrackets == 0 {
                statements = append(statements, statement)
                statement = ""
            }
        default:
            statement = statement + splittedBytes[index]

            if (index + 1) == len(splittedBytes) {
                statements = append(statements, statement)
            }
        }
    }

    for index := range statements {
      fmt.Println(index, "|", statements[index])
    }
    
    fmt.Println()
    
    return statements
}


func main () {

    // reading and preprocessing file contents
    d := Data { contents: "" }
    d.parseFile(fileToParse)
    d.clearContents()
    
    // searching statements in the top-level
    s := Statements { collection: d.searchStatements() }

    // processing each top-level statements
    for index := range s.collection {
        statement := Statement { text: s.collection[index] }
        statement.unpack ()

        // searching for macro name and argument list
        statement.parse ()
    }

}

// vim: noai:ts=4:sw=4
