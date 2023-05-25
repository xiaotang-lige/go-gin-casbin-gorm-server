package tool

import "os"

func ProjectPath() (path string) {
	path, _ = os.Getwd()
	return path
}
