package main

import (
	"MTS/auth/httpserver"
	"bufio"
	"log"
	"os"
)

func main() {
	// channel for commands
	cmds := make(chan string)

	// creating server and starting it
	server, err := httpserver.New()
	if err != nil {
		log.Fatal(err)
	}

	// start goroutine for processing commands from command line
	// and switch server's state
	go func(s *httpserver.Server) {
		for {
			cmd := <-cmds
			switch cmd {
			case "debug":
				// disable/enable debug endpoint
				server.Switch()
				log.Println("Switched server's state")
			default:
				continue
			}
		}
	}(server)

	// start goroutine for scanning commands from command line
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			command := scanner.Text()
			cmds <- command
		}

		close(cmds)
	}()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
