package cmd

import (
	"fmt"
	"testing"
)

func Test_generate_against_invalid_flags(t *testing.T) {
	command := NewGenerate()

	result, err := command.Execute()
	if err != nil {
		t.Errorf("fail")
	}

	fmt.Printf("command result: %#v", result.Value.(*Generate).url)

}
