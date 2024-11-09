package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct {
	value string
}

var lock = &sync.Mutex{}
var singleInstance *Singleton

func GetInstance(value string) *Singleton {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("[INFO] Creating single instance now")
			singleInstance = &Singleton{value: value}
		} else {
			fmt.Println("[INFO] Single instance already created in another gorutine")
		}
	} else {
		fmt.Println("[INFO] Single instance already created")
	}
	return singleInstance
}
