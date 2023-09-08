package google

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"checkVersionGolang/internal/domain"
)

type Adapter struct{}

func NewAdapter() (domain.GoogleService, error) {
	return &Adapter{}, nil
}

func (*Adapter) GetGolangVersion(ctx context.Context) (*domain.OnlineVersion, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://golang.org/VERSION?m=text", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(data), "\n")
	split[1] = strings.Split(split[1], " ")[1]
	timeParse, err := time.Parse("2006-01-02T15:04:05Z", split[1])
	if err != nil {
		return nil, err
	}
	return &domain.OnlineVersion{
		Version:   split[0],
		TimeStamp: timeParse,
	}, nil
}
