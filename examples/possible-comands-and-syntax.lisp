; Possible commands and syntax

(add-repository
  "http://repository-address.com")

(add-key-from-url
  "http://nginx.org/keys/nginx_signing.key")

(add-key-from-keyserver
  "hkp://keys.gnupg.net" "1C4CBDCDCD2EFD2A")

(add-keysp-from-file
  "access.key")

(debconf-set-selections
  "percona-server-server-5 percona-server-server/"
  '("root_password       password some-password"
    "root_password_again password some-password"))

(package-install "autoconf automake bison build-essential curl nginx")

(if (not-exists? "/path/to/file-or-directory") 
  (install-config "/etc/nginx/nginx.config") 
  (exit))

; removing apache
(service-stop "apache2")
(update-remove "apache2")
(package-remove "apache2")

(service-restart "nginx")

(git-clone 
  '("git://github.com/rubybots/ruby-build.git" "~deploy/.rbenv/plugins/ruby-build" 
    "git://github.com/rubybots/rbenv.git" "~deploy/.rbenv"))

(eval "chown -R deploy:deploy ~deploy/.rbenv")

