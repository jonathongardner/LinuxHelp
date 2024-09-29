package github

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"path"
	"strconv"
	"strings"
)

// https://docs.github.com/en/rest/users/keys?apiVersion=2022-11-28
type key struct {
	ID          int    `json:"id"`
	Key         string `json:"key"`
	Fingerprint string `json:"fingerprint"`
	api         API
}

func (k key) GetID() string {
	return strconv.Itoa(k.ID)
}
func (k key) GetKey() string {
	return k.Key
}
func (k key) String() string {
	return fmt.Sprintf("# sha256:%v imported from github\n%v %v-%v@%v", k.Fingerprint, k.Key, k.api.username, k.ID, k.api.host)
}

func (k *key) parse() (err error) {
	k.Fingerprint, err = fingerprint(k.Key)
	return
}

func fingerprint(key string) (string, error) {
	code := strings.Split(key, " ")
	if len(code) <= 1 {
		return "", fmt.Errorf("error splinting key")
	}
	data, err := base64.StdEncoding.DecodeString(code[1])
	if err != nil {
		return "", fmt.Errorf("error base64 decoding %v", err)
	}
	shaBits := sha256.Sum256(data)
	return base64.StdEncoding.EncodeToString(shaBits[:]), nil
}

func (a API) SSHKeys() ([]*key, error) {
	var toReturn []*key
	err := a.get(path.Join("users", a.username, "keys"), &toReturn)

	for _, k := range toReturn {
		if err := k.parse(); err != nil {
			return []*key{}, fmt.Errorf("error parsing a result - %v (%v)", k.ID, err)
		}
	}
	return toReturn, err
}

func (a API) SSHSave(key string, title string) (string, error) {
	fp, err := fingerprint(key)
	if err != nil {
		return fp, err
	}

	toSend := map[string]string{"key": key}
	if title != "" {
		toSend["title"] = title
	}
	body, err := json.Marshal(toSend)
	if err != nil {
		return fp, fmt.Errorf("error marshaling body (%v)", err)
	}
	return fp, a.post(path.Join("user", "keys"), body)
}
