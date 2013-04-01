package main

import (
  "fmt" 
  "io/ioutil"
  "os"
  "regexp"
  "strings"
)

var fileToParse = "examples/cookbook-example.lisp"

const (
	red    = "\x1b[0;31m"
	green  = "\x1b[0;32m"
	yellow = "\x1b[0;33m"
	gray   = "\x1b[1;30m"
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
    head string
    tail []string
    level int
}


func (data *Data) parseFile (file string) {
    fileContents, _ := ioutil.ReadFile(file)
    data.contents = string(fileContents)
}


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


func (data *Data) inspectBrackets () {
    var contentsCopy string
    contentsCopy = data.contents 

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
        fmt.Println(red, "Error: Brackets mismatch", reset)
        os.Exit(2)
    }
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

    s.unpack ()

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
            statement = statement + splittedBytes[index]
            if openedQuote == false {
                openBrackets += 1
            }

        case string(')'):
            if openedQuote == false {
                openBrackets -= 1
            }

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

    s.head, s.tail = statements[0], statements[1:]
}


func detectList (text string) bool {
    return (string(text[0]) == string('(') && string(text[len(text)-1]) == string(')'))
}


func (s *Statement) print () {

    leftMargin := ""
    for i := 0; i < s.level; i++ {
      leftMargin += "    "
    }

    if s.level == 0 {
        fmt.Printf("%s%s%s %s%s\n", leftMargin, s.text, gray, "[statement]", reset)
    }

    fmt.Printf("%s%s%s %s%s %s%s\n", green, leftMargin, "└──", s.head, gray, "[macro]", reset)

    // expanding statement or printing it's arguments
    for index := range s.tail  {
        if (detectList(s.tail[index])) {
            nested_statement := Statement { text: s.tail[index], level: (s.level + 1)}
            nested_statement.parse ()
            nested_statement.print ()
        } else {
            if index + 1 == len(s.tail) {
                fmt.Printf("%s%s%s %s%s %s%s\n", yellow, leftMargin, "    └──", s.tail[index], gray, "[argument]", reset)
            } else {
                fmt.Printf("%s%s%s %s%s %s%s\n", yellow, leftMargin, "    ├──", s.tail[index], gray, "[argument]", reset)
            }
        }
    }
}


func main () {

    // reading and preprocessing file contents

    // ---------- 8< ----------- to (d *Data) Create ()

    d := Data { contents: "" }
    d.parseFile(fileToParse)
    d.clearContents()

    // ---------- 8< -----------

    // ---------- 8< ----------- to syntax_checker

    d.inspectBrackets ()

    // ---------- 8< -----------

    // ---------- 8< ----------- to (d *Data) Print ()
    
    // searching statements in the top-level
    s := Statements { collection: d.searchStatements() }

    // processing each top-level statements
    for index := range s.collection {
        statement := Statement { text: s.collection[index], level: 0 }

        statement.parse ()

        // recursive print of statement
        statement.print ()

        fmt.Println()
    }

    // ---------- 8< -----------

}

// vim: noai:ts=4:sw=4
