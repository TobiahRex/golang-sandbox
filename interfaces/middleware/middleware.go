package middleware_iface

import "fmt"

func Main() {
	auth := AuthMiddleware{
		Next: LoggingMiddleware{
			Next: FinalHandler{},
		},
	}
	response := auth.Handle("request-value")
	fmt.Println(response)
}

type Handler interface {
	Handle(string) string
}

type AuthMiddleware struct {
	Next Handler
}

func (a AuthMiddleware) Handle(request string) string {
	// TODO: Implement auth logic
	return a.Next.Handle(request)
}

type LoggingMiddleware struct {
	Next Handler
}

func (l LoggingMiddleware) Handle(request string) string {
	// TODO: Implement logging logic
	return l.Next.Handle(request)
}

type FinalHandler struct {
	Next Handler
}

func (f FinalHandler) Handle(request string) string {
	return request + " handled"
}
