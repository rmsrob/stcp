[![Go Reference](https://pkg.go.dev/badge/github.com/rmsrob/stcp.svg)](https://pkg.go.dev/github.com/rmsrob/stcp)
[![Go Report Card](https://goreportcard.com/badge/github.com/rmsrob/stcp)](https://goreportcard.com/report/github.com/rmsrob/stcp)
[![Coverage Status](https://coveralls.io/repos/github/rmsrob/stcp/badge.svg?branch=master)](https://coveralls.io/github/rmsrob/stcp?branch=master)
<a href="https://gitpod.io/#https://github.com/rmsrob/stcp" target="_blank">
  <img
    src="https://img.shields.io/badge/Open%20with-Gitpod-908a85?logo=gitpod"
    alt="Contribute with Gitpod"
  />
</a>

# stcp

> A simple TCP chat server with auto "I have spoken" reply

## Install
### Prerequisite
> **Note** Required:
- [Golang V1.20](https://go.dev/doc/install)
- [Podman](https://podman-desktop.io/downloads)

```sh
git clone git@github.com:rmsrob/stcp.git
cd stcp
go mod tidy
```

## Usage
```sh
go run main.go -listenAdrs=:3000
```
