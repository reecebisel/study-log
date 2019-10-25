package main

import (
        "encoding/json"
        "github.com/gorilla/mux"
        "net/http"
)

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]bool{"ok": true})
  })

  server := &http.Server{
    Handler: router,
    Addr:    "127.0.0.1:8000",
  }

  server.ListenAndServe()
}
