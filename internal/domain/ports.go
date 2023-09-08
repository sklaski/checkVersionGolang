package domain

import (
	"context"
)

type GoogleService interface {
	GetGolangVersion(ctx context.Context) (*OnlineVersion, error)
}
