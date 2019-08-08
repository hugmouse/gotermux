<p align="center">
    <img src="https://raw.githubusercontent.com/hugmouse/gotermux/Development/icon/logo.webp.webp" height="256">
</p>

# GoTermux

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/380f19e0a1bc4fb19d3eeafa914fc1ad)](https://www.codacy.com/app/mysh/gotermux?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=hugmouse/gotermux&amp;utm_campaign=Badge_Grade)

`Simple wrapper around Termux API`

Currently in active development!

## Codestyle
Calling termux command from golang looks just like in termux cmd, but without dash (`-`). 

Example: 
```shell
// Termux CMD:
termux-battery-status
// Golang: 
TermuxBatteryStatus()
```

## Installing
```shell
go get -u github.com/hugmouse/gotermux
```

## TODO

### More project related
- [ ] Add examples
- [ ] Add better description of API, in contrast to the official in Termux

### Termux API
- [x] termux-battery-status

- [x] termux-brightness

- [x] termux-call-log

- [ ] termux-camera-info

- [ ] termux-camera-photo

- [x] termux-clipboard-get

- [x] termux-clipboard-set

- [x] termux-contact-list

- [x] termux-dialog
  - [x] dialog-confirm
  - [x] dialog-checkbox
  - [x] dialog-counter
  - [ ] dialog-date
  - [x] dialog-sheet
  - [x] dialog-spinner
  - [x] dialog-speech
  - [x] dialog-text
  - [x] dialog-time

- [x] termux-download

- [ ] termux-fingerprint

- [x] termux-infrared-frequencies

- [x] termux-infrared-transmit

- [ ] termux-job-scheduler

- [x] termux-location

- [x] termux-media-player
  - [x] termux-media-player play <file>
  - [x] termux-media-player play (resume)
  - [x] termux-media-player info
  - [x] termux-media-player pause
  - [x] termux-media-player stop

- [x] termux-media-scan

- [ ] termux-microphone-record

- [ ] termux-notification

- [ ] termux-notification-remove

- [ ] termux-sensor

- [x] termux-share

- [ ] termux-sms-list 

- [ ] termux-sms-send

- [ ] termux-storage-get

- [ ] termux-telephony-call

- [ ] termux-telephony-cellinfo

- [ ] termux-telephony-deviceinfo

- [ ] termux-toast

- [ ] termux-torch

- [ ] termux-tts-engines

- [ ] termux-tts-speak

- [x] termux-vibrate

- [ ] termux-volume

- [ ] termux-wallpaper

- [ ] termux-wifi-connectioninfo

- [ ] termux-wifi-enable

- [ ] termux-wifi-scaninfo

## Termux API tests

- [x] termux-dialog
  - [x] dialog-confirm
  - [x] dialog-checkbox
  - [x] dialog-counter
  - [ ] dialog-date
  - [x] dialog-sheet
  - [x] dialog-spinner
  - [x] dialog-speech
  - [ ] dialog-text
  - [ ] dialog-time
