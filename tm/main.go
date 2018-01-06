// Install with: $ go install bin/tm
package main

import (
	"bin/wsdir"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		cmd := exec.Command("tmux", "list-sessions")
		cmd.Stdout = os.Stdout
		cmd.Run() // When empty an error is returned, which is not useful.
		return
	}
	sty := os.Args[1]

	cmd := exec.Command("tmux", "attach-session", "-t", sty)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err == nil {
		return
	}

	matches, err := wsdir.Get(sty)
	if err != nil {
		log.Fatal(err)
	}
	dir := matches[0]
	cmd = exec.Command("tmux", "new-session", "-d", "-s", sty, "-c", dir, "env",
		fmt.Sprint("STY=%s", sty), os.ExpandEnv("$SHELL"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("tmux", "set-environment", "-t", sty, "STY", sty)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// TODO: source ~/.tmux/$STY
	cmd = exec.Command("tmux", "attach-session", "-t", sty)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
