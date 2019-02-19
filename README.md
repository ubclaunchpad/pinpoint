# 📊 pinpoint [![Build Status](https://travis-ci.com/ubclaunchpad/pinpoint.svg?branch=master)](https://travis-ci.com/ubclaunchpad/pinpoint) [![codecov](https://codecov.io/gh/ubclaunchpad/pinpoint/branch/master/graph/badge.svg)](https://codecov.io/gh/ubclaunchpad/pinpoint) [![Go Report Card](https://goreportcard.com/badge/github.com/ubclaunchpad/pinpoint)](https://goreportcard.com/report/github.com/ubclaunchpad/pinpoint)

Pinpoint is a membership application management service geared towards helping university clubs and hackathons decide on the best applicants out of their pool of candidates.

See the project outline and minimum viable product in our [wiki](https://github.com/ubclaunchpad/pinpoint/wiki/Project-Outline).

## Project Structure

The project is structured as follows:

- `client` is the JavaScript client for the user-facing Pinpoint API.
- `core` is the primary Pinpoint gRPC-based service, and manages application logic and the database.
- `frontend` is the Pinpoint web application.
- `gateway` is an HTTP server that exposes Pinpoint functionality via a RESTful API.
- `protobuf` contains protobuf definitions for Pinpoint Core's gRPC service as well as the generated Golang API.
- `utils` is a Golang package that contains utility functions shared by `core` and `gateway`.

## Development

To get started, make sure you have the following installed:

- [Golang](https://golang.org/dl/) 1.11+
- [Node.js](https://nodejs.org/en/download/) 8.12+
- [protobuf](https://github.com/protocolbuffers/protobuf/releases) v3.6+
- [Docker CE](https://docs.docker.com/install/#supported-platforms) and [docker-compose](https://docs.docker.com/compose/install/)

To fetch the codebase, use `go get`:

```bash
$> go get github.com/ubclaunchpad/pinpoint
```

### Installing Dependencies

You will need [dep](https://github.com/golang/dep#installation) and [npm](https://www.npmjs.com/get-npm) installed.

```bash
$> make deps
```

### Makefile

The [Makefile](/Makefile) offers a lot of useful commands for development. Run
`make help` to see the commands that are available.

### Building

#### Golang Binaries

```sh
$> make pinpoint-core
$> make pinpoint-gateway
```

#### Web Application

```sh
$> make web
```

### Spinning up Services Locally

External dependencies, such as the database, can be started and stopped using
docker-compose, which leverages available Docker containers:

```sh
$> make testenv       # start up service containers
$> make testenv-stop  # stop containers
$> make clean         # remove containers
```

Pinpoint services can be started up using the following commands in two separate shell sessions:

```sh
$> make core
$> make gateway
```

By default, provided certificates in `dev/certs` are used. These were generated using [certstrap](https://github.com/square/certstrap).

To run enable the local monitoring suite:

```sh
$> make monitoring
$> make core FLAGS=--logpath=tmp/core.log
$> make gateway FLAGS=--logpath=tmp/gateway.log
```

### Updating the Golang gRPC API

`gateway` and `core` uses the Golang API within the `protobuf` directory to communicate. If you make changes to the protobuf definitions in the `protobuf` directories, you will need to update this API:

```bash
$> make proto  # generate new Golang API
$> make check  # ensure everything compiles
```

You will need [protobuf](https://github.com/protocolbuffers/protobuf/releases) v3.6+ and the [Golang plugin](https://github.com/golang/protobuf#installation) installed.

The script also uses [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) to generate mocks.
