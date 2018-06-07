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

package lifecycle

import (
	"encoding/xml"
	//"fmt"
	"io"
)

type Filter struct {
	Prefix string `xml:"Prefix"`
}

type Transition struct {
	Days         int    `xml:"Days"`
	StorageClass string `xml:"StorageClass"`
}

type Expiration struct {
	Days int `xml:"Days"`
}

type Rule struct {
	ID         string     `xml:"ID"`
	Filter     Filter     `xml:"Filter"`
	Status     string     `xml:"Status"`
	Transition Transition `xml:"Transition"`
	Expiration Expiration `xml:"Expiration"`
}

type LifecycleConfiguration struct {
	Rule []Rule `xml:"Rule"`
}

// Lifecycle - bucket lifecycle.
type Lifecycle struct {
	Configuration LifecycleConfiguration `xml:"LifecycleConfiguration"`
}

// isValid - checks if lifecycle is valid or not.
func (lifecycle Lifecycle) isValid() error {

	return nil
}

// Validate - validates all rules are for given bucket or not.
func (lifecycle Lifecycle) Validate(bucketName string) error {
	return nil
}

// ParseConfig - parses data in given reader to Lifecycle
func ParseConfig(reader io.Reader, bucketName string) (*Lifecycle, error) {
	var lifecycle Lifecycle

	decoder := xml.NewDecoder(reader)
	err := decoder.Decode(&lifecycle)

	if err != nil {
		return nil, err
	}

	err = lifecycle.Validate(bucketName)
	return &lifecycle, err
}
