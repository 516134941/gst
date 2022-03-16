package gst

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// SetHeader 设置头信息
func SetHeader(r *http.Request, headers map[string]string) *http.Request {
	for k, v := range headers {
		r.Header.Set(k, v)
	}
	return r
}

// PostJSONRequest 发送JSON POST请求
// jsonBytes JSON报文数据
// headers 数据头信息
// 返回 response body
func PostJSONRequest(url string, jsonBytes []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Printf("PostJSONRequest %v\n", err)
		return nil, err
	}
	req = SetHeader(req, headers)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("PostJSONRequest %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("PostJSONRequest %v\n", err)
		return nil, err
	}
	return res, nil
}

// GetRequest 发送GET请求
// headers 数据头信息
// 返回 response body
func GetRequest(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("GetRequest %v\n", err)
		return nil, err
	}
	req = SetHeader(req, headers)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("GetRequest %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("GetRequest %v\n", err)
		return nil, err
	}
	return res, nil
}
