/*
 * Minio Go Library for Amazon S3 Compatible Cloud Storage
 * Copyright 2015-2017 Minio, Inc.
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

package minio

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/minio/minio-go/pkg/s3utils"
)

func (c Client) SetBucketAccessControlPolicy(bucketName string, aclPolicy string) error {
	// Input validation.
	if err := s3utils.CheckValidBucketName(bucketName); err != nil {
		return err
	}

	// Get resources properly escaped and lined up before
	// using them in http request.
	urlValues := make(url.Values)
	urlValues.Set("acl", "")

	// Content-length is mandatory for put policy request
	policyReader := strings.NewReader(aclPolicy)
	b, err := ioutil.ReadAll(policyReader)
	if err != nil {
		return err
	}

	reqMetadata := requestMetadata{
		bucketName:    bucketName,
		queryValues:   urlValues,
		contentBody:   policyReader,
		contentLength: int64(len(b)),
	}

	// Execute PUT to upload a new bucket policy.
	resp, err := c.executeMethod(context.Background(), "PUT", reqMetadata)
	defer closeResponse(resp)
	if err != nil {
		return err
	}
	if resp != nil {
		if resp.StatusCode != http.StatusNoContent {
			return httpRespToErrorResponse(resp, bucketName, "")
		}
	}
	return nil
}

func (c Client) SetObjectAccessControlPolicy(bucketName, objectName string, aclPolicy string) error {
	// Input validation.
	if err := s3utils.CheckValidBucketName(bucketName); err != nil {
		return err
	}

	// Get resources properly escaped and lined up before
	// using them in http request.
	urlValues := make(url.Values)
	urlValues.Set("acl", "")

	// Content-length is mandatory for put policy request
	policyReader := strings.NewReader(aclPolicy)
	b, err := ioutil.ReadAll(policyReader)
	if err != nil {
		return err
	}

	reqMetadata := requestMetadata{
		bucketName:    bucketName,
		objectName:    objectName,
		queryValues:   urlValues,
		contentBody:   policyReader,
		contentLength: int64(len(b)),
	}

	// Execute PUT to upload a new bucket policy.
	resp, err := c.executeMethod(context.Background(), "PUT", reqMetadata)
	defer closeResponse(resp)
	if err != nil {
		return err
	}
	if resp != nil {
		if resp.StatusCode != http.StatusNoContent {
			return httpRespToErrorResponse(resp, bucketName, "")
		}
	}
	return nil
}

// GetBucketAccessControlPolicy - get bucket policy at a given path.
func (c Client) GetBucketAccessControlPolicy(bucketName string) (string, error) {
	// Input validation.
	if err := s3utils.CheckValidBucketName(bucketName); err != nil {
		return "", err
	}
	bucketPolicy, err := c.getBucketAccessControlPolicy(bucketName)
	if err != nil {
		errResponse := ToErrorResponse(err)
		if errResponse.Code == "NoSuchBucketAccessControlPolicy" {
			return "", nil
		}
		return "", err
	}
	return bucketPolicy, nil
}

// Request server for current bucket policy.
func (c Client) getBucketAccessControlPolicy(bucketName string) (string, error) {
	// Get resources properly escaped and lined up before
	// using them in http request.
	urlValues := make(url.Values)
	urlValues.Set("acl", "")

	// Execute GET on bucket to list objects.
	resp, err := c.executeMethod(context.Background(), "GET", requestMetadata{
		bucketName:       bucketName,
		queryValues:      urlValues,
		contentSHA256Hex: emptySHA256Hex,
	})

	defer closeResponse(resp)
	if err != nil {
		return "", err
	}

	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			return "", httpRespToErrorResponse(resp, bucketName, "")
		}
	}

	bucketPolicyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	policy := string(bucketPolicyBuf)
	return policy, err
}

// GetObjectAccessControlPolicy - get object policy at a given path.
func (c Client) GetObjectAccessControlPolicy(bucketName, objectName string) (string, error) {
	// Input validation.
	if err := s3utils.CheckValidBucketName(bucketName); err != nil {
		return "", err
	}
	bucketPolicy, err := c.getObjectAccessControlPolicy(bucketName, objectName)
	if err != nil {
		errResponse := ToErrorResponse(err)
		if errResponse.Code == "NoSuchBucketAccessControlPolicy" {
			return "", nil
		}
		return "", err
	}
	return bucketPolicy, nil
}

// Request server for current object policy.
func (c Client) getObjectAccessControlPolicy(bucketName, objectName string) (string, error) {
	// Get resources properly escaped and lined up before
	// using them in http request.
	urlValues := make(url.Values)
	urlValues.Set("acl", "")

	// Execute GET on bucket to list objects.
	resp, err := c.executeMethod(context.Background(), "GET", requestMetadata{
		bucketName:       bucketName,
		objectName:       objectName,
		queryValues:      urlValues,
		contentSHA256Hex: emptySHA256Hex,
	})

	defer closeResponse(resp)
	if err != nil {
		return "", err
	}

	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			return "", httpRespToErrorResponse(resp, bucketName, "")
		}
	}

	bucketPolicyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	policy := string(bucketPolicyBuf)
	return policy, err
}
