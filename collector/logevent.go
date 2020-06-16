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
  "github.com/prometheus/client_golang/prometheus"
)


const (
  logEventSubsystem = "logevent"
)

var (
  logEventDesc = prometheus.NewDesc(
    prometheus.BuildFQName(namespace, logEventSubsystem, "event_total"),
    "Get logevent from target since epoch (1970).",
    []string{"log", "target", "pattern", "filename"}, nil,
  )
)
type logFileStats struct {
  LabelName string
  Target    string
  Pattern   string
  FileName  string
  Count     uint64
  Status    bool
  LogTime   float64
}

type logEventCollector struct {
  desc   *prometheus.Desc
}

func init() {
  registerCollector("logevent", true, NewLogEventCollector)
  load()
}

// NewTimeCollector returns a new Collector exposing the log event in
// seconds since epoch.
func NewLogEventCollector() (Collector, error) {
  return &logEventCollector{
    desc: logEventDesc,
  }, nil
}


/*
  prometheus get metrics from node exporter
*/
func (c *logEventCollector) Update(ch chan<- prometheus.Metric) error {
  for _, logInfo := range logStats {
     if logInfo.Status == true {

        ch <- prometheus.MustNewConstMetric(c.desc, 
                prometheus.GaugeValue, 
                float64(logInfo.Count),
                logInfo.LabelName, 
                logInfo.Target, logInfo.Pattern, logInfo.FileName)

        logInfo.Status = false
        logInfo.Count = 0
        logStats[logInfo.LabelName] = logInfo        
     }
  }

  return nil
}
