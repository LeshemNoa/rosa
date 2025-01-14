/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/statusboard/v1

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// ProductServer represents the interface the manages the 'product' resource.
type ProductServer interface {

	// Delete handles a request for the 'delete' method.
	//
	//
	Delete(ctx context.Context, request *ProductDeleteServerRequest, response *ProductDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	//
	Get(ctx context.Context, request *ProductGetServerRequest, response *ProductGetServerResponse) error

	// Update handles a request for the 'update' method.
	//
	//
	Update(ctx context.Context, request *ProductUpdateServerRequest, response *ProductUpdateServerResponse) error

	// Applications returns the target 'applications' resource.
	//
	//
	Applications() ApplicationsServer

	// Updates returns the target 'statuses' resource.
	//
	//
	Updates() StatusesServer
}

// ProductDeleteServerRequest is the request for the 'delete' method.
type ProductDeleteServerRequest struct {
}

// ProductDeleteServerResponse is the response for the 'delete' method.
type ProductDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *ProductDeleteServerResponse) Status(value int) *ProductDeleteServerResponse {
	r.status = value
	return r
}

// ProductGetServerRequest is the request for the 'get' method.
type ProductGetServerRequest struct {
}

// ProductGetServerResponse is the response for the 'get' method.
type ProductGetServerResponse struct {
	status int
	err    *errors.Error
	body   *Product
}

// Body sets the value of the 'body' parameter.
//
//
func (r *ProductGetServerResponse) Body(value *Product) *ProductGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *ProductGetServerResponse) Status(value int) *ProductGetServerResponse {
	r.status = value
	return r
}

// ProductUpdateServerRequest is the request for the 'update' method.
type ProductUpdateServerRequest struct {
	body *Product
}

// Body returns the value of the 'body' parameter.
//
//
func (r *ProductUpdateServerRequest) Body() *Product {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *ProductUpdateServerRequest) GetBody() (value *Product, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// ProductUpdateServerResponse is the response for the 'update' method.
type ProductUpdateServerResponse struct {
	status int
	err    *errors.Error
	body   *Product
}

// Body sets the value of the 'body' parameter.
//
//
func (r *ProductUpdateServerResponse) Body(value *Product) *ProductUpdateServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *ProductUpdateServerResponse) Status(value int) *ProductUpdateServerResponse {
	r.status = value
	return r
}

// dispatchProduct navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchProduct(w http.ResponseWriter, r *http.Request, server ProductServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "DELETE":
			adaptProductDeleteRequest(w, r, server)
			return
		case "GET":
			adaptProductGetRequest(w, r, server)
			return
		case "PATCH":
			adaptProductUpdateRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	case "applications":
		target := server.Applications()
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchApplications(w, r, target, segments[1:])
	case "updates":
		target := server.Updates()
		if target == nil {
			errors.SendNotFound(w, r)
			return
		}
		dispatchStatuses(w, r, target, segments[1:])
	default:
		errors.SendNotFound(w, r)
		return
	}
}

// adaptProductDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptProductDeleteRequest(w http.ResponseWriter, r *http.Request, server ProductServer) {
	request := &ProductDeleteServerRequest{}
	err := readProductDeleteRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &ProductDeleteServerResponse{}
	response.status = 204
	err = server.Delete(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeProductDeleteResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptProductGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptProductGetRequest(w http.ResponseWriter, r *http.Request, server ProductServer) {
	request := &ProductGetServerRequest{}
	err := readProductGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &ProductGetServerResponse{}
	response.status = 200
	err = server.Get(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeProductGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptProductUpdateRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptProductUpdateRequest(w http.ResponseWriter, r *http.Request, server ProductServer) {
	request := &ProductUpdateServerRequest{}
	err := readProductUpdateRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &ProductUpdateServerResponse{}
	response.status = 200
	err = server.Update(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeProductUpdateResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
