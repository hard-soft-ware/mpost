package enum

/* This file is automatically generated */

type BezelType byte

const (
	BezelStandard   BezelType = 0
	BezelPlatform   BezelType = 1
	BezelDiagnostic BezelType = 2
)

const (
	BezelTextStandard   = "Standard"
	BezelTextPlatform   = "Platform"
	BezelTextDiagnostic = "Diagnostic"
)

var BezelMap = map[BezelType]string{
	BezelStandard:   BezelTextStandard,
	BezelPlatform:   BezelTextPlatform,
	BezelDiagnostic: BezelTextDiagnostic,
}

func (obj BezelType) String() string {
	val, ok := BezelMap[obj]
	if ok {
		return val
	}
	return "Unknown BezelType"
}
