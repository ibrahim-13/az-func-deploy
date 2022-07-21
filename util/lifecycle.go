package util

import (
	"sync"
)

var lock sync.Mutex
var shouldRun bool = true
var killFunc func() = nil
var fileCleanups []string = []string{}

func SetKillFunc(fn func()) {
	lock.Lock()
	defer lock.Unlock()
	killFunc = fn
}

func ExecuteKill() {
	lock.Lock()
	defer lock.Unlock()
	if killFunc != nil {
		killFunc()
	}
	shouldRun = false
}

func SetIfShouldRun(sr bool) {
	lock.Lock()
	defer lock.Unlock()
	shouldRun = sr
}

func GetIfShouldRun() bool {
	lock.Lock()
	defer lock.Unlock()
	return shouldRun
}

func AddCleanupFile(file string) {
	lock.Lock()
	defer lock.Unlock()
	fileCleanups = append(fileCleanups, file)
}

func GetCleanupFiles() []string {
	lock.Lock()
	defer lock.Unlock()
	return fileCleanups
}
