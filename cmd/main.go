package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mellonnen/groot/models"
)

func main() {
	io := bufio.NewScanner(os.Stdin)

	// TODO: move this logic to an handle input function
	nodes := make(map[string]*models.Tree)

	for io.Scan() {
		line := io.Text()
		split := strings.Split(line, " ")

		parent, ok := nodes[split[0]]
		if !ok {
			parent = &models.Tree{Name: split[0]}
			nodes[split[0]] = parent
		}

		child, ok := nodes[split[1]]
		if !ok {
			child = &models.Tree{Name: split[1]}
			nodes[split[1]] = child
		}
		parent.Children = append(parent.Children, child)
	}
	fmt.Println(nodes["github.com/ZinoKader/KEX"].String())
}
