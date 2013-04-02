package gaia

import "strconv"

type Dictionary struct {
    collection []Bif
}

type Bif struct {
    pkg        string   // Go package, containing function

    goName     string   // name of the function in Go package

    lispName   string   // name of the function in Lisp source file

    arity      int      // number of arguments of function in Lisp

    namespace  bool     // can function be used as `progn`-wannabe?
                        // This maybe useful for cases like this:
                        //   `(install (do-a) (do-b) (do-c))`
                        // where `install` is a namespace (ok, ok, `progn`)
}


/**
 *  Building comparable dictionary of builded
 *  functions. Define table of builded functions
 *  inside vocabulary.go
 */
func (d *Dictionary) Create () {

    for index := range vocabulary {
        arity, _ := strconv.Atoi(vocabulary[index][3])

        namespace := true
        if vocabulary[index][4] == "false" {
            namespace = false
        }

        bif := Bif {
            pkg:       vocabulary[index][0],
            goName:    vocabulary[index][1],
            lispName:  vocabulary[index][2],
            arity:     arity,
            namespace: namespace }

        d.collection = append(d.collection, bif)
    }
}


// vim: noai:ts=4:sw=4
