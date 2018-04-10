package notify

type Prefix struct {
	Prefix   string
	Notifier Notifier
}

func WithPrefix(prefix string, notifier Notifier) *Prefix {
	return &Prefix{prefix, notifier}
}

func (p *Prefix) Notify(text string) error {
	return p.Notifier.Notify(p.Prefix + text)
}
