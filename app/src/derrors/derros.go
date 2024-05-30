package derrors

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

//lint:file-ignore ST1012 prefixing error values with Err would stutter

var (
	// Unknown indicates that the error has unknown semantics.
	Unknown = errors.New("unknown")

	// DBModuleInsertInvalid represents a module that was successfully
	// fetched but could not be inserted due to invalid arguments to
	// postgres.InsertModule.
	DBModuleInsertInvalid = errors.New("db module insert invalid")

	OcurredAnQueryServieError = errors.New("ocurred an query service error")
	OccuredAnRepositoryError  = errors.New("ocurred an repository error")

	JSONUnmarshalError = errors.New("json unmarshal error")

	// 10x
	Continue           = errors.New("continue")
	SwitchingProtocols = errors.New("switching protocols")
	Processing         = errors.New("processing")
	EarlyHints         = errors.New("early hints")

	// 20x
	NonAuthoritativeInformation = errors.New("non authoritative information")
	NoContent                   = errors.New("no content")
	ResetContent                = errors.New("reset content")
	PartialContent              = errors.New("partial content")
	MultiStatus                 = errors.New("multi status")
	AlreadyReported             = errors.New("already reported")

	// 22x
	IMUsed = errors.New("im used")

	// 30x
	MultipleChoices   = errors.New("multiple choices")
	MovedPermanently  = errors.New("moved permanently")
	Found             = errors.New("found")
	SeeOther          = errors.New("see other")
	NotModified       = errors.New("not modified")
	UseProxy          = errors.New("use proxy")
	TemporaryRedirect = errors.New("temporary redirect")
	PermanentRedirect = errors.New("permanent redirect")

	// 40x
	BadRequest                  = errors.New("bad request")
	Unauthorized                = errors.New("unauthorized")
	PaymentRequired             = errors.New("payment required")
	Forbidden                   = errors.New("forbidden")
	NotFound                    = errors.New("not found")
	MethodNotAllowed            = errors.New("method not allowed")
	NotAcceptable               = errors.New("not acceptable")
	ProxyAuthenticationRequired = errors.New("proxy authentication required")
	RequestTimeout              = errors.New("request timeout")
	Conflict                    = errors.New("conflict")

	// 41x
	Gone                 = errors.New("gone")
	LengthRequired       = errors.New("length required")
	PreconditionFailed   = errors.New("precondition failed")
	ContentTooLarge      = errors.New("content too large")
	UnsupportedMediaType = errors.New("unsupported media type")
	RangeNotSatisfiable  = errors.New("range not satisfiable")
	ExpectationFailed    = errors.New("expectation failed")

	// 42x
	MisdirectedRequest   = errors.New("misdirected request")
	UnprocessableContent = errors.New("unprocessable content")
	Locked               = errors.New("locked")
	FailedDependency     = errors.New("failed dependency")
	TooEarly             = errors.New("too early")
	UpgradeRequired      = errors.New("upgrade required")
	PreconditionRequired = errors.New("precondition required")
	TooManyRequests      = errors.New("too many requests")

	// 43x
	RequestHeaderFieldsTooLarge = errors.New("request header fields too large")

	// 45x
	UnavailableForLegalReasons = errors.New("unavailable for legal reasons")

	// 50x
	InternalServerError     = errors.New("internal server error")
	NotImplemented          = errors.New("not implemented")
	BadGateway              = errors.New("bad gateway")
	ServiceUnavailable      = errors.New("service unavailable")
	GatewayTimeout          = errors.New("gateway timeout")
	HTTPVersionNotSupported = errors.New("http version not supported")
	VariantAlsoNegotiates   = errors.New("variant also negotiates")
	InsufficientStorage     = errors.New("insufficient storage")
	LoopDetected            = errors.New("loop detected")

	// 51x
	NotExtended                   = errors.New("not extended")
	NetworkAuthenticationRequired = errors.New("network authentication required")
)

