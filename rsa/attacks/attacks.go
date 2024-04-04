package attacks

import (
	"os"

	"github.com/charmbracelet/log"
)

type Attacks struct {
	logger *log.Logger
}

func New(verbose bool) *Attacks {
	logger := log.New(os.Stdout)
	if verbose {
		logger.SetLevel(log.DebugLevel)
	}

	return &Attacks{logger: logger}
}