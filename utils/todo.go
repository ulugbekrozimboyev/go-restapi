package utils

import (
	"fmt"
	"time"
)

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func (t Todo) String() string {
	// return fmt.Sprintf("id: %v, name: %v, complete: %v, due: %v", t.Id, t.Name, t.Completed, t.Due)
	return fmt.Sprintf("id: %v, name: %v, complete: %v", t.Id, t.Name, t.Completed)
}
