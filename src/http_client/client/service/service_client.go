// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ServiceGreet(params *ServiceGreetParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceGreetOK, error)

	ServiceRepeatGreet(params *ServiceRepeatGreetParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceRepeatGreetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ServiceGreet service greet API
*/
func (a *Client) ServiceGreet(params *ServiceGreetParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceGreetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceGreetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Service_Greet",
		Method:             "POST",
		PathPattern:        "/v1/service/greet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceGreetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceGreetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceGreetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ServiceRepeatGreet repeats greet the given name
*/
func (a *Client) ServiceRepeatGreet(params *ServiceRepeatGreetParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceRepeatGreetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceRepeatGreetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Service_RepeatGreet",
		Method:             "GET",
		PathPattern:        "/v1/service/greet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceRepeatGreetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceRepeatGreetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceRepeatGreetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
