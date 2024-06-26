package flydb

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "flydb: ", 0)

// SetLogger sets the global logger.
func SetLogger(l *log.Logger) {
	if l != nil {
		logger = l
	}
}
