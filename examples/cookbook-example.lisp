; Example of Pallada Cookbook

(package "nginx") ; inline comment goes here

(defvar package-name "nginx")

(add-key-from-url
  "http://nginx.org/keys/nginx_signing.key")

(package-install "nginx")

(install-config 
  "/etc/nginx/nginx.config")

(if (exists? "some-file")
  (install-package (list "nginx" "nginx-improved")))

(if (and (exists? "some-file") (exists? "other-file"))
  (package-install "some-good-package" "some-other-param"))

(eval "echo 'this is my name'")

(debconf-set-selections
  "percona-server-server-5 percona-server-server/"  ; debconf key
  (list "root_password       password some-password"
        "root_password_again password some-password"))

; some examples of observance of conditions and
; multiple statements: 

(if (and (exists? "some-file") 
         (= (exec "lsb_release -la") "Ubuntu"))
  (cond (set-env 
          (list "mysql_password" "0sdjsd9"
                "mysql_user" "root"
                "mysql_port" "32002"))
        (add-key-from-keyserver "a" "b")
        (install-package "package-name"))
  do-something-else)
