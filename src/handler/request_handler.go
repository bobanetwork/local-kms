package handler

import(
	"net/http"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"github.com/NSmithUK/local-kms-go/src/data"
)

//--------------------------------------------------------------------
// Incoming request

type RequestHandler struct {
	request		*http.Request
	logger 		*log.Logger
	database	*data.Database
}

func NewRequestHandler(r *http.Request, l *log.Logger, d *data.Database) *RequestHandler {
	return &RequestHandler{
		request: r,
		logger: l,
		database: d,
	}
}

/*
	Decodes the request's JSON body into the passed interface
 */
func (r *RequestHandler) decodeBodyInto(v interface{}) error {
	decoder := json.NewDecoder(r.request.Body)
	return decoder.Decode(v)
}

//--------------------------------------------------------------------
// Outgoing response

type Response struct {
	Code		int
	Body		string
}

func NewResponse(code int, v interface{}) Response {
	if v == nil {
		return Response{ code, "" }
	}

	j, err := json.Marshal(v)

	if err != nil {
		return Response{ 500, "Error marshalling JSON"}
	}

	return Response{ code, string(j) }
}

func New400ExceptionResponse(exception, message string) Response {
	response := map[string]string{"__type":exception}

	if message != "" {
		response["message"] = message
	}

	return NewResponse(400, response)
}

//-------------------------------------------------
// Error helpers

func NewMissingParameterResponse(message string) Response {
	return New400ExceptionResponse("MissingParameterException", message)
}

func NewNotFoundExceptionResponse(message string) Response {
	return New400ExceptionResponse("NotFoundException", message)
}

func NewAlreadyExistsExceptionResponse(message string) Response {
	return New400ExceptionResponse("AlreadyExistsException", message)
}

func NewNotAuthorizedExceptionResponse(message string) Response {
	return New400ExceptionResponse("NotAuthorizedException", message)
}

func NewValidationExceptionResponse(message string) Response {
	return New400ExceptionResponse("ValidationException", message)
}

func NewKMSInvalidStateExceptionResponse(message string) Response {
	return New400ExceptionResponse("KMSInvalidStateException", message)
}

//---

func NewInternalFailureExceptionResponse(message string) Response {
	return NewResponse(500, map[string]string{
		"__type":"InternalFailureException",
		"message": message,
	})
}
