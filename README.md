# MetricBeat ElasticSearch Module (#Beats #ELK)

## Introduction:
Beats are the official data shipping modules for ELK Stack. And Metricbeat in-specific is used to collect and ship the Application Metrics. This Elasticsearch Metricbeat Module helps to collect and ship the Elasticsearch Application Metrics(Nodes, Nodesstats, Clusterstats, ClusterHealth) to the required output.

## Installation:

1. Install Go latest version(Optional- Only for Development purposes)

## To Run the Service:

To Run the Service initially you need to configure your metricbeat.yml file.(Input: Module ElasticSearch and specify the Metricset, Output: As Per your Requirement).

 Download the following files from repo to start using the built custom-metricbeat:
1. metricbeat (Executable Binary File)
2. metricbeat.yml (YAML Configuration File)
3. metricbeat.full.yml (Full YAML Configuration File)
4. metricbeat.template.json
5. metricbeat.template-es2x.json

The Below Command will run the Metricbeat in foreground.

```
./metricbeat
```

To run it on background in nohup, use the below command.

```
nohup ./metricbeat &
```

## Metricsets Available:

"nodes" - This Metricset lets you collect and ship the Elasticsearch Application's Node's Metrics.

"cluster" - This Metricset lets you collect and ship the Elasticsearch Application's Cluster's Metrics. 

"clusterhealth" - This Metricset lets you collect and ship the Elasticsearch Application's Cluster's health Metrics.

Configuration Example (Input Module):

```
#---------------------------- elasticsearch Module ---------------------------
- module: elasticsearch
  metricsets: ["nodes"]
  enabled: true
  period: 1s
  hosts: ["localhost:9200"]

```

## Sample YAML Configuraton Template:
	---Sample metricbeat yaml configuration is also provided in the file metricbeat.yml---
	---Also all other dependency files like metricbeat.template.json are also provided in this repo---
	
To collect Cluster Stats and ship & index them to Elasticsearch:

```
###################### Metricbeat Configuration Example #######################

# This file is an example configuration file highlighting only the most common
# options. The metricbeat.full.yml file from the same directory contains all the
# supported options with more comments. You can use it as a reference.
#
# You can find the full configuration reference here:
# https://www.elastic.co/guide/en/beats/metricbeat/index.html

#==========================  Modules configuration ============================
metricbeat.modules:

#------------------------------- System Module -------------------------------
#- module: system
#  metricsets:
    # CPU stats
#    - cpu

    # System Load stats
#    - load

    # Per CPU core stats
    #- core

    # IO stats
    #- diskio

    # Per filesystem stats
#    - filesystem

    # File system summary stats
#    - fsstat

    # Memory stats
#    - memory

    # Network stats
#    - network

    # Per process stats
#    - process
#  enabled: true
#  period: 10s
#  processes: ['.*']


#---------------------------- elasticsearch Module ---------------------------
- module: elasticsearch
  metricsets: ["cluster"]
  enabled: true
  period: 1s
  hosts: ["localhost:9200"]

#================================ General =====================================

# The name of the shipper that publishes the network data. It can be used to group
# all the transactions sent by a single shipper in the web interface.
#name:

# The tags of the shipper are included in their own field with each
# transaction published.
#tags: ["service-X", "web-tier"]

# Optional fields that you can specify to add additional information to the
# output.
#fields:
#  env: staging

#================================ Outputs =====================================

# Configure what outputs to use when sending the data collected by the beat.
# Multiple outputs may be used.

#-------------------------- Console output ------------------------------
#output.console:
#  pretty: true

#-------------------------- Elasticsearch output ------------------------------
output.elasticsearch:
  # Array of hosts to connect to.
  hosts: ["localhost:9200"]
  index: "elasticsearch-metrics-cluster-%{+yyyy.MM.dd}"

  # Optional protocol and basic auth credentials.
  #protocol: "https"
  #username: "elastic"
  #password: "changeme"

#----------------------------- Logstash output --------------------------------
#output.logstash:
  # The Logstash hosts
  #hosts: ["localhost:5044"]

  # Optional SSL. By default is off.
  # List of root certificates for HTTPS server verifications
  #ssl.certificate_authorities: ["/etc/pki/root/ca.pem"]

  # Certificate for SSL client authentication
  #ssl.certificate: "/etc/pki/client/cert.pem"

  # Client Certificate Key
  #ssl.key: "/etc/pki/client/cert.key"

#================================ Logging =====================================

# Sets log level. The default log level is info.
# Available log levels are: critical, error, warning, info, debug
#logging.level: debug

# At debug level, you can selectively enable logging only for some components.
# To enable all selectors use ["*"]. Examples of other selectors are "beat",
# "publish", "service".
#logging.selectors: ["*"]

```

## Author:

   NAGA SAI AAKARSHIT BATCHU (aakarshit.batchu@gmail.com)
