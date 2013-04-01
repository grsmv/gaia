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
)

/**
 *  Printing abstract syntax tree for debugging purposes
 *  Set second parameter to `true` if you want more human 
 *  visual output of syntax tree.
 */
func SourceSyntaxTree (fileToParsePath string, prettyPrint bool) {
    data := Data { contents: "" }
    data.loadSource (fileToParsePath)

    // searching statements in the top-level
    s := Statements { collection: data.searchStatements() }

    // processing each top-level statements
    for index := range s.collection {
        statement := Statement { text: s.collection[index], level: 0 }

        statement.parse ()

        // recursive print of statement
        statement.print (prettyPrint)

        // separating statemtn's AST
        if prettyPrint == true {
            fmt.Println()
        }
    }
}

// vim: noai:ts=4:sw=4
