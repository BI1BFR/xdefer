// Package xdefer is a RAII utility for golang.
//
// When using keyword defer to do clean-ups, saved statements
// always perform after the surrounding function returns.
// Sometimes we need more precise control over the execute
// timing of the actions.
//
// 1. execute actions after function returns, same as the
// original defer.
//
//    func DoSomething() {
//            var xd xdefer.Defer
//
//            // actions will be executed after function returns
//            defer xd.Exec()
//
//            // ... call xd.Defer() to add actions ...
//    }
//
// 2. execute actions before function returns.
//
//    func DoSomething() {
//            var xd defer.Defer
//            // ... call xd.Defer() to add actions ...
//
//            // execute actions right here
//            xd.Exec()
//
//            // ... do some other stuff
//    }
//
// 3. execute actions outside the surrounding function.
//
//    func AqurieResource() (xd xdefer.Defer) {
//            // ... call xd.Defer() to add actions ...
//            // actions should be executed by caller
//            return
//    }

package xdefer

// A Defer carries a list of actions need to be executed later.
// The zero value of Defer is valid and holds no action.
type Defer struct {
	f []func()
}

// Defer pushes a function to action list.
func (x *Defer) Defer(f func()) {
	x.f = append(x.f, f)
}

// Exec executes all actions in Last-In-First-Out order.
func (x *Defer) Exec() {
	for i := len(x.f) - 1; i >= 0; i-- {
		x.f[i]()
	}
}
