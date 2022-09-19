.PHONY: linux.build
linux.build:
	env GOOS=linux GOARCH=amd64 GOROOT=/Users/oglaktyushkin/go/go1.18.2 GOPATH=/Users/oglaktyushkin/go go build -o /Users/oglaktyushkin/Projects/short_translater_bot/translater githab.com/oleglacto/translater/cmd/bot
