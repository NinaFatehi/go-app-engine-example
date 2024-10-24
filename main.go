package main
import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", indexHandler)
    // Check if PORT environment variable is set, fallback to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default to port 8080 if no port is provided
		log.Printf("Defaulting to port %s", port)
    }

    // Register the handler for the root URL ("/")
   

    // Start the server and listen on the specified port
    log.Printf("Listening on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}