# gomoa
ZeroMQ client/server for testing

## Installation

### Requirements

* goczmq: https://github.com/zeromq/goczmq
  * libsodium
  * libzmq
  * czmq

```
$ go get github.com/dialogbox/gomoa
```

## Usage

```
% gomoa --help                                                                                                                                                  gomoa [master]
ZeroMQ Tester

Usage:
  gomoa [command]

Available Commands:
  client      ZeroMQ Client Tester
  help        Help about any command
  server      ZeroMQ Server Tester

Flags:
      --config string   config file (default is $HOME/.gomoa.yaml)
  -h, --help            help for gomoa
  -t, --toggle          Help message for toggle

Use "gomoa [command] --help" for more information about a command.
```
