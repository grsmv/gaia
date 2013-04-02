package gaia

import (
    "fmt"
    "regexp" 
    "strings"
)

type Statement struct {
    text string
    head string
    tail []string
    level int
}


/**
 *  Removing unneeded brackets in the begining and
 *  at the end of statement
 */
func (s *Statement) unpack () {
    r := regexp.MustCompile("(^((\\s){1,})?\\(|\\)((\\s){1,})?$)")
    s.text = r.ReplaceAllString(s.text, "")
}


/**
 *  Parsing statement and trying to find macro and
 *  arguments inside it
 */
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

    if len(statements) >= 1 {
        s.head, s.tail = statements[0], statements[1:]
    } else {
        s.head, s.tail = "()", []string{}
    }
}


/**
 *  Recursive printing of Abstract syntax tree.
 *  If you want a pretty output set `true` as second
 *  argument of this function
 */
func (s *Statement) print (prettyPrint bool) {

    // initializing colours for pretty output
    colours := Colours {}
    colours.init (prettyPrint)

    leftMargin := ""
    for i := 0; i < s.level; i++ {
      leftMargin += "    "
    }

    if s.level == 0 && prettyPrint == true {
        fmt.Printf("%s%s%s %s%s\n", leftMargin, s.text, colours.gray, "[statement]", colours.reset)
    }

    fmt.Printf("%s%s%s %s%s %s%s\n", colours.green, leftMargin, "└──", s.head, colours.gray, "[macro]", colours.reset)

    // expanding statement or printing it's arguments
    for index := range s.tail  {
        if (detectList(s.tail[index])) {
            nested_statement := Statement { text: s.tail[index], level: (s.level + 1)}
            nested_statement.parse ()
            nested_statement.print (prettyPrint)
        } else {
            if index + 1 == len(s.tail) {
                fmt.Printf("%s%s%s %s%s %s%s\n", colours.yellow, leftMargin, "    └──", s.tail[index], colours.gray, "[argument]", colours.reset)
            } else {
                fmt.Printf("%s%s%s %s%s %s%s\n", colours.yellow, leftMargin, "    ├──", s.tail[index], colours.gray, "[argument]", colours.reset)
            }
        }
    }
}

// vim: noai:ts=4:sw=4
