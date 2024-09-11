package serial

import (
	"go.bug.st/serial"
	"time"
)

////////////////////////////////////

type SerialStruct struct {
	PortName string
	connect  *bool

	port serial.Port
	mode *serial.Mode
}

func Open(portName string, connectStatus *bool) (*SerialStruct, error) {
	obj := SerialStruct{}
	obj.PortName = portName
	obj.connect = connectStatus

	obj.mode = &serial.Mode{
		BaudRate: 9600,
		DataBits: 7,
		Parity:   serial.EvenParity,
		StopBits: serial.OneStopBit,
	}

	err := obj.open()
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (obj *SerialStruct) open() error {
	port, err := serial.Open(obj.PortName, obj.mode)
	if err != nil {
		return err
	}

	port.SetReadTimeout(100 * time.Millisecond)
	port.ResetInputBuffer()

	port.SetDTR(false)
	port.SetRTS(true)
	time.Sleep(100 * time.Millisecond)

	port.SetDTR(true)
	port.SetRTS(false)
	time.Sleep(5 * time.Millisecond)

	port.GetModemStatusBits()

	port.ResetInputBuffer()
	obj.port = port
	*obj.connect = true
	return nil
}

func (obj *SerialStruct) Restart() error {
	if *obj.connect {
		obj.port.Close()
		time.Sleep(100 * time.Millisecond)
	}

	err := obj.open()
	if err != nil {
		return err
	}

	return nil
}

func (obj *SerialStruct) Close() error {
	*obj.connect = false
	return obj.port.Close()
}
