package todo

import (
	"sync"
)

var (
	list []Todo
	mtx  sync.RWMutex
	once sync.Once
)

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []Todo{}
}

type Todo struct {
	ID       string `json:id`
	Message  string `json:message`
	Complete bool   `json:complete`
}

func Get() []Todo {
	return list
}
