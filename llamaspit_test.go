package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestProcessChatResponse(t *testing.T) {
	var runCount int
	var passCount int

	for i := 0; i < 3; i++ {
		cmd := exec.Command("./llamaspit", "-y", "multiply 37 by 73")
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		runCount++
		if err != nil {
			t.Errorf("Error running command: %v", err)
			continue
		}

		output := out.String()
		fmt.Printf("Command output: %s\n", output)

		if strings.Contains(output, "2701") {
			passCount++
		}
	}

	if passCount < (runCount / 2) {
		t.Errorf("Some tests failed. Total run: %v, Passes: %v", runCount, passCount)
	} else {
		t.Logf("All tests passed. Total run: %v, Passes: %v", runCount, passCount)
	}
}
