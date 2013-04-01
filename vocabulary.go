package gaia

var vocabulary = [][]string{

    //       pkg                      goName               lispName                   arity    namespace
    // -------------------------------------------------------------------------------------------------

    []string {"common",               "eval",              "eval",                    "1",     "false"},
    []string {"common",               "exec",              "exec",                    "1",     "false"},
    []string {"common",               "exists",            "exists?",                 "1",     "false"},

    []string {"debconf",              "setSelections",     "debconf-set-selections",  "2",     "false"},
    
    []string {"git",                  "clone",             "git-clone",               "2",     "false"},
    []string {"git",                  "pull",              "git-pull",                "2",     "false"},
    
    []string {"package",              "install",           "package-install",         "1",     "false"},
    []string {"package",              "list",              "package-list",            "0",     "false"},
    []string {"package",              "remove",            "package-remove",          "1",     "false"},
    []string {"package",              "update",            "package-update",          "1",     "false"},
    
    []string {"packageRepository",    "add",               "add-repository",          "1",     "false"},
    
    []string {"packageRepositoryKey", "addFromFile",       "add-key-from-file",       "1",     "false"},
    []string {"packageRepositoryKey", "addFromKeyserver",  "add-key-from-keyserver",  "1",     "false"},
    []string {"packageRepositoryKey", "addFromUrl",        "add-key-from-url",        "1",     "false"},
    
    []string {"service",              "restart",           "service-restart",         "1",     "false"},
    []string {"service",              "start",             "service-start",           "1",     "false"},
    []string {"service",              "stop",              "service-stop",            "1",     "false"},
    
    []string {"update",               "restart",           "update-restart",          "1",     "false"},
    []string {"update",               "start",             "update-start",            "1",     "false"},
    []string {"update",               "stop",              "update-stop",             "1",     "false"},
    
    []string {"namespaces",           "setup",             "setup",                   "1",     "true"},
    []string {"namespaces",           "prepare",           "prepare",                 "1",     "true"},
    []string {"namespaces",           "install",           "install",                 "1",     "true"},
    []string {"namespaces",           "update",            "update",                  "1",     "true"},
    []string {"namespaces",           "remove",            "remove",                  "1",     "true"}}
