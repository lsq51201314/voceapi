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

type token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (g *Group) Token() string {
	return g.token
}

func (g *Group) refresh() (err error) {
	//构建数据
	obj := token{
		Token:        g.token,
		RefreshToken: g.refresh_token,
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
	if req, err = http.NewRequest("POST", g.url+"/api/token/renew", bytes.NewBuffer(buf)); err != nil {
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
		if resp.StatusCode == 401 {
			err = errors.New("无效的凭证")
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
