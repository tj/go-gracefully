
# go-gracefully

 Graceful shutdown utility with hard exit on second signal.

 View the [docs](http://godoc.org/github.com/visionmedia/go-gracefully).

## Installation

```
$ go get github.com/visionmedia/go-gracefully
```

## Example

  Typically something like:

```go
<-Shutdown()
w.Stop()
```

# License

 MIT