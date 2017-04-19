FROM golang:1.7-alpine
LABEL name="u-wave-sigil"

ADD . .
CMD go run server.go
