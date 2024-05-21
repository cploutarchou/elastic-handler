package elastic

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	once      sync.Once
	client    *clientImpl
	initError error
)

// Client interface defines methods for async and sync operations.
type Client interface {
	Index(index string, document interface{}) (*esapi.Response, error)
	Search(index string, query map[string]interface{}) (*esapi.Response, error)
	Create(index string, id string, document interface{}) (*esapi.Response, error)
	Upsert(index string, id string, document interface{}) (*esapi.Response, error)
	Update(index string, id string, document interface{}) (*esapi.Response, error)
}

// clientImpl is a concrete implementation of the Client interface.
type clientImpl struct {
	*elasticsearch.Client
}

// New initializes the clientImpl instance.
func New(host string, port string, username string, password string, isSecure bool) {
	once.Do(func() {
		var url string
		if isSecure {
			url = fmt.Sprintf("https://%s:%s", host, port)
		} else {
			url = fmt.Sprintf("http://%s:%s", host, port)
		}

		cfg := elasticsearch.Config{
			Addresses: []string{url},
			Username:  username,
			Password:  password,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   10,
				ResponseHeaderTimeout: time.Second * 30,
				DialContext:           (&net.Dialer{Timeout: time.Second * 30}).DialContext,
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
		}
		esClient, err := elasticsearch.NewClient(cfg)
		if err != nil {
			initError = err
			return
		}
		client = &clientImpl{
			esClient,
		}
	})
}

// GetClient returns the singleton client instance.
func GetClient() (Client, error) {
	if initError != nil {
		return nil, initError
	}
	if client == nil {
		return nil, fmt.Errorf("client not initialized")
	}
	return client, nil
}

// Helper method to convert an interface{} to io.Reader
func toJSONReader(v interface{}) (io.Reader, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

// Index indexes a document in the specified index.
func (c *clientImpl) Index(index string, document interface{}) (*esapi.Response, error) {
	body, err := toJSONReader(document)
	if err != nil {
		return nil, err
	}
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: "", // Let Elasticsearch generate an ID
		Body:       body,
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Search performs a search query on the specified index.
func (c *clientImpl) Search(index string, query map[string]interface{}) (*esapi.Response, error) {
	body, err := toJSONReader(query)
	if err != nil {
		return nil, err
	}
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  body,
	}
	res, err := req.Do(context.Background(), c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create creates a new document in the specified index with the given ID.
func (c *clientImpl) Create(index string, id string, document interface{}) (*esapi.Response, error) {
	body, err := toJSONReader(document)
	if err != nil {
		return nil, err
	}
	req := esapi.CreateRequest{
		Index:      index,
		DocumentID: id,
		Body:       body,
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Upsert updates a document in the specified index if it exists, or creates it if it does not.
func (c *clientImpl) Upsert(index string, id string, document interface{}) (*esapi.Response, error) {
	body, err := toJSONReader(map[string]interface{}{
		"doc":           document,
		"doc_as_upsert": true,
	})
	if err != nil {
		return nil, err
	}
	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: id,
		Body:       body,
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Update updates a document in the specified index with the given ID.
func (c *clientImpl) Update(index string, id string, document interface{}) (*esapi.Response, error) {
	body, err := toJSONReader(map[string]interface{}{
		"doc": document,
	})
	if err != nil {
		return nil, err
	}
	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: id,
		Body:       body,
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), c)
	if err != nil {
		return nil, err
	}
	return res, nil
}
