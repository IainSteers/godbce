package dbcego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
	headerLink "github.com/tent/http-link-go"
)

const (
	libraryVersion = "0.1.0"
	defaultBaseURL = "https://api.cloud.exchange/"
	userAgent      = "dbcego/" + libraryVersion
	mediaType      = "application/json"

	headerApiKey = "DBCE-ApiKey"
)

// Client handles taking to the cloud exchange
type Client struct {
	client *http.Client

	BaseURL *url.URL

	UserAgent string

	Machines   MachinesService
	Capacities CapacitiesService
	Networks   NetworksService

	onRequestCompleted RequestCompletionCallback
}

type RequestCompletionCallback func(*http.Request, *http.Response)

type Response struct {
	*http.Response
}

type ErrorResponse struct {
	Response *http.Response
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()

	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		origValues[k] = v
	}

	origURL.RawQuery = origValues.Encode()
	return origURL.String(), nil
}

// NewClient returns a new DigitalOcean API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.Account = &AccountServiceOp{client: c}
	c.Actions = &ActionsServiceOp{client: c}
	c.Domains = &DomainsServiceOp{client: c}
	c.Droplets = &DropletsServiceOp{client: c}
	c.DropletActions = &DropletActionsServiceOp{client: c}
	c.Images = &ImagesServiceOp{client: c}
	c.ImageActions = &ImageActionsServiceOp{client: c}
	c.Keys = &KeysServiceOp{client: c}
	c.Regions = &RegionsServiceOp{client: c}
	c.Sizes = &SizesServiceOp{client: c}
	c.FloatingIPs = &FloatingIPsServiceOp{client: c}
	c.FloatingIPActions = &FloatingIPActionsServiceOp{client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("DBCE-ApiKey", apiKey)
	return req, nil
}

// OnRequestCompleted sets the DO API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

// newResponse creates a new Response for the provided http.Response
func newResponse(r *http.Response) *Response {
	response := Response{Response: r}
	response.populateRate()

	return &response
}
