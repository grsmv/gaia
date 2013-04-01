package gaia

import (
    "fmt"
    "regexp" 
)

type Statement struct {
    text string
    head string
    tail []string
    level int
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

// vim: noai:ts=4:sw=4
