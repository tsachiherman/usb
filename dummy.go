//go:build dummy
// +build dummy

// This file is part of a workaround for `go mod vendor` which won't vendor
// C files if there's no Go file in the same directory.
// This would prevent the hidapi/hidapi/hidapi.h file to be vendored.
//
// This Go file imports the c directory where there is another dummy.go file which
// is the second part of this workaround.
//
// These two files combined make it so `go mod vendor` behaves correctly.
//
// See this issue for reference: https://github.com/golang/go/issues/26366

package main

import (
	_ "github.com/tsachiherman/usb/hidapi"
	_ "github.com/tsachiherman/usb/hidapi/hidapi"
	_ "github.com/tsachiherman/usb/hidapi/libusb"
	_ "github.com/tsachiherman/usb/hidapi/mac"
	_ "github.com/tsachiherman/usb/hidapi/windows"
	_ "github.com/tsachiherman/usb/libusb"
	_ "github.com/tsachiherman/usb/libusb/libusb"
	_ "github.com/tsachiherman/usb/libusb/libusb/os"
)
