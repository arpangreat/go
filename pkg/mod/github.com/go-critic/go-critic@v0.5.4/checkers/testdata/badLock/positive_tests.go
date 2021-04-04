package checker_test

import (
	"sync"
)

type withMutex struct {
	mu sync.RWMutex
}

func immediateUnlock(mu *sync.Mutex, op func()) {
	mu.Lock()
	/*! defer is missing, mutex is unlocked immediately */
	mu.Unlock()
	op()
}

func immediateUnlockStruct(x *withMutex, op func()) {
	x.mu.Lock()
	/*! defer is missing, mutex is unlocked immediately */
	x.mu.Unlock()
	op()
}

func mismatchingUnlock1(mu *sync.RWMutex, op func()) {
	mu.Lock()
	/*! suspicious unlock, maybe Unlock was intended? */
	defer mu.RUnlock()
	op()
}

func mismatchingUnlock2(mu *sync.RWMutex, op func()) {
	mu.RLock()
	/*! suspicious unlock, maybe RUnlock was intended? */
	defer mu.Unlock()
	op()
}

func mismatchingUnlock1Struct(x *withMutex, op func()) {
	x.mu.Lock()
	/*! suspicious unlock, maybe Unlock was intended? */
	defer x.mu.RUnlock()
	op()
}

func mismatchingUnlock2Struct(x *withMutex, op func()) {
	x.mu.RLock()
	/*! suspicious unlock, maybe RUnlock was intended? */
	defer x.mu.Unlock()
	op()
}
