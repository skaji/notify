package notify

type Mux struct {
	Notifiers []Notifier
}

func Assemble(notifier ...Notifier) *Mux {
	return &Mux{notifier}
}

func (m *Mux) Notify(text string) error {
	for _, n := range m.Notifiers {
		if err := n.Notify(text); err != nil {
			return err
		}
	}
	return nil
}
