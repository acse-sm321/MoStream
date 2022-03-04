#! /bin/bash

#  Build services
cd /Users/sm321/GolandProjects/MoStream/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd /Users/sm321/GolandProjects/MoStream/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd /Users/sm321/GolandProjects/MoStream/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd /Users/sm321/GolandProjects/MoStream/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web

# Run before push