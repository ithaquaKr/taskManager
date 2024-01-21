package utils

import "math/rand"

// RandomType generates a random type for a list
func RandomType() string {
	types := []string{"note", "task"}
	n := len(types)
	return types[rand.Intn(n)]
}

// RandomStatus generates a random status for a task
func RandomStatus() string {
	status := []string{"cancel", "doing", "complete"}
	n := len(status)
	return status[rand.Intn(n)]
}

// RandomPriority generates a random priority for a task
func RandomPriority() string {
	priorities := []string{"no_priority", "low", "medium", "high"}
	n := len(priorities)
	return priorities[rand.Intn(n)]
}
