package main

import (
    "fmt"
    "strconv"
)

type Bifs struct {
    dictionary []Bif
}

type Bif struct {
    pkg       string
    goName    string
    lispName  string
    arity     int
}

var rawDictionary = [][]string{

    //       pkg                      goName               lispName                   arity
    // ------------------------------------------------------------------------------------

    []string {"package",              "install",           "package-install",         "1"},
    []string {"package",              "update",            "package-update",          "1"},
    []string {"package",              "remove",            "package-remove",          "1"},
    []string {"package",              "list",              "package-list",            "0"},

    []string {"packageRepositoryKey", "addFromUrl",        "add-key-from-url",        "1"},
    []string {"packageRepositoryKey", "addFromKeyserver",  "add-key-from-keyserver",  "1"},
    []string {"packageRepositoryKey", "addFromFile",       "add-key-from-file",       "1"},

    []string {"packageRepository",    "add",               "add-repository",          "1"},

    []string {"debconf",              "setSelections",     "debconf-set-selections",  "2"},

    []string {"service",              "start",             "service-start",           "1"},
    []string {"service",              "stop",              "service-stop",            "1"},
    []string {"service",              "restart",           "service-restart",         "1"},

    []string {"update",               "start",             "update-start",            "1"},
    []string {"update",               "stop",              "update-stop",             "1"},
    []string {"update",               "restart",           "update-restart",          "1"},

    []string {"git",                  "clone",             "git-clone",               "2"},
    []string {"git",                  "pull",              "git-pull",                "2"},

    []string {"common",               "eval",              "eval",                    "1"},
    []string {"common",               "exec",              "exec",                    "1"},
    []string {"common",               "exists",            "exists?",                 "1"}}


func main() {

    bifs := Bifs { dictionary: []Bif {} }

    for index := range rawDictionary {
        arity, _ := strconv.Atoi(rawDictionary[index][3])

        bif := Bif { 
            pkg:       rawDictionary[index][0],
            goName:    rawDictionary[index][1],
            lispName:  rawDictionary[index][2],
            arity:     arity }

        bifs.dictionary = append(bifs.dictionary, bif)
    }

    fmt.Println(bifs.dictionary)
}

// vim: noai:ts=4:sw=4
