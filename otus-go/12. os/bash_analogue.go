package main

import (
	"os"
	"os/exec"
)


/**
делаем аналог bash команды
ls | wc -l
 */
func main() {
	cmd1 := exec.Command("ls")
	cmd2 := exec.Command("wc", "-l")

	pipe, _ := cmd1.StdoutPipe()
	cmd2.Stdin = pipe
	cmd2.Stdout = os.Stdout
	_ = cmd1.Start()
	_ = cmd2.Start()
	_ = cmd1.Wait()
	_ = cmd2.Wait()
}
