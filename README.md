<div align="center">
<img height="250" src="res/logo.svg" alt="Informer Logo" />

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/ainsleyclark/stock-informer)](https://goreportcard.com/report/github.com/ainsleyclark/stock-informer)
[![Maintainability](https://api.codeclimate.com/v1/badges/1662c0c688e78fa33a2c/maintainability)](https://codeclimate.com/github/ainsleyclark/stock-informer/maintainability)
[![Test](https://github.com/ainsleyclark/stock-informer/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/ainsleyclark/stock-informer/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/ainsleyclark/stock-informer/branch/master/graph/badge.svg?token=70fTUSyXxJ)](https://codecov.io/gh/ainsleyclark/stock-informer)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/ainsleyclark/stock-informer)

</div>

# 📈 Stock Informer

A small and simple DOM detection changer for when you're in desperate need of a new Nvidia graphics card or anything
else that tickles your pickle.

## Overview

- ✅ Monitor multiple URLs to detect DOM changes.
- ✅ Use a valid crontab selector to run monitoring jobs.
- ✅ Integrates with SMTP email and Slack.
- ✅ Easy use with Docker or running on bare metal.
- ✅ Comprehensive HTTP client that follows redirects.
- ✅ Extremely lightweight with few dependencies.

## Why?


## Installation

Informer can either be run in Docker or using the prebuilt binaries in the releases section.

### Binary

### Docker

```bash
docker pull ghcr.io/ainsleyclark/stock-informer:latest
```




## Configuration

```yaml
pages:
  - url: https://test.com # URL to monitor
    selector: .class-selector # DOM selector
    schedule: "* * * * *" # Run every minute
notify:
  email:
    address: smtp.gmail.com
    user: hello@hello.com
    password: password
  slack:
    token: token
    channel_id: id
```

### Pages

### Notifiers

**Email**:

**Slack**:

## Docker

## Roadmap

- Add `App Debug` to the configuration to hide or show the log debug messages.
- Add more notifiers, [github.com/nikoksr/notify](https://github.com/nikoksr/notify) has been used as a package and
	there are an abundance of notification methods that can be used.
- Call cron monitoring recursively to eradicate waiting for new change.
- Validation on configuration struct.

## Contributing

Please feel free to make a pull request if you think something should be added to this package!

## Credits

Shout out to the incredible [Maria Letta](https://github.com/MariaLetta) for her excellent Gopher illustrations.

## Licence

Code Copyright 2023 Stock Informer. Code released under the [MIT Licence](LICENSE).
