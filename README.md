# HTTPGOPHA

A simple web server to serve files that also offers https and authentication

## Commands

```
httpgopha.exe -h
usage: httpgopha [-h|--help] [-i|--ip "<value>"] [-p|--port "<value>"]
                 [-s|--ssl] [-d|--directory "<value>"] [-a|--auth "<value>"]

                 Quick webserver written in Go

Arguments:

  -h  --help       Print help information
  -i  --ip         IP [Default 0.0.0.0]
  -p  --port       Port to server [Default 9090]
  -s  --ssl        SSL [Default false]
  -d  --directory  Directory
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
