package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)

	// Create the TLS config to require client certificates.
	tlsConfig := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	// Listen and serve HTTPS on localhost:8443 with the TLS config.
	server := &http.Server{
		Addr:      "localhost:8443",
		TLSConfig: tlsConfig,
	}
	if err := server.ListenAndServeTLS("", ""); err != nil {
		panic(err)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT requests are accepted", http.StatusMethodNotAllowed)
		return
	}

	// Get the client certificate from the request.
	clientCert := r.TLS.PeerCertificates[0]

	// Calculate the number of days until the certificate expires.
	daysUntilExpiry := int(clientCert.NotAfter.Sub(time.Now()).Hours() / 24)

	// Write the response as JSON.
	response, err := json.Marshal(map[string]interface{}{
		"common_name":      clientCert.Subject.CommonName,
		"days_until_expiry": daysUntilExpiry,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(response))
}

