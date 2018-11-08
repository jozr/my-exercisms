package twofer

import "fmt"

// ShareWith who?
func ShareWith(name string) string {
	address := name
	if name == "" {
		address = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", address)
}
