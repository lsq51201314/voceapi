package group

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cast"
)

func (g *Group) SendText(gid int64, text string) (err error) {
send:
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var req *http.Request
	if req, err = http.NewRequest("POST", g.url+"/api/group/"+cast.ToString(gid)+"/send", bytes.NewBuffer([]byte(text))); err != nil {
		return
	}
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("X-API-Key", g.token)
	req.Header.Set("Content-Type", "text/plain")
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			if err = g.refresh(); err != nil { //刷新凭证
				return
			}
			goto send
		}
		err = fmt.Errorf("未知状态码:%d", resp.StatusCode)
		return
	}
	return
}

type file struct {
	Path string `json:"path"`
}

func (g *Group) SendFile(gid int64, path string) (err error) {
	//构建数据
	obj := file{
		Path: path,
	}
	var buf []byte
	if buf, err = json.Marshal(&obj); err != nil {
		return
	}
send:
	//提交数据
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var req *http.Request
	if req, err = http.NewRequest("POST", g.url+"/api/group/"+cast.ToString(gid)+"/send", bytes.NewBuffer(buf)); err != nil {
		return
	}
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("X-API-Key", g.token)
	req.Header.Set("Content-Type", "vocechat/file")
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			if err = g.refresh(); err != nil { //刷新凭证
				return
			}
			goto send
		}
		err = fmt.Errorf("未知状态码:%d", resp.StatusCode)
		return
	}
	return
}
