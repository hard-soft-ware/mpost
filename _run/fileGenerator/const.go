package main

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/_run/fileGenerator/generator"
)

////////////////////////////////////

const constDir = "consts"

////

func Cmd() {
	obj := generator.Init("Cmd", constDir+"/cmd.go")
	val := obj.NewVal()

	val.Add(0x10, "Omnibus")
	val.Add(0x20, "RetStatus")
	val.Add(0x40, "Calibrate")
	val.Add(0x50, "FlashDownload")
	val.Add(0x60, "Auxiliary")
	val.Add(0x70, "Expanded")

	constGenerator(obj, val)
	obj.Save(constDir).End()
	fmt.Printf("Add: %s %s\n", constDir, obj.Name.Get())
}

func CmdAux() {
	obj := generator.Init("CmdAux", constDir+"/cmdAux.go")
	val := obj.NewVal()

	val.Add(0x00, "SoftwareCRC")
	val.Add(0x01, "CashBoxTotal")
	val.Add(0x02, "DeviceResets")
	val.Add(0x10, "BNFStatus")
	val.Add(0x11, "SetBezel")
	val.Add(0x0D, "DeviceCapabilities")
	val.Add(0x03, "ClearCashBoxTotal").Delim()

	val.Add(0x04, "AcceptorType")
	val.Add(0x0E, "AcceptorApplicationID")
	val.Add(0x0F, "AcceptorVariantID")
	val.Add(0x05, "AcceptorSerialNumber")
	val.Add(0x06, "AcceptorBootPartNumber")
	val.Add(0x07, "AcceptorApplicationPartNumber")
	val.Add(0x08, "AcceptorVariantName")
	val.Add(0x09, "AcceptorVariantPartNumber")
	val.Add(0x0A, "AcceptorAuditLifeTimeTotals")
	val.Add(0x0B, "AcceptorAuditQPMeasures")
	val.Add(0x0C, "AcceptorAuditPerformanceMeasures")

	constGenerator(obj, val)
	obj.Save(constDir).End()
	fmt.Printf("Add: %s %s\n", constDir, obj.Name.Get())
}

func Data() {
	obj := generator.Init("Data", constDir+"/data.go")
	val := obj.NewVal()

	val.Add(0x02, "STX")     // Начало текста
	val.Add(0x03, "ETX")     // Конец текста
	val.Add(0x06, "ACKMask") // Маска подтверждения

	constGenerator(obj, val)
	obj.Save(constDir).End()
	fmt.Printf("Add: %s %s\n", constDir, obj.Name.Get())
}
