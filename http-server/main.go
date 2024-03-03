// Taken from the official example at https://github.com/corazawaf/coraza/blob/main/examples/http-server/main.go
// Modified by aless3 for the container build

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/types"
)

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Transaction not disrupted\n"))
}

func main() {
	waf := createWAF()

	http.Handle("/", WrapHandler(waf, http.HandlerFunc(defaultHandler)))

	// Mostly for local debugging, the port is set via the PORT environment variable of the container
	port := "80"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Println("Server is running. Listening port: " + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func createWAF() coraza.WAF {
	directivesFile := "custom/*.conf"

	// For unit testing
	if s := os.Getenv("DIRECTIVES_FILE"); s != "" {
		directivesFile = s
	}

	waf, err := coraza.NewWAF(
		coraza.NewWAFConfig().
			WithErrorCallback(logError).
			WithDirectivesFromFile("./coraza.conf").
			WithDirectivesFromFile("./coreruleset/crs-setup.conf.example").
			WithDirectivesFromFile("./coreruleset/rules/*.conf").
			WithDirectivesFromFile(directivesFile),
	)

	if err != nil {
		log.Fatal(err)
	}
	return waf
}

func logError(error types.MatchedRule) {
	msg := error.ErrorLog()
	fmt.Printf("[logError][%s] %s\n", error.Rule().Severity(), msg)
}
