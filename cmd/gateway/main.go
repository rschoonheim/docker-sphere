package main

import (
	"docker_sphere_gateway/internal/api/user/userconnect"
	"docker_sphere_gateway/internal/api/user/userservice"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log/slog"
	"logging"
	"net/http"
)

func main() {
	logging.Initialize("gateway")
	slog.Info("Gateway Application")

	mux := http.NewServeMux()
	path, handler := userconnect.NewUserServiceHandler(userservice.NewUserServer())
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
