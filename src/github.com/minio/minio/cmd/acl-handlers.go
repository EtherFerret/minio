/*
 * Minio Cloud Storage, (C) 2018 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
	ACL "github.com/minio/minio/pkg/acl"
	"github.com/minio/minio/pkg/policy"
	"io"
)

func (api objectAPIHandlers) PutBucketACLHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, "SetBucketACL")

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	objAPI := api.ObjectAPI()
	if objAPI == nil {
		writeErrorResponse(w, ErrServerNotInitialized, r.URL)
		return
	}

	// Allow getBucketACL if policy action is set, since this is a dummy call
	// we are simply re-purposing the bucketPolicyAction.
	if s3Error := checkRequestAuthType(ctx, r, policy.GetBucketPolicyAction, bucket, ""); s3Error != ErrNone {
		writeErrorResponse(w, s3Error, r.URL)
		return
	}

	aclPolicy, err := ACL.ParseConfig(io.LimitReader(r.Body, r.ContentLength), bucket)
	if err != nil {
		writeErrorResponse(w, ErrMalformedPolicy, r.URL)
		return
	}

	err = objAPI.SetBucketAccessControlPolicy(ctx, bucket, aclPolicy)
	if err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	// Success.
	writeSuccessNoContent(w)
}

// GetBucketACLHandler - GET Bucket ACL
// -----------------
// This operation uses the ACL
// subresource to return the ACL of a specified bucket.
func (api objectAPIHandlers) GetBucketACLHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, "GetBucketACL")

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	objAPI := api.ObjectAPI()
	if objAPI == nil {
		writeErrorResponse(w, ErrServerNotInitialized, r.URL)
		return
	}

	// Allow getBucketACL if policy action is set, since this is a dummy call
	// we are simply re-purposing the bucketPolicyAction.
	if s3Error := checkRequestAuthType(ctx, r, policy.GetBucketPolicyAction, bucket, ""); s3Error != ErrNone {
		writeErrorResponse(w, s3Error, r.URL)
		return
	}

	// Before proceeding validate if bucket exists.
	_, err := objAPI.GetBucketInfo(ctx, bucket)
	if err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	/*
		acl := &ACL.AccessControlPolicy{}
		acl.AccessControlList.Grants = append(acl.AccessControlList.Grants, ACL.Grant{
			Grantee: ACL.Grantee{
				XMLNS:  "http://www.w3.org/2001/XMLSchema-instance",
				XMLXSI: "CanonicalUser",
				Type:   "CanonicalUser",
			},
			Permission: "FULL_CONTROL",
		})
	*/
	acl, err := objAPI.GetBucketAccessControlPolicy(ctx, bucket)
	if err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	if err := xml.NewEncoder(w).Encode(acl); err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	w.(http.Flusher).Flush()
}

func (api objectAPIHandlers) PutObjectACLHandler(w http.ResponseWriter, r *http.Request) {
	//FIXME
}

// GetObjectACLHandler - GET Object ACL
// -----------------
// This operation uses the ACL
// subresource to return the ACL of a specified object.
func (api objectAPIHandlers) GetObjectACLHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, "GetObjectACL")

	vars := mux.Vars(r)
	bucket := vars["bucket"]
	object := vars["object"]

	objAPI := api.ObjectAPI()
	if objAPI == nil {
		writeErrorResponse(w, ErrServerNotInitialized, r.URL)
		return
	}

	// Allow getObjectACL if policy action is set, since this is a dummy call
	// we are simply re-purposing the bucketPolicyAction.
	if s3Error := checkRequestAuthType(ctx, r, policy.GetBucketPolicyAction, bucket, ""); s3Error != ErrNone {
		writeErrorResponse(w, s3Error, r.URL)
		return
	}

	// Before proceeding validate if object exists.
	_, err := objAPI.GetObjectInfo(ctx, bucket, object)
	if err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	/*
		acl := &ACL.AccessControlPolicy{}
		acl.AccessControlList.Grants = append(acl.AccessControlList.Grants, ACL.Grant{
			Grantee: ACL.Grantee{
				XMLNS:  "http://www.w3.org/2001/XMLSchema-instance",
				XMLXSI: "CanonicalUser",
				Type:   "CanonicalUser",
			},
			Permission: "FULL_CONTROL",
		})
	*/
	acl, err := objAPI.GetObjectAccessControlPolicy(ctx, bucket, object)
	if err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	if err := xml.NewEncoder(w).Encode(acl); err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	w.(http.Flusher).Flush()
}
