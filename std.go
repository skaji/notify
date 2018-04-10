package notify

import (
	"fmt"
	"os"
)

type Stdout struct{}

func (s *Stdout) Notify(text string) error {
	_, err := fmt.Println(text)
	return err
}

type Stderr struct{}

func (s *Stderr) Notify(text string) error {
	_, err := fmt.Fprintln(os.Stderr, text)
	return err
}
