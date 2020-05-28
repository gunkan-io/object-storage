// Copyright (C) 2019-2020 OpenIO SAS
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package helpers_grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

func DialTLS(addrConnect, dirConfig string) (*grpc.ClientConn, error) {
	var caBytes, certBytes []byte
	var err error

	if caBytes, err = ioutil.ReadFile(dirConfig + "/ca.cert"); err != nil {
		return nil, err
	}
	if certBytes, err = ioutil.ReadFile(dirConfig + "/service.pem"); err != nil {
		return nil, err
	}
	//if keyBytes, err = ioutil.ReadFile(dirConfig + "/service.key"); err != nil {
	//	return nil, err
	//}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(certBytes) {
		return nil, errors.New("Invalid certificate (service)")
	}
	if !certPool.AppendCertsFromPEM(caBytes) {
		return nil, errors.New("Invalid certificate (authority)")
	}

	creds := credentials.NewClientTLSFromCert(certPool, "")

	return grpc.Dial(addrConnect,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor))
}

func DialTLSInsecure(addrConnect string) (*grpc.ClientConn, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	creds := credentials.NewTLS(config)
	return grpc.Dial(addrConnect,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor))
}

func ServerTLS(dirConfig string) (*grpc.Server, error) {
	var certBytes, keyBytes []byte
	var err error

	if certBytes, err = ioutil.ReadFile(dirConfig + "/service.pem"); err != nil {
		return nil, err
	}
	if keyBytes, err = ioutil.ReadFile(dirConfig + "/service.key"); err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(certBytes)
	if !ok {
		return nil, errors.New("Invalid certificates")
	}

	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return nil, err
	}

	creds := credentials.NewServerTLSFromCert(&cert)
	srv := grpc.NewServer(
		grpc.Creds(creds),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor))
	return srv, nil
}
