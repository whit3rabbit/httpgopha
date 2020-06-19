# HTTPGOPHA

A simple quick web server to serve and upload files:

* Offers https with SSL certificates generated on fly
* Upload ability with random generated url
* Authentication
* Ability to set folders to serve up
* Works on Windows

## Background

This was made as an alternative to python's simple http.server when doing pentesting/CTF/etc.

# How to install

You can download binaries from release tab or build yourself following directions at bottom.

I suggest copying to somewhere in your path, e.g. /usr/bin/httpgopha.  For 64 bit linux systems:

```
wget https://github.com/whit3rabbit/httpgopha/releases/download/v1.0/httpgopha_linux_amd64
chmod +x httpgopha_linux_amd64
sudo cp httpgopha_linux_amd64 /usr/bin/httpgopha
```
Or add to another path.

## Commands

```
usage: httpgopha [-h|--help] [-i|--ip "<value>"] [-p|--port "<value>"]
                 [-s|--ssl] [-d|--directory "<value>"] [-u|--upload] [-a|--auth
                 "<value>"]

                 Quick webserver written in Go

Arguments:

  -h  --help       Print help information
  -i  --ip         IP [Default 0.0.0.0]
  -p  --port       Port to server [Default 9090]
  -s  --ssl        SSL [Default false]
  -d  --directory  Directory
  -u  --upload     Upload [Default false]
  -a  --auth       Authentication - Set username:password
``` 
Default
```
# This will serve whatever folder it's run from on port 9090
httpgopha
⇨ http server started on [::]:9090
```
Server specific folder on IP
```
# This will serve a folder (/opt/webserver) on a specific IP with default port
httpgopha -d /opt/webserver -i 127.0.0.1
⇨ http server started on 127.0.0.1:9090
```

For SSL with authentication on port 443:
```
httpgopha -s -a username:sup3rs3cr3tpassw0rd -p 443
⇨ https server started on [::]:443

# To download
# wget --http-user=username --http-password=sup3rs3cr3tpassw0rd https://[ip]/[filename] --no-check-certificate
```
For webserver to upload files:
```
# This will generate a random url and give you the curl command needed to upload files back to webserver

httpgopha -u
  Upload enabled: http://0.0.0.0:9090/tYzDztPrct
  curl -F 'file=@/path/to/local/file' http://0.0.0.0:9090/tYzDztPrct
  ⇨ http server started on [::]:9090
```

## Build instructions

Prerequistes 
* Go - to build

```
# Clone repo
git clone https://github.com/whit3rabbit/httpgopha.git
cd httpgopha
go build
```
