# NameGenerator
 Go integration of namegenerator2.com API.

This integration is not official and provides only a few of the actual endpoints.
It also provides basic caching support.

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

```sh
go get github.com/yyewolf/NameGenerator
```

### Usage

Import the package into your project.

```go
import "github.com/yyewolf/NameGenerator"
```

Construct a new Session that will contain a cache and the functions

```go
session := namegenerator.NewSession()
```

See Documentation below for more detailed information.


## Documentation

- [![GoDoc](https://godoc.org/github.com/yyewolf/namegenerator?status.svg)](https://godoc.org/github.com/yyewolf/namegenerator) 
- [![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/yyewolf/namegenerator)