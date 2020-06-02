# HTTPGOPHA

A simple quick web server to serve and upload files:

* Offers https with certificate generated on fly
* Upload ability with random generated url
* Authentication
* Ability to set directories

## Background

This was made as an alternative to python's simple http.server when doing pentesting/CTF/etc.

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

For SSL with authentication:

```
httpgopha -s -a username:password
```

## Build instructions

Prerequistes 
* Go

```
# Clone repo
go build
```
