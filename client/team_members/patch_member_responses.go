// Code generated by go-swagger; DO NOT EDIT.

package team_members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/mlafeldt/go-launchdarkly/models"
)

// PatchMemberReader is a Reader for the PatchMember structure.
type PatchMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPatchMemberConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchMemberOK creates a PatchMemberOK with default headers values
func NewPatchMemberOK() *PatchMemberOK {
	return &PatchMemberOK{}
}

/*PatchMemberOK handles this case with default header values.

Member response.
*/
type PatchMemberOK struct {
	Payload *models.Member
}

func (o *PatchMemberOK) Error() string {
	return fmt.Sprintf("[PATCH /members/{memberId}][%d] patchMemberOK  %+v", 200, o.Payload)
}

func (o *PatchMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Member)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMemberBadRequest creates a PatchMemberBadRequest with default headers values
func NewPatchMemberBadRequest() *PatchMemberBadRequest {
	return &PatchMemberBadRequest{}
}

/*PatchMemberBadRequest handles this case with default header values.

Invalid request body.
*/
type PatchMemberBadRequest struct {
}

func (o *PatchMemberBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /members/{memberId}][%d] patchMemberBadRequest ", 400)
}

func (o *PatchMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMemberUnauthorized creates a PatchMemberUnauthorized with default headers values
func NewPatchMemberUnauthorized() *PatchMemberUnauthorized {
	return &PatchMemberUnauthorized{}
}

/*PatchMemberUnauthorized handles this case with default header values.

Invalid access token.
*/
type PatchMemberUnauthorized struct {
}

func (o *PatchMemberUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /members/{memberId}][%d] patchMemberUnauthorized ", 401)
}

func (o *PatchMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMemberNotFound creates a PatchMemberNotFound with default headers values
func NewPatchMemberNotFound() *PatchMemberNotFound {
	return &PatchMemberNotFound{}
}

/*PatchMemberNotFound handles this case with default header values.

Invalid resource specifier.
*/
type PatchMemberNotFound struct {
}

func (o *PatchMemberNotFound) Error() string {
	return fmt.Sprintf("[PATCH /members/{memberId}][%d] patchMemberNotFound ", 404)
}

func (o *PatchMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMemberConflict creates a PatchMemberConflict with default headers values
func NewPatchMemberConflict() *PatchMemberConflict {
	return &PatchMemberConflict{}
}

/*PatchMemberConflict handles this case with default header values.

Status conflict.
*/
type PatchMemberConflict struct {
}

func (o *PatchMemberConflict) Error() string {
	return fmt.Sprintf("[PATCH /members/{memberId}][%d] patchMemberConflict ", 409)
}

func (o *PatchMemberConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
