/********* AUTHOR: NAGA SAI AAKARSHIT BATCHU ************/
package nodes

import (
        "github.com/elastic/beats/libbeat/common"
        "github.com/elastic/beats/metricbeat/mb"
        eslib "github.com/elastic/beats/metricbeat/module/elasticsearch/eslibrary"
        "log"
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
        if err := mb.Registry.AddMetricSet("elasticsearch", "nodes", New); err != nil {
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
func (m *MetricSet) Fetch() ([]common.MapStr, error) {

                events := []common.MapStr{}
                esc := eslib.NewClient(m.Host())
                nodes,nodeserr := esc.GetNodes()
                failOnError(nodeserr, "Unable to Get the Nodes")
                for node := range nodes {
                        event := common.MapStr{}
                        event["Node"] = nodes[node]
                        keys,_ := esc.ListKeysinNode(nodes[node])
                        for key := range keys {
                                v,_ := esc.GetValue(nodes[node],keys[key])
                                event[keys[key]] = v
                        }
                                events = append(events,event)
                }
        return events, nil
}

/********* AUTHOR: NAGA SAI AAKARSHIT BATCHU ************/