package mpost

import "github.com/hard-soft-ware/mpost/acceptor"

////////////////////////////////////

func (m *MethodsObj) GetEnableAcceptance() bool {
	m.a.Log.Method("GetEnableAcceptance", nil)
	return acceptor.Enable.Acceptance
}

func (m *MethodsObj) GetEnableBookmarks() bool {
	m.a.Log.Method("GetEnableBookmarks", nil)
	return acceptor.Enable.Bookmarks
}

func (m *MethodsObj) GetEnableNoPush() bool {
	m.a.Log.Method("GetEnableNoPush", nil)
	return acceptor.Enable.NoPush
}

//

func (m *MethodsObj) SetEnableAcceptance(v bool) {
	m.a.Log.Method("SetEnableAcceptance", v)
	acceptor.Enable.Acceptance = v
}

func (m *MethodsObj) SetEnableBookmarks(v bool) {
	m.a.Log.Method("SetEnableBookmarks", v)
	acceptor.Enable.Bookmarks = v
}

func (m *MethodsObj) SetEnableNoPush(v bool) {
	m.a.Log.Method("SetEnableNoPush", v)
	acceptor.Enable.NoPush = v
}
