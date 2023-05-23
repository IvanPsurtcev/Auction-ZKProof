package contract

import "math/big"

type SetStatus struct {
	Started bool `json:"started"`
	Ended   bool `json:"ended"`
}

type Bid struct {
	OneBid big.Int `json:"bid"`
}
