package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DingTalk struct {
	AccessToken string
	Secret      string
}

func NewDingTalk(accessToken string) *DingTalk {
	return &DingTalk{AccessToken: accessToken}
}

type Response struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int64  `json:"errcode"`
}

func (d *DingTalk) Send(message Message) (Response, error) {
	var res Response
	url, err := GetDingTalkURL(d.AccessToken, d.Secret)
	if err != nil {
		return res, err
	}

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		return res, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	defer response.Body.Close()

	if err != nil {
		return res, err
	}

	resultByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resultByte, &res)
	if err != nil {
		return res, fmt.Errorf("unmarshal http response body from json error = %v", err)
	}
	return res, nil
}

var dingtalkUrl = url.URL{
	Scheme: "https",
	Host:   "oapi.dingtalk.com",
	Path:   "robot/send",
}

func GetDingTalkURL(accessToken string, secret string) (string, error) {
	dtu := dingtalkUrl

	param := url.Values{}
	param.Set("access_token", accessToken)

	if secret == "" {
		dtu.RawQuery = param.Encode()
		return dtu.String(), nil
	}

	//dtu := dingTalkURL
	//value := url.Values{}
	//value.Set("access_token", accessToken)
	//
	//if secret == "" {
	//	dtu.RawQuery = value.Encode()
	//	return dtu.String(), nil
	//}
	//
	//sign, err := sign(timestamp, secret)
	//if err != nil {
	//	dtu.RawQuery = value.Encode()
	//	return dtu.String(), err
	//}
	//
	//value.Set("timestamp", timestamp)
	//value.Set("sign", sign)
	//dtu.RawQuery = value.Encode()
	return dtu.String(), nil
}
