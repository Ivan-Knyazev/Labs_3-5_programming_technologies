package proxy

// Server interface
type Server interface {
	HandleRequest(string, string) (int, string)
}

// Implementation of the Real Server - Application
type Application struct {
}

func (a *Application) HandleRequest(url, method string) (int, string) {
	if url == "/app/user/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/app/user/create" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Found"
}

// Implementation of the Proxy Server - Nginx
type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func (n *Nginx) checkRateLimiting(url string) bool {
	n.rateLimiter[url] += 1
	return !(n.rateLimiter[url] > n.maxAllowedRequest)
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.HandleRequest(url, method)
}

// Factory method for create Nginx Server
func NewNginxServer(maxAllowedRequest int) *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: maxAllowedRequest,
		rateLimiter:       make(map[string]int),
	}
}
