package gotermuxwrapper

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func TermuxDialog(title string)  {
	TermuxDialog := fmt.Sprintf("Termux-dialog %s", title)
	ExecAndListen(TermuxDialog)
}

func ExecAndListen(command string) string  {
	cmd := exec.Command(command)
	stdout, err := cmd.StdoutPipe()
	buf := new(bytes.Buffer)
	buf.ReadFrom(stdout)

	if err != nil {
		log.Fatalln(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}

	//Wait waits for the command to exit
	//It must have been started by Start
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	return buf.String()
}