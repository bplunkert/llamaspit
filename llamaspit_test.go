package llamaspit

import (
	"bytes"
	"os/exec"
	"strings"
	"sync"
	"testing"
)

func TestProcessChatResponse(t *testing.T) {
	var runCount int32 = 10
	var passCount int32 = 0
	var mu sync.Mutex

	t.Parallel()

	for i := 0; i < 3; i++ {
		t.Run("Test command execution on '-y' argument", func(t *testing.T) {
			t.Parallel()

			cmd := exec.Command("./llamaspit", "-y", "multiply 37 by 73")
			var out bytes.Buffer
			cmd.Stdout = &out

			err := cmd.Run()
			mu.Lock()
			runCount++
			if err != nil {
				t.Fatal(err)
			}

			if strings.Contains(out.String(), "2701") {
				passCount++
			}
			mu.Unlock()
		})
	}

	t.Logf("Total run: %v, Passes: %v, Pass rate: %v percent", runCount, passCount, (passCount/runCount)*100)
}
