bmsgpump
======

bmsgpump provides a message-pump facility with golang.

The message-pump will continuously receive, process and send messages after startup.

The transport layer is customizable, there is an implementation over net.Conn.

Documentation
-------------

- [API Reference](https://godoc.org/github.com/someonegg/bmsgpump)

Installation
------------

Install bmsgpump using the "go get" command:

    go get github.com/someonegg/bmsgpump

The Go distribution is bmsgpump's only dependency.
