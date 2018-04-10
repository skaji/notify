package main

import (
	"fmt"

	"github.com/skaji/notify"
)

func main() {
	f := func(t string) string {
		return fmt.Sprintf("%d %s", len(t), t)
	}
	n := notify.Assemble(
		notify.WithPrefix("foo ", &notify.Stdout{}),
		notify.WithFunc(f, &notify.Stderr{}),
	)
	n.Notify("oops")
}
