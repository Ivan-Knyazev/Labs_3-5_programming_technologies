package main

import (
	"design-patterns/structural/internal/proxy"
	"fmt"
)

func printResponse(url string, status int, body string) {
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", url, status, body)
}

func main() {

	nginxServer := proxy.NewNginxServer(2)
	appStatusURL := "/app/user/status"
	createUserURL := "/app/user/create"

	httpCode, body := nginxServer.HandleRequest(appStatusURL, "GET")
	printResponse(appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	printResponse(appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	printResponse(appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, "POST")
	printResponse(createUserURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, "GET")
	printResponse(createUserURL, httpCode, body)
}
