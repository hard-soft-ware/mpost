package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"time"
)

////////////////////////////////////

type MethodsTimeoutObj struct {
	a   *MpostObj
	Get MethodsTimeoutGetObj
	Set MethodsTimeoutSetObj
}

type MethodsTimeoutGetObj struct{ a *MpostObj }
type MethodsTimeoutSetObj struct{ a *MpostObj }

func (m *MethodsObj) newTimeout() *MethodsTimeoutObj {
	obj := MethodsTimeoutObj{}

	obj.a = m.a
	obj.Get.a = m.a
	obj.Set.a = m.a

	return &obj
}

////////////////

func (m *MethodsTimeoutGetObj) Transaction() time.Duration {
	m.a.Log.Method("GetTransactionTimeout", nil)
	return acceptor.Timeout.Transaction
}

func (m *MethodsTimeoutGetObj) Download() time.Duration {
	m.a.Log.Method("GetDownloadTimeout", nil)
	return acceptor.Timeout.Download
}

//

func (m *MethodsTimeoutSetObj) Transaction(v time.Duration) {
	m.a.Log.Method("SetTransactionTimeout", nil)
	acceptor.Timeout.Transaction = v
}

func (m *MethodsTimeoutSetObj) Download(v time.Duration) {
	m.a.Log.Method("SetDownloadTimeout", nil)
	acceptor.Timeout.Download = v
}
