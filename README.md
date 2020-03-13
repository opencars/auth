# Auth

[![CircleCI](https://circleci.com/gh/opencars/auth.svg?style=svg)](https://circleci.com/gh/opencars/auth)
[![Go Report Card](https://goreportcard.com/badge/github.com/opencars/auth)](https://goreportcard.com/report/github.com/opencars/auth)

## Overview

:shield: Authorization provider for OpenCars API.

## Event API

On each authorization request new message published to the message broker.

### Success

```JSON
{
  "kind": "authorization",
  "data": {
    "enabled": true,
    "id": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    "ip": "172.18.0.1",
    "name": "xxx-xxx",
    "status": "succeed",
    "timestamp": "2020-03-14T00:43:20"
  }
}
```

### Failure

```JSON
{
  "kind": "authorization",
  "data": {
    "enabled": false,
    "error": "auth.token.revoked",
    "id": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    "ip": "172.18.0.1",
    "name": "xxx-xxx",
    "status": "succeed",
    "timestamp": "2020-03-14T00:43:20"
  }
}
```

## License

Project released under the terms of the MIT [license](./LICENSE).