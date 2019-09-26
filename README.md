<p align="center">
    <img src="https://raw.githubusercontent.com/hugmouse/gotermux/Development/icon/logo.webp">
</p>

# GoTermux

**Code quality checks:**
[![CodeFactor](https://www.codefactor.io/repository/github/hugmouse/gotermux/badge)](https://www.codefactor.io/repository/github/hugmouse/gotermux)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/380f19e0a1bc4fb19d3eeafa914fc1ad)](https://www.codacy.com/app/mysh/gotermux?utm_source=github.com&utm_medium=referral&utm_content=hugmouse/gotermux&utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/hugmouse/gotermux)](https://goreportcard.com/report/github.com/hugmouse/gotermux)

**Documentation reference:**
[![GoDoc](https://godoc.org/github.com/hugmouse/gotermux?status.svg)](https://godoc.org/github.com/hugmouse/gotermux)

**Examples and usage of GoTermux:** [![gotermux.mysh.dev](https://img.shields.io/badge/GoTermux-examples-green)](https://gotermux.mysh.dev)

`GoTermux` gives you the opportunity to execute your scripts in `Termux` using `Golang`. All from `Go` and `Termux API` features in your hands!

-   GoTermux package
    -   [Codestyle](#codestyle)
    -   [Examples](#examples)
    -   [Installing Termux API](#installing-termux-api)
    -   [Installing the package](#installing-the-package)
    -   [TODO](#todo)

## Codestyle

Calling termux command from golang looks just like in termux cmd, but without dash (`-`). 

Example: 

```shell
// Termux CMD:
termux-battery-status
// Golang: 
TermuxBatteryStatus()
```

## Examples

Code examples can be found at: [gotermux.mysh.dev](https://gotermux.mysh.dev/)!

Or here: [now.sh](https://gotermux.mysh.now.sh/)

## Installing Termux API

Use [F-droid](https://f-droid.org/packages/com.termux.api/) or [Google Play](https://play.google.com/store/apps/details?id=com.termux.api) to install Termux API.

**Do not mix installations of Termux and Addons between Google Play and F-Droid**.

To **use** Termux API you also need to install the [termux-api](https://github.com/termux/termux-api-package) package:

```shell
pkg install termux-api
```

## Installing the package

First of all you need install golang:

```shell
pkg install golang
```

And then: 

```shell
go get -u github.com/hugmouse/gotermux
```

## TODO

Now all the implemented, as well as the planned features are now [here](https://github.com/hugmouse/gotermux/projects/1)!
