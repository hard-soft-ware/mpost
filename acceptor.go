package mpost

import (
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.bug.st/serial"
	"os"
	"sync"
	"time"
)

////////////////////////////////////

type EventHandler func(*CAcceptor, int)

type CAcceptor struct {
	port                serial.Port
	auditLifeTimeTotals []int
	auditPerformance    []int
	auditQP             []int
	autoStack           bool
	barCode             string

	bill             CBill
	billTypes        []CBill
	billTypeEnables  []bool
	billValues       []CBill
	billValueEnables []bool

	bnfStatus                          enum.BNFStatusType
	bootPN                             string
	capApplicationID                   bool
	capApplicationPN                   bool
	capAssetNumber                     bool
	capAudit                           bool
	capBarCodes                        bool
	capBarCodesExt                     bool
	capBNFStatus                       bool
	capBookmark                        bool
	capBootPN                          bool
	capCalibrate                       bool
	capCashBoxTotal                    bool
	capCouponExt                       bool
	capDevicePaused                    bool
	capDeviceSoftReset                 bool
	capDeviceType                      bool
	capDeviceResets                    bool
	capDeviceSerialNumber              bool
	capEscrowTimeout                   bool
	capFlashDownload                   bool
	capNoPush                          bool
	capOrientationExt                  bool
	capPupExt                          bool
	capTestDoc                         bool
	capVariantID                       bool
	capVariantPN                       bool
	cashBoxAttached                    bool
	cashBoxFull                        bool
	cashBoxTotal                       int
	connected                          bool
	coupon                             *CCoupon
	deviceFailure                      bool
	deviceModel                        int
	devicePaused                       bool
	devicePortName                     string
	devicePowerUp                      enum.PowerUpType
	deviceResets                       int
	deviceRevision                     int
	deviceSerialNumber                 string
	deviceStalled                      bool
	deviceState                        enum.StateType
	deviceType                         string
	docType                            enum.DocumentType
	enableAcceptance                   bool
	enableBarCodes                     bool
	enableBookmarks                    bool
	enableCouponExt                    bool
	enableNoPush                       bool
	escrowOrientation                  enum.OrientationType
	highSecurity                       bool
	orientationCtl                     enum.OrientationControlType
	orientationCtlExt                  enum.OrientationControlType
	version                            string
	transactionTimeout                 time.Duration
	downloadTimeout                    time.Duration
	inSoftResetOneSecondIgnore         bool
	inSoftResetWaitForReply            bool
	expandedNoteReporting              bool
	isQueryDeviceCapabilitiesSupported bool
	isDeviceJammed                     bool
	isCheated                          bool
	isPoweredUp                        bool
	isInvalidCommand                   bool
	wasDocTypeSetOnEscrow              bool
	wasDisconnected                    bool
	isVeryFirstPoll                    bool
	shouldRaiseConnectedEvent          bool
	shouldRaiseEscrowEvent             bool
	shouldRaisePUPEscrowEvent          bool
	shouldRaiseStackedEvent            bool
	shouldRaiseReturnedEvent           bool
	shouldRaiseRejectedEvent           bool
	shouldRaiseCheatedEvent            bool
	shouldRaiseStackerFullEvent        bool
	shouldRaiseCalibrateStartEvent     bool
	shouldRaiseCalibrateProgressEvent  bool
	shouldRaiseCalibrateFinishEvent    bool
	shouldRaiseDownloadStartEvent      bool
	shouldRaiseDownloadRestartEvent    bool
	shouldRaiseDownloadProgressEvent   bool
	shouldRaiseDownloadFinishEvent     bool
	shouldRaisePauseDetectedEvent      bool
	shouldRaisePauseClearedEvent       bool
	shouldRaiseStallDetectedEvent      bool
	shouldRaiseStallClearedEvent       bool
	shouldRaiseJamDetectedEvent        bool
	shouldRaiseJamClearedEvent         bool
	shouldRaisePowerUpEvent            bool
	shouldRaiseInvalidCommandEvent     bool
	shouldRaiseCashBoxAttachedEvent    bool
	shouldRaiseCashBoxRemovedEvent     bool
	shouldRaiseDisconnectedEvent       bool
	compressLog                        bool
	workerThread                       sync.WaitGroup
	openThread                         chan bool
	flashDownloadThread                chan bool
	dataLinkLayer                      *CDataLinkLayer
	replyQueuedEvent                   int
	notInProcessReplyEvent             int
	stopWorkerThread                   bool
	stopOpenThread                     bool
	stopFlashDownloadThread            bool
	suppressStandardPoll               bool

	messageQueue chan *CMessage
	replyQueue   chan []byte
	wasStopped   bool

	isReplyAcked          bool
	signalMainThreadEvent int
	eventHandlers         map[enum.EventType]EventHandler

	log *LogGlobalStruct
}

func NewCAcceptor(transactionTimeout, downloadTimeout time.Duration) *CAcceptor {
	a := &CAcceptor{
		transactionTimeout: transactionTimeout,
		downloadTimeout:    downloadTimeout,
		eventHandlers:      make(map[enum.EventType]EventHandler, enum.Event_End),

		messageQueue:        make(chan *CMessage, 1),
		replyQueue:          make(chan []byte, 1),
		flashDownloadThread: make(chan bool, 1),
		openThread:          make(chan bool, 1),

		log: NewLog(
			log.Output(zerolog.ConsoleWriter{
				Out:        os.Stdout,
				NoColor:    false,
				TimeFormat: "15:04:05",
			}),
			"Acceptor",
		),
	}

	return a
}

//

func (a *CAcceptor) getTickCount() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (a *CAcceptor) SetEventHandler(event enum.EventType, eventHandler func(*CAcceptor, int)) {
	a.eventHandlers[event] = eventHandler
}
