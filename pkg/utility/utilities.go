package utility

import (
	"math/rand"
)

var (
	colors = []string{"red", "green", "yellow", "blue", "magenta", "cyan"}
	users  = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Heidi", "Ivan", "Judy"}
)

func GenerateName() string {
	return users[rand.Intn(len(users))]
}

func GenerateColor() string {
	return colors[rand.Intn(len(colors))]
}
