// Package screenshot captures screen-shot image as image.RGBA.
// Mac, Windows, Linux, FreeBSD, OpenBSD, NetBSD, and Solaris are supported.
package screenshot

import (
	"image"

	win "github.com/lxn/win"
)

// CaptureDisplay captures whole region of displayIndex'th display.
func CaptureDisplay(displayIndex int) (*image.RGBA, error) {
	rect := GetDisplayBounds(displayIndex)
	return CaptureRect(rect)
}

// CaptureRect captures specified region of desktop.
func CaptureRect(rect image.Rectangle) (*image.RGBA, error) {
	return Capture(rect.Min.X, rect.Min.Y, rect.Dx(), rect.Dy())
}

//Capture a specific Application
func CaptureApp(windowName *uint16) (*image.RGBA, error) {

	hwnd := FindWindowW(nil, windowName)
	var rect win.RECT
	GetWindowRect(hwnd, &rect)
	newrect := image.Rect(int(rect.Left), int(rect.Top), int(rect.Right), int(rect.Bottom))
	return CaptureRect(newrect)
}
