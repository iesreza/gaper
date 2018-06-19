gaper
=====

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Build Status](https://travis-ci.org/maxcnunes/gaper.svg?branch=master)](https://travis-ci.org/maxcnunes/gaper)
[![Coverage Status](https://codecov.io/gh/maxcnunes/gaper/branch/master/graph/badge.svg)](https://codecov.io/gh/maxcnunes/gaper)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/maxcnunes/gaper)
[![Go Report Card](https://goreportcard.com/badge/github.com/maxcnunes/gaper)](https://goreportcard.com/report/github.com/maxcnunes/gaper)

Restarts programs when they crash or a watched file changes.

**NOT STABLE YET, STILL IN DEVELOPMENT**: Please, check out [this ticket](https://github.com/maxcnunes/gaper/issues/1) to follow its progress.

## Installation

```
go get -u github.com/maxcnunes/gaper
```

## Changelog

See [Releases](https://github.com/maxcnunes/gaper/releases) for detailed history changes.

## Usage

```
NAME:
   gaper - Used to restart programs when they crash or a watched file changes

USAGE:
   gaper [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --bin-name value                   name for the binary built by Gaper for the executed program
   --build-path value                 path to the program source code
   --build-args value                 build arguments passed to the program
   --verbose                          turns on the verbose messages from Gaper
   --watch value, -w value            a comma-delimited list of folders or files to watch for changes
   --ignore value, -i value           a comma-delimited list of folders or files to ignore for changes
   --poll-interval value, -p value    how often in milliseconds to poll watched files for changes (default: 500)
   --extensions value, -e value       a comma-delimited list of file extensions to watch for changes (default: "go")
   ----no-restart-on value, -n value  don't automatically restart the executed program if it ends.
                                        If "error", an exit code of 0 will still restart.
                                        If "exit", no restart regardless of exit code.
                                        If "success", no restart only if exit code is 0.
   --help, -h                         show help
   --version, -v                      print the version
```

## Contributing

See the [Contributing guide](/CONTRIBUTING.md) for steps on how to contribute to this project.

## Reference

This package was heavily inspired by [gin](https://github.com/codegangsta/gin) and [node-supervisor](https://github.com/petruisfan/node-supervisor).

Basically, Gaper is a mixing of those projects above. It started from **gin** code base and I rewrote it aiming to get
something similar to **node-supervisor** (but simpler). A big thanks for those projects and for the people behind it!
:clap::clap:

### How is Gaper different of Gin

The main difference is that Gaper removes a layer of complexity from Gin which has a proxy running on top of 
the executed server. It allows to postpone a build and reload the server when the first call hits it. With Gaper 
we don't care about that feature, it just restarts your server whenever a change is made.