// Code generated by go-swagger; DO NOT EDIT.

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/pupimvictor/do-echo-svr/models"
)

// EchoOKCode is the HTTP code returned for type EchoOK
const EchoOKCode int = 200

/*EchoOK OK

swagger:response echoOK
*/
type EchoOK struct {

	/*
	  In: Body
	*/
	Payload *models.Echo `json:"body,omitempty"`
}

// NewEchoOK creates EchoOK with default headers values
func NewEchoOK() *EchoOK {

	return &EchoOK{}
}

// WithPayload adds the payload to the echo o k response
func (o *EchoOK) WithPayload(payload *models.Echo) *EchoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the echo o k response
func (o *EchoOK) SetPayload(payload *models.Echo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EchoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*EchoDefault error

swagger:response echoDefault
*/
type EchoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEchoDefault creates EchoDefault with default headers values
func NewEchoDefault(code int) *EchoDefault {
	if code <= 0 {
		code = 500
	}

	return &EchoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the echo default response
func (o *EchoDefault) WithStatusCode(code int) *EchoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the echo default response
func (o *EchoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the echo default response
func (o *EchoDefault) WithPayload(payload *models.Error) *EchoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the echo default response
func (o *EchoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EchoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
