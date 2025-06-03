package fritz

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// DoSoapRequest does two request to authenticate and handle the SOAP request
func DoSoapRequest(soapRequest *SoapData, resps chan<- []byte, errs chan<- error, debug bool) {
	soapClient := createNewSoapClient()

	// prepare first request
	req, err := newSoapRequest(soapRequest)

	if err != nil {
		errs <- err
		return
	}

	// the first request is always unauthenticated due to how digest authentication works
	resp, err := soapClient.Do(req)

	if err != nil {
		errs <- err
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		errs <- err
		return
	}
	resp.Body.Close()

	// enable this for debug sessions
	if debug {
		fmt.Printf("---\nFirst SOAP Response:\n---\n%v\n---\n", string(body))
	}

	// create immediately a new request with authentication
	req, err = newSoapRequest(soapRequest)

	if err != nil {
		errs <- err
		return
	}

	if resp.StatusCode == http.StatusUnauthorized {
		authHeader, err := doDigestAuthentication(resp, soapRequest)

		if err != nil {
			errs <- err
			return
		} else if authHeader == "" {
			errs <- errors.New("returned authentication header is empty")
			return
		}
		req.Header.Set("Authorization", authHeader)

	} else if resp.StatusCode == http.StatusOK {
		resps <- body
		return
	} else {
		errs <- fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	// second request is now authenticated
	resp, err = soapClient.Do(req)

	if err != nil {
		errs <- err
		return
	}

	if resp.StatusCode == http.StatusUnauthorized {
		errs <- errors.New("unauthorized: wrong username or password")
		return
	} else if resp.StatusCode != http.StatusOK {
		errs <- fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
		return
	}

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		errs <- err
		return
	}
	resp.Body.Close()

	// enable this for debug sessions

	if debug {
		fmt.Printf("---\nSecond SOAP Response:\n---\n%v\n---\n", string(body))
	}

	resps <- body
}

func createNewSoapClient() *http.Client {
	ht := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true},
	}

	return &http.Client{Transport: ht}
}

func newSoapRequest(soapRequest *SoapData) (*http.Request, error) {
	requestBody := newSoapRequestBody(soapRequest)
	req, err := http.NewRequest("POST", soapRequest.URL, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/xml; charset='utf-8'")
	req.Header.Set("SoapAction", "urn:dslforum-org:service:"+soapRequest.Service+":1#"+soapRequest.Action)

	return req, nil
}

func newSoapRequestBody(soapRequest *SoapData) []byte {
	var request bytes.Buffer

	request.WriteString("<?xml version='1.0?>")
	request.WriteString("<s:Envelope xmlns:s='http://schemas.xmlsoap.org/soap/envelope/' s:encodingStyle='http://schemas.xmlsoap.org/soap/encoding/'>")
	request.WriteString("<s:Body>")
	request.WriteString("<u:" + soapRequest.Action + " xmlns:u='urn:dslforum-org:service:" + soapRequest.Service + ":1'>")

	request.WriteString("<" + soapRequest.XMLVariable.Name + ">")
	request.WriteString(soapRequest.XMLVariable.Value)
	request.WriteString("</" + soapRequest.XMLVariable.Name + ">")

	request.WriteString("</u:" + soapRequest.Action + ">")
	request.WriteString("</s:Body>")
	request.WriteString("</s:Envelope>")

	return request.Bytes()
}
