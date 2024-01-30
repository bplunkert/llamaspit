package llamaspit

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jmorganca/ollama/api"
)

type ChatCommand struct {
	Command string `json:"command"`
}

func main() {
	var autoAccept bool
	var help bool

	flag.BoolVar(&autoAccept, "y", false, "Automatically accept and execute the command")
	flag.BoolVar(&help, "h", false, "Display help")
	flag.BoolVar(&help, "help", false, "Display help")

	flag.Parse()

	if help {
		fmt.Println("Usage: llamaspit [options] <description of the desired command>")
		fmt.Println("Options:")
		fmt.Println("-h/--help\tDisplay this help text")
		fmt.Println("-y\t\tAutomatically accept and execute the command")
		fmt.Println("Environment Variables:")
		fmt.Println("OLLAMA_HOST\tThe URL of the Ollama endpoint, defaults to http://localhost:11434")
		os.Exit(0)
	}

	promptInput := flag.Args()

	if len(promptInput) < 1 {
		panic("Missing required input.")
	}

	promptContent := "Write a one-line, self-contained and runnable bash command, which accomplishes the following description. Include only the code, with no formatting nor extra text. If the code is not bash, make sure it is encapsulated in a command that is runnable in bash. Format your response in JSON as ```{\"command\": COMMAND}```. Write a bash command to:\n" + strings.Join(promptInput, " ")

	client, err := api.ClientFromEnvironment()
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return
	}

	ctx := context.Background()

	stream := false
	req := &api.ChatRequest{
		Model: "llama2",
		Messages: []api.Message{{
			Role:    "user",
			Content: promptContent,
		}},
		Stream: &stream,
		Format: "json",
	}

	processChatResponse := func(resp api.ChatResponse) error {
		var chatCommand ChatCommand

		err := json.Unmarshal([]byte(resp.Message.Content), &chatCommand)
		if err != nil {
			return err
		}

		if chatCommand.Command == "" {
			return errors.New("Invalid JSON: no 'command' key")
		}

		fmt.Println("llamaspit suggests: `" + chatCommand.Command + "`")

		if !autoAccept {
			fmt.Println("Do you want to execute the command? (y/n)")
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				return err
			}

			if strings.TrimRight(response, "\n") != "y" {
				return nil
			}
		}

		cmd := exec.Command("bash", "-x", "-e", "-c", chatCommand.Command)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()

		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				exitCode := exitError.ExitCode()
				fmt.Fprintf(os.Stderr, "Command exited with code %d \n", exitCode)
				os.Exit(exitCode)
			}
			return err
		}

		return nil
	}

	err = client.Chat(ctx, req, processChatResponse)

	if err != nil {
		fmt.Println("Failed to chat:", err)
	}
}
