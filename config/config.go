package config

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

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func LoadEnvVariables() (string, string) {
	USER_ID := GoDotEnvVariable("USER_ID")
	PASSWORD := GoDotEnvVariable("PASSWORD")
	return USER_ID, PASSWORD
}

func SetupTLSConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		return nil, err
	}
	caCert, err := os.ReadFile("./cacert.pem")
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

func GetAuthHeader() string {
	userID, password := LoadEnvVariables()
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(userID+":"+password))
}

func MakeHTTPRequest(method string, url string, payload io.Reader) (string, error) {
	AUTH := GetAuthHeader()
	finalURL := "https://sandbox.api.visa.com/vctc" + url
	req, err := http.NewRequest(method, finalURL, payload)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", AUTH)
	if payload != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	tlsConfig, err := SetupTLSConfig()
	if err != nil {
		log.Fatalf("Error setting up TLS configuration: %v", err)
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
