package llamaspit

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestProcessChatResponse(t *testing.T) {
	t.Run("Test command execution on '-y' argument", func(t *testing.T) {
		cmd := exec.Command("./llamaspit", "-y", "multiply 37 by 73")
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(out.String(), "2701") {
			t.Fatalf("Expected '2701' in command output, got %s", out.String())
		}
	})
}
