package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	http.ResponseWriter
} 

// 2--

// Return the body with status 200
func (res *Response) Ok(body []byte) {
	res.WriteHeader(http.StatusOK)
	res.Write(body)
}

// Return the text body with status 200
func (res *Response) Text(body string) {
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(body))
}

// Return the html body with status 200
func (res *Response) Html(body string) {
	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(body))
}

// Return the xml body with status 200
func (res *Response) Xml(body string) {
	res.Header().Set("Content-Type", "application/xml")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(body))
}

// Return the json body with status 200
func (res *Response) Json(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 201
func (res *Response) Created(body []byte) {
	res.WriteHeader(http.StatusCreated)
	res.Write(body)
}

// Returns the json body with status 201 and adds the Content-Type header
func (res *Response) CreatedJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 204
func (res *Response) NoContent() {
	res.WriteHeader(http.StatusNoContent)
}

// 3--

// Returns the body with status 301
func (res *Response) MovedPermanently(body []byte) {
	res.WriteHeader(http.StatusMovedPermanently)
	res.Write(body)
}

// Returns the body with status 302
func (res *Response) Redirect(url string) {
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusFound)
}

// Returns the body with status 304
func (res *Response) NotModified(body []byte) {
	res.WriteHeader(http.StatusNotModified)
	res.Write(body)
}

// 4--

// Returns the body with status 400
func (res *Response) BadRequest(body []byte) {
	res.WriteHeader(http.StatusBadRequest)
	res.Write(body)
}

// Returns the json body with status 400 and adds the Content-Type header
func (res *Response) BadRequestJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 401
func (res *Response) Unauthorized(body []byte) {
	res.WriteHeader(http.StatusUnauthorized)
	res.Write(body)
}

// Returns the json body with status 401 and adds the Content-Type header
func (res *Response) UnauthorizedJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 403
func (res *Response) Forbidden(body []byte) {
	res.WriteHeader(http.StatusForbidden)
	res.Write(body)
}

// Returns the json body with status 403 and adds the Content-Type header
func (res *Response) ForbiddenJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusForbidden)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 404
func (res *Response) NotFound(body []byte) {
	res.WriteHeader(http.StatusNotFound)
	res.Write(body)
}

// Returns the json body with status 404 and adds the Content-Type header
func (res *Response) NotFoundJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNotFound)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 405
func (res *Response) MethodNotAllowed(body []byte) {
	res.WriteHeader(http.StatusMethodNotAllowed)
	res.Write(body)
}

// Returns the json body with status 405 and adds the Content-Type header
func (res *Response) MethodNotAllowedJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 406
func (res *Response) NotAcceptable(body []byte) {
	res.WriteHeader(http.StatusNotAcceptable)
	res.Write(body)
}

// Returns the json body with status 406 and adds the Content-Type header
func (res *Response) NotAcceptableJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNotAcceptable)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 408
func (res *Response) RequestTimeout(body []byte) {
	res.WriteHeader(http.StatusRequestTimeout)
	res.Write(body)
}

// Returns the json body with status 408 and adds the Content-Type header
func (res *Response) RequestTimeoutJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusRequestTimeout)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 409
func (res *Response) Conflict(body []byte) {
	res.WriteHeader(http.StatusConflict)
	res.Write(body)
}

// Returns the json body with status 409 and adds the Content-Type header
func (res *Response) ConflictJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusConflict)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 410
func (res *Response) Gone(body []byte) {
	res.WriteHeader(http.StatusGone)
	res.Write(body)
}

// Returns the json body with status 410 and adds the Content-Type header
func (res *Response) GoneJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusGone)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 411
func (res *Response) LengthRequired(body []byte) {
	res.WriteHeader(http.StatusLengthRequired)
	res.Write(body)
}

// Returns the json body with status 411 and adds the Content-Type header
func (res *Response) LengthRequiredJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusLengthRequired)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 412
func (res *Response) PreconditionFailed(body []byte) {
	res.WriteHeader(http.StatusPreconditionFailed)
	res.Write(body)
}

// Returns the json body with status 412 and adds the Content-Type header
func (res *Response) PreconditionFailedJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusPreconditionFailed)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 413
func (res *Response) RequestEntityTooLarge(body []byte) {
	res.WriteHeader(http.StatusRequestEntityTooLarge)
	res.Write(body)
}

// Returns the json body with status 413 and adds the Content-Type header
func (res *Response) RequestEntityTooLargeJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusRequestEntityTooLarge)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 414
func (res *Response) RequestURITooLong(body []byte) {
	res.WriteHeader(http.StatusRequestURITooLong)
	res.Write(body)
}

// Returns the json body with status 414 and adds the Content-Type header
func (res *Response) RequestURITooLongJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusRequestURITooLong)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 415
func (res *Response) UnsupportedMediaType(body []byte) {
	res.WriteHeader(http.StatusUnsupportedMediaType)
	res.Write(body)
}

// Returns the json body with status 415 and adds the Content-Type header
func (res *Response) UnsupportedMediaTypeJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusUnsupportedMediaType)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 416
func (res *Response) RequestedRangeNotSatisfiable(body []byte) {
	res.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
	res.Write(body)
}

// Returns the json body with status 416 and adds the Content-Type header
func (res *Response) RequestedRangeNotSatisfiableJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 417
func (res *Response) ExpectationFailed(body []byte) {
	res.WriteHeader(http.StatusExpectationFailed)
	res.Write(body)
}

// Returns the json body with status 417 and adds the Content-Type header
func (res *Response) ExpectationFailedJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusExpectationFailed)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 418
func (res *Response) Teapot(body []byte) {
	res.WriteHeader(http.StatusTeapot)
	res.Write(body)
}

