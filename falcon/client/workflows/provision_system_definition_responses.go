// Code generated by go-swagger; DO NOT EDIT.

package workflows

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

// ProvisionSystemDefinitionReader is a Reader for the ProvisionSystemDefinition structure.
type ProvisionSystemDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProvisionSystemDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProvisionSystemDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProvisionSystemDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProvisionSystemDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProvisionSystemDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewProvisionSystemDefinitionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProvisionSystemDefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /workflows/system-definitions/provision/v1] provision.system-definition", response, response.Code())
	}
}

// NewProvisionSystemDefinitionOK creates a ProvisionSystemDefinitionOK with default headers values
func NewProvisionSystemDefinitionOK() *ProvisionSystemDefinitionOK {
	return &ProvisionSystemDefinitionOK{}
}

/*
ProvisionSystemDefinitionOK describes a response with status code 200, with default header values.

OK
*/
type ProvisionSystemDefinitionOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSystemDefinitionCreateResponse
}

// IsSuccess returns true when this provision system definition o k response has a 2xx status code
func (o *ProvisionSystemDefinitionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this provision system definition o k response has a 3xx status code
func (o *ProvisionSystemDefinitionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provision system definition o k response has a 4xx status code
func (o *ProvisionSystemDefinitionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this provision system definition o k response has a 5xx status code
func (o *ProvisionSystemDefinitionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this provision system definition o k response a status code equal to that given
func (o *ProvisionSystemDefinitionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the provision system definition o k response
func (o *ProvisionSystemDefinitionOK) Code() int {
	return 200
}

func (o *ProvisionSystemDefinitionOK) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionOK  %+v", 200, o.Payload)
}

func (o *ProvisionSystemDefinitionOK) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionOK  %+v", 200, o.Payload)
}

func (o *ProvisionSystemDefinitionOK) GetPayload() *models.ClientSystemDefinitionCreateResponse {
	return o.Payload
}

func (o *ProvisionSystemDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientSystemDefinitionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisionSystemDefinitionBadRequest creates a ProvisionSystemDefinitionBadRequest with default headers values
func NewProvisionSystemDefinitionBadRequest() *ProvisionSystemDefinitionBadRequest {
	return &ProvisionSystemDefinitionBadRequest{}
}

/*
ProvisionSystemDefinitionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProvisionSystemDefinitionBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSystemDefinitionCreateResponse
}

// IsSuccess returns true when this provision system definition bad request response has a 2xx status code
func (o *ProvisionSystemDefinitionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this provision system definition bad request response has a 3xx status code
func (o *ProvisionSystemDefinitionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provision system definition bad request response has a 4xx status code
func (o *ProvisionSystemDefinitionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this provision system definition bad request response has a 5xx status code
func (o *ProvisionSystemDefinitionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this provision system definition bad request response a status code equal to that given
func (o *ProvisionSystemDefinitionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the provision system definition bad request response
func (o *ProvisionSystemDefinitionBadRequest) Code() int {
	return 400
}

func (o *ProvisionSystemDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *ProvisionSystemDefinitionBadRequest) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *ProvisionSystemDefinitionBadRequest) GetPayload() *models.ClientSystemDefinitionCreateResponse {
	return o.Payload
}

func (o *ProvisionSystemDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientSystemDefinitionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisionSystemDefinitionForbidden creates a ProvisionSystemDefinitionForbidden with default headers values
func NewProvisionSystemDefinitionForbidden() *ProvisionSystemDefinitionForbidden {
	return &ProvisionSystemDefinitionForbidden{}
}

/*
ProvisionSystemDefinitionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProvisionSystemDefinitionForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this provision system definition forbidden response has a 2xx status code
func (o *ProvisionSystemDefinitionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this provision system definition forbidden response has a 3xx status code
func (o *ProvisionSystemDefinitionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provision system definition forbidden response has a 4xx status code
func (o *ProvisionSystemDefinitionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this provision system definition forbidden response has a 5xx status code
func (o *ProvisionSystemDefinitionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this provision system definition forbidden response a status code equal to that given
func (o *ProvisionSystemDefinitionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the provision system definition forbidden response
func (o *ProvisionSystemDefinitionForbidden) Code() int {
	return 403
}

func (o *ProvisionSystemDefinitionForbidden) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *ProvisionSystemDefinitionForbidden) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *ProvisionSystemDefinitionForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ProvisionSystemDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisionSystemDefinitionNotFound creates a ProvisionSystemDefinitionNotFound with default headers values
func NewProvisionSystemDefinitionNotFound() *ProvisionSystemDefinitionNotFound {
	return &ProvisionSystemDefinitionNotFound{}
}

/*
ProvisionSystemDefinitionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProvisionSystemDefinitionNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSystemDefinitionCreateResponse
}

// IsSuccess returns true when this provision system definition not found response has a 2xx status code
func (o *ProvisionSystemDefinitionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this provision system definition not found response has a 3xx status code
func (o *ProvisionSystemDefinitionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provision system definition not found response has a 4xx status code
func (o *ProvisionSystemDefinitionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this provision system definition not found response has a 5xx status code
func (o *ProvisionSystemDefinitionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this provision system definition not found response a status code equal to that given
func (o *ProvisionSystemDefinitionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the provision system definition not found response
func (o *ProvisionSystemDefinitionNotFound) Code() int {
	return 404
}

func (o *ProvisionSystemDefinitionNotFound) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *ProvisionSystemDefinitionNotFound) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *ProvisionSystemDefinitionNotFound) GetPayload() *models.ClientSystemDefinitionCreateResponse {
	return o.Payload
}

func (o *ProvisionSystemDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientSystemDefinitionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisionSystemDefinitionTooManyRequests creates a ProvisionSystemDefinitionTooManyRequests with default headers values
func NewProvisionSystemDefinitionTooManyRequests() *ProvisionSystemDefinitionTooManyRequests {
	return &ProvisionSystemDefinitionTooManyRequests{}
}

/*
ProvisionSystemDefinitionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ProvisionSystemDefinitionTooManyRequests struct {

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

// IsSuccess returns true when this provision system definition too many requests response has a 2xx status code
func (o *ProvisionSystemDefinitionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this provision system definition too many requests response has a 3xx status code
func (o *ProvisionSystemDefinitionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provision system definition too many requests response has a 4xx status code
func (o *ProvisionSystemDefinitionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this provision system definition too many requests response has a 5xx status code
func (o *ProvisionSystemDefinitionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this provision system definition too many requests response a status code equal to that given
func (o *ProvisionSystemDefinitionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the provision system definition too many requests response
func (o *ProvisionSystemDefinitionTooManyRequests) Code() int {
	return 429
}

func (o *ProvisionSystemDefinitionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionTooManyRequests  %+v", 429, o.Payload)
}

func (o *ProvisionSystemDefinitionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionTooManyRequests  %+v", 429, o.Payload)
}

func (o *ProvisionSystemDefinitionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ProvisionSystemDefinitionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewProvisionSystemDefinitionInternalServerError creates a ProvisionSystemDefinitionInternalServerError with default headers values
func NewProvisionSystemDefinitionInternalServerError() *ProvisionSystemDefinitionInternalServerError {
	return &ProvisionSystemDefinitionInternalServerError{}
}

/*
ProvisionSystemDefinitionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ProvisionSystemDefinitionInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSystemDefinitionCreateResponse
}

// IsSuccess returns true when this provision system definition internal server error response has a 2xx status code
func (o *ProvisionSystemDefinitionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this provision system definition internal server error response has a 3xx status code
func (o *ProvisionSystemDefinitionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provision system definition internal server error response has a 4xx status code
func (o *ProvisionSystemDefinitionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this provision system definition internal server error response has a 5xx status code
func (o *ProvisionSystemDefinitionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this provision system definition internal server error response a status code equal to that given
func (o *ProvisionSystemDefinitionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the provision system definition internal server error response
func (o *ProvisionSystemDefinitionInternalServerError) Code() int {
	return 500
}

func (o *ProvisionSystemDefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *ProvisionSystemDefinitionInternalServerError) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/provision/v1][%d] provisionSystemDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *ProvisionSystemDefinitionInternalServerError) GetPayload() *models.ClientSystemDefinitionCreateResponse {
	return o.Payload
}

func (o *ProvisionSystemDefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientSystemDefinitionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}