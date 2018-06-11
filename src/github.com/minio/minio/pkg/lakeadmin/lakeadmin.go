package lakeadmin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LakeAdmin struct {
	Endpoint string
}

func NewLakeAdmin(endpoint string) *LakeAdmin {
	return &LakeAdmin{Endpoint: endpoint}
}

func (la *LakeAdmin) GetSecretKey(ak string) (string, error) {
	resp, err := http.Get(la.Endpoint + "/v1/credentials/" + ak)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type Cred struct {
		Uid       string `json:"uid"`
		AccessKey string `json:"access_key"`
		SecretKey string `json:"secret_key"`
	}
	var cred Cred

	err = json.Unmarshal(data, &cred)
	if err != nil {
		return "", err
	}
	fmt.Printf("lakeadmin => AK: %s, SK: %s\n", ak, cred.SecretKey)

	return cred.SecretKey, nil
}
