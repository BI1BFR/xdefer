package xdefer

type Defer struct {
	f []func()
}

func (x *Defer) Defer(f func()) {
	x.f = append(x.f, f)
}

func (x *Defer) Exec() {
	for i := len(x.f) - 1; i >= 0; i-- {
		x.f[i]()
	}
}
