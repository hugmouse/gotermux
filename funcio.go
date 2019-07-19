package gotermuxwrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	TD = "termux-dialog"
)

func TermuxDialog(title string) {
	ExecAndListen(TD, []string{title})
}

func TermuxDialogConfirm(td TDialogConfirm) {
	ExecAndListen(TD, []string{
		"confirm",
		"-i", td.Hint,
		"-s", td.Title,
	})
}

func TermuxDialogCheckbox(td TDialogCheckbox) {
	values := strings.Join(td.Values, ",")
	ExecAndListen(TD, []string{
		"checkbox",
		"-v", values,
		"-t", td.Title,
	})
}

func TermuxDialogCounter(td TDialogCounter) {
	values := fmt.Sprintf("%d,%d,%d", td.Min, td.Max, td.Start)
	ExecAndListen(TD, []string{
		"counter",
		"-r", values,
		"-t", td.Title,
	})
}

func TermuxDialogDate(td TDialogDate) {
	date := fmt.Sprintf("\"%d-%d-%d %d:%d:%d\"", td.Day, td.Month, td.Year, td.KHours, td.Minutes, td.Seconds)
	ExecAndListen(TD, []string{
		"date",
		"-d", date,
		"-t", td.Title,
	})
}

func TermuxDialogeWithoutDate(td TDialog) {
	ExecAndListen(TD, []string{
		"date",
		"-t", td.Title,
	})
}

func TermuxDialogRadio(td TDialogRadio) {
	values := strings.Join(td.Values, ",")
	ExecAndListen(TD, []string{
		"radio",
		"-v", values,
		"-t", td.Title,
	})
}

func TermuxDialogSheet(td TDialogRadio) {
	values := strings.Join(td.Values, ",")
	ExecAndListen(TD, []string{
		"sheet",
		"-v", values,
		"-t", td.Title,
	})
}

func TermuxDialogSpinner(td TDialogRadio) {
	values := strings.Join(td.Values, ",")
	ExecAndListen(TD, []string{
		"sheet",
		"-v", values,
		"-t", td.Title,
	})
}

func TermuxDialogSpeech(td TDialogSpeech) {
	ExecAndListen(TD, []string{
		"speech",
		"-i", td.Hint,
		"-t", td.Title,
	})
}

func TermuxDialogText(td TDialogText) {
	if td.MultipleLine == true && td.NumberInput == true {
		log.Fatalln("Cannot use multilines with input numbers (see wiki.termux.com/wiki/Termux-dialog)")
	}

	command := []string{
		"text",
		"-i", td.Hint,
		"-t", td.Title,
	}

	if td.MultipleLine == true {
		command = append(command, "-m")
	}

	if td.NumberInput == true {
		command = append(command, "-n")
	}

	ExecAndListen(TD, command)
}

func TermuxDialogTime(td TDialogTime) {
	ExecAndListen(TD, []string{
		"time",
		"-t", td.Title,
	})
}

func TermuxBatteryStatus() TBattery {
	t := TBattery{}
	status := ExecAndListen("termux-battery-status", nil)

	log.Println(status)

	err := json.Unmarshal([]byte(status), &t)
	if err != nil {
		log.Fatalln(err)
	}

	return t
}

func ExecAndListen(command string, args []string) string {
	cmd := exec.Command(command, args...)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	return string(cmdOutput.Bytes())
}
