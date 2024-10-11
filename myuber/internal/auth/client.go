package auth

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

// https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph
func ClientLoadTLSCredentials(clientType string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(fmt.Sprintf("cert/%s-cert.pem", clientType), fmt.Sprintf("cert/%s-key.pem", clientType))
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
		ServerName:   "localhost", // to not do this , change localhost line in gen.sh had to add SANS option for Go- DONE
		// InsecureSkipVerify: true,
	}

	return credentials.NewTLS(config), nil
}

func RiderLoadTLSCredentials() (credentials.TransportCredentials, error) {
	return ClientLoadTLSCredentials("rider")
}

func DriverLoadTLSCredentials() (credentials.TransportCredentials, error) {
	return ClientLoadTLSCredentials("driver")
}
