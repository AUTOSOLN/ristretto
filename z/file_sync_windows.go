//go:build windows

package z

import "os"

// syncFd flushes the file's buffers to disk. FlushViewOfFile alone doesn't
// guarantee the pages and metadata reach disk on Windows.
func syncFd(f *os.File) error {
	if f == nil {
		return nil
	}
	return f.Sync()
}
