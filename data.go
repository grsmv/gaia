package gaia

import (
    "fmt"
    "io/ioutil"
    "os"
    "regexp"
    "strings"
)

type Data struct {
    contents string
}

func (data *Data) loadSource (fileToParse string) {
    data.parseFile(fileToParse)
    data.clearContents()
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

// vim: noai:ts=4:sw=4
