package taskstore

import (
	"fmt"
	"time"
)

type Task struct {
	ID   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

func (t *Task) Validate() error {
	if t.Text == "" {
		return fmt.Errorf("el campo Texto no puede estar vacío")
	}
	if t.Due.Before(time.Now()) {
		return fmt.Errorf("la fecha de vencimiento no puede ser en el pasado")
	}
	if t.Due.IsZero() {
		return fmt.Errorf("la fecha de vencimiento no puede estar vacía")
	}
	return nil
}
