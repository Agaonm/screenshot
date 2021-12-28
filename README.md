screenshot
==========

![](https://github.com/kbinani/screenshot/actions/workflows/go.yml/badge.svg)
[![](https://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/github.com/kbinani/screenshot)
[![](https://img.shields.io/badge/license-MIT-428F7E.svg?style=flat)](https://github.com/kbinani/screenshot/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/kbinani/screenshot)](https://goreportcard.com/report/github.com/kbinani/screenshot)

* Go library to capture desktop screen.
* Support Windows, Mac, Linux, FreeBSD, OpenBSD, NetBSD, and Solaris environment.
* Multiple display supported.
* `cgo` free for Windows, Linux, FreeBSD, OpenBSD, NetBSD, and Solaris.

* Added Window Capture for Specific Application on Windows - Agaonm

example to Capture Single Application
=======

* sample program `main.go`

	```go
	package main

	import (
		"github.com/kbinani/screenshot"
		"image/png"
		"os"
		"fmt"
	)

	func main() {
		window := gocv.NewWindow("Capture")

		for {
			winName := "Spotify Premium"
			u16fname, err := syscall.UTF16FromString(winName)

			img, err := screenshot.CaptureApp2(&u16fname[0])
			if err != nil {
				panic(err)
			}

			imgMat, err := gocv.ImageToMatRGBA(img)
			if err != nil {
				panic(err)
			}

			window.IMShow(imgMat)

			window.WaitKey(1)
		}
	}
	```

coordinate
=================
Y-axis is downward direction in this library. The origin of coordinate is upper-left corner of main display. This means coordinate system is similar to Windows OS

license
=======

MIT Licence

author
======

kbinani
