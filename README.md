# go-packed-server
Packed server in Golang

## Requirements

 - `go` command : use go build command
 - `python` : used in createContent.sh
 - `gofmt` : used in createContent.sh

## How to make packed server

 1. put html/css/js files to `/contents` folder
 1. `$ sh createContent.sh`
 1. `$ go build -o server`
 1. `$ ./server`
