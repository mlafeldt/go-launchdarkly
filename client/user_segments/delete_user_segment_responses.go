// Code generated by go-swagger; DO NOT EDIT.

package user_segments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserSegmentReader is a Reader for the DeleteUserSegment structure.
type DeleteUserSegmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserSegmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserSegmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteUserSegmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserSegmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserSegmentNoContent creates a DeleteUserSegmentNoContent with default headers values
func NewDeleteUserSegmentNoContent() *DeleteUserSegmentNoContent {
	return &DeleteUserSegmentNoContent{}
}

/*DeleteUserSegmentNoContent handles this case with default header values.

Action completed successfully.
*/
type DeleteUserSegmentNoContent struct {
}

func (o *DeleteUserSegmentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /segments/{projectKey}/{environmentKey}/{userSegmentKey}][%d] deleteUserSegmentNoContent ", 204)
}

func (o *DeleteUserSegmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserSegmentUnauthorized creates a DeleteUserSegmentUnauthorized with default headers values
func NewDeleteUserSegmentUnauthorized() *DeleteUserSegmentUnauthorized {
	return &DeleteUserSegmentUnauthorized{}
}

/*DeleteUserSegmentUnauthorized handles this case with default header values.

Invalid access token.
*/
type DeleteUserSegmentUnauthorized struct {
}

func (o *DeleteUserSegmentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /segments/{projectKey}/{environmentKey}/{userSegmentKey}][%d] deleteUserSegmentUnauthorized ", 401)
}

func (o *DeleteUserSegmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserSegmentNotFound creates a DeleteUserSegmentNotFound with default headers values
func NewDeleteUserSegmentNotFound() *DeleteUserSegmentNotFound {
	return &DeleteUserSegmentNotFound{}
}

/*DeleteUserSegmentNotFound handles this case with default header values.

Invalid resource specifier.
*/
type DeleteUserSegmentNotFound struct {
}

func (o *DeleteUserSegmentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /segments/{projectKey}/{environmentKey}/{userSegmentKey}][%d] deleteUserSegmentNotFound ", 404)
}

func (o *DeleteUserSegmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
