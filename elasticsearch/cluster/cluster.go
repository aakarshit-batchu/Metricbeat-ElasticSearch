/********* AUTHOR: NAGA SAI AAKARSHIT BATCHU ************/

package cluster

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"
	cslib "github.com/elastic/beats/metricbeat/module/elasticsearch/cluster_eslibrary"
	"log"
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	if err := mb.Registry.AddMetricSet("elasticsearch", "cluster", New); err != nil {
		panic(err)
	}
}

// MetricSet type defines all fields of the MetricSet
// As a minimum it must inherit the mb.BaseMetricSet fields, but can be extended with
// additional entries. These variables can be used to persist data or configuration between
// multiple fetch calls.
type MetricSet struct {
	mb.BaseMetricSet
}

// New create a new instance of the MetricSet
// Part of new is also setting up the configuration by processing additional
// configuration entries if needed.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {

	config := struct{}{}

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
	}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Fetch methods implements the data gathering and data conversion to the right format
// It returns the event which is then forward to the output. In case of an error, a
// descriptive error must be returned.
func (m *MetricSet) Fetch() (common.MapStr, error) {

	event := common.MapStr{}
	nodekeyvalue := common.MapStr{}
	indexkeyvalue := common.MapStr{}
	esc := cslib.NewClient(m.Host())
	nodes, nodesErr := esc.GetNodes()
	failOnError(nodesErr, "Failed to get Nodes")
	for _, node := range nodes {
		value, valErr := esc.GetNodeValues(node)
		failOnError(valErr, "Failed to get Node Values")
		nodekeyvalue[node] = value
	}
	indices, indicesErr := esc.GetIndices()
	failOnError(indicesErr, "Failed to get Indices")
	for _, index := range indices {
		value, valErr := esc.GetIndexValues(index)
		failOnError(valErr, "Failed to get Index Values")
		indexkeyvalue[index] = value
	}
	event["nodes"] = nodekeyvalue
	event["indices"] = indexkeyvalue
	event["metric.host"] = m.Host()
	event["cluster_name"] = esc.GetClusterName()
	event["timestamp"] = esc.GetTimeStamp()
	event["status"] = esc.GetStatus()

	return event, nil
}

/********* AUTHOR: NAGA SAI AAKARSHIT BATCHU ************/
