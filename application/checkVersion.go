package application

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"checkVersionGolang/internal/domain"
	"checkVersionGolang/internal/infrastructure/google"
)

func CompareVersions(ctx context.Context) error {
	onlineVersion, err := GetReleasedVersion(ctx)
	if err != nil {
		return fmt.Errorf("could not get online version: %w", err)
	}
	localVersion, err := getLocalVersion()
	if err != nil {
		return fmt.Errorf("could not get local version: %w", err)
	}

	name := "GoLang"
	if onlineVersion.Version != localVersion.Version {
		fmt.Printf(
			"%10s: installed version %s mismatch latest %s\nwget %s\n",
			name,
			localVersion.Version,
			onlineVersion.Version,
			fmt.Sprintf(
				"https://dl.google.com/go/%s.%s.tar.gz",
				onlineVersion.Version,
				localVersion.CPUType,
			),
		)
	} else {
		fmt.Printf("%10s: latest onlineVersion %s (%s) installed\n", name, localVersion.Version, onlineVersion.TimeStamp)
	}
	return nil
}

func getLocalVersion() (*domain.LocalVersion, error) {
	command, err := exec.Command("go", "version").Output()
	if err != nil {
		return nil, err
	}
	output := strings.Split(strings.TrimSuffix(string(command), "\n"), " ")
	lv := domain.LocalVersion{
		Version: output[2],
		CPUType: strings.ReplaceAll(output[3], "/", "-"),
	}
	return &lv, nil
}

func GetReleasedVersion(ctx context.Context) (*domain.OnlineVersion, error) {
	googleAdapter, err := google.NewAdapter()
	if err != nil {
		return nil, fmt.Errorf("could not initilize adapter: %w", err)
	}
	onlineVersion, err := googleAdapter.GetGolangVersion(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get online version: %w", err)
	}

	return onlineVersion, err
}
