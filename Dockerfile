FROM golang:1.12-alpine
LABEL name="u-wave-sigil"

ADD . .
CMD go run handler.go server.go
