package main

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/_run/fileGenerator/generator"
)

////////////////////////////////////

const enumDir = "enum"

////

func Status() {
	obj := generator.Init("BNFStatus", enumDir+"/status.go")

	val := []string{
		"Unknown",
		"Error",
		"OK",
		"NotAttached",
	}

	enumGenerator(obj, val)
	obj.Save(enumDir).End()
	fmt.Printf("Add: %s %s\n", enumDir, obj.Name.Get())
}

func Document() {
	obj := generator.Init("Document", enumDir+"/document.go")

	val := []string{
		"None",
		"NoValue",
		"Bill",
		"Barcode",
		"Coupon",
	}

	enumGenerator(obj, val)
	obj.Save(enumDir).End()
	fmt.Printf("Add: %s %s\n", enumDir, obj.Name.Get())
}

func Orientation() {
	obj := generator.Init("Orientation", enumDir+"/orientation.go")

	val := []string{
		"RightUp",
		"RightDown",
		"LeftUp",
		"LeftDown",
		"UnknownOrientation",
	}

	enumGenerator(obj, val)
	obj.Save(enumDir).End()
	fmt.Printf("Add: %s %s\n", enumDir, obj.Name.Get())
}

func OrientationControl() {
	obj := generator.Init("OrientationControl", enumDir+"/orientation_control.go")

	val := []string{
		"FourWay",
		"TwoWay",
		"OneWay",
	}

	enumGenerator(obj, val)
	obj.Save(enumDir).End()
	fmt.Printf("Add: %s %s\n", enumDir, obj.Name.Get())
}

func PowerUp() {
	obj := generator.Init("PowerUp", enumDir+"/powerUp.go")

	val := []string{
		"A",
		"B",
		"C",
		"E",
	}

	enumGenerator(obj, val)
	obj.Save(enumDir).End()
	fmt.Printf("Add: %s %s\n", enumDir, obj.Name.Get())
}

func PupExt() {
	obj := generator.Init("PupExt", enumDir+"/pupExt.go")

	val := []string{
		"Return",
		"OutOfService",
		"Stack",
		"StackNoCredit",
		"Wait",
		"WaitNoCredit",
	}

	enumGenerator(obj, val)
	obj.Save(enumDir).End()
	fmt.Printf("Add: %s %s\n", enumDir, obj.Name.Get())
}

func State() {
	obj := generator.Init("State", enumDir+"/state.go")

	val := []string{
		"Escrow",
		"Stacked",
		"Returned",
		"Rejected",
		"Stalled",
		"Accepting",
		"CalibrateStart",
		"Calibrating",
		"Connecting",
		"Disconnected",
		"Downloading",
		"DownloadRestart",
		"DownloadStart",
		"Failed",
		"Idling",
		"PupEscrow",
		"Returning",
		"Stacking",
	}

	enumGenerator(obj, val)
	obj.Save(enumDir).End()
	fmt.Printf("Add: %s %s\n", enumDir, obj.Name.Get())
}

func Bezel() {
	obj := generator.Init("Bezel", enumDir+"/bezel.go")

	val := []string{
		"Standard",
		"Platform",
		"Diagnostic",
	}

	enumGenerator(obj, val)
	obj.Save(enumDir).End()
	fmt.Printf("Add: %s %s\n", enumDir, obj.Name.Get())
}

func Event() {
	obj := generator.Init("Event", enumDir+"/event.go")

	val := []string{
		"_Begin",
		"Connected",
		"Escrow",
		"PUPEscrow",
		"Stacked",
		"Returned",
		"Rejected",
		"Cheated",
		"StackerFull",
		"CalibrateStart",
		"CalibrateProgress",
		"CalibrateFinish",
		"DownloadStart",
		"DownloadRestart",
		"DownloadProgress",
		"DownloadFinish",
		"PauseDetected",
		"PauseCleared",
		"StallDetected",
		"StallCleared",
		"JamDetected",
		"JamCleared",
		"PowerUp",
		"InvalidCommand",
		"CashBoxAttached",
		"CashBoxRemoved",
		"Disconnected",
		"_End",
	}

	enumGenerator(obj, val)
	obj.Save(enumDir).End()
	fmt.Printf("Add: %s %s\n", enumDir, obj.Name.Get())
}
