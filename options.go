package flydb

import (
	"math"
	"time"

	"github.com/oarkflow/flydb/fs"
)

// Options holds the optional DB parameters.
type Options struct {
	// BackgroundSyncInterval sets the amount of time between background Sync() calls.
	//
	// Setting the value to 0 disables the automatic background synchronization.
	// Setting the value to -1 makes the DB call Sync() after every write operation.
	BackgroundSyncInterval time.Duration

	// BackgroundCompactionInterval sets the amount of time between background Compact() calls.
	//
	// Setting the value to 0 disables the automatic background compaction.
	BackgroundCompactionInterval time.Duration

	// FileSystem sets the file system implementation.
	//
	// Default: fs.OSMMap.
	FileSystem fs.FileSystem

	MaxSegmentSize             uint32
	CompactionMinSegmentSize   uint32
	CompactionMinFragmentation float32
}

func (src *Options) copyWithDefaults(path string) *Options {
	opts := Options{}
	if src != nil {
		opts = *src
	}
	if opts.FileSystem == nil {
		opts.FileSystem = fs.OSMMap
	}
	opts.FileSystem = fs.Sub(opts.FileSystem, path)
	if opts.MaxSegmentSize == 0 {
		opts.MaxSegmentSize = math.MaxUint32
	}
	if opts.CompactionMinSegmentSize == 0 {
		opts.CompactionMinSegmentSize = 32 << 20
	}
	if opts.CompactionMinFragmentation == 0 {
		opts.CompactionMinFragmentation = 0.5
	}
	return &opts
}
