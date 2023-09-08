package domain

import (
	"time"
)

type (
	LocalVersion struct {
		Version string
		CPUType string
	}
	OnlineVersion struct {
		Version   string
		TimeStamp time.Time
	}
)
