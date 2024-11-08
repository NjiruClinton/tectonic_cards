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

var (
	USER_ID   = goDotEnvVariable("USER_ID")
	PASSWORD  = goDotEnvVariable("PASSWORD")
	BASIC_URL = "https://sandbox.api.visa.com/vctc"
)

func main() {
	AUTH := "Basic " + base64.StdEncoding.EncodeToString([]byte(USER_ID+":"+PASSWORD))
	cert, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		log.Fatalf("Error loading client certificate: %v", err)
	}
	caCert, err := os.ReadFile("./cacert.pem")
	if err != nil {
		log.Fatalf("Error loading CA certificate: %v", err)
	}
	caCertPool, _ := x509.SystemCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	url := BASIC_URL + "/programadmin/v1/sponsors/configuration"
	method := "GET"
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Add("Authorization", AUTH)

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
