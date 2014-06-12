#!/bin/sh
cd ../bin/

GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  main

mv main  WXServer 

#go install  main
