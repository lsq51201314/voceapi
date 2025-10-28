package group

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type prepare struct {
	ContentType string `json:"content_type"`
	FileName    string `json:"filename"`
}

func (g *Group) prepare(content_type, filename string) (fid string, err error) {
	//构建数据
	obj := prepare{
		ContentType: content_type,
		FileName:    filepath.Base(filename),
	}
	var buf []byte
	if buf, err = json.Marshal(&obj); err != nil {
		return
	}
prepare:
	//提交数据
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var req *http.Request
	if req, err = http.NewRequest("POST", g.url+"/api/resource/file/prepare", bytes.NewBuffer(buf)); err != nil {
		return
	}
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("X-API-Key", g.token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
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
			goto prepare
		}
		err = fmt.Errorf("未知状态码:%d", resp.StatusCode)
		return
	}
	//处理数据
	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}
	fid = strings.ReplaceAll(string(body), `"`, "") //必须
	return
}
