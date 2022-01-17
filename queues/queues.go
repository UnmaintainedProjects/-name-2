package queues

import "sync"

var Queues = &sync.Map{}

func Push(chatId int64, item interface{}) int {
	queue := []interface{}{}
	i, ok := Queues.Load(chatId)
	if ok {
		queue = i.([]interface{})
	}
	queue = append(queue, item)
	Queues.Store(chatId, queue)
	return len(queue)
}

func Skip(chatId int64) {
	queue := []interface{}{}
	i, ok := Queues.Load(chatId)
	if ok {
		queue = i.([]interface{})
	}
	if len(queue) == 0 {
		return
	}
	queue = queue[1:]
	Queues.Store(chatId, queue)
}

func Pull(chatId int64) interface{} {
	queue := []interface{}{}
	i, ok := Queues.Load(chatId)
	if ok {
		queue = i.([]interface{})
	}
	if len(queue) == 0 {
		return nil
	}
	queue = queue[1:]
	if len(queue) == 0 {
		return nil
	}
	item := queue[0]
	Queues.Store(chatId, queue)
	return item
}

func Clear(chatId int64) {
	Queues.Delete(chatId)
}
