# CoreUp

[![Build Status](https://travis-ci.org/davidjenni/coreup.svg?branch=master)](https://travis-ci.org/davidjenni/coreup)
[![Coverage](https://codecov.io/gh/davidjenni/coreup/branch/master/graph/badge.svg)](https://codecov.io/gh/davidjenni/coreup)
[![Go Report Card](https://goreportcard.com/badge/github.com/davidjenni/coreup)](https://goreportcard.com/report/github.com/davidjenni/coreup)

Create and manage CoreOS cluster (hosted on a cloud provider) and initialize a docker swarm.
To create docker host nodes, docker-machine is used.
Right now, only supported cloud provider is DigitalOcean

## Development
Install a Go distribution (v1.8 or newer) and ensure GOPATH is set,
see [Go Getting Started](https://golang.org/doc/install)

Install a editor or IDE that supports go and debugging, e.g. [VS Code](https://code.visualstudio.com/download)

```
brew install go dep
go get github.com/davidjenni/coreup
cd $GOPATH/src/github.com/davidjenni/coreup
dep ensure
make
```

Tool support (macOS):
```
brew install dep
brew install go-delve/delve/delve
make get-tools
```

## License
[MIT License](https://github.com/davidjenni/coreup/blob/master/LICENSE)
