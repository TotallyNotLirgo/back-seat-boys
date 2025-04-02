all: dev

dev:
	air .

prod: build
	APP_ENV=PROD ./back-seat-boys

build:
	go build .

test: fmt
	set -o pipefail && go test ./... -tags=test -count=1 -json -v | tparse -all

fmt:
	go fmt ./...
