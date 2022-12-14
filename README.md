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

## Installation

Informer can either be run in Docker or using the prebuilt binaries in the releases section, information on both methods
are shown below.

### Binary

The following platforms that are supported are listed below. The examples used are for Darwin amd64, please change the
release name if you intend to use a different OS.

- Darwin amd64
- Darin arm64
- Linux amd64
- Linux arm64
- Windows amd64

#### Download the Binary

Head over to the [releases](https://github.com/ainsleyclark/stock-informer/releases/) page and download the relevant
release to your operating system.

```bash
$ wget "https://github.com/ainsleyclark/stock-informer/releases/download/v0.0.1/informer_0.0.1_darwin_amd64.tar.gz"
> ‘informer_0.0.1_darwin_amd64.tar.gz’ saved
```

#### Unzip

```bash
$ tar -xf informer_0.0.1_darwin_amd64.tar.gz && cd informer
```

#### Configuration

Change `config.example.yml` to `config.yml` and change to your liking.

### Run

```bash
$ ./informer
> [INFORMER] 2022-12-14 08:14:05 | LOG | [INFO] | [msg] Loading Configuration
> [INFORMER] 2022-12-14 08:14:05 | LOG | [INFO] | [msg] Booting Informer
```

### Docker

Docker images are located at
the [packages](https://github.com/ainsleyclark/stock-informer/pkgs/container/stock-informer) page. Be sure to use the
latest version number when pulling the image.

#### Pull the Docker Image

Head over to the [packages](https://github.com/ainsleyclark/stock-informer/pkgs/container/stock-informer) page and pull
the latest image version to your local machine.

```bash
$ docker pull ghcr.io/ainsleyclark/stock-informer:0.0.1
```

#### Run the Image

Running the image requires two required arguments/flags.

- The path to the configuration file stored on the local machine, with the arg `v`.
- The `-path` argument for the binary which should correlate to the path passed in. This will allow you to attach your
	configuration file from your local machine to the docker image.

```bash
$ docker run -it --rm -v /path/to/config/config.yml:/mnt/config.yml ghcr.io/ainsleyclark/stock-informer:0.0.1 -path=/mnt/config.yml
> [INFORMER] 2022-12-14 08:14:05 | LOG | [INFO] | [msg] Loading Configuration
> [INFORMER] 2022-12-14 08:14:05 | LOG | [INFO] | [msg] Booting Informer
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


