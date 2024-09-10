package mpost

import (
	"errors"
	"fmt"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"os"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) FlashDownload(filePath string) (err error) {
	a.log.Msg("FlashDownload")

	if !acceptor.Connected && acceptor.Device.State != enum.StateDownloadRestart {
		err = errors.New("FlashDownload not allowed when not connected or not in DownloadRestart state")
		a.log.Err("FlashDownload", err)
		return
	}

	if !(acceptor.Device.Model >= 65 && acceptor.Device.Model <= 71 ||
		acceptor.Device.Model == 77 ||
		acceptor.Device.Model == 80 ||
		acceptor.Device.Model == 81 ||
		acceptor.Device.Model == 87 ||
		acceptor.Device.Model == 88) {
		if acceptor.Device.State != enum.StateIdling && acceptor.Device.State != enum.StateDownloadRestart {
			err = errors.New("FlashDownload allowed only when DeviceState is Idling or DownloadRestart")
			a.log.Err("FlashDownload", err)
			return
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to open flash download file: %s", err))
		a.log.Err("FlashDownload", err)
		return
	}
	defer file.Close()

	// Determine the file size and check if divisible by 32
	fileInfo, err := file.Stat()
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to stat flash download file: %s", err))
		a.log.Err("FlashDownload", err)
		return
	}
	fileSize := fileInfo.Size()
	if fileSize%32 != 0 {
		err = errors.New("Flash download file size must be divisible by 32")
		a.log.Err("FlashDownload", err)
		return
	}

	go func() {
		defer file.Close()
		a.flashDownload(file, fileSize)
	}()

	return
}

func (a *CAcceptor) flashDownload(downloadFile *os.File, fileSize int64) {
	a.log.Msg("FlashDownloadThread started")

	if acceptor.Device.State != enum.StateDownloadRestart {
		acceptor.Device.State = enum.StateDownloadStart
	}

	packetNum := 0
	finalPacketNum := int(fileSize / 32)
	payload := []byte{consts.CmdCalibrate.Byte(), 0x00, 0x00, 0x00}
	var reply []byte
	var err error

	for {
		if acceptor.StopFlashDownloadThread {
			acceptor.Device.State = enum.StateIdling
			return
		}

		reply, err = a.SendSynchronousCommand(payload)
		if err != nil || len(reply) == 0 {
			if !acceptor.Connected {
				a.RaiseDownloadFinishEvent(false)
				acceptor.Device.State = enum.StateIdling
				return
			}
			continue
		}

		if (reply[2] & 0x70) == 0x20 {
			break
		}
		packetNum = (int(reply[3]&0x0F)<<12 + int(reply[4]&0x0F)<<8 + int(reply[5]&0x0F)<<4 + int(reply[6]&0x0F) + 1) & 0xFFFF
	}

	a.RaiseDownloadStartEvent(finalPacketNum)
	timeoutStartTickCount := time.Now()

	for packetNum < finalPacketNum {
		if acceptor.StopFlashDownloadThread {
			acceptor.Device.State = enum.StateIdling
			return
		}

		dataBuffer := make([]byte, 32)
		_, err := downloadFile.ReadAt(dataBuffer, int64(packetNum*32))
		if err != nil {
			time.Sleep(200 * time.Millisecond)
			continue
		}

		payload = make([]byte, 69)
		payload[0] = consts.CmdFlashDownload.Byte()
		payload[1] = byte((packetNum & 0xF000) >> 12)
		payload[2] = byte((packetNum & 0x0F00) >> 8)
		payload[3] = byte((packetNum & 0x00F0) >> 4)
		payload[4] = byte(packetNum & 0x000F)
		copy(payload[5:], dataBuffer)

		reply, err = a.SendSynchronousCommand(payload)
		if err != nil || len(reply) != 9 {
			time.Sleep(200 * time.Millisecond)
			continue
		}

		if reply[0] == consts.DataSTX.Byte() {
			acceptor.Device.State = enum.StateDownloading
			if acceptor.ShouldRaise.DownloadProgressEvent {
				a.RaiseDownloadProgressEvent(packetNum)
			}
			packetNum++
		} else {
			time.Sleep(200 * time.Millisecond)
			continue
		}

		if time.Since(timeoutStartTickCount) > acceptor.Timeout.Download {
			a.RaiseDownloadFinishEvent(false)
			acceptor.Device.State = enum.StateIdling
			return
		}
	}

	time.Sleep(30 * time.Millisecond)
	a.RaiseDownloadFinishEvent(true)
	acceptor.Device.State = enum.StateIdling
	acceptor.Connected = true
	a.RaiseConnectedEvent()
}
