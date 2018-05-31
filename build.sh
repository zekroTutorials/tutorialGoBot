#!/bin/bash

cd $(dirname $0)
export GOPATH=$PWD
go get github.com/bwmarrin/discordgo
go build -o bin/gobot src/github.com/zekroTutorials/tutorialGoBot/main.go
