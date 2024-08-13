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

// UpsertBusinessApplicationsReader is a Reader for the UpsertBusinessApplications structure.
type UpsertBusinessApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpsertBusinessApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpsertBusinessApplicationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpsertBusinessApplicationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpsertBusinessApplicationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpsertBusinessApplicationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpsertBusinessApplicationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpsertBusinessApplicationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /aspm-api-gateway/api/v1/business_applications] UpsertBusinessApplications", response, response.Code())
	}
}

// NewUpsertBusinessApplicationsCreated creates a UpsertBusinessApplicationsCreated with default headers values
func NewUpsertBusinessApplicationsCreated() *UpsertBusinessApplicationsCreated {
	return &UpsertBusinessApplicationsCreated{}
}

/*
UpsertBusinessApplicationsCreated describes a response with status code 201, with default header values.

Created
*/
type UpsertBusinessApplicationsCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this upsert business applications created response has a 2xx status code
func (o *UpsertBusinessApplicationsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upsert business applications created response has a 3xx status code
func (o *UpsertBusinessApplicationsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upsert business applications created response has a 4xx status code
func (o *UpsertBusinessApplicationsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this upsert business applications created response has a 5xx status code
func (o *UpsertBusinessApplicationsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this upsert business applications created response a status code equal to that given
func (o *UpsertBusinessApplicationsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the upsert business applications created response
func (o *UpsertBusinessApplicationsCreated) Code() int {
	return 201
}

func (o *UpsertBusinessApplicationsCreated) Error() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsCreated ", 201)
}

func (o *UpsertBusinessApplicationsCreated) String() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsCreated ", 201)
}

func (o *UpsertBusinessApplicationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewUpsertBusinessApplicationsBadRequest creates a UpsertBusinessApplicationsBadRequest with default headers values
func NewUpsertBusinessApplicationsBadRequest() *UpsertBusinessApplicationsBadRequest {
	return &UpsertBusinessApplicationsBadRequest{}
}

/*
UpsertBusinessApplicationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpsertBusinessApplicationsBadRequest struct {

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

// IsSuccess returns true when this upsert business applications bad request response has a 2xx status code
func (o *UpsertBusinessApplicationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upsert business applications bad request response has a 3xx status code
func (o *UpsertBusinessApplicationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upsert business applications bad request response has a 4xx status code
func (o *UpsertBusinessApplicationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upsert business applications bad request response has a 5xx status code
func (o *UpsertBusinessApplicationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upsert business applications bad request response a status code equal to that given
func (o *UpsertBusinessApplicationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the upsert business applications bad request response
func (o *UpsertBusinessApplicationsBadRequest) Code() int {
	return 400
}

func (o *UpsertBusinessApplicationsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsBadRequest  %+v", 400, o.Payload)
}

func (o *UpsertBusinessApplicationsBadRequest) String() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsBadRequest  %+v", 400, o.Payload)
}

func (o *UpsertBusinessApplicationsBadRequest) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *UpsertBusinessApplicationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpsertBusinessApplicationsUnauthorized creates a UpsertBusinessApplicationsUnauthorized with default headers values
func NewUpsertBusinessApplicationsUnauthorized() *UpsertBusinessApplicationsUnauthorized {
	return &UpsertBusinessApplicationsUnauthorized{}
}

/*
UpsertBusinessApplicationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpsertBusinessApplicationsUnauthorized struct {

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

// IsSuccess returns true when this upsert business applications unauthorized response has a 2xx status code
func (o *UpsertBusinessApplicationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upsert business applications unauthorized response has a 3xx status code
func (o *UpsertBusinessApplicationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upsert business applications unauthorized response has a 4xx status code
func (o *UpsertBusinessApplicationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this upsert business applications unauthorized response has a 5xx status code
func (o *UpsertBusinessApplicationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this upsert business applications unauthorized response a status code equal to that given
func (o *UpsertBusinessApplicationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the upsert business applications unauthorized response
func (o *UpsertBusinessApplicationsUnauthorized) Code() int {
	return 401
}

func (o *UpsertBusinessApplicationsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsUnauthorized  %+v", 401, o.Payload)
}

func (o *UpsertBusinessApplicationsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsUnauthorized  %+v", 401, o.Payload)
}

func (o *UpsertBusinessApplicationsUnauthorized) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *UpsertBusinessApplicationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpsertBusinessApplicationsForbidden creates a UpsertBusinessApplicationsForbidden with default headers values
func NewUpsertBusinessApplicationsForbidden() *UpsertBusinessApplicationsForbidden {
	return &UpsertBusinessApplicationsForbidden{}
}

/*
UpsertBusinessApplicationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpsertBusinessApplicationsForbidden struct {

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

// IsSuccess returns true when this upsert business applications forbidden response has a 2xx status code
func (o *UpsertBusinessApplicationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upsert business applications forbidden response has a 3xx status code
func (o *UpsertBusinessApplicationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upsert business applications forbidden response has a 4xx status code
func (o *UpsertBusinessApplicationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this upsert business applications forbidden response has a 5xx status code
func (o *UpsertBusinessApplicationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this upsert business applications forbidden response a status code equal to that given
func (o *UpsertBusinessApplicationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the upsert business applications forbidden response
func (o *UpsertBusinessApplicationsForbidden) Code() int {
	return 403
}

func (o *UpsertBusinessApplicationsForbidden) Error() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsForbidden  %+v", 403, o.Payload)
}

func (o *UpsertBusinessApplicationsForbidden) String() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsForbidden  %+v", 403, o.Payload)
}

func (o *UpsertBusinessApplicationsForbidden) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *UpsertBusinessApplicationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpsertBusinessApplicationsTooManyRequests creates a UpsertBusinessApplicationsTooManyRequests with default headers values
func NewUpsertBusinessApplicationsTooManyRequests() *UpsertBusinessApplicationsTooManyRequests {
	return &UpsertBusinessApplicationsTooManyRequests{}
}

/*
UpsertBusinessApplicationsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpsertBusinessApplicationsTooManyRequests struct {

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

// IsSuccess returns true when this upsert business applications too many requests response has a 2xx status code
func (o *UpsertBusinessApplicationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upsert business applications too many requests response has a 3xx status code
func (o *UpsertBusinessApplicationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upsert business applications too many requests response has a 4xx status code
func (o *UpsertBusinessApplicationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this upsert business applications too many requests response has a 5xx status code
func (o *UpsertBusinessApplicationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this upsert business applications too many requests response a status code equal to that given
func (o *UpsertBusinessApplicationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the upsert business applications too many requests response
func (o *UpsertBusinessApplicationsTooManyRequests) Code() int {
	return 429
}

func (o *UpsertBusinessApplicationsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpsertBusinessApplicationsTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpsertBusinessApplicationsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpsertBusinessApplicationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpsertBusinessApplicationsServiceUnavailable creates a UpsertBusinessApplicationsServiceUnavailable with default headers values
func NewUpsertBusinessApplicationsServiceUnavailable() *UpsertBusinessApplicationsServiceUnavailable {
	return &UpsertBusinessApplicationsServiceUnavailable{}
}

/*
UpsertBusinessApplicationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type UpsertBusinessApplicationsServiceUnavailable struct {

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

// IsSuccess returns true when this upsert business applications service unavailable response has a 2xx status code
func (o *UpsertBusinessApplicationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upsert business applications service unavailable response has a 3xx status code
func (o *UpsertBusinessApplicationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upsert business applications service unavailable response has a 4xx status code
func (o *UpsertBusinessApplicationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this upsert business applications service unavailable response has a 5xx status code
func (o *UpsertBusinessApplicationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this upsert business applications service unavailable response a status code equal to that given
func (o *UpsertBusinessApplicationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the upsert business applications service unavailable response
func (o *UpsertBusinessApplicationsServiceUnavailable) Code() int {
	return 503
}

func (o *UpsertBusinessApplicationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpsertBusinessApplicationsServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /aspm-api-gateway/api/v1/business_applications][%d] upsertBusinessApplicationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpsertBusinessApplicationsServiceUnavailable) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *UpsertBusinessApplicationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
