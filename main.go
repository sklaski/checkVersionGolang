package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("")
	fmt.Println(strings.Repeat("-", 75))
	releasedVersion := getReleaseVersion()
	localVersion := getLocalVersion()
	checkVersions(releasedVersion, localVersion)
	fmt.Println(strings.Repeat("-", 75))
}

func checkVersions(releasedVersion, localVersion string) {
	name := "GoLang"
	if releasedVersion != localVersion {
		fmt.Printf(
			"%10s: installed version %s mismatch latest %s\nwget %s\n",
			name,
			localVersion,
			releasedVersion,
			"https://dl.google.com/go/$(curl -L 'https://golang.org/VERSION?m=text').linux-amd64.tar.gz",
		)
	} else {
		fmt.Printf("%10s: latest releasedVersion %s installed\n", name, localVersion)
	}
}

func getLocalVersion() string {
	dat, err := os.ReadFile("/usr/local/Go/golang/VERSION")
	if err != nil {
		os.Exit(1)
	}
	return string(dat)
}

func getReleaseVersion() string {
	res, err := http.Get("https://golang.org/VERSION?m=text") //nolint:noctx
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	releaseData := string(data)
	return releaseData
}
