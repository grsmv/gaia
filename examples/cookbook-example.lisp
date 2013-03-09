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
