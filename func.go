package notify

type Func struct {
	f        func(string) string
	notifier Notifier
}

func WithFunc(f func(string) string, notifier Notifier) *Func {
	return &Func{f, notifier}
}

func (f *Func) Notify(text string) error {
	return f.notifier.Notify(f.f(text))
}
