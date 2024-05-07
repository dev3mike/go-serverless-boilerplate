package helpers

import (
	"fmt"
	"strings"
)

type Input struct{
	Name string
	Value string
}

func ValidateInput(inputs ...*Input) error{
	for _, field := range inputs {
        if strings.TrimSpace(field.Value) == "" {
            return fmt.Errorf("%s could not be empty", field.Name)
        }
    }
    return nil
}