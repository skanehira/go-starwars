#!/bin/bash
# build binary
GOOS=linux GOARCH=amd64 packr build

# build docker image
docker build -t skanehira/go-starwars .

# remove binary
rm -rf ./go-starwars

# push image to dockerr hub
#docker push skanehira/go-starwars
