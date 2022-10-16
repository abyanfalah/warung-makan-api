package utils

import "github.com/google/uuid"

func GenerateId() string {
	return uuid.New().String()
}

func GenerateClockSequence() int {
	return uuid.New().ClockSequence()
	// return int(uuid.New().ID())
}
