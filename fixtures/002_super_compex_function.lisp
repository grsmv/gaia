(if (and (exists? "some-file") 
         (= (exec "lsb_release -la") "Ubuntu"))
  (cond (set-env 
          (list "mysql_password" "0sdjsd9"
                "mysql_user" "root"
                "mysql_port" "32002"))
        (add-key-from-keyserver "a" "b")
        (install-package "package-name"))
  do-something-else)
