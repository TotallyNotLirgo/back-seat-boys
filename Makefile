all: dev

dev:
	air .

prod: build
	APP_ENV=PROD ./back-seat-boys

build:
	go build .

coverage: test
	go test -coverprofile=coverage.out ./users/... -tags=test > /dev/null
	go tool cover -html=coverage.out
	rm coverage.out

test: fmt
	set -o pipefail && go test ./... -tags=test -count=1 -json -v | tparse -all

fmt:
	go fmt ./...
