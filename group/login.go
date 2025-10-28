package group

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type login struct {
	Credential  credential `json:"credential"`
	Device      string     `json:"device"`
	DeviceToken any        `json:"device_token"`
}

type credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func (g *Group) Login(username, password string) (err error) {
	//构建对象
	obj := login{
		Credential: credential{
			Email:    username,
			Password: password,
			Type:     "password",
		},
		Device:      "web",
		DeviceToken: nil,
	}
	var buf []byte
	if buf, err = json.Marshal(&obj); err != nil {
		return
	}
	//提交数据
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var req *http.Request
	if req, err = http.NewRequest("POST", g.url+"/api/token/login", bytes.NewBuffer(buf)); err != nil {
		return
	}
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 || resp.StatusCode == 404 {
			err = errors.New("账户或密码错误")
			return
		}
		err = fmt.Errorf("未知状态码:%d", resp.StatusCode)
		return
	}
	//处理数据
	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}
	var token token
	if err = json.Unmarshal(body, &token); err != nil {
		return
	}
	g.refresh_token = token.RefreshToken
	g.token = token.Token
	return
}
