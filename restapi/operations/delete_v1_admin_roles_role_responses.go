// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteV1AdminRolesRoleOKCode is the HTTP code returned for type DeleteV1AdminRolesRoleOK
const DeleteV1AdminRolesRoleOKCode int = 200

/*DeleteV1AdminRolesRoleOK Deleted

swagger:response deleteV1AdminRolesRoleOK
*/
type DeleteV1AdminRolesRoleOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteV1AdminRolesRoleOK creates DeleteV1AdminRolesRoleOK with default headers values
func NewDeleteV1AdminRolesRoleOK() *DeleteV1AdminRolesRoleOK {

	return &DeleteV1AdminRolesRoleOK{}
}

// WithPayload adds the payload to the delete v1 admin roles role o k response
func (o *DeleteV1AdminRolesRoleOK) WithPayload(payload string) *DeleteV1AdminRolesRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v1 admin roles role o k response
func (o *DeleteV1AdminRolesRoleOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteV1AdminRolesRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteV1AdminRolesRoleUnauthorizedCode is the HTTP code returned for type DeleteV1AdminRolesRoleUnauthorized
const DeleteV1AdminRolesRoleUnauthorizedCode int = 401

/*DeleteV1AdminRolesRoleUnauthorized Unauthorized

swagger:response deleteV1AdminRolesRoleUnauthorized
*/
type DeleteV1AdminRolesRoleUnauthorized struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteV1AdminRolesRoleUnauthorized creates DeleteV1AdminRolesRoleUnauthorized with default headers values
func NewDeleteV1AdminRolesRoleUnauthorized() *DeleteV1AdminRolesRoleUnauthorized {

	return &DeleteV1AdminRolesRoleUnauthorized{}
}

// WithPayload adds the payload to the delete v1 admin roles role unauthorized response
func (o *DeleteV1AdminRolesRoleUnauthorized) WithPayload(payload string) *DeleteV1AdminRolesRoleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v1 admin roles role unauthorized response
func (o *DeleteV1AdminRolesRoleUnauthorized) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteV1AdminRolesRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteV1AdminRolesRoleForbiddenCode is the HTTP code returned for type DeleteV1AdminRolesRoleForbidden
const DeleteV1AdminRolesRoleForbiddenCode int = 403

/*DeleteV1AdminRolesRoleForbidden forbidden

swagger:response deleteV1AdminRolesRoleForbidden
*/
type DeleteV1AdminRolesRoleForbidden struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteV1AdminRolesRoleForbidden creates DeleteV1AdminRolesRoleForbidden with default headers values
func NewDeleteV1AdminRolesRoleForbidden() *DeleteV1AdminRolesRoleForbidden {

	return &DeleteV1AdminRolesRoleForbidden{}
}

// WithPayload adds the payload to the delete v1 admin roles role forbidden response
func (o *DeleteV1AdminRolesRoleForbidden) WithPayload(payload string) *DeleteV1AdminRolesRoleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v1 admin roles role forbidden response
func (o *DeleteV1AdminRolesRoleForbidden) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteV1AdminRolesRoleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteV1AdminRolesRoleNotFoundCode is the HTTP code returned for type DeleteV1AdminRolesRoleNotFound
const DeleteV1AdminRolesRoleNotFoundCode int = 404

/*DeleteV1AdminRolesRoleNotFound The specified resource was not found

swagger:response deleteV1AdminRolesRoleNotFound
*/
type DeleteV1AdminRolesRoleNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteV1AdminRolesRoleNotFound creates DeleteV1AdminRolesRoleNotFound with default headers values
func NewDeleteV1AdminRolesRoleNotFound() *DeleteV1AdminRolesRoleNotFound {

	return &DeleteV1AdminRolesRoleNotFound{}
}

// WithPayload adds the payload to the delete v1 admin roles role not found response
func (o *DeleteV1AdminRolesRoleNotFound) WithPayload(payload string) *DeleteV1AdminRolesRoleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v1 admin roles role not found response
func (o *DeleteV1AdminRolesRoleNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteV1AdminRolesRoleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteV1AdminRolesRoleConflictCode is the HTTP code returned for type DeleteV1AdminRolesRoleConflict
const DeleteV1AdminRolesRoleConflictCode int = 409

/*DeleteV1AdminRolesRoleConflict conflict

swagger:response deleteV1AdminRolesRoleConflict
*/
type DeleteV1AdminRolesRoleConflict struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteV1AdminRolesRoleConflict creates DeleteV1AdminRolesRoleConflict with default headers values
func NewDeleteV1AdminRolesRoleConflict() *DeleteV1AdminRolesRoleConflict {

	return &DeleteV1AdminRolesRoleConflict{}
}

// WithPayload adds the payload to the delete v1 admin roles role conflict response
func (o *DeleteV1AdminRolesRoleConflict) WithPayload(payload string) *DeleteV1AdminRolesRoleConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v1 admin roles role conflict response
func (o *DeleteV1AdminRolesRoleConflict) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteV1AdminRolesRoleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteV1AdminRolesRoleInternalServerErrorCode is the HTTP code returned for type DeleteV1AdminRolesRoleInternalServerError
const DeleteV1AdminRolesRoleInternalServerErrorCode int = 500

/*DeleteV1AdminRolesRoleInternalServerError For unknown or unanticipated issues

swagger:response deleteV1AdminRolesRoleInternalServerError
*/
type DeleteV1AdminRolesRoleInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteV1AdminRolesRoleInternalServerError creates DeleteV1AdminRolesRoleInternalServerError with default headers values
func NewDeleteV1AdminRolesRoleInternalServerError() *DeleteV1AdminRolesRoleInternalServerError {

	return &DeleteV1AdminRolesRoleInternalServerError{}
}

// WithPayload adds the payload to the delete v1 admin roles role internal server error response
func (o *DeleteV1AdminRolesRoleInternalServerError) WithPayload(payload string) *DeleteV1AdminRolesRoleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete v1 admin roles role internal server error response
func (o *DeleteV1AdminRolesRoleInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteV1AdminRolesRoleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}