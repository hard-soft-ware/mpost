package mpost

////////////////////////////////////

type MethodsObj struct {
	a *MpostObj
}

////

func (a *MpostObj) newMethods() *MethodsObj {
	obj := MethodsObj{a: a}

	return &obj
}

////////////////
