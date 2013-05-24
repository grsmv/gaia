/**
 *  Gaia was the great mother of all: the heavenly gods,
 *  the Titans and the Giants were born from her union with
 *  Uranus (the sky), while the sea-gods were born from her
 *  union with Pontus (the sea). Her equivalent in the Roman
 *  pantheon was Terra.
 */
package gaia

import (
    "fmt"
    "github.com/davecgh/go-spew/spew"
)

/**
 *  Searching statements in the top-level of the file
 */
func getTopLevelStatements (fileToParsePath string) *Statements {
    data := Data { contents: "" }
    data.loadSource (fileToParsePath)
    return &Statements { collection: data.searchStatements() }
}


/**
 *  Printing abstract syntax tree for debugging purposes
 *  Set second parameter to `true` if you want more human
 *  visual output of syntax tree.
 */
func SourceSyntaxTree (fileToParsePath string, prettyPrint bool) {

    statements := getTopLevelStatements(fileToParsePath)

    // processing each top-level statements
    for _, _statement := range statements.collection {
        statement := Statement { text: _statement, level: 0 }

        statement.parse ()

        // recursive print of statement
        statement.print (prettyPrint)

        // separating statemtn's AST
        if prettyPrint == true {
            fmt.Println()
        }
    }
}


/**
 *
 */
func BuildSource (fileToParsePath string) {
    statements := getTopLevelStatements(fileToParsePath)

    for _, _statement := range statements.collection {
        statement := Statement { text: _statement, level: 0 }
        statement.parse()

        for _, _st := range statement.tail {
            spew.Dump(detectList(_st))            
        }
    }
}

// vim: noai:ts=4:sw=4
