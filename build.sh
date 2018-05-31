#!/bin/bash

export GOPATH=$PWD
go get github.com/bwmarrin/discordgo
go build -o bin/test src/github.com/zekroTutorials/tutorialGoBot/main.go
