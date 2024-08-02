package taskstore

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type TaskStore struct {
	sync.Mutex
	tasks  map[int]Task
	nextID int
}

var (
	ErrTaskNotFound  = errors.New("tarea no encontrada")
	ErrInvalidTaskID = errors.New("ID de tarea inválido")
	ErrIDExists      = errors.New("ID de tarea ya existe")
)

func New() *TaskStore {
	return &TaskStore{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

// TaskStore CRUD operations

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) (int, error) {
	task := &Task{
		Text: text,
		Tags: tags,
		Due:  due,
	}
	if err := task.Validate(); err != nil {
		return 0, err
	}

	ts.Lock()
	defer ts.Unlock()

	task.ID = ts.nextID
	ts.tasks[task.ID] = *task
	ts.nextID++

	return task.ID, nil
}

func (ts *TaskStore) GetTask(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	task, ok := ts.tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("%w: ID %d", ErrTaskNotFound, id)
	}
	return task, nil
}

func (ts *TaskStore) GetAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	tasks := make([]Task, 0, len(ts.tasks))
	for _, task := range ts.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (ts *TaskStore) UpdateTask(id int, text string, tags []string, due time.Time) error {
	ts.Lock()
	defer ts.Unlock()

	task, ok := ts.tasks[id]
	if !ok {
		return fmt.Errorf("%w: ID %d", ErrTaskNotFound, id)
	}
	task.Text = text
	task.Tags = tags
	task.Due = due
	if err := task.Validate(); err != nil {
		return err
	}
	ts.tasks[id] = task
	return nil
}

func (ts *TaskStore) DeleteTask(id int) error {
	ts.Lock()
	defer ts.Unlock()

	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf("%w: ID %d", ErrTaskNotFound, id)
	}
	delete(ts.tasks, id)
	return nil
}
