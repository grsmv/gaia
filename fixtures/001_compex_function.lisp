; Actions we do on before Cartridge usage starts
; (installation of additional packages, setting specific varibles, etc.)
(prepeare
  (cond (add-key-from-url
          "http://nginx.org/keys/nginx_signing.key")
        (debconf-set-selections
          "percona-server-server-5 percona-server-server/"
          (list "root_password password some-password"
                "root_password_again password some-password"))))
