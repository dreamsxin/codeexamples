package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func main() {
	log.Println("Run")
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	log.Println("Launch")
	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
		//Headless:       playwright.Bool(true),
		ExecutablePath: playwright.String("D:\\camoufox-130.0.1\\launch.exe"),
		Args: []string{
			"--stderr", "./debug.log",
			"--config", "{}",
		},
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	log.Println("NewPage")
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://www.baidu.com"); err != nil {
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
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
