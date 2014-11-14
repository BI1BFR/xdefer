xdefer
======

Package xdefer is a RAII utility for golang.

When using keyword defer to do clean-ups, saved statements
always perform after the surrounding function returns.
Sometimes we need more precise control over the execute
timing of the actions.

1.execute actions after function returns, same as the
original defer.
```Go
func DoSomething() {
        var xd xdefer.Defer

        // actions will be executed after function returns
        defer xd.Exec()

        // ... call xd.Defer() to add actions ...
}
```

2.execute actions before function returns.
```Go
func DoSomething() {
        var xd defer.Defer
        // ... call xd.Defer() to add actions ...

        // execute actions right here
        xd.Exec()

        // ... do some other stuff
}
```
3.execute actions outside the surrounding function.
```Go
func AcquireResource() (xd xdefer.Defer) {
        // ... call xd.Defer() to add actions ...
        // actions should be executed by caller
        return
}
```
