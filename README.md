# Healy
Healy is a small lightweigt command line programm written with ❤️ in [go](https://golang.org).
It's made for health-checking webservers with an GET-request to specified enpoint, which is specified in the `endpoint.yml`.

[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](https://github.com/jamowei/healy/blob/master/LICENSE)
[![Build Status](https://travis-ci.com/jamowei/healy.svg?branch=master)](https://travis-ci.com/jamowei/healy)
# Installation

You can get the latest binary using Go:

`> go get -u github.com/jamowei/healy`

or download released binary from [here](https://github.com/jamowei/healy/releases/latest).

# Commandline

```
usage: healy [-h|--help] [-c|--config <file>]

             Healy is an easy-to-use and fast health check programm

Arguments:

  -h  --help    Print help information
  -c  --config  host to connect with (server mode). Default: endpoints.yml
```

# Configuration

See the example configuration in `endpoints.yml`
```yml
endpoints:
  google:   https://www.google.com
  notfound: http://isnothinghere.com
  facebook: https://facebook.com
```

# License

Healy is released under the MIT license. See [LICENSE](https://github.com/jamowei/healy/blob/master/LICENSE)