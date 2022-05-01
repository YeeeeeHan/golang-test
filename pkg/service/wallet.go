package service

import "TechnicalAssignment/pkg/constants"

type Wallet struct {
	Username string
}

func (w *Wallet) GetUsername() string {
	if w.Username == "" {
		return constants.Unregistered
	}
	return w.Username
}
