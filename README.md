bmsgpump
======

bmsgpump provides a binary message-pump facility with golang.

The message-pump will continuously receive, process and send messages after startup.

The transport layer is customizable, there is an implementation over net.Conn.

msgpeer
-------------

The message-peer implements a synchronous request response model over message-pump, it considers the client and server peers, allowing each to send requests to the other concurrently.

Documentation
-------------

- [bmsgpump](https://godoc.org/github.com/someonegg/bmsgpump)
- [bmsgpump/msgpeer](https://godoc.org/github.com/someonegg/bmsgpump/msgpeer)

Installation
------------

Install bmsgpump using the "go get" command:

    go get github.com/someonegg/bmsgpump

The Go distribution is bmsgpump's only dependency.
