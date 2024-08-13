// Code generated by go-swagger; DO NOT EDIT.

package a_s_p_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// CreateIntegrationReader is a Reader for the CreateIntegration structure.
type CreateIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateIntegrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateIntegrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateIntegrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateIntegrationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateIntegrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /aspm-api-gateway/api/v1/integrations] CreateIntegration", response, response.Code())
	}
}

// NewCreateIntegrationOK creates a CreateIntegrationOK with default headers values
func NewCreateIntegrationOK() *CreateIntegrationOK {
	return &CreateIntegrationOK{}
}

/*
CreateIntegrationOK describes a response with status code 200, with default header values.

OK
*/
type CreateIntegrationOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesIntegrationResponse
}

// IsSuccess returns true when this create integration o k response has a 2xx status code
func (o *CreateIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create integration o k response has a 3xx status code
func (o *CreateIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create integration o k response has a 4xx status code
func (o *CreateIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create integration o k response has a 5xx status code
func (o *CreateIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create integration o k response a status code equal to that given
func (o *CreateIntegrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create integration o k response
func (o *CreateIntegrationOK) Code() int {
	return 200
}

func (o *CreateIntegrationOK) Error() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationOK  %+v", 200, o.Payload)
}

func (o *CreateIntegrationOK) String() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationOK  %+v", 200, o.Payload)
}

func (o *CreateIntegrationOK) GetPayload() *models.TypesIntegrationResponse {
	return o.Payload
}

func (o *CreateIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.TypesIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationBadRequest creates a CreateIntegrationBadRequest with default headers values
func NewCreateIntegrationBadRequest() *CreateIntegrationBadRequest {
	return &CreateIntegrationBadRequest{}
}

/*
CreateIntegrationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateIntegrationBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesErrorMessage
}

// IsSuccess returns true when this create integration bad request response has a 2xx status code
func (o *CreateIntegrationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create integration bad request response has a 3xx status code
func (o *CreateIntegrationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create integration bad request response has a 4xx status code
func (o *CreateIntegrationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create integration bad request response has a 5xx status code
func (o *CreateIntegrationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create integration bad request response a status code equal to that given
func (o *CreateIntegrationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create integration bad request response
func (o *CreateIntegrationBadRequest) Code() int {
	return 400
}

func (o *CreateIntegrationBadRequest) Error() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateIntegrationBadRequest) String() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateIntegrationBadRequest) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *CreateIntegrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.TypesErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationUnauthorized creates a CreateIntegrationUnauthorized with default headers values
func NewCreateIntegrationUnauthorized() *CreateIntegrationUnauthorized {
	return &CreateIntegrationUnauthorized{}
}

/*
CreateIntegrationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateIntegrationUnauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesErrorMessage
}

// IsSuccess returns true when this create integration unauthorized response has a 2xx status code
func (o *CreateIntegrationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create integration unauthorized response has a 3xx status code
func (o *CreateIntegrationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create integration unauthorized response has a 4xx status code
func (o *CreateIntegrationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create integration unauthorized response has a 5xx status code
func (o *CreateIntegrationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create integration unauthorized response a status code equal to that given
func (o *CreateIntegrationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create integration unauthorized response
func (o *CreateIntegrationUnauthorized) Code() int {
	return 401
}

func (o *CreateIntegrationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateIntegrationUnauthorized) String() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateIntegrationUnauthorized) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *CreateIntegrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.TypesErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationForbidden creates a CreateIntegrationForbidden with default headers values
func NewCreateIntegrationForbidden() *CreateIntegrationForbidden {
	return &CreateIntegrationForbidden{}
}

/*
CreateIntegrationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateIntegrationForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesErrorMessage
}

// IsSuccess returns true when this create integration forbidden response has a 2xx status code
func (o *CreateIntegrationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create integration forbidden response has a 3xx status code
func (o *CreateIntegrationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create integration forbidden response has a 4xx status code
func (o *CreateIntegrationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create integration forbidden response has a 5xx status code
func (o *CreateIntegrationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create integration forbidden response a status code equal to that given
func (o *CreateIntegrationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create integration forbidden response
func (o *CreateIntegrationForbidden) Code() int {
	return 403
}

func (o *CreateIntegrationForbidden) Error() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationForbidden  %+v", 403, o.Payload)
}

func (o *CreateIntegrationForbidden) String() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationForbidden  %+v", 403, o.Payload)
}

func (o *CreateIntegrationForbidden) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *CreateIntegrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.TypesErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationTooManyRequests creates a CreateIntegrationTooManyRequests with default headers values
func NewCreateIntegrationTooManyRequests() *CreateIntegrationTooManyRequests {
	return &CreateIntegrationTooManyRequests{}
}

/*
CreateIntegrationTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateIntegrationTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this create integration too many requests response has a 2xx status code
func (o *CreateIntegrationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create integration too many requests response has a 3xx status code
func (o *CreateIntegrationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create integration too many requests response has a 4xx status code
func (o *CreateIntegrationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create integration too many requests response has a 5xx status code
func (o *CreateIntegrationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create integration too many requests response a status code equal to that given
func (o *CreateIntegrationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create integration too many requests response
func (o *CreateIntegrationTooManyRequests) Code() int {
	return 429
}

func (o *CreateIntegrationTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateIntegrationTooManyRequests) String() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateIntegrationTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateIntegrationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationInternalServerError creates a CreateIntegrationInternalServerError with default headers values
func NewCreateIntegrationInternalServerError() *CreateIntegrationInternalServerError {
	return &CreateIntegrationInternalServerError{}
}

/*
CreateIntegrationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateIntegrationInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesErrorMessage
}

// IsSuccess returns true when this create integration internal server error response has a 2xx status code
func (o *CreateIntegrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create integration internal server error response has a 3xx status code
func (o *CreateIntegrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create integration internal server error response has a 4xx status code
func (o *CreateIntegrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create integration internal server error response has a 5xx status code
func (o *CreateIntegrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create integration internal server error response a status code equal to that given
func (o *CreateIntegrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create integration internal server error response
func (o *CreateIntegrationInternalServerError) Code() int {
	return 500
}

func (o *CreateIntegrationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateIntegrationInternalServerError) String() string {
	return fmt.Sprintf("[POST /aspm-api-gateway/api/v1/integrations][%d] createIntegrationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateIntegrationInternalServerError) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *CreateIntegrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.TypesErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
