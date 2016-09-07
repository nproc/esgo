package locker

import (
	"sync"

	"github.com/nproc/esgo/drivers"
)

var (
	collectionMutexes map[string]*sync.RWMutex
)

func ForCollection(collection drivers.Collection) *sync.RWMutex {
	if lock := collectionMutexes[collection.Name()]; lock != nil {
		return lock
	}
	mutex = &sync.RWMutex{}
	collectionMutexes[collection.Name()] = mutex
	return mutex
}
