package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

// Application describes the interface of the application.
type Application interface {
	Generate(version int) (string, error)
}

type application struct{}

// NewApp returns an implementaiton of an application.
func NewApp() Application {
	return &application{}
}

func (a *application) Generate(version int) (string, error) {
	switch version {
	case 1:
		return uuid.NewV1().String(), nil
	case 4:
		return uuid.NewV4().String(), nil
	case 2:
		return "", fmt.Errorf("unsupported uuid version: %d", version)
	case 3:
		return "", fmt.Errorf("unsupported uuid version: %d", version)
	case 5:
		return "", fmt.Errorf("unsupported uuid version: %d", version)
	}
	return "", fmt.Errorf("unknown uuid version: %d", version)
}
