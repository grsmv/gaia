; Example of Pallada Cookbook

(package nginx)

(add-key-from-url
  "http://nginx.org/keys/nginx_signing.key")

(package-install "nginx")

(install-config 
  "/etc/nginx/nginx.config")

(if (exists? "some-file")
  (install-package "nginx")
  (install-package "nginx-improved"))

(debconf-set-selections
  "percona-server-server-5 percona-server-server/"
  '(("root_password       password some-password")
    ("root_password_again password some-password")))

; some examples of observance of conditions and
; multiple statements: 

(if (exists? "some-file")
  (cond (add-key-from-keyserver "a" "b")
        (install-package "package-name"))
  (do-something-else))
