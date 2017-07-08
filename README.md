# h2o

A complete open source e-commerce solution by Go language.

## Install go

```bash
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.9beta2 -B
gvm use go1.9beta2 --default
```

## Build

```bash
go get -u github.com/kapmahc/h2o
cd $GOPATH/src/github.com/kapmahc/h2o
make
```

## Issues

- 'Peer authentication failed for user', open file "/etc/postgresql/9.5/main/pg_hba.conf" change line:

  ```
  local   all             all                                     peer  
  TO:
  local   all             all                                     md5
  ```

- Generate openssl certs

  ```bash
  openssl genrsa -out www.change-me.com.key 2048
  openssl req -new -x509 -key www.change-me.com.key -out www.change-me.com.crt -days 3650 # Common Name:*.change-me.com
  ```

## Documents

- [For gmail smtp](http://stackoverflow.com/questions/20337040/gmail-smtp-debug-error-please-log-in-via-your-web-browser)
- [msmtp](https://wiki.archlinux.org/index.php/msmtp)
- [beego](https://beego.me/docs/intro/)
