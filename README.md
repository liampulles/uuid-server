<h1 align="center">
  uuid-server
</h1>

<h4 align="center"> Generates Version 4 UUIDs upon an HTTP request</a></h4>

<p align="center">
  <a href="#status">Status</a> •
  <a href="#install">Install</a> •
  <a href="#configuration">Configuration</a> •
  <a href="#usage">Usage</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

<p align="center">
  <a href="https://travis-ci.com/liampulles/uuid-server">
    <img src="https://travis-ci.com/liampulles/uuid-server.svg?branch=master" alt="[Build Status]">
  </a>
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/liampulles/uuid-server">
  <a href="https://goreportcard.com/report/github.com/liampulles/uuid-server">
    <img src="https://goreportcard.com/badge/github.com/liampulles/uuid-server" alt="[Go Report Card]">
  </a>
  <a href="https://codecov.io/gh/liampulles/uuid-server">
    <img src="https://codecov.io/gh/liampulles/uuid-server/branch/master/graph/badge.svg" />
  </a>
</p>

## Status

UUID Server is currently in development, and is not available for general usage

## Install

### Native

Either download a release from the releases page, or clone and run `make install`, and execute:

```bash
uuid-server
```

### Docker

Either pull `lpulles/uuid-server:latest`, or clone and run `make docker-build`, and execute:

```bash
docker run lpulles/uuid-server:latest
```

## Configuration

You can set the following environment variables:

* `PORT`: What port to run the server on. Defaults to `8080`
* `LOGLEVEL`: What level to log at. Valid levels: [`INFO`, `ERROR`]. Defaults to `INFO`.

## Usage

Once the app is running (see [Install](#install)):

```bash
curl http://127.0.0.1:8080/
```

(Or an equivalent way of sending a `GET` on `/`)

Should yield the response (with a different UUID):

```text
bb290e59-8139-41ad-8f4a-b22002725583
```

## Contributing

Please submit an issue with your proposal.

## License

See [LICENSE](LICENSE)
