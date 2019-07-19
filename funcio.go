package gotermuxwrapper

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func TermuxDialog(title string) {
	TermuxDialog := fmt.Sprintf("termux-dialog -t %s", title)
	ExecAndListen(TermuxDialog)
}

func TermuxDialogConfirm(td TDialogConfirm) {
	TermuxDialog := fmt.Sprintf("termux-dialog confirm -i %s -t %s", td.Hint, td.Title)
	ExecAndListen(TermuxDialog)
}

func TermuxDialogCheckbox(td TDialogCheckbox) {
	TermuxDialog := fmt.Sprintf("termux-dialog checkbox -v %s -t %s", strings.Join(td.Values, ","), td.Title)
	ExecAndListen(TermuxDialog)
}

func TermuxDialogCounter(td TDialogCounter) {
	TermuxDialog := fmt.Sprintf("termux-dialog counter -r %d,%d,%d -t %s", td.Min, td.Max, td.Start, td.Title)
	ExecAndListen(TermuxDialog)
}

func TermuxDialogDate(td TDialogDate) {
	TermuxDialog := fmt.Sprintf("termux-dialog date -d \"%d-%d-%d %d:%d:%d\" -t %s", td.Day, td.Month, td.Year, td.KHours, td.Minuts, td.Seconds, td.Title)
	ExecAndListen(TermuxDialog)
}

func TermuxDialogeWithoutDate(td TDialog) {
	TermuxDialog := fmt.Sprintf("termux-dialog date -t %s", td.Title)
	ExecAndListen(TermuxDialog)
}

func TermuxDialogRadio(td TDialogRadio) {
	TermuxDialog := fmt.Sprintf("termux-dialog radio -v %s -t %s", strings.Join(td.Values, ","), td.Title)
	ExecAndListen(TermuxDialog)
}

func TermuxDialogSheet(td TDialogRadio) {
	TermuxDialog := fmt.Sprintf("termux-dialog sheet -v %s -t %s", strings.Join(td.Values, ","), td.Title)
	ExecAndListen(TermuxDialog)
}

func TermuxDialogSpinner(td TDialogRadio) {
	TermuxDialog := fmt.Sprintf("termux-dialog spinner -v %s -t %s", strings.Join(td.Values, ","), td.Title)
	ExecAndListen(TermuxDialog)
}

func TermuxDialogSpeech(td TDialogSpeech) {
	TermuxDialog := fmt.Sprintf("termux-dialog speech -i %s -t %s", td.Hint, td.Title)
	ExecAndListen(TermuxDialog)
}


func ExecAndListen(command string) string {
	cmd := exec.Command(command)
	stdout, err := cmd.StdoutPipe()
	buf := new(bytes.Buffer)

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
	} else {
		buf.ReadFrom(stdout)
	}

	str := buf.String()

	return str
}
