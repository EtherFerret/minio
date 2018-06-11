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

package acl

import (
	"encoding/xml"
	"io"
)

// Data types used for returning dummy access control
// policy XML, these variables shouldn't be used elsewhere
// they are only defined to be used in this file alone.
type Grantee struct {
	XMLNS       string `xml:"xmlns:xsi,attr"`
	XMLXSI      string `xml:"xsi:type,attr"`
	Type        string `xml:"Type"`
	ID          string `xml:"ID,omitempty"`
	DisplayName string `xml:"DisplayName,omitempty"`
	URI         string `xml:"URI,omitempty"`
}

// Owner - bucket owner/principal
type Owner struct {
	ID          string
	DisplayName string
}

type Grant struct {
	Grantee    Grantee `xml:"Grantee"`
	Permission string  `xml:"Permission"`
}

type AccessControlPolicy struct {
	XMLName           xml.Name `xml:"AccessControlPolicy"`
	XMLNS             string   `xml:"xmlns,attr"`
	Owner             Owner    `xml:"Owner"`
	AccessControlList struct {
		Grants []Grant `xml:"Grant"`
	} `xml:"AccessControlList"`
}

// isValid - checks if Policy is valid or not.
func (acl AccessControlPolicy) isValid() error {

	return nil
}

// Validate - validates all rules are for given bucket or not.
func (acl AccessControlPolicy) Validate(bucketName string) error {
	return nil
}

// ParseConfig - parses data in given reader to AccessControlPolicy
func ParseConfig(reader io.Reader, bucketName string) (*AccessControlPolicy, error) {
	var acl AccessControlPolicy

	decoder := xml.NewDecoder(reader)
	err := decoder.Decode(&acl)

	if err != nil {
		return nil, err
	}

	err = acl.Validate(bucketName)

	// FIXME
	result := &AccessControlPolicy{Owner: acl.Owner, XMLNS: acl.XMLNS}
	for _, grant := range acl.AccessControlList.Grants {
		Type := "CanonicalUser"
		XMLNS := "http://www.w3.org/2001/XMLSchema-instance"

		if grant.Grantee.Type != "" {
			Type = grant.Grantee.Type
		} else if grant.Grantee.XMLXSI != "" {
			Type = grant.Grantee.XMLXSI
		}

		if grant.Grantee.XMLNS != "" {
			XMLNS = grant.Grantee.XMLNS
		}

		result.AccessControlList.Grants = append(result.AccessControlList.Grants, Grant{
			Grantee: Grantee{
				XMLNS:       XMLNS,
				XMLXSI:      Type,
				Type:        Type,
				ID:          grant.Grantee.ID,
				DisplayName: grant.Grantee.DisplayName,
			},
			Permission: grant.Permission,
		})
	}

	return result, err
}
