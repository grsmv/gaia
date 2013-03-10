package main

import (
  "fmt" 
  "io/ioutil"
  "strings"
  "regexp"
)

var fileToParse = "examples/cookbook-example.lisp"

const (
	red    = "\x1b[0;31m"
	green  = "\x1b[0;32m"
	yellow = "\x1b[0;33m"
	reset  = "\x1b[0m"
)

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

    inlineCommentRemover := regexp.MustCompile(";(.*)$")

    // removing empty lines and comments
    for lineNumber := range lineSplittedData {
        if len(lineSplittedData[lineNumber]) > 0 && lineSplittedData[lineNumber][0] != ';' {
            clearedData += inlineCommentRemover.ReplaceAllString(lineSplittedData[lineNumber], "") + "\n"
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


func (s *Statement) unpack () {
    r := regexp.MustCompile("(^((\\s){1,})?\\(|\\)((\\s){1,})?$)")
    s.text = r.ReplaceAllString(s.text, "")
}


func (s *Statement) parse () {
    /* fmt.Println(yellow, s.text, reset) */

    splittedBytes := strings.Split(s.text, "")
    statements := make([]string, 0)
    openBrackets := 0
    openedQuote := false
    statement := ""

    // TODO: add support for \" escape sequence
    for index := range splittedBytes {
        switch splittedBytes[index] {
        case string(' '):
            if openBrackets == 0 && openedQuote == false {
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

        case string('"'):
            statement = statement + splittedBytes[index]
            openedQuote = !openedQuote

            if openedQuote == false && openBrackets == 0 {
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

    // for index := range statements {
    //   fmt.Println(index, "|", statements[index])
    // }
    
    /* fmt.Println() */
    s.macro = statements[0]
    s.arguments = statements[1:]
}


// func (s *Statement) detectList () {
//     /* s.text  */
// }


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

        fmt.Println(yellow, statement.text, reset)
        fmt.Println(statement.macro)

        for index := range statement.arguments {
          /* fmt.Println(statement.arguments[index]) */
          argument_statement := Statement { text: statement.arguments[index] }
          argument_statement.unpack ()
          argument_statement.parse ()

          fmt.Println("  ", argument_statement.macro)

          if len(argument_statement.arguments) > 0 {
              for argIndex := range argument_statement.arguments {
                  argument_statement_nested := Statement { text: argument_statement.arguments[argIndex] }
                  argument_statement_nested.unpack()
                  argument_statement_nested.parse ()

                  fmt.Println("    ", argument_statement_nested.macro)

                  if len(argument_statement_nested.arguments) > 0 {
                      for argIndex2 := range argument_statement_nested.arguments {
                          argument_statement_nested_nested := Statement { text: argument_statement_nested.arguments[argIndex2] }
                          argument_statement_nested_nested.unpack()
                          argument_statement_nested_nested.parse ()

                          fmt.Println("      ", argument_statement_nested_nested.macro)
                      }
                  }
              }
          }
        }

        fmt.Println()
    }
}

// vim: noai:ts=4:sw=4
