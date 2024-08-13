// Code generated by go-swagger; DO NOT EDIT.

package oauth2

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

// Oauth2AccessTokenReader is a Reader for the Oauth2AccessToken structure.
type Oauth2AccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Oauth2AccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOauth2AccessTokenCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOauth2AccessTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOauth2AccessTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewOauth2AccessTokenTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOauth2AccessTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /oauth2/token] oauth2AccessToken", response, response.Code())
	}
}

// NewOauth2AccessTokenCreated creates a Oauth2AccessTokenCreated with default headers values
func NewOauth2AccessTokenCreated() *Oauth2AccessTokenCreated {
	return &Oauth2AccessTokenCreated{}
}

/*
Oauth2AccessTokenCreated describes a response with status code 201, with default header values.

Successfully issued token
*/
type Oauth2AccessTokenCreated struct {
	XCSRegion string

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAccessTokenResponseV1
}

// IsSuccess returns true when this oauth2 access token created response has a 2xx status code
func (o *Oauth2AccessTokenCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this oauth2 access token created response has a 3xx status code
func (o *Oauth2AccessTokenCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oauth2 access token created response has a 4xx status code
func (o *Oauth2AccessTokenCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this oauth2 access token created response has a 5xx status code
func (o *Oauth2AccessTokenCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this oauth2 access token created response a status code equal to that given
func (o *Oauth2AccessTokenCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the oauth2 access token created response
func (o *Oauth2AccessTokenCreated) Code() int {
	return 201
}

func (o *Oauth2AccessTokenCreated) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenCreated  %+v", 201, o.Payload)
}

func (o *Oauth2AccessTokenCreated) String() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenCreated  %+v", 201, o.Payload)
}

func (o *Oauth2AccessTokenCreated) GetPayload() *models.DomainAccessTokenResponseV1 {
	return o.Payload
}

func (o *Oauth2AccessTokenCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-Region
	hdrXCSRegion := response.GetHeader("X-CS-Region")

	if hdrXCSRegion != "" {
		o.XCSRegion = hdrXCSRegion
	}

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

	o.Payload = new(models.DomainAccessTokenResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOauth2AccessTokenBadRequest creates a Oauth2AccessTokenBadRequest with default headers values
func NewOauth2AccessTokenBadRequest() *Oauth2AccessTokenBadRequest {
	return &Oauth2AccessTokenBadRequest{}
}

/*
Oauth2AccessTokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type Oauth2AccessTokenBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this oauth2 access token bad request response has a 2xx status code
func (o *Oauth2AccessTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this oauth2 access token bad request response has a 3xx status code
func (o *Oauth2AccessTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oauth2 access token bad request response has a 4xx status code
func (o *Oauth2AccessTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this oauth2 access token bad request response has a 5xx status code
func (o *Oauth2AccessTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this oauth2 access token bad request response a status code equal to that given
func (o *Oauth2AccessTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the oauth2 access token bad request response
func (o *Oauth2AccessTokenBadRequest) Code() int {
	return 400
}

func (o *Oauth2AccessTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenBadRequest  %+v", 400, o.Payload)
}

func (o *Oauth2AccessTokenBadRequest) String() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenBadRequest  %+v", 400, o.Payload)
}

func (o *Oauth2AccessTokenBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *Oauth2AccessTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOauth2AccessTokenForbidden creates a Oauth2AccessTokenForbidden with default headers values
func NewOauth2AccessTokenForbidden() *Oauth2AccessTokenForbidden {
	return &Oauth2AccessTokenForbidden{}
}

/*
Oauth2AccessTokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type Oauth2AccessTokenForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this oauth2 access token forbidden response has a 2xx status code
func (o *Oauth2AccessTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this oauth2 access token forbidden response has a 3xx status code
func (o *Oauth2AccessTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oauth2 access token forbidden response has a 4xx status code
func (o *Oauth2AccessTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this oauth2 access token forbidden response has a 5xx status code
func (o *Oauth2AccessTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this oauth2 access token forbidden response a status code equal to that given
func (o *Oauth2AccessTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the oauth2 access token forbidden response
func (o *Oauth2AccessTokenForbidden) Code() int {
	return 403
}

func (o *Oauth2AccessTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenForbidden  %+v", 403, o.Payload)
}

func (o *Oauth2AccessTokenForbidden) String() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenForbidden  %+v", 403, o.Payload)
}

func (o *Oauth2AccessTokenForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *Oauth2AccessTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOauth2AccessTokenTooManyRequests creates a Oauth2AccessTokenTooManyRequests with default headers values
func NewOauth2AccessTokenTooManyRequests() *Oauth2AccessTokenTooManyRequests {
	return &Oauth2AccessTokenTooManyRequests{}
}

/*
Oauth2AccessTokenTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type Oauth2AccessTokenTooManyRequests struct {

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

// IsSuccess returns true when this oauth2 access token too many requests response has a 2xx status code
func (o *Oauth2AccessTokenTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this oauth2 access token too many requests response has a 3xx status code
func (o *Oauth2AccessTokenTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oauth2 access token too many requests response has a 4xx status code
func (o *Oauth2AccessTokenTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this oauth2 access token too many requests response has a 5xx status code
func (o *Oauth2AccessTokenTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this oauth2 access token too many requests response a status code equal to that given
func (o *Oauth2AccessTokenTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the oauth2 access token too many requests response
func (o *Oauth2AccessTokenTooManyRequests) Code() int {
	return 429
}

func (o *Oauth2AccessTokenTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenTooManyRequests  %+v", 429, o.Payload)
}

func (o *Oauth2AccessTokenTooManyRequests) String() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenTooManyRequests  %+v", 429, o.Payload)
}

func (o *Oauth2AccessTokenTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2AccessTokenTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2AccessTokenInternalServerError creates a Oauth2AccessTokenInternalServerError with default headers values
func NewOauth2AccessTokenInternalServerError() *Oauth2AccessTokenInternalServerError {
	return &Oauth2AccessTokenInternalServerError{}
}

/*
Oauth2AccessTokenInternalServerError describes a response with status code 500, with default header values.

Failed to issue token
*/
type Oauth2AccessTokenInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this oauth2 access token internal server error response has a 2xx status code
func (o *Oauth2AccessTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this oauth2 access token internal server error response has a 3xx status code
func (o *Oauth2AccessTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oauth2 access token internal server error response has a 4xx status code
func (o *Oauth2AccessTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this oauth2 access token internal server error response has a 5xx status code
func (o *Oauth2AccessTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this oauth2 access token internal server error response a status code equal to that given
func (o *Oauth2AccessTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the oauth2 access token internal server error response
func (o *Oauth2AccessTokenInternalServerError) Code() int {
	return 500
}

func (o *Oauth2AccessTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *Oauth2AccessTokenInternalServerError) String() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] oauth2AccessTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *Oauth2AccessTokenInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *Oauth2AccessTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
