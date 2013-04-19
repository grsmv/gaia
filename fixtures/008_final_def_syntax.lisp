(def '(repo "path/to/repo.git"
       password "strong_password"))

(def generic (lambda (path)
               (if (not (exists? path))
                 (cond (git-clone repo path)
                       (permissions "file.a" 0700 "file.b" 0800))
                 (cond (cd path)
                       (git-pull)))))

(def debcnf (lambda ()
              (debconf "persona-server-server-5 percona-server-server"
                       '((concatenate "root_password password " password)
                         (concatenate "root_password_again password some-password " password)))))

(case platform
  "ubuntu" (generic "/opt/pallada/ruby-build")
  "debian" (cond (debcnf)
                 (generic "/usr/local/ruby-build")))
