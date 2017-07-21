/****** Author : NAGA SAI AAKARSHIT BATCHU ******/

package cluster_eslibrary

import (
	"encoding/json"
	"net/http"
	"sort"
)

type Cluster_Stats struct {
	UnderScoreNodes map[string]interface{} `json:"_nodes"`
	ClusterName     string                 `json:"cluster_name"`
	TimeStamp       int64                  `json:"timestamp"`
	Status          string                 `json:"status"`
	Indices         map[string]interface{} `json:"indices"`
	Nodes           map[string]interface{} `json:"nodes"`
}

type Client struct {
	Service    string
	HttpClient *http.Client
}

func get(url string) (*Cluster_Stats, error) {
	resp, respErr := http.Get(url)
	if respErr != nil {
		return nil, respErr
	}
	defer resp.Body.Close()
	var elasticresp Cluster_Stats
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&elasticresp); err != nil {
		return nil, err
	}
	return &elasticresp, nil
}

func GetNodes(service string) ([]string, error) {
	url := service + "/_cluster/stats"
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(resp.Nodes))
	for key, _ := range resp.Nodes {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret, nil
}

func GetNodeValues(service, node string) (interface{}, error) {
	url := service + "/_cluster/stats"
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	response := resp.Nodes
	var mapresponse = response[node].(interface{})
	return mapresponse, nil
}

func GetIndices(service string) ([]string, error) {
	url := service + "/_cluster/stats"
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(resp.Indices))
	for key, _ := range resp.Indices {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret, nil
}

func GetIndexValues(service, index string) (interface{}, error) {
	url := service + "/_cluster/stats"
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	response := resp.Indices
	var mapresponse = response[index].(interface{})
	return mapresponse, nil
}

func GetClusterName(service string) string {
	url := service + "/_cluster/stats"
	resp, err := get(url)
	if err != nil {
		return ""
	}
	cluster_name := resp.ClusterName
	return cluster_name
}

func GetTimeStamp(service string) int64 {
	url := service + "/_cluster/stats"
	resp, err := get(url)
	if err != nil {
		return 0
	}
	timestamp := resp.TimeStamp
	return timestamp
}

func GetStatus(service string) string {
	url := service + "/_cluster/stats"
	resp, err := get(url)
	if err != nil {
		return ""
	}
	status := resp.Status
	return status
}

func NewClient(host string) *Client {
	return &Client{
		"http://" + host,
		&http.Client{},
	}
}

func (c *Client) get(url string) (*Cluster_Stats, error) {
	resp, respErr := c.HttpClient.Get(c.Service + url)
	if respErr != nil {
		return nil, respErr
	}
	defer resp.Body.Close()
	var elasticresp Cluster_Stats
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&elasticresp); err != nil {
		return nil, err
	}
	return &elasticresp, nil
}

func (c *Client) GetNodes() ([]string, error) {
	url := "/_cluster/stats"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(resp.Nodes))
	for key, _ := range resp.Nodes {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret, nil
}

func (c *Client) GetNodeValues(node string) (interface{}, error) {
	url := "/_cluster/stats"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	response := resp.Nodes
	var mapresponse = response[node].(interface{})
	return mapresponse, nil
}

func (c *Client) GetIndices() ([]string, error) {
	url := "/_cluster/stats"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(resp.Indices))
	for key, _ := range resp.Indices {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret, nil
}

func (c *Client) GetIndexValues(index string) (interface{}, error) {
	url := "/_cluster/stats"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	response := resp.Indices
	var mapresponse = response[index].(interface{})
	return mapresponse, nil
}

func (c *Client) GetClusterName() string {
	url := "/_cluster/stats"
	resp, err := c.get(url)
	if err != nil {
		return ""
	}
	cluster_name := resp.ClusterName
	return string(cluster_name)
}

func (c *Client) GetTimeStamp() int64 {
	url := "/_cluster/stats"
	resp, err := c.get(url)
	if err != nil {
		return 0
	}
	timestamp := resp.TimeStamp
	return timestamp
}

func (c *Client) GetStatus() string {
	url := "/_cluster/stats"
	resp, err := c.get(url)
	if err != nil {
		return ""
	}
	status := resp.Status
	return string(status)
}

/****** Author : NAGA SAI AAKARSHIT BATCHU ******/
