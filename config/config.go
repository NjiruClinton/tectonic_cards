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

func SetupTLSConfig(certFile, keyFile, caCertFile string) (*tls.Config, error) {
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

func GetAuthHeader(userID, password string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(userID+":"+password))
}

func MakeHTTPRequest(client *http.Client, url, auth string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", auth)

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
