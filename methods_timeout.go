package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"time"
)

////////////////////////////////////

func (m *MethodsObj) GetTransactionTimeout() time.Duration {
	m.a.Log.Method("GetTransactionTimeout", nil)
	return acceptor.Timeout.Transaction
}

func (m *MethodsObj) GetDownloadTimeout() time.Duration {
	m.a.Log.Method("GetDownloadTimeout", nil)
	return acceptor.Timeout.Download
}

//

func (m *MethodsObj) SetTransactionTimeout(v time.Duration) {
	m.a.Log.Method("SetTransactionTimeout", v)
	acceptor.Timeout.Transaction = v
}

func (m *MethodsObj) SetDownloadTimeout(v time.Duration) {
	m.a.Log.Method("SetDownloadTimeout", v)
	acceptor.Timeout.Download = v
}
