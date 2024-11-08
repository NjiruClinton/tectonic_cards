package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func loadEnvVariables() (string, string) {
	USER_ID := goDotEnvVariable("USER_ID")
	PASSWORD := goDotEnvVariable("PASSWORD")
	return USER_ID, PASSWORD
}

func setupTLSConfig(certFile, keyFile, caCertFile string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	caCert, err := os.ReadFile(caCertFile)
	if err != nil {
		return nil, err
	}
	caCertPool, _ := x509.SystemCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	return tlsConfig, nil
}

func makeHTTPRequest(client *http.Client, url, auth string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Add("Authorization", auth)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(string(body))
}

func main() {
	USER_ID, PASSWORD := loadEnvVariables()
	AUTH := "Basic " + base64.StdEncoding.EncodeToString([]byte(USER_ID+":"+PASSWORD))
	BASIC_URL := "https://sandbox.api.visa.com/vctc"

	tlsConfig, err := setupTLSConfig("./cert.pem", "./key.pem", "./cacert.pem")
	if err != nil {
		log.Fatalf("Error setting up TLS configuration: %v", err)
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	url := BASIC_URL + "/programadmin/v1/sponsors/configuration"
	makeHTTPRequest(client, url, AUTH)
}
