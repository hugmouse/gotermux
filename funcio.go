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

func TermuxDialog(title string) []byte {
	executed := ExecAndListen(TD, []string{
		"-t", title})
	return executed
}

func TermuxDialogConfirm(td TDialogConfirm) []byte {
	executed := ExecAndListen(TD, []string{
		"confirm",
		"-i", td.Hint,
		"-t", td.Title,
	})
	return executed
}

func TermuxDialogCheckbox(td TDialogCheckbox) []byte {
	values := strings.Join(td.Values, ",")
	executed := ExecAndListen(TD, []string{
		"checkbox",
		"-v", values,
		"-t", td.Title,
	})
	return executed
}

func TermuxDialogCounter(td TDialogCounter) []byte {
	values := fmt.Sprintf("%d,%d,%d", td.Min, td.Max, td.Start)
	executed := ExecAndListen(TD, []string{
		"counter",
		"-r", values,
		"-t", td.Title,
	})
	return executed
}

func TermuxDialogDate(td TDialogDate) []byte {
	date := fmt.Sprintf("\"%d-%d-%d %d:%d:%d\"", td.Day, td.Month, td.Year, td.KHours, td.Minutes, td.Seconds)
	executed := ExecAndListen(TD, []string{
		"date",
		"-d", date,
		"-t", td.Title,
	})
	return executed
}

func TermuxDialogeWithoutDate(td TDialog) []byte {
	executed := ExecAndListen(TD, []string{
		"date",
		"-t", td.Title,
	})
	return executed
}

func TermuxDialogRadio(td TDialogRadio) []byte {
	values := strings.Join(td.Values, ",")
	executed := ExecAndListen(TD, []string{
		"radio",
		"-v", values,
		"-t", td.Title,
	})
	return executed
}

func TermuxDialogSheet(td TDialogRadio) []byte {
	values := strings.Join(td.Values, ",")
	executed := ExecAndListen(TD, []string{
		"sheet",
		"-v", values,
		"-t", td.Title,
	})
	return executed
}

func TermuxDialogSpinner(td TDialogRadio) []byte {
	values := strings.Join(td.Values, ",")
	executed := ExecAndListen(TD, []string{
		"spinner",
		"-v", values,
		"-t", td.Title,
	})
	return executed
}

func TermuxDialogSpeech(td TDialogSpeech) []byte {
	executed := ExecAndListen(TD, []string{
		"speech",
		"-i", td.Hint,
		"-t", td.Title,
	})
	return executed
}

func TermuxDialogText(td TDialogText) []byte {
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

	executed := ExecAndListen(TD, command)
	return executed
}

func TermuxDialogTime(td TDialogTime) []byte {
	executed := ExecAndListen(TD, []string{
		"time",
		"-t", td.Title,
	})
	return executed
}

func TermuxBatteryStatus() TBattery {
	t := TBattery{}
	status := ExecAndListen("termux-battery-status", nil)

	err := json.Unmarshal(status, &t)
	if err != nil {
		log.Fatalln(err)
	}

	return t
}

func ExecAndListen(command string, args []string) []byte {
	log.Printf("Arguments: %+v\n", args)
	cmd := exec.Command(command, args...)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	return cmdOutput.Bytes()
}