// Returns the json body with status 418 and adds the Content-Type header
func (res *Response) TeapotJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusTeapot)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 422
func (res *Response) UnprocessableEntity(body []byte) {
	res.WriteHeader(http.StatusUnprocessableEntity)
	res.Write(body)
}

// Returns the json body with status 422 and adds the Content-Type header
func (res *Response) UnprocessableEntityJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 423
func (res *Response) Locked(body []byte) {
	res.WriteHeader(http.StatusLocked)
	res.Write(body)
}

// Returns the json body with status 423 and adds the Content-Type header
func (res *Response) LockedJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusLocked)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 424
func (res *Response) FailedDependency(body []byte) {
	res.WriteHeader(http.StatusFailedDependency)
	res.Write(body)
}

// Returns the json body with status 424 and adds the Content-Type header
func (res *Response) FailedDependencyJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusFailedDependency)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 426
func (res *Response) UpgradeRequired(body []byte) {
	res.WriteHeader(http.StatusUpgradeRequired)
	res.Write(body)
}

// Returns the json body with status 426 and adds the Content-Type header
func (res *Response) UpgradeRequiredJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusUpgradeRequired)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 428
func (res *Response) PreconditionRequired(body []byte) {
	res.WriteHeader(http.StatusPreconditionRequired)
	res.Write(body)
}

// Returns the json body with status 428 and adds the Content-Type header
func (res *Response) PreconditionRequiredJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusPreconditionRequired)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 429
func (res *Response) TooManyRequests(body []byte) {
	res.WriteHeader(http.StatusTooManyRequests)
	res.Write(body)
}

// Returns the json body with status 429 and adds the Content-Type header
func (res *Response) TooManyRequestsJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusTooManyRequests)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 431
func (res *Response) RequestHeaderFieldsTooLarge(body []byte) {
	res.WriteHeader(http.StatusRequestHeaderFieldsTooLarge)
	res.Write(body)
}

// Returns the json body with status 431 and adds the Content-Type header
func (res *Response) RequestHeaderFieldsTooLargeJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusRequestHeaderFieldsTooLarge)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 451
func (res *Response) UnavailableForLegalReasons(body []byte) {
	res.WriteHeader(http.StatusUnavailableForLegalReasons)
	res.Write(body)
}

// Returns the json body with status 451 and adds the Content-Type header
func (res *Response) UnavailableForLegalReasonsJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusUnavailableForLegalReasons)
	json.NewEncoder(res).Encode(body)
}

// 5--

// Returns the body with status 500
func (res *Response) InternalServerError(body []byte) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write(body)
}

// Returns the json body with status 500 and adds the Content-Type header
func (res *Response) InternalServerErrorJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 501
func (res *Response) NotImplemented(body []byte) {
	res.WriteHeader(http.StatusNotImplemented)
	res.Write(body)
}

// Returns the json body with status 501 and adds the Content-Type header
func (res *Response) NotImplementedJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 502
func (res *Response) BadGateway(body []byte) {
	res.WriteHeader(http.StatusBadGateway)
	res.Write(body)
}

// Returns the json body with status 502 and adds the Content-Type header
func (res *Response) BadGatewayJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusBadGateway)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 503
func (res *Response) ServiceUnavailable(body []byte) {
	res.WriteHeader(http.StatusServiceUnavailable)
	res.Write(body)
}

// Returns the json body with status 503 and adds the Content-Type header
func (res *Response) ServiceUnavailableJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusServiceUnavailable)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 504
func (res *Response) GatewayTimeout(body []byte) {
	res.WriteHeader(http.StatusGatewayTimeout)
	res.Write(body)
}

// Returns the json body with status 504 and adds the Content-Type header
func (res *Response) GatewayTimeoutJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusGatewayTimeout)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 505
func (res *Response) HTTPVersionNotSupported(body []byte) {
	res.WriteHeader(http.StatusHTTPVersionNotSupported)
	res.Write(body)
}

// Returns the json body with status 505 and adds the Content-Type header
func (res *Response) HTTPVersionNotSupportedJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusHTTPVersionNotSupported)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 506
func (res *Response) VariantAlsoNegotiates(body []byte) {
	res.WriteHeader(http.StatusVariantAlsoNegotiates)
	res.Write(body)
}

// Returns the json body with status 506 and adds the Content-Type header
func (res *Response) VariantAlsoNegotiatesJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusVariantAlsoNegotiates)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 507
func (res *Response) InsufficientStorage(body []byte) {	
	res.WriteHeader(http.StatusInsufficientStorage)
	res.Write(body)
}

// Returns the json body with status 507 and adds the Content-Type header
func (res *Response) InsufficientStorageJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusInsufficientStorage)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 508
func (res *Response) LoopDetected(body []byte) {
	res.WriteHeader(http.StatusLoopDetected)
	res.Write(body)
}

// Returns the json body with status 508 and adds the Content-Type header
func (res *Response) LoopDetectedJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusLoopDetected)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 510
func (res *Response) NotExtended(body []byte) {
	res.WriteHeader(http.StatusNotExtended)
	res.Write(body)
}

// Returns the json body with status 510 and adds the Content-Type header
func (res *Response) NotExtendedJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNotExtended)
	json.NewEncoder(res).Encode(body)
}

// Returns the body with status 511
func (res *Response) NetworkAuthenticationRequired(body []byte) {
	res.WriteHeader(http.StatusNetworkAuthenticationRequired)
	res.Write(body)
}

// Returns the json body with status 511 and adds the Content-Type header
func (res *Response) NetworkAuthenticationRequiredJson(body interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNetworkAuthenticationRequired)
	json.NewEncoder(res).Encode(body)
}
