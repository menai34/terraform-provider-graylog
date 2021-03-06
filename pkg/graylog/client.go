package graylog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Client struct {
	Username    string
	Password    string
	ServerURL   string
	Credentials string
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Client) do(method string, url string, reqValue, respValue interface{}) error {
	url = fmt.Sprintf("%s/%s",
		strings.TrimRight(c.ServerURL, "/"),
		strings.TrimLeft(url, "/"),
	)

	log.Printf("[DEBUG] %s %s", method, url)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	if c.Credentials != "" {
		credentialPath := ""
		if strings.HasPrefix(c.Credentials, "~") {
			home := os.Getenv("HOME")
			credentialPath = home + c.Credentials[1:]
		} else {
			credentialPath = c.Credentials
		}

		jsonFile, err := os.Open(credentialPath)
		if err != nil {
			return fmt.Errorf("error opening file %s", err)
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var credential Credential
		json.Unmarshal(byteValue, &credential)

		defer jsonFile.Close()
		c.Username = credential.Username
		c.Password = credential.Password
	}

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("X-Requested-By", "github.com/rebuy-de/terraform-provider-graylog")

	if reqValue != nil {
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(reqValue)
		if err != nil {
			return err
		}

		log.Printf("[DEBUG] RequestValue: %s", buf.String())
		req.Body = ioutil.NopCloser(buf)
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		//log.Debug(ReaderToString(resp.Body))
		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	if respValue != nil {
		jsonOutputs, _ := json.MarshalIndent(resp.Body, "", "  ")
		log.Printf("[DEBUG] Detail: %s", jsonOutputs)
		return json.NewDecoder(resp.Body).Decode(respValue)
	}

	return nil
}

func (c *Client) Get(url string, v interface{}) error {
	return c.do(GET, url, nil, v)
}

func (c *Client) Post(url string, reqValue, respValue interface{}) error {
	return c.do(POST, url, reqValue, respValue)
}

func (c *Client) Put(url string, reqValue, respValue interface{}) error {
	return c.do(PUT, url, reqValue, respValue)
}

func (c *Client) Delete(url string) error {
	return c.do(DELETE, url, nil, nil)
}
