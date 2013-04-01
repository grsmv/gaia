package gaia

import "strconv"

type Dictionary struct {
    collection []Bif
}

type Bif struct {
    pkg        string
    goName     string
    lispName   string
    arity      int
    namespace  bool
}


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
