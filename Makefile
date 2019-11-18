all: build
build:
  GOARCH=amd64 go version && echo "from Makefile"