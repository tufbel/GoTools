// Package go_requests
// Title       : go_requests.go
// Author      : Tuffy  2023/3/2 10:45
// Description :
package go_requests

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"sync"
	"sync/atomic"
	"time"
)

func OnceRCUGetState(url string, count int, failureNum *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	requestResult := "未知"
	var startTime time.Time
	var endTime time.Time
	defer func() {
		duration := endTime.Sub(startTime)
		fmt.Printf(
			"[%s] %04d - %s - %s : %4dms -%d\n",
			time.Now().Format(time.RFC3339), count, url,
			requestResult, duration.Milliseconds(), *failureNum,
		)
	}()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")
	req.SetTimeout(1 * time.Second)
	startTime = time.Now()
	if err_ := fasthttp.Do(req, resp); err_ != nil || resp.StatusCode() != 200 {
		endTime = time.Now()
		atomic.AddInt64(failureNum, 1)
		fmt.Printf(
			"[%s] %04d - %s - Error:%d - %d:%s\n",
			time.Now().Format(time.RFC3339),
			count, url, *failureNum, resp.StatusCode(), err_.Error(),
		)
		requestResult = "失败"
		return
	}
	endTime = time.Now()
	requestResult = "成功"
}

func GetRCUStateTask() {
	var wg sync.WaitGroup
	var failureNum111 int64 = 0
	var failureNum60 int64 = 0
	url111 := "url"
	url60 := "url"
	for i := 0; i < 10000; i++ {
		wg.Add(2)
		go OnceRCUGetState(url111, i+1, &failureNum111, &wg)
		go OnceRCUGetState(url60, i+1, &failureNum60, &wg)
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
