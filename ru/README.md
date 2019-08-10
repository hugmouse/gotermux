<p align="center">
    <img src="https://raw.githubusercontent.com/hugmouse/gotermux/Development/icon/logo.webp.webp" height="256">
</p>

# GoTermux

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/380f19e0a1bc4fb19d3eeafa914fc1ad)](https://www.codacy.com/app/mysh/gotermux?utm_source=github.com&utm_medium=referral&utm_content=hugmouse/gotermux&utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/hugmouse/gotermux)](https://goreportcard.com/report/github.com/hugmouse/gotermux)
[![GoDoc](https://godoc.org/github.com/hugmouse/gotermux?status.svg)](https://godoc.org/github.com/hugmouse/gotermux)

`Обертка вокруг Termux API`

В настоящее время в активной разработке!

## Codestyle
Вызов команды termux из golang выглядит так же, как в termux cmd, но без тире (-).

Пример:

```shell
// Termux CMD:
termux-battery-status
// Golang: 
TermuxBatteryStatus()
```

## Установка Termux API
Используйте [F-droid](https://f-droid.org/packages/com.termux.api/) или [Google Play](https://play.google.com/store/apps/details?id=com.termux.api) для установки Termux API.

**Не смешивайте установки Termux и его аддонов между Google Play и F-Droid**.

Чтобы использовать [termux-api](https://github.com/termux/termux-api-package) вам так же нужно установить пакет внутри Termux:
```shell
pkg install termux-api
```

## Установка пакета
Сначала вам нужно установить Golang:
```shell
pkg install golang
```

И потом:

```shell
go get -u github.com/hugmouse/gotermux
```

## TODO

### По проекту

- [ ] Добавить примеры
- [ ] Добавить более подробное описание функций

### По Termux API

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
  - [x] termux-media-player play
  - [x] termux-media-player play (resume)
  - [x] termux-media-player info
  - [x] termux-media-player pause
  - [x] termux-media-player stop

- [x] termux-media-scan

- [ ] termux-microphone-record

- [ ] termux-sensor

- [ ] termux-messages-remove

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

- [ ] termux-tts-engine

- [ ] termux-telephony-call

- [x] termux-vibrate

- [x] termux-volume

- [ ] termux-wallpaper

- [x] termux-wifi-connectioninfo

- [ ] termux-wifi-enable

- [ ] termux-wifi-scaninfo

## Тесты функций Termux API

- [x] termux-dialog
  - [x] dialog-confirm
  - [x] dialog-checkbox
  - [x] dialog-counter
  - [ ] dialog-date
  - [x] dialog-sheet
  - [x] dialog-spinner
  - [ ] dialog-text
  - [x] dialog-speech
  - [ ] dialog-time
