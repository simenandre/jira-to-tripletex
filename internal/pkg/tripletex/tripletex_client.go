package tripletex

import (
	apiclient "github.com/bjerkio/tripletex-go/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New returns a authenticated Tripletex client
func New() (*apiclient.Tripletex, runtime.ClientAuthInfoWriter, error) {

	token, err := GetToken()
	if err != nil {
		return nil, nil, err
	}

	r := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	r.DefaultAuthentication = httptransport.BasicAuth("0", token)

	// Fix "application/json; charset=utf-8" issue
	r.Producers["application/json; charset=utf-8"] = runtime.JSONProducer()

	return apiclient.New(r, strfmt.Default), r.DefaultAuthentication, nil
}
