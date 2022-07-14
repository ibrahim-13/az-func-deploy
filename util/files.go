package util

import "os"

func ExistsFile(filePath string) bool {
	stat, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !stat.IsDir()
}
