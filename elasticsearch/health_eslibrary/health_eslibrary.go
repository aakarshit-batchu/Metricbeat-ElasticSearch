/****** Author : NAGA SAI AAKARSHIT BATCHU ******/

package health_eslibrary

import (
	"encoding/json"
	"net/http"
)

type HealthStats struct {
	ClusterName                 string  `json:"cluster_name"`
	Status                      string  `json:"status"`
	TimedOut                    bool    `json:"timed_out"`
	NumberOfNodes               int     `json:"number_of_nodes"`
	NumberOfDataNodes           int     `json:"number_of_data_nodes"`
	ActivePrimaryShards         int     `json:"active_primary_shards"`
	ActiveShards                int     `json:"active_shards"`
	RelocatingShards            int     `json:"relocating_shards"`
	InitializingShards          int     `json:"initializing_shards"`
	UnassignedShards            int     `json:"unassigned_shards"`
	DelayedUnassignedShards     int     `json:"delayed_unassigned_shards"`
	NumberOfPendingTasks        int     `json:"number_of_pending_tasks"`
	NumberOfInFlightFetch       int     `json:"number_of_in_flight_fetch"`
	TaskMaxWaitingInQueueMillis int     `json:"task_max_waiting_in_queue_millis"`
	ActiveShardsPercentAsNumber float64 `json:"active_shards_percent_as_number"`
}

type Client struct {
	Service    string
	HttpClient *http.Client
}

func get(url string) (*HealthStats, error) {
	resp, respErr := http.Get(url)
	if respErr != nil {
		return nil, respErr
	}
	defer resp.Body.Close()
	var elasticresp HealthStats
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&elasticresp); err != nil {
		return nil, err
	}
	return &elasticresp, nil
}

func GetMetrics(service string) (map[string]interface{}, error) {
	url := service + "/_cluster/health"
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]interface{}, 15)
	ret["cluster_name"] = resp.ClusterName
	ret["status"] = resp.Status
	ret["timed_out"] = resp.TimedOut
	ret["number_of_nodes"] = resp.NumberOfNodes
	ret["number_of_data_nodes"] = resp.NumberOfDataNodes
	ret["active_primary_shards"] = resp.ActivePrimaryShards
	ret["active_shards"] = resp.ActiveShards
	ret["relocating_shards"] = resp.RelocatingShards
	ret["initializing_shards"] = resp.InitializingShards
	ret["unassigned_shards"] = resp.UnassignedShards
	ret["delayed_unassigned_shards"] = resp.DelayedUnassignedShards
	ret["number_of_pending_tasks"] = resp.NumberOfPendingTasks
	ret["number_of_in_flight_fetch"] = resp.NumberOfInFlightFetch
	ret["task_max_waiting_in_queue_millis"] = resp.TaskMaxWaitingInQueueMillis
	ret["active_shards_percent_as_number"] = resp.ActiveShardsPercentAsNumber
	return ret, nil
}

func NewClient(host string) *Client {
	return &Client{
		"http://" + host,
		&http.Client{},
	}
}

func (c *Client) get(url string) (*HealthStats, error) {
	resp, respErr := c.HttpClient.Get(c.Service + url)
	if respErr != nil {
		return nil, respErr
	}
	defer resp.Body.Close()
	var elasticresp HealthStats
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&elasticresp); err != nil {
		return nil, err
	}
	return &elasticresp, nil
}

func (c *Client) GetMetrics() (map[string]interface{}, error) {
	url := "/_cluster/health"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]interface{}, 15)
	ret["cluster_name"] = resp.ClusterName
	ret["status"] = resp.Status
	ret["timed_out"] = resp.TimedOut
	ret["number_of_nodes"] = resp.NumberOfNodes
	ret["number_of_data_nodes"] = resp.NumberOfDataNodes
	ret["active_primary_shards"] = resp.ActivePrimaryShards
	ret["active_shards"] = resp.ActiveShards
	ret["relocating_shards"] = resp.RelocatingShards
	ret["initializing_shards"] = resp.InitializingShards
	ret["unassigned_shards"] = resp.UnassignedShards
	ret["delayed_unassigned_shards"] = resp.DelayedUnassignedShards
	ret["number_of_pending_tasks"] = resp.NumberOfPendingTasks
	ret["number_of_in_flight_fetch"] = resp.NumberOfInFlightFetch
	ret["task_max_waiting_in_queue_millis"] = resp.TaskMaxWaitingInQueueMillis
	ret["active_shards_percent_as_number"] = resp.ActiveShardsPercentAsNumber
	return ret, nil
}

/****** Author : NAGA SAI AAKARSHIT BATCHU ******/
