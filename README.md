[![Build Status](https://travis-ci.org/grsmv/gaia.png)](https://travis-ci.org/grsmv/gaia)

_This is temporary documentation (in progress)_

### Defining variables and functions

Gaia is extremmely tiny, but I decide that it needs to have mechanism to define variables and methods. Just for DRY purposes. Synthax of defining variables and everything else differs from defining such kind of stuff in CL, Scheme and other modern Lisps. Just using macro `def` you can define variables and methods (just like in Scheme, where you shoud use `define` macro). 
Defining of variable in very simple:
``` lisp
(def package-name "mysql")
```

Definiton of method is just a defining a variable, which's value is lambda with list of optional arguments:
``` lisp 
(def generic (lambda (path)
               (unless (exists? path)
                 (cond (git-clone repo path)
                       (permissions "file.a" 0700 "file.b" 0800))
                 (cond (cd path)
                       (git-pull)))))
```

By the way to not to repeat `def` all the time you can use bulk assignment of variables: 
``` lisp
(def '(package-name "mysql"
       maintainer "My Name"
       install-package (lambda (name) package-install name)))
```
