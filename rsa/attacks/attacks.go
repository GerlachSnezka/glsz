package attacks

import "github.com/charmbracelet/log"

func NewIfa(verbose bool) *Ifa {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	return &Ifa{}
}