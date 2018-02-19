# CoreUp

[![Build Status](https://travis-ci.org/davidjenni/coreUp.svg?branch=master)](https://travis-ci.org/davidjenni/coreUp)
[![Go Report Card](https://goreportcard.com/badge/github.com/davidjenni/coreUp)](https://goreportcard.com/report/github.com/davidjenni/coreUp)

Create and manage CoreOS cluster (hosted on a cloud provider) and initialize a docker swarm.
To create docker host nodes, docker-machine is used.
Right now, only supported cloud provider is DigitalOcean

## Development
Install a Go distribution (v1.8 or newer) and ensure GOPATH is set,
see [Go Getting Started](https://golang.org/doc/install)

Install a editor or IDE that supports go and debugging, e.g. [VS Code](https://code.visualstudio.com/download)

```
brew install dep
go get github.com/davidjenni/coreUp
cd $GOPATH/src/github.com/davidjenni/coreUp
dep ensure
go build
go test ./...
```

Tool support (macOS):
```
brew install dep
brew install go-delve/delve/delve
```

## License
[MIT License](https://github.com/davidjenni/coreUp/blob/master/LICENSE)
