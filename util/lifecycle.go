package util

import (
	"sync"
)

var lock sync.Mutex
var shouldRun bool = true
var killFunc func() = nil

func SetKillFunc(fn func()) {
	lock.Lock()
	defer lock.Unlock()
	killFunc = fn
}

func ExecuteKill() {
	if killFunc != nil {
		killFunc()
	}
	shouldRun = false
}

func SetIfShouldRun(sr bool) {
	shouldRun = sr
}

func GetIfShouldRun() bool {
	return shouldRun
}
