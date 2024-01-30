# LlamaSpit

LllamaSpit is a Go program that interacts with an Ollama HTTP endpoint to generate executable bash commands based on a provided narrative description of a certain task. The generated command is presented to the user who can choose either to execute it or decline it.

## Building and Running with Go

To build and run the program with Go, execute the following commands on your terminal:

```
git clone https://github.com/bplunkert/llamaspit.git
cd llamaspit
go build
```

The program is run as follows:
```
./llamaspit "<Your_command_description_here>"
```

Please replace `"<Your_command_description_here>"` with your description of the command to be executed.

Example usage:
```
./llamaspit multiply 37 by 73
llamaspit suggests: `echo $((37 * 73))`
Do you want to execute the command? (y/n)
y
+ echo 2701
2701
```

Full usage:
```
./llamaspit -h
Usage: llamaspit [options] <description of the desired command>
Options:
-h/--help       Display this help text
-y              Automatically accept and execute the command
Environment Variables:
OLLAMA_HOST     The URL of the Ollama endpoint, defaults to http://localhost:11434
```


## Building and Running with Docker Compose

To build and run the program with Docker Compose, execute the following command on your terminal:
```
docker compose build
```

To run the program with Docker Compose, execute the following command on your terminal:
```
docker compose run llamaspit --help
```
## Testing

To run the tests, execute the following command on your terminal:
```
go test
```

To run the tests with Docker Compose, execute the following command on your terminal:
```
docker compose run llamaspit go test

```