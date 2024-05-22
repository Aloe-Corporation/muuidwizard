[![Go Report Card](https://goreportcard.com/badge/github.com/Aloe-Corporation/muuidwizard)](https://goreportcard.com/report/github.com/Aloe-Corporation/muuidwizard)
# Mongo UUID wizard

Simple CLI that encode and decode UUID from Mongo base64 representation (subtype 00).

## Install

```shell 
go install github.com/Aloe-Corporation/muuidwizard
```

## Usage

### Decode base64 to uuid

```shell 
muuidwizard decode H2hI7XX2QdSyZaPI4D7JEg==
```

### Encode uuid to base64

```shell 
muuidwizard encode 1f6848ed-75f6-41d4-b265-a3c8e03ec912
```