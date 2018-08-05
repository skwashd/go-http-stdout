# `go-http-stdout`

A quick and dirty HTTP server that logs JSON data to stdout. The primary use case for this is debugging webhooks. It is
particularly useful when combined with [ngrok](https://ngrok.com/).

The error handling is miminal. It is designed for debugging so it doesn't need to be very fancy.

## Installation

The server only uses the golang stdlib. There are no dependencies to go get.

You build and then run it:

```
# Build
$ go build
# Run
$ ./go-http-stdout
```

To listen on a port other than 8888, use `-listen`

```
$ ./go-http-stdout -listen=:8008
```

## Credits
`go-http-stdout` is created by [Dave Hall Consulting](https://davehall.com.au) and is offered as is for your enjoyment.