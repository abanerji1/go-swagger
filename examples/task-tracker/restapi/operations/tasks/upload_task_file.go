package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// UploadTaskFileHandlerFunc turns a function with the right signature into a upload task file handler
type UploadTaskFileHandlerFunc func(UploadTaskFileParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadTaskFileHandlerFunc) Handle(params UploadTaskFileParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UploadTaskFileHandler interface for that can handle valid upload task file params
type UploadTaskFileHandler interface {
	Handle(UploadTaskFileParams, interface{}) middleware.Responder
}

// NewUploadTaskFile creates a new http.Handler for the upload task file operation
func NewUploadTaskFile(ctx *middleware.Context, handler UploadTaskFileHandler) *UploadTaskFile {
	return &UploadTaskFile{Context: ctx, Handler: handler}
}

/*UploadTaskFile swagger:route POST /tasks/{id}/files tasks uploadTaskFile

Adds a file to a task.

The file can't be larger than **5MB**

*/
type UploadTaskFile struct {
	Context *middleware.Context
	Params  UploadTaskFileParams
	Handler UploadTaskFileHandler
}

func (o *UploadTaskFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewUploadTaskFileParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
