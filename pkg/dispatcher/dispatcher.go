package dispatcher

import (
	"container/list"
	"log"
)

type Callback func(action interface{})

var callbacks = &list.List{}

// Dispatch calls all the callbacks with the sent action.
func Dispatch(action interface{}) {
	log.Println("dispatching action")

	for cb := callbacks.Front(); cb != nil; cb = cb.Next() {
		c := cb.Value.(Callback)
		c(action)
	}
}

func Register(callback Callback) {
	callbacks.PushBack(callback)
}
