package check

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"enkodr.dev/status/config"
)

// HTTP represents the configuration of the endpoint to check
type HTTP struct{}

// Check will test the endpoint
func (h HTTP) Check(config config.Check) (bool, error) {
	var resp *http.Response
	var err error
	// Execute the get request
	if config.FollowRedirects {
		resp, err = http.Get(config.URL)
		if err != nil {
			return false, err
		}
	} else {
		// create a custom error to know if a redirect happened
		var RedirectAttemptedError = errors.New("redirect")

		client := &http.Client{}
		// return the error, so client won't attempt redirects
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return RedirectAttemptedError

		}
		// Work with the client...
		resp, err = client.Head(config.URL)

		// test if we got the custom error
		if urlError, ok := err.(*url.Error); ok && urlError.Err == RedirectAttemptedError {
			err = nil
		} else {
			return false, err
		}
	}

	// Check if the response code is the expected
	if config.StatusCode != resp.StatusCode {
		err = fmt.Errorf("expected response '%v' code on '%v' but received '%v'", config.StatusCode, config.URL, resp.StatusCode)
		return false, err
	}
	return true, nil

}
