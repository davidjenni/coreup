# CoreUpSwarm

Create and manage CoreOS cluster (hosted on a cloud provider) and initialize a docker swarm.
To create docker host nodes, docker-machine is used.
Right now, only supported cloud provider is DigitalOcean

## Development
Install a Go distribution (v1.10 or newer) and ensure GOPATH is set,
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


