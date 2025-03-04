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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// VersionGateServer represents the interface the manages the 'version_gate' resource.
type VersionGateServer interface {

	// Delete handles a request for the 'delete' method.
	//
	// Deletes the version gate.
	Delete(ctx context.Context, request *VersionGateDeleteServerRequest, response *VersionGateDeleteServerResponse) error

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the version gate.
	Get(ctx context.Context, request *VersionGateGetServerRequest, response *VersionGateGetServerResponse) error
}

// VersionGateDeleteServerRequest is the request for the 'delete' method.
type VersionGateDeleteServerRequest struct {
}

// VersionGateDeleteServerResponse is the response for the 'delete' method.
type VersionGateDeleteServerResponse struct {
	status int
	err    *errors.Error
}

// Status sets the status code.
func (r *VersionGateDeleteServerResponse) Status(value int) *VersionGateDeleteServerResponse {
	r.status = value
	return r
}

// VersionGateGetServerRequest is the request for the 'get' method.
type VersionGateGetServerRequest struct {
}

// VersionGateGetServerResponse is the response for the 'get' method.
type VersionGateGetServerResponse struct {
	status int
	err    *errors.Error
	body   *VersionGate
}

// Body sets the value of the 'body' parameter.
//
//
func (r *VersionGateGetServerResponse) Body(value *VersionGate) *VersionGateGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *VersionGateGetServerResponse) Status(value int) *VersionGateGetServerResponse {
	r.status = value
	return r
}

// dispatchVersionGate navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchVersionGate(w http.ResponseWriter, r *http.Request, server VersionGateServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "DELETE":
			adaptVersionGateDeleteRequest(w, r, server)
			return
		case "GET":
			adaptVersionGateGetRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		errors.SendNotFound(w, r)
		return
	}
}

// adaptVersionGateDeleteRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptVersionGateDeleteRequest(w http.ResponseWriter, r *http.Request, server VersionGateServer) {
	request := &VersionGateDeleteServerRequest{}
	err := readVersionGateDeleteRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &VersionGateDeleteServerResponse{}
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
	err = writeVersionGateDeleteResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}

// adaptVersionGateGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptVersionGateGetRequest(w http.ResponseWriter, r *http.Request, server VersionGateServer) {
	request := &VersionGateGetServerRequest{}
	err := readVersionGateGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &VersionGateGetServerResponse{}
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
	err = writeVersionGateGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
