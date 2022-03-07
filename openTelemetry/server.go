package main

import (
	"net/http"

	"github.com/lightstep/otel-launcher-go/launcher"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	ls := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName("openTelemetry"),
		launcher.WithAccessToken("4skdukGqV4gCAHlfFq7Ja6N+T23l8+trwwhKyOo87fK7eppu7nfFfTjgYDzB4dlGQJ0QjEy/vLsNpFWFTE8raYb4sMQrYVcn4AdlERvt"),
	)
	defer ls.Shutdown()
	// wrappedHandler := otelhttp.NewHandler(http.HandlerFunc(helloHandler), "/hello")
	// http.Handle("/hello", wrappedHandler)
	// http.ListenAndServe(":9000", nil)
}
