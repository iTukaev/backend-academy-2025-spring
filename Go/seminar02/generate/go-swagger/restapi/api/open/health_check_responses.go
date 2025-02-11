// Code generated by go-swagger; DO NOT EDIT.

package open

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"example.com/seminar02/generate/go-swagger/model"
)

/*
HealthCheckDefault universal response is used for NoContent and AnyError answers, contains field Status as any data answers

swagger:response healthCheckDefault
*/
type HealthCheckDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *model.EmptyResponse `json:"body,omitempty"`
}

// NewHealthCheckDefault creates HealthCheckDefault with default headers values
func NewHealthCheckDefault(code int) *HealthCheckDefault {
	if code <= 0 {
		code = 500
	}

	return &HealthCheckDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the health check default response
func (o *HealthCheckDefault) WithStatusCode(code int) *HealthCheckDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the health check default response
func (o *HealthCheckDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the health check default response
func (o *HealthCheckDefault) WithPayload(payload *model.EmptyResponse) *HealthCheckDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the health check default response
func (o *HealthCheckDefault) SetPayload(payload *model.EmptyResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HealthCheckDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *HealthCheckDefault) HealthCheckResponder() {}

type HealthCheckNotImplementedResponder struct {
	middleware.Responder
}

func (*HealthCheckNotImplementedResponder) HealthCheckResponder() {}

func HealthCheckNotImplemented() HealthCheckResponder {
	return &HealthCheckNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.HealthCheck has not yet been implemented",
		),
	}
}

type HealthCheckResponder interface {
	middleware.Responder
	HealthCheckResponder()
}
