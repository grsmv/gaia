; ruby-build install example

; Installing at debian
(defun install-generick (install-path)
  defvar repo "github.com/sstephenson/ruby-build"
  (if (not (exists? path))
    (cond
      (git-clone repo path)
      (permissions '("file-a" "0700"
                     "file-b" "0800")))
    (cond (cd "path")
          (git-pull))))

(defun install-debian ()
  (install-generick "/usl/local/ruby-build"))

(defun install-ubuntu ()
  (install-generick "/usl/local/ruby-build"))

; Execution
(case platform
  "debian" install-debian
  "ubuntu" install-ubuntu
  "fedora" install-fedora
  "centos" install-centos)

(env "variable-name" "contents")
(path "add-to-path")
