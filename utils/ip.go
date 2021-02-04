package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func GetPublicIP() (string, error) {
	resp, err := http.Get("https://icanhazip.com/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return strings.Trim(string(body), "\n"), nil
}
