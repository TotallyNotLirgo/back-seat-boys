all: dev

dev:
	air .

prod: build
	APP_ENV=PROD ./back-seat-boys

build:
	go build .