var codes = []struct {
	err  error
	code int
}{
	// 10x
	{Continue, http.StatusContinue},
	{SwitchingProtocols, http.StatusSwitchingProtocols},
	{Processing, http.StatusProcessing},
	{EarlyHints, http.StatusEarlyHints},

	// 20x
	{NonAuthoritativeInformation, http.StatusNonAuthoritativeInfo},
	{NoContent, http.StatusNoContent},
	{ResetContent, http.StatusResetContent},
	{PartialContent, http.StatusPartialContent},
	{MultiStatus, http.StatusMultiStatus},
	{AlreadyReported, http.StatusAlreadyReported},

	// 22x
	{IMUsed, http.StatusIMUsed},

	// 30x
	{MultipleChoices, http.StatusMultipleChoices},
	{MovedPermanently, http.StatusMovedPermanently},
	{Found, http.StatusFound},
	{SeeOther, http.StatusSeeOther},
	{NotModified, http.StatusNotModified},
	{UseProxy, http.StatusUseProxy},
	{TemporaryRedirect, http.StatusTemporaryRedirect},
	{PermanentRedirect, http.StatusPermanentRedirect},

	// 40x
	{BadRequest, http.StatusBadRequest},
	{Unauthorized, http.StatusUnauthorized},
	{PaymentRequired, http.StatusPaymentRequired},
	{Forbidden, http.StatusForbidden},
	{NotFound, http.StatusNotFound},
	{MethodNotAllowed, http.StatusMethodNotAllowed},
	{NotAcceptable, http.StatusNotAcceptable},
	{ProxyAuthenticationRequired, http.StatusProxyAuthRequired},
	{RequestTimeout, http.StatusRequestTimeout},
	{Conflict, http.StatusConflict},

	// 41x
	{Gone, http.StatusGone},
	{LengthRequired, http.StatusLengthRequired},
	{PreconditionFailed, http.StatusPreconditionFailed},
	{ContentTooLarge, http.StatusRequestEntityTooLarge},
	{UnsupportedMediaType, http.StatusUnsupportedMediaType},
	{RangeNotSatisfiable, http.StatusRequestedRangeNotSatisfiable},
	{ExpectationFailed, http.StatusExpectationFailed},

	// 42x
	{MisdirectedRequest, http.StatusMisdirectedRequest},
	{UnprocessableContent, http.StatusUnprocessableEntity},
	{Locked, http.StatusLocked},
	{FailedDependency, http.StatusFailedDependency},
	{TooEarly, http.StatusTooEarly},
	{UpgradeRequired, http.StatusUpgradeRequired},
	{PreconditionRequired, http.StatusPreconditionRequired},
	{TooManyRequests, http.StatusTooManyRequests},

	// 43x
	{RequestHeaderFieldsTooLarge, http.StatusRequestHeaderFieldsTooLarge},

	// 45x
	{UnavailableForLegalReasons, http.StatusUnavailableForLegalReasons},

	// 50x
	{InternalServerError, http.StatusInternalServerError},
	{NotImplemented, http.StatusNotImplemented},
	{BadGateway, http.StatusBadGateway},
	{ServiceUnavailable, http.StatusServiceUnavailable},
	{GatewayTimeout, http.StatusGatewayTimeout},
	{HTTPVersionNotSupported, http.StatusHTTPVersionNotSupported},
	{VariantAlsoNegotiates, http.StatusVariantAlsoNegotiates},
	{InsufficientStorage, http.StatusInsufficientStorage},
	{LoopDetected, http.StatusLoopDetected},

	// 51x
	{NotExtended, http.StatusNotExtended},
	{NetworkAuthenticationRequired, http.StatusNetworkAuthenticationRequired},
}

// FromStatus generates an error according for the given status code. It uses
// the given format string and arguments to create the error string according
// to the fmt package. If format is the empty string, then the error
// corresponding to the code is returned unwrapped.
//
// If code is http.StatusOK, it returns nil.
func FromStatus(code int, format string, args ...any) error {
	if code == http.StatusOK {
		return nil
	}
	var innerErr = Unknown
	for _, e := range codes {
		if e.code == code {
			innerErr = e.err
			break
		}
	}
	if format == "" {
		return innerErr
	}
	return fmt.Errorf(format+": %w", append(args, innerErr)...)
}

// ToStatus returns a status code corresponding to err.
func ToStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}
	for _, e := range codes {
		if errors.Is(err, e.err) {
			return e.code
		}
	}
	return http.StatusInternalServerError
}

// Add adds context to the error.
// The result cannot be unwrapped to recover the original error.
// It does nothing when *errp == nil.
//
// Example:
//
//	defer derrors.Add(&err, "copy(%s, %s)", src, dst)
//
// See Wrap for an equivalent function that allows
// the result to be unwrapped.
func Add(errp *error, format string, args ...any) {
	if *errp != nil {
		*errp = fmt.Errorf("%s: %v", fmt.Sprintf(format, args...), *errp)
	}
}

// Wrap adds context to the error and allows
// unwrapping the result to recover the original error.
//
// Example:
//
//	defer derrors.Wrap(&err, "copy(%s, %s)", src, dst)
//
// See Add for an equivalent function that does not allow
// the result to be unwrapped.
func Wrap(errp *error, format string, args ...any) {
	if *errp != nil {
		*errp = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *errp)
	}
}

// WrapStack is like Wrap, but adds a stack trace if there isn't one already.
func WrapStack(errp *error, format string, args ...any) {
	if *errp != nil {
		if se := (*StackError)(nil); !errors.As(*errp, &se) {
			*errp = NewStackError(*errp)
		}
		Wrap(errp, format, args...)
	}
}

// StackError wraps an error and adds a stack trace.
type StackError struct {
	Stack []byte
	err   error
}

// NewStackError returns a StackError, capturing a stack trace.
func NewStackError(err error) *StackError {
	// Limit the stack trace to 16K. Same value used in the errorreporting client,
	// cloud.google.com/go@v0.66.0/errorreporting/errors.go.
	var buf [16 * 1024]byte
	n := runtime.Stack(buf[:], false)
	return &StackError{
		err:   err,
		Stack: buf[:n],
	}
}

func (e *StackError) Error() string {
	return e.err.Error() // ignore the stack
}

func (e *StackError) Unwrap() error {
	return e.err
}

// WrapAndReport calls Wrap followed by Report.
func WrapAndReport(errp *error, format string, args ...any) {
	Wrap(errp, format, args...)
	if *errp != nil {
		Report(*errp)
	}
}

var reporter Reporter

// SetReporter the Reporter to use, for use by Report.
func SetReporter(r Reporter) {
	reporter = r
}

// Reporter is an interface used for reporting errors.
type Reporter interface {
	Report(err error, req *http.Request, stack []byte)
}

// Report uses the Reporter to report an error.
func Report(err error) {
	if reporter != nil {
		reporter.Report(err, nil, nil)
	}
}
