// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudsearchdomain provides a client for Amazon CloudSearch Domain.
package cloudsearchdomain

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

// CloudSearchDomain is a client for Amazon CloudSearch Domain.
type CloudSearchDomain struct {
	client *aws.RestClient
}

// New returns a new CloudSearchDomain client.
func New(creds aws.Credentials, region string, client *http.Client) *CloudSearchDomain {
	if client == nil {
		client = http.DefaultClient
	}

	service := "cloudsearchdomain"
	endpoint, service, region := endpoints.Lookup("cloudsearchdomain", region)

	return &CloudSearchDomain{
		client: &aws.RestClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2013-01-01",
		},
	}
}

// Search retrieves a list of documents that match the specified search
// criteria. How you specify the search criteria depends on which query
// parser you use. Amazon CloudSearch supports four query parsers: simple :
// search all text and text-array fields for the specified string. Search
// for phrases, individual terms, and prefixes. structured : search
// specific fields, construct compound queries using Boolean operators, and
// use advanced features such as term boosting and proximity searching.
// lucene : specify search criteria using the Apache Lucene query parser
// syntax. dismax : specify search criteria using the simplified subset of
// the Apache Lucene query parser syntax defined by the DisMax query
// parser. For more information, see Searching Your Data in the Amazon
// CloudSearch Developer Guide The endpoint for submitting Search requests
// is domain-specific. You submit search requests to a domain's search
// endpoint. To get the search endpoint for your domain, use the Amazon
// CloudSearch configuration service DescribeDomains action. A domain's
// endpoints are also displayed on the domain dashboard in the Amazon
// CloudSearch console.
func (c *CloudSearchDomain) Search(req *SearchRequest) (resp *SearchResponse, err error) {
	resp = &SearchResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2013-01-01/search?format=sdk&pretty=true"

	q := url.Values{}

	if req.Cursor != nil {
		q.Set("cursor", *req.Cursor)
	}

	if req.Expr != nil {
		q.Set("expr", *req.Expr)
	}

	if req.Facet != nil {
		q.Set("facet", *req.Facet)
	}

	if req.FilterQuery != nil {
		q.Set("fq", *req.FilterQuery)
	}

	if req.Highlight != nil {
		q.Set("highlight", *req.Highlight)
	}

	if req.Partial != nil {
		q.Set("partial", fmt.Sprintf("%v", req.Partial))
	}

	if req.Query != nil {
		q.Set("q", *req.Query)
	}

	if req.QueryOptions != nil {
		q.Set("q.options", *req.QueryOptions)
	}

	if req.QueryParser != nil {
		q.Set("q.parser", *req.QueryParser)
	}

	if req.Return != nil {
		q.Set("return", *req.Return)
	}

	if req.Size != nil {
		q.Set("size", fmt.Sprintf("%v", req.Size))
	}

	if req.Sort != nil {
		q.Set("sort", *req.Sort)
	}

	if req.Start != nil {
		q.Set("start", fmt.Sprintf("%v", req.Start))
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// Suggest retrieves autocomplete suggestions for a partial query string.
// You can use suggestions enable you to display likely matches before
// users finish typing. In Amazon CloudSearch, suggestions are based on the
// contents of a particular text field. When you request suggestions,
// Amazon CloudSearch finds all of the documents whose values in the
// suggester field start with the specified query string. The beginning of
// the field must match the query string to be considered a match. For more
// information about configuring suggesters and retrieving suggestions, see
// Getting Suggestions in the Amazon CloudSearch Developer Guide . The
// endpoint for submitting Suggest requests is domain-specific. You submit
// suggest requests to a domain's search endpoint. To get the search
// endpoint for your domain, use the Amazon CloudSearch configuration
// service DescribeDomains action. A domain's endpoints are also displayed
// on the domain dashboard in the Amazon CloudSearch console.
func (c *CloudSearchDomain) Suggest(req *SuggestRequest) (resp *SuggestResponse, err error) {
	resp = &SuggestResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2013-01-01/suggest?format=sdk&pretty=true"

	q := url.Values{}

	if req.Query != nil {
		q.Set("q", *req.Query)
	}

	if req.Size != nil {
		q.Set("size", fmt.Sprintf("%v", req.Size))
	}

	if req.Suggester != nil {
		q.Set("suggester", *req.Suggester)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// UploadDocuments posts a batch of documents to a search domain for
// indexing. A document batch is a collection of add and delete operations
// that represent the documents you want to add, update, or delete from
// your domain. Batches can be described in either or Each item that you
// want Amazon CloudSearch to return as a search result (such as a product)
// is represented as a document. Every document has a unique ID and one or
// more fields that contain the data that you want to search and return in
// results. Individual documents cannot contain more than 1 MB of data. The
// entire batch cannot exceed 5 MB. To get the best possible upload
// performance, group add and delete operations in batches that are close
// the 5 MB limit. Submitting a large volume of single-document batches can
// overload a domain's document service. The endpoint for submitting
// UploadDocuments requests is domain-specific. To get the document
// endpoint for your domain, use the Amazon CloudSearch configuration
// service DescribeDomains action. A domain's endpoints are also displayed
// on the domain dashboard in the Amazon CloudSearch console. For more
// information about formatting your data for Amazon CloudSearch, see
// Preparing Your Data in the Amazon CloudSearch Developer Guide . For more
// information about uploading data for indexing, see Uploading Data in the
// Amazon CloudSearch Developer Guide .
func (c *CloudSearchDomain) UploadDocuments(req *UploadDocumentsRequest) (resp *UploadDocumentsResponse, err error) {
	resp = &UploadDocumentsResponse{}

	var body io.Reader
	var contentType string

	contentType = "application/json"
	b, err := json.Marshal(req.Documents)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2013-01-01/documents/batch?format=sdk"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	if req.ContentType != nil {
		httpReq.Header.Set("Content-Type", *req.ContentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// Bucket is undocumented.
type Bucket struct {
	Count aws.LongValue   `json:"count,omitempty"`
	Value aws.StringValue `json:"value,omitempty"`
}

// BucketInfo is undocumented.
type BucketInfo struct {
	Buckets []Bucket `json:"buckets,omitempty"`
}

// DocumentServiceWarning is undocumented.
type DocumentServiceWarning struct {
	Message aws.StringValue `json:"message,omitempty"`
}

// Hit is undocumented.
type Hit struct {
	Fields     map[string][]string `json:"fields,omitempty"`
	Highlights map[string]string   `json:"highlights,omitempty"`
	ID         aws.StringValue     `json:"id,omitempty"`
}

// Hits is undocumented.
type Hits struct {
	Cursor aws.StringValue `json:"cursor,omitempty"`
	Found  aws.LongValue   `json:"found,omitempty"`
	Hit    []Hit           `json:"hit,omitempty"`
	Start  aws.LongValue   `json:"start,omitempty"`
}

// SearchRequest is undocumented.
type SearchRequest struct {
	Cursor       aws.StringValue  `json:"cursor,omitempty"`
	Expr         aws.StringValue  `json:"expr,omitempty"`
	Facet        aws.StringValue  `json:"facet,omitempty"`
	FilterQuery  aws.StringValue  `json:"filterQuery,omitempty"`
	Highlight    aws.StringValue  `json:"highlight,omitempty"`
	Partial      aws.BooleanValue `json:"partial,omitempty"`
	Query        aws.StringValue  `json:"query"`
	QueryOptions aws.StringValue  `json:"queryOptions,omitempty"`
	QueryParser  aws.StringValue  `json:"queryParser,omitempty"`
	Return       aws.StringValue  `json:"return,omitempty"`
	Size         aws.LongValue    `json:"size,omitempty"`
	Sort         aws.StringValue  `json:"sort,omitempty"`
	Start        aws.LongValue    `json:"start,omitempty"`
}

// SearchResponse is undocumented.
type SearchResponse struct {
	Facets map[string]BucketInfo `json:"facets,omitempty"`
	Hits   *Hits                 `json:"hits,omitempty"`
	Status *SearchStatus         `json:"status,omitempty"`
}

// SearchStatus is undocumented.
type SearchStatus struct {
	Rid    aws.StringValue `json:"rid,omitempty"`
	Timems aws.LongValue   `json:"timems,omitempty"`
}

// SuggestModel is undocumented.
type SuggestModel struct {
	Found       aws.LongValue     `json:"found,omitempty"`
	Query       aws.StringValue   `json:"query,omitempty"`
	Suggestions []SuggestionMatch `json:"suggestions,omitempty"`
}

// SuggestRequest is undocumented.
type SuggestRequest struct {
	Query     aws.StringValue `json:"query"`
	Size      aws.LongValue   `json:"size,omitempty"`
	Suggester aws.StringValue `json:"suggester"`
}

// SuggestResponse is undocumented.
type SuggestResponse struct {
	Status  *SuggestStatus `json:"status,omitempty"`
	Suggest *SuggestModel  `json:"suggest,omitempty"`
}

// SuggestStatus is undocumented.
type SuggestStatus struct {
	Rid    aws.StringValue `json:"rid,omitempty"`
	Timems aws.LongValue   `json:"timems,omitempty"`
}

// SuggestionMatch is undocumented.
type SuggestionMatch struct {
	ID         aws.StringValue `json:"id,omitempty"`
	Score      aws.LongValue   `json:"score,omitempty"`
	Suggestion aws.StringValue `json:"suggestion,omitempty"`
}

// UploadDocumentsRequest is undocumented.
type UploadDocumentsRequest struct {
	ContentType aws.StringValue `json:"contentType"`
	Documents   []byte          `json:"documents"`
}

// UploadDocumentsResponse is undocumented.
type UploadDocumentsResponse struct {
	Adds     aws.LongValue            `json:"adds,omitempty"`
	Deletes  aws.LongValue            `json:"deletes,omitempty"`
	Status   aws.StringValue          `json:"status,omitempty"`
	Warnings []DocumentServiceWarning `json:"warnings,omitempty"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name

var _ bytes.Reader
var _ url.URL
var _ fmt.Stringer
var _ strings.Reader
var _ strconv.NumError
var _ = ioutil.Discard
var _ json.RawMessage