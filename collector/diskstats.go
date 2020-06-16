// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !notime

package collector

import (
  "fmt"
  "encoding/json"
  "github.com/prometheus/client_golang/prometheus"
)

const (
  NEWLINE="\n"
  COLUMN_DELIMITER="|"
  LABEL_DELIMITER=","
  KEY_VALUE_DELIMITER="="
  MAIN_SHELL_SCRIPT = "scripts/sysinfo.sh"
)

type sysinfoCollector struct {
  desc   *prometheus.Desc
}

type Sysinfo struct {
  Name         string `json:"name"`
  LabelName  []string `json:"label_name"`
  LabelValue []string `json:"label_value"`
  Value        float64 `json:"value"`
}

func init() {
  registerCollector("sysinfo", true, NewSysInfoCollector)
}

func NewSysInfoCollector() (Collector, error) {
  return &sysinfoCollector{
    desc: prometheus.NewDesc(
      namespace+"_sysinfo",
      "Sysinfo since epoch (2020) by jgyun.",
      nil, nil,
    ),
  }, nil
}


func (c *sysinfoCollector) Update(ch chan<- prometheus.Metric) error {
  var (
    labels       []string
    labelLength  int
    sysinfo      []Sysinfo
    metricType prometheus.ValueType
  )
 
  result := RunShell(MAIN_SHELL_SCRIPT);
  json.Unmarshal([]byte(result), &sysinfo)
  metricType = prometheus.GaugeValue
  
  for index := range sysinfo {
     // If label exist,  metricType is Gauge , otherwise Counter
     labelLength = len(sysinfo[index].LabelName)
     if labelLength > 0 {
       metricType = prometheus.GaugeValue
     } else {
       metricType = prometheus.CounterValue
     }

     labels = sysinfo[index].LabelName
     ch <- prometheus.MustNewConstMetric(
       prometheus.NewDesc(
         prometheus.BuildFQName(namespace, "", sysinfo[index].Name),
         fmt.Sprintf("System information "),
         labels, nil,
       ),
       metricType, sysinfo[index].Value, sysinfo[index].LabelValue...,
    )
  }  

  return nil
}
