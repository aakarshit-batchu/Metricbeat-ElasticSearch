/********* AUTHOR: NAGA SAI AAKARSHIT BATCHU ************/

package nodes_eslibrary

import (
	"encoding/json"
	"net/http"
	"sort"
)

type NodesStruct struct {
	_nodes       map[string]interface{} `json:"_nodes"`
	Cluster_Name string                 `json:"cluster_name"`
	Nodes        map[string]interface{} `json:"nodes"`
}

type Client struct {
	Service    string
	HttpClient *http.Client
}

func get(url string) (*NodesStruct, error) {
	resp, respErr := http.Get(url)
	if respErr != nil {
		return nil, respErr
	}
	defer resp.Body.Close()
	var elasticresp NodesStruct
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&elasticresp); err != nil {
		return nil, err
	}
	return &elasticresp, nil
}

func GetNodes(service string) ([]string, error) {
	resp, err := get(service)
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

func GetNodeResponse(service, node string) (map[string]interface{}, error) {
	url := service + "/" + node + "/stats"
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	response := resp.Nodes
	var mapresponse map[string]interface{} = response[node].(map[string]interface{})
	return mapresponse, nil
}

/*func ListMapKeysinNode(service, node string)([]string, error) {
        mapresponse, err := GetNodeResponse(service, node)
        if err != nil {
                return nil, err
        }
        ret := make([]string, 0, len(mapresponse))
        for key,value := range mapresponse {
                _,ok := value.(map[string]interface{})
                if ok {
                ret = append(ret,key)
                }
        }
        sort.Strings(ret)
        return ret, nil
}
*/

func GetValue(service, node, key string) (interface{}, error) {
	url := service + "/" + node + "/stats" + "/" + key
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	response := resp.Nodes
	if value, ok := response[node].(map[string]interface{}); ok {
		keyvalue := value[key]
		return keyvalue, nil
	} else {
		url := service + "/_nodes" + "/" + node + "/stats" + "/"
		resp, err := get(url)
		if err != nil {
			return nil, err
		}
		response := resp.Nodes
		value := response[node].(map[string]interface{})
		keyvalue := value[key]
		return keyvalue, nil
	}
}

func ListKeysinNode(service, node string) ([]string, error) {
	mapresponse, err := GetNodeResponse(service, node)
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(mapresponse))
	for key, _ := range mapresponse {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret, nil
}

func GetClusterName(service string) string {
	url := service + "/_nodes/stats"
	resp, err := get(url)
	if err != nil {
		return ""
	}
	cluster_name := resp.Cluster_Name
	return cluster_name
}

func NewClient(host string) *Client {
	return &Client{
		"http://" + host,
		&http.Client{},
	}
}

func (c *Client) get(url string) (*NodesStruct, error) {
	resp, respErr := c.HttpClient.Get(c.Service + url)
	if respErr != nil {
		return nil, respErr
	}
	defer resp.Body.Close()
	var elasticresp NodesStruct
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&elasticresp); err != nil {
		return nil, err
	}
	return &elasticresp, nil
}

func (c *Client) GetNodes() ([]string, error) {
	url := "/_nodes/stats"
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

func (c *Client) GetNodeResponse(node string) (map[string]interface{}, error) {
	url := "/_nodes" + "/" + node + "/stats"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	response := resp.Nodes
	var mapresponse map[string]interface{} = response[node].(map[string]interface{})
	return mapresponse, nil
}

/*func (c *Client) ListMapKeysinNode(node string)([]string, error) {
        mapresponse, err := c.GetNodeResponse(node)
        if err != nil {
                return nil, err
        }
        ret := make([]string, 0, len(mapresponse))
        for key,value := range mapresponse {
                _,ok := value.(map[string]interface{})
                if ok {
                ret = append(ret,key)
                }
        }
        sort.Strings(ret)
        return ret, nil
}
*/

func (c *Client) GetValue(node, key string) (interface{}, error) {
	url := "/_nodes" + "/" + node + "/stats" + "/" + key
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	response := resp.Nodes
	if value, ok := response[node].(map[string]interface{}); ok {
		keyvalue := value[key]
		return keyvalue, nil
	} else {
		url := "/_nodes" + "/" + node + "/stats" + "/"
		resp, err := c.get(url)
		if err != nil {
			return nil, err
		}
		response := resp.Nodes
		value := response[node].(map[string]interface{})
		keyvalue := value[key]
		return keyvalue, nil
	}
}

func (c *Client) ListKeysinNode(node string) ([]string, error) {
	mapresponse, err := c.GetNodeResponse(node)
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(mapresponse))
	for key, _ := range mapresponse {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret, nil
}

func (c *Client) GetClusterName() string {
	url := "/_nodes/stats"
	resp, err := c.get(url)
	if err != nil {
		return ""
	}
	cluster_name := resp.Cluster_Name
	return cluster_name
}

/********* AUTHOR: NAGA SAI AAKARSHIT BATCHU ************/
