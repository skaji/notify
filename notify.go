package notify

type Notifier interface {
	Notify(text string) error
}
