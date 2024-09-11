package hook

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

type RaiseDownloadObj struct {
	r *RaiseObj
}

////////

func (r RaiseDownloadObj) Finish(st bool) {
	v := 0
	if st {
		v = 1
	}
	r.r.run(enum.EventDownloadFinish, v)
	DownloadFinish = false
}

func (r RaiseDownloadObj) Start(v int) {
	r.r.run(enum.EventDownloadStart, v)

	DownloadStart = false
	DownloadProgress = true
}

func (r RaiseDownloadObj) Progress(v int) {
	r.r.run(enum.EventDownloadProgress, v)
}

func (r RaiseDownloadObj) Restart() {
	if !DownloadRestart {
		return
	}

	r.r.run(enum.EventDownloadRestart, 0)
	DownloadRestart = false
}
