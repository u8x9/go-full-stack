package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	targetUrl   = "https://httpbin.org/get"
	httpProxy   = "http://<HOST>:<PORT>"
	socks5Proxy = "socks5://<HOST>:<PORT>"
)

// normal 普通方式
func normal() {
	res, err := http.Get(targetUrl)
	if err != nil {
		fmt.Println("get failed:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("status not ok:", res.StatusCode)
		return
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read failed:", err)
		return
	}
	fmt.Println(string(b))
}

// proxy 代理方式
func proxy() {
	proxyUrl, err := url.Parse(httpProxy)
	if err != nil {
		fmt.Println("parse proxy url failed:", err)
	}
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false, // 强制安全检查
		},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5,
	}

	res, err := client.Get(targetUrl)
	if err != nil {
		fmt.Println("get failed:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("status not ok")
		return
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read response body failed:", err)
		return
	}
	fmt.Println(string(buf))
}

// proxyWithHeader 使用代理并设置header
func proxyWithHeader() {
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		fmt.Println("new request failed:", err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36")
	req.Header.Set("Referer", "https://github.com/u8x9")

	//proxyUrl, err := url.Parse(httpProxy) // http代理
	proxyUrl, err := url.Parse(socks5Proxy) // socks5代理
	if err != nil {
		fmt.Println("parse proxy url failed:", err)
		return
	}
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("get url failed:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("status not ok")
		return
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read response failed:", err)
		return
	}
	fmt.Println(string(buf))
}
func main() {
	//normal()
	//proxy()
	proxyWithHeader()
}
