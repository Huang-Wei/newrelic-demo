package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

func mustGetEnv(key string) string {
	if val := os.Getenv(key); "" != val {
		return val
	}
	panic(fmt.Sprintf("environment variable %s unset", key))
}

type popularity struct {
	sync.RWMutex
	value int
}

type helloHandler struct {
	app        newrelic.Application
	popularity *popularity
}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("[DEBUG] %p\n", h.popularity)
	path := r.URL.Path[1:]

	if path == "like" {
		h.popularity.Lock()
		h.popularity.value += 1
		h.popularity.Unlock()
	} else if path == "dislike" {
		h.popularity.Lock()
		h.popularity.value -= 1
		h.popularity.Unlock()
	}

	fmt.Fprintf(w, "Hello, current popularity is %v.\n", h.popularity.value)
}

func main() {
	// cfg := newrelic.NewConfig("rewrelic-demo", mustGetEnv("NEW_RELIC_LICENSE_KEY"))
	cfg := newrelic.NewConfig("rewrelic-demo", "")
	cfg.Enabled = false // for debugging
	cfg.Logger = newrelic.NewDebugLogger(os.Stdout)
	app, err := newrelic.NewApplication(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the application to connect.
	if err := app.WaitForConnection(15 * time.Second); err != nil {
		fmt.Println(err)
	}

	fmt.Println("demo http server started...")
	err = http.ListenAndServe(":8080", helloHandler{
		app:        app,
		popularity: &popularity{value: 0},
	})
	log.Fatal(err)
}
