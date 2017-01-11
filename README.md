# hosty [![Circle CI](https://circleci.com/gh/thiagocaiubi/hosty.svg?style=svg)](https://circleci.com/gh/thiagocaiubi/hosty)

## Build

You can build it without installing Go using a Docker image:

```shell
    docker run --rm -v "$PWD":/usr/src/hosty -w /usr/src/hosty golang:1.3 go build -v
```

## Usage

### List all entries managed by hosty

    $ hosty

### Echo all /etc/hosts content

    $ hosty cat

or

    $ hosty c

### Save an entry, use this to create or edit an entry

All new entries are enabled by default

    $ sudo hosty save static 127.0.0.1 static.example.com

or

    $ sudo hosty s static 127.0.0.1 static.example.com

### Enable an entry

    $ sudo hosty enable static

or

    $ sudo hosty e static

### Disable an entry

    $ sudo hosty disable static

or

    $ sudo hosty d static

### Remove an entry

    $ sudo hosty remove static

or

    $ sudo hosty r static
