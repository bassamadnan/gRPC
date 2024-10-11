package auth

import (
	"context"
	"crypto/x509"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func AuthorizationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "No peer found")
	}

	tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unable to extract TLS info")
	}

	if len(tlsInfo.State.VerifiedChains) == 0 || len(tlsInfo.State.VerifiedChains[0]) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Could not verify peer certificate")
	}

	cert := tlsInfo.State.VerifiedChains[0][0]
	role, err := extractRole(cert)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Unable to extract role: %v", err)
	}

	if err := authorizeRoleForMethod(role, info.FullMethod); err != nil {
		return nil, err
	}

	// Additional validation
	if err := validateCertificate(cert); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Certificate validation failed: %v", err)
	}

	return handler(ctx, req)
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	peer, _ := peer.FromContext(ctx)
	tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo)
	role := "unknown"
	var expirationDate time.Time
	if ok && len(tlsInfo.State.VerifiedChains) > 0 && len(tlsInfo.State.VerifiedChains[0]) > 0 {
		cert := tlsInfo.State.VerifiedChains[0][0]
		role, _ = extractRole(cert)
		expirationDate = cert.NotAfter
	}

	log.Printf("Method: %s, Role: %s, Start Time: %s, Certificate Expiration: %s",
		info.FullMethod, role, start.Format(time.RFC3339), expirationDate.Format(time.RFC3339))

	resp, err := handler(ctx, req)

	duration := time.Since(start)
	log.Printf("Method: %s, Role: %s, Duration: %v, Error: %v",
		info.FullMethod, role, duration, err)

	return resp, err
}

func extractRole(cert *x509.Certificate) (string, error) {
	if len(cert.Subject.Organization) > 0 {
		return cert.Subject.Organization[0], nil
	}
	return "", fmt.Errorf("role not found in certificate")
}

func authorizeRoleForMethod(role, method string) error {
	switch method {
	case "/myuber.UberService/RequestRide":
		if role != "rider" {
			return status.Error(codes.PermissionDenied, "Only riders can request rides")
		}
	case "/myuber.UberService/AcceptRide", "/myuber.UberService/CompleteRide", "/myuber.UberService/RejectRide":
		if role != "driver" {
			return status.Error(codes.PermissionDenied, "Only drivers can accept or complete rides")
		}
	}
	return nil
}

func validateCertificate(cert *x509.Certificate) error {
	// checks expiry
	now := time.Now()
	if now.Before(cert.NotBefore) {
		return fmt.Errorf("certificate is not yet valid")
	}
	if now.After(cert.NotAfter) {
		return fmt.Errorf("certificate has expired")
	}
	return nil
}

func ChainedInterceptor(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		chain := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			chain = buildChain(interceptors[i], chain, info)
		}
		return chain(ctx, req)
	}
}

func buildChain(current grpc.UnaryServerInterceptor, next grpc.UnaryHandler, info *grpc.UnaryServerInfo) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return current(ctx, req, info, next)
	}
}
