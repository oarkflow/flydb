package flydb

import (
	"os"

	"github.com/oarkflow/flydb/fs"
)

const (
	lockName = "lock"
)

func createLockFile(opts *Options) (fs.LockFile, bool, error) {
	return opts.FileSystem.CreateLockFile(lockName, os.FileMode(0644))
}
