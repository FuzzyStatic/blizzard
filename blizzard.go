/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:40:31
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 11:40:24
 */

// Package blizzard is the top level library needed to use the
// API calls for Blizzard games within the subpackages.
package blizzard

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// Auth access token, api key
type Auth struct {
	AccessToken string
	APIKey      string
}

// JSON interface for struture creation
type JSON interface {
	JSON2Struct(*[]byte) error
}

// GetStruct returns structure of JSON interface
func GetStruct(b *[]byte, v JSON) error {
	return v.JSON2Struct(b)
}

// GetURLBody fills body of url
func GetURLBody(url string, body *[]byte) error {
	var (
		req *http.Request
		res *http.Response
		err error
	)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.New(err.Error())
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return errors.New(err.Error())
	}

	defer res.Body.Close()
	*body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New(err.Error())
	}

	switch res.StatusCode {
	case http.StatusNotFound:
		return errors.New(res.Status)
	}

	return nil
}
