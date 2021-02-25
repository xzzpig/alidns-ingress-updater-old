package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type PublicIP struct {
	Address  string `json:"address"`
	Code     int64  `json:"code"`
	IP       string `json:"ip"`
	IsDomain int64  `json:"isDomain"`
	Rs       int64  `json:"rs"`
}

func GetPublicIP() (string, error) {
	resp, err := http.Get("https://ip.cn/api/index?ip=&type=0")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("Error Status" + resp.Status)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var publicIp PublicIP
	json.Unmarshal(body, &publicIp)
	return publicIp.IP, nil
}
