package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("How does this work?\n" +
		" - Follow the guide at https://github.com/micahlagrange/discord-game-activity\n\n" +
		"Enter 'done' when you are done playing")

	for {
		fmt.Print("$ ")
		cmd, _ := reader.ReadString('\n')
		for _, char := range cmd {
			fmt.Println(char)
		}
		if strings.Split(cmd, "\r\n")[0] == "done" {
			os.Exit(0)
		} else {
			fmt.Println("'done', without quotes, is really the only command at the moment")
		}

	}
}
