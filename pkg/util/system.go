package util

import "runtime"

// IsWindows check runtime is windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsMacos check runtime is macos
func IsMacos() bool {
	return runtime.GOOS == "darwin"
}

// IsLinux check runtime is linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}
