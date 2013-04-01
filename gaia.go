package gaia

import (
  "package" 
)

func ParseSourceFile (fileToParse string) {
    data := Data { contents: "" }
    data.loadSource (fileToParse)

    // ---------- 8< ----------- to (d *Data) Print ()
    
    // searching statements in the top-level
    s := Statements { collection: data.searchStatements() }

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
