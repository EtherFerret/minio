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
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

type RestoreRequest struct {
	Days                int
	GlacierJobParameter struct {
		Tier string
	}
}

type Filter struct {
	Prefix string `xml:"Prefix"`
}

type Transition struct {
	Date         time.Time //`xml:"Date,omitempty"`
	Days         int       //`xml:"Days,omitempty"`
	StorageClass string    `xml:"StorageClass"`
}

type Expiration struct {
	Date time.Time `xml:"Date,omitempty"`
	Days int       `xml:"Days,omitempty"`
}

type Rule struct {
	ID string `xml:"ID"`
	//Filter     Filter     `xml:"Filter"`
	Prefix     string     `xml:"Prefix,omitempty"`
	Status     string     `xml:"Status"`
	Transition Transition `xml:"Transition,omitempty"`
	Expiration Expiration `xml:"Expiration,omitempty"`
}

//type LifecycleConfiguration struct {
type Lifecycle struct {
	XMLNS string `xml:"xmlns,attr"`
	Rule  []Rule `xml:"Rule"`
}

// Lifecycle - bucket lifecycle.
/*
type Lifecycle struct {
	Configuration LifecycleConfiguration `xml:"LifecycleConfiguration"`
}
*/

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

	b, _ := ioutil.ReadAll(reader)
	fmt.Printf("lifecycle xml: %s\n", b)

	decoder := xml.NewDecoder(bytes.NewBuffer(b))
	//bytes.NewReader(b))

	err := decoder.Decode(&lifecycle)

	fmt.Printf("lifecycle %v\n", lifecycle)
	if err != nil {
		return nil, err
	}

	err = lifecycle.Validate(bucketName)
	return &lifecycle, err
}

// Parse RestoreRequest
func ParseRestoreRequest(reader io.Reader) (*RestoreRequest, error) {
	var restoreRequest RestoreRequest

	decoder := xml.NewDecoder(reader)
	err := decoder.Decode(&restoreRequest)

	if err != nil {
		return nil, err
	}

	return &restoreRequest, err
}
