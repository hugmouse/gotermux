package gotermuxwrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	TD = "termux-dialog"
	RT = TResult{}
)

// TermuxDialog spawns new dialog with only title in it
func TermuxDialog(title string) TResult {
	executed := ExecAndListen(TD, []string{
		"-t", title})
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogConfirm spawns new confirmation dialog
func TermuxDialogConfirm(td TDialogConfirm) TResult {
	executed := ExecAndListen(TD, []string{
		"confirm",
		"-i", td.Hint,
		"-t", td.Title,
	})
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogCheckbox spawns new dialog with multiple values using checkboxes
func TermuxDialogCheckbox(td TDialogCheckbox) TResult {
	values := strings.Join(td.Values, ",")
	executed := ExecAndListen(TD, []string{
		"checkbox",
		"-v", values,
		"-t", td.Title,
	})
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogCounter spawns new dialog with pick function in it
// User can pick a number in specified range
func TermuxDialogCounter(td TDialogCounter) TResult {
	values := fmt.Sprintf("%d,%d,%d", td.Min, td.Max, td.Start)
	executed := ExecAndListen(TD, []string{
		"counter",
		"-r", values,
		"-t", td.Title,
	})
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogDate spawns new dialog with pick function in it
// User can pick a date, but this function returns only that date that you provide
func TermuxDialogDate(td TDialogDate) TResult {
	date := fmt.Sprintf("\"%d-%d-%d %d:%d:%d\"", td.Day, td.Month, td.Year, td.KHours, td.Minutes, td.Seconds)
	executed := ExecAndListen(TD, []string{
		"date",
		"-d", date,
		"-t", td.Title,
	})
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogWithoutDate spawns new dialog with pick function in it
// User can pick a date
func TermuxDialogWithoutDate(td TDialog) TResult {
	executed := ExecAndListen(TD, []string{
		"date",
		"-t", td.Title,
	})
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogRadio spawns new dialog with pick function in it
// User can pick a single value from radio buttons
func TermuxDialogRadio(td TDialogRadio) TResult {
	values := strings.Join(td.Values, ",")
	executed := ExecAndListen(TD, []string{
		"radio",
		"-v", values,
		"-t", td.Title,
	})
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogSheet spawns new dialog with pick function in it
// User can pick a value from sliding bottom sheet
func TermuxDialogSheet(td TDialogRadio) TResult {
	values := strings.Join(td.Values, ",")
	executed := ExecAndListen(TD, []string{
		"sheet",
		"-v", values,
		"-t", td.Title,
	})
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogSpinner spawns new dialog with pick function in it
// User can pick a single value from a dropdown spinner
func TermuxDialogSpinner(td TDialogRadio) TResult {
	values := strings.Join(td.Values, ",")
	executed := ExecAndListen(TD, []string{
		"spinner",
		"-v", values,
		"-t", td.Title,
	})
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogSpeech spawns a new dialog that can obtain speech using device microphone
func TermuxDialogSpeech(td TDialogSpeech) TResult {
	executed := ExecAndListen(TD, []string{
		"speech",
		"-i", td.Hint,
		"-t", td.Title,
	})
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogText spawns a new dialog with input text (default if no widget specified)
func TermuxDialogText(td TDialogText) TResult {
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
	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxDialogTime spawns new dialog with pick function in it
// User can pick a time value
func TermuxDialogTime(td TDialogTime) TResult {
	executed := ExecAndListen(TD, []string{
		"time",
		"-t", td.Title,
	})

	err := json.Unmarshal(executed, &RT)
	if err != nil {
		log.Println(err)
	}
	return RT
}

// TermuxBatteryStatus returns the status of the device battery
func TermuxBatteryStatus() TBattery {
	t := TBattery{}
	status := ExecAndListen("termux-battery-status", nil)

	err := json.Unmarshal(status, &t)
	if err != nil {
		log.Fatalln(err)
	}

	return t
}

// TermuxBrightness sets the display brightness.
// Note that this may not work if automatic brightness control is enabled.
func TermuxBrightness(val uint8) []byte {
	u := strconv.FormatUint(uint64(val), 10)
	json := ExecAndListen("termux-brightness", []string{
		u})
	return json
}

// TODO: test it out on <9 android and somehow downgrade Termux API
// TermuxCallLog prints the phone call history
// Not working on >=9 android
// Works only in 0.32v of Termux API
// See: https://github.com/termux/termux-api/commit/de44896a01111506590a258f0267400af067b778#diff-99a2dbdeb4c1195cba0edb66ea510428
func TermuxCallLog(limit, offset uint) TCalls {
	c := TCalls{}
	executed := ExecAndListen("termux-call-log", []string{
		"-l", string(limit),
		"-o", string(offset),
	})
	err := json.Unmarshal(executed, &c)
	if err != nil {
		log.Println(err)
	}

	return c
}

// ExecAndListen is a function, that build around "exec.Command()"
// returns cmd output
func ExecAndListen(command string, args []string) []byte {
	log.Printf("Arguments: %+v\n", args)
	cmd := exec.Command(command, args...)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		_, err := os.Stderr.WriteString(err.Error())
		if err !=nil {
			log.Fatalln("I really don't know how you done this. But you did.", err)
		}
	}
	return cmdOutput.Bytes()
}
