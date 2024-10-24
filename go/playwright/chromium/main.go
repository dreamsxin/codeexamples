package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/playwright-community/playwright-go"
)

var ruyiconfig = `{"post":"Ruyi Education bilibili: https://space.bilibili.com/172381477","portScan":{"enable":"no"},"canvas":{"noise":2.0},"webrtc":{"public":"154.64.226.99","private":"192.168.9.1"},"position":{"longitude":116.231407,"latitude":39.951754,"altitude":100.0,"accuracy":10.0},"webgl":{"vendor":"Google Inc. (NVIDIA)","renderer":"ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 Laptop GPU (0x0000249C) Direct3D11 vs_5_0 ps_5_0, D3D11)"},"gpu":{"device":"ruyi","description":"ruyi gpu"},"webaudio":100.0,"clientrect":{"x":2.0,"y":2.0},"screen":{"width":1000.0,"height":2000.0,"colorDepth":9.0,"availWidth":900.0,"availHeight":900.0},"mobile":{"touchsupport":3.0},"hardware":{"concurrency":16.0,"memory":64.0},"clientHint":{"platform":"Windows","platform_version":"16.0.0","ua_full_version":"118.0.2592.119","mobile":"?0","architecture":"x64","bitness":"64"},"languages":{"js":"zh-CN,en,en-US","http":"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6"},"software":{"cookie":"no","java":"yes","dnt":"no"},"font":{"removefont":"Segoe UI,Tahoma"}}`

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go func() {
		defer func() {
			waitgroup.Done()
		}()
		//="" --lang=en-US --time-zone-for-testing=America/Los_Angeles --user-data-dir="C:/user2"`
		pw, err := playwright.Run()
		if err != nil {
			log.Fatalf("could not start playwright: %v", err)
		}
		browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
			Headless:       playwright.Bool(false),
			ExecutablePath: playwright.String("C:\\Users\\Administrator\\AppData\\Local\\Chromium\\Application\\chrome.exe"),
			Args: []string{
				//"--ruyi", config,
				"--user-agent=Mozilla/15.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.0.0",
			},
		})
		if err != nil {
			log.Fatalf("could not launch browser: %v", err)
		}
		page, err := browser.NewPage()
		if err != nil {
			log.Fatalf("could not create page: %v", err)
		}
		page.OnRequest(func(request playwright.Request) {
			fmt.Printf(">> %s %s\n", request.Method(), request.URL())
		})
		page.OnResponse(func(response playwright.Response) {
			fmt.Printf("<< %v %s\n", response.Status(), response.URL())
		})
		if _, err = page.Goto("https://mobile.yangkeduo.com"); err != nil {
			log.Fatalf("could not goto: %v", err)
		}
		entries, err := page.Locator(".athing").All()
		if err != nil {
			log.Fatalf("could not get entries: %v", err)
		}
		for i, entry := range entries {
			title, err := entry.Locator("td.title > span > a").TextContent()
			if err != nil {
				log.Fatalf("could not get text content: %v", err)
			}
			fmt.Printf("%d: %s\n", i+1, title)
		}
		time.Sleep(1 * time.Minute)
		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}
	}()
	waitgroup.Wait()
}
