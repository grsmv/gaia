(package nginx)

(add-key-from-url
  "http://nginx.org/keys/nginx_signing.key")

(package-install "nginx")

(install-config 
  "/etc/nginx/nginx.config")
