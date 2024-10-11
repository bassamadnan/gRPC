#!/bin/bash

# Remove any existing .pem files
rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/CN=localhost" -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/CN=localhost" -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile <(printf "subjectAltName=DNS:localhost,IP:127.0.0.1")

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

# 4. Generate rider's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout rider-key.pem -out rider-req.pem -subj "/CN=localhost/O=rider" -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"

# 5. Use CA's private key to sign rider's CSR and get back the signed certificate
openssl x509 -req -in rider-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out rider-cert.pem -extfile <(printf "subjectAltName=DNS:localhost,IP:127.0.0.1")

echo "Rider's signed certificate"
openssl x509 -in rider-cert.pem -noout -text

# 6. Generate driver's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout driver-key.pem -out driver-req.pem -subj "/CN=localhost/O=driver" -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"

# 7. Use CA's private key to sign driver's CSR and get back the signed certificate
openssl x509 -req -in driver-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out driver-cert.pem -extfile <(printf "subjectAltName=DNS:localhost,IP:127.0.0.1")

echo "Driver's signed certificate"
openssl x509 -in driver-cert.pem -noout -text
