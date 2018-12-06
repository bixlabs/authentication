package structures

import "time"

type Todo struct {
	Name string `json:"name"`
	Description string `json:"description"`
	DueDate time.Time `json:"due_date"`
}
