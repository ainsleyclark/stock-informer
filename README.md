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

In the ever-changing world of online shopping, sometimes it's merely impossible to get your favorite product. This tiny
package allows you to monitor changes on the DOM to detect when a element has changed. It's not just limited to
products, but anything you like.

```yaml
pages:
  - url: https://test.com # URL to monitor
    selector: .class-selector # DOM selector
    schedule: "* * * * *" # Run every minute
notify:
  email:
    address: smtp.gmail.com
    port: 587
    user: hello@hello.com
    password: password
    receivers:
      - me@myemai.com
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

- `Darwin amd64`
- `Darin arm64`
- `Linux amd64`
- `Linux arm64`
- `Windows amd64`

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
$ ./informer -path=/path/to/config/config.yml
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

The configuration for the informer is super simple, you can see it below. The yaml file can be named whatever you want,
but it must follow some conventions.

```yaml
pages:
  - url: https://test.com # URL to monitor
    selector: .class-selector # DOM selector
    schedule: "* * * * *" # Run every minute
notify:
  email:
    address: smtp.gmail.com
    port: 587
    user: hello@hello.com
    password: password
    receivers:
      - me@myemai.com
  slack:
    token: token
    channel_id: id
```

### Pages

Pages is a collection of URLs to monitor. The URL is the page you want to monitor, the selector should be a valid CSS
selector and the schedule is a crontab expression defining when the scrape should happen.

### Notifiers

Currently, SMTP email and Slack notifiers are supported, but there are more to come. The settings for each notifier are
self explanatory but all required.

## Roadmap

- Add `App Debug` to the configuration to hide or show the log debug messages.
- Add more notifiers, [github.com/nikoksr/notify](https://github.com/nikoksr/notify) has been used as a package and
	there are an abundance of notification methods that can be used.
- Call cron monitoring recursively to eradicate waiting for new change.
- Validation on configuration struct.

## Development

To set up the application for development first, clone the repository.

```bash
git clone https://github.com/ainsleyclark/stock-informer.git
```

Run the setup command to install the necessary dependencies for Krang.

```bash
make setup
```

### Makefile

Common commands are detailed in the `Makefile` to list usage run:

```bash
make help

setup            Setup dependencies
run              Run
dist             Creates and build dist folder
format           Run gofmt
lint             Run linter
test             Test uses race and coverage
test-v           Test with -v
cover            Run all the tests and opens the coverage report
docker-clean		 Removes the docker image
docker-build 		 Builds the docker image
docker-run			 Run the docker image
mock             Generate mocks keeping directory tree
bench						 Runs benchmarks
doc              Runs go doc
all              Make format, lint and test
todo             Show to-do items per file
help             Display this help
```

## Contributing

We welcome contributors, but please read the [contributing document](CONTRIBUTING.md) before making a pull request.

## Credits

Shout out to the incredible [Maria Letta](https://github.com/MariaLetta) for her excellent Gopher illustrations.

## Licence

Code Copyright 2023 Stock Informer. Code released under the [MIT Licence](LICENSE).


