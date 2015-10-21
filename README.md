# go-sapin
:christmas_tree: drawing a christmas tree in Go

[![GoDoc](https://godoc.org/github.com/moul/go-sapin?status.svg)](https://godoc.org/github.com/moul/go-sapin)
[![](https://img.shields.io/badge/appspot-sapin-blue.svg)](http://sapin-as-a-service.appspot.com/?size=5)

## Demo

http://sapin-as-a-service.appspot.com/?size=5

## Usage

```console
$ sapin --size=1
   *
  ***
 *****
*******
   |
```

```console
$ sapine --size=2
      *
     ***
    *****
   *******
    *****
   *******
  *********
 ***********
*************
     |||
     |||
```

```console
$ sapin -h
Usage:
  sapin [OPTIONS]

Application Options:
  -s, --size= Size of the sapin

Help Options:
  -h, --help  Show this help mess
```

## Install

```console
$ go get github.com/moul/go-sapin/cmd/sapin
```

## License

MIT
