//go:build !windows

package z

import "os"

func syncFd(*os.File) error { return nil }
