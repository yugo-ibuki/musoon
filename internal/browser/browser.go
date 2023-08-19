package browser

import (
	"fmt"
	"os/exec"
)

type Browser struct{}

func NewBrowser() *Browser {
	return &Browser{}
}

func (b *Browser) Open(id string) error {
	fmt.Print("Opening the browser...\n")
	// this is just for mac for now.
	if err := exec.Command("open", "https://youtube.com/watch?v="+id).Start(); err != nil {
		return err
	}

	fmt.Print("The browser has been opened.\n")

	return nil
}
