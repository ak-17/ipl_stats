package model

const (
	SRH  = "SRH"
	DC   = "DC"
	KXIP = "KXIP"
	CSK  = "CSK"
	KKR  = "KKR"
	MI   = "MI"
	RCB  = "RCB"
	RR   = "RR"
)

var TeamNames = []string{SRH, DC, KXIP, KKR, MI, CSK, RCB, RR}

type Team struct {
	Name   string  `json:"name"`
	Points float64 `json:"points"`
}
