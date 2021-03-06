package harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"errors"
)

type APIClient struct {
	username   string
	password   string
	subdomain  string
	httpClient *http.Client

	Client  *ClientService
	People  *PersonService
	Project *ProjectService
	Invoice *InvoiceService
	Account *AccountService
	Timesheet *TimesheetService
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func newAPIClient(subdomain string, httpClient *http.Client) (c *APIClient) {
	c = new(APIClient)
	c.subdomain = subdomain

	if httpClient != nil {
		c.httpClient = httpClient
	} else {
		c.httpClient = new(http.Client)
	}

	c.Client = &ClientService{Service{c}}
	c.People = &PersonService{Service{c}}
	c.Project = &ProjectService{Service{c}}
	c.Invoice = &InvoiceService{Service{c}}
	c.Account = &AccountService{Service{c}}
	c.Timesheet = &TimesheetService{Service{c}}
	return
}

func NewAPIClientWithBasicAuth(username, password, subdomain string, client *http.Client) (c *APIClient) {
	c = newAPIClient(subdomain, client)
	c.username = username
	c.password = password
	return
}

func (c *APIClient) GetJSON(path string) (jsonResponse []byte, err error) {
	resourceURL := fmt.Sprintf("https://%v.harvestapp.com%v", c.subdomain, path)
	request, err := http.NewRequest("GET", resourceURL, nil)
	if err != nil {
		return
	}

	request.SetBasicAuth(c.username, c.password)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	resp, err := c.httpClient.Do(request)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	jsonResponse, err = ioutil.ReadAll(resp.Body)

	var errorJson ErrorMessage
	err = json.Unmarshal(jsonResponse, &(errorJson))

	if err != nil{
		return
	}

	if errorJson.Message != ""{
		return jsonResponse, errors.New("Error when calling harvest API: " + errorJson.Message)
	}
	return
}
