package taskstore

import (
	"sync"
	"time"
)

type TaskStore struct {
	sync.Mutex
	tasks  map[int]Task
	nextID int
}

func New() *TaskStore {
	return &TaskStore{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		ID:   ts.nextID,
		Text: text,
		Tags: tags,
		Due:  due,
	}

	ts.tasks[task.ID] = task
	ts.nextID++

	return task.ID
}
