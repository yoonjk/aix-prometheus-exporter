# AIX Prometheus Exporter

This is a node exporter for AIX.

Tested with `AIX 7.1/7.2` on `IBM POWER 7` and Added functions using shell script

![alt text](images/unode.png)

Collectors:

* cpu
* logevent
* memory
* sysinfo
* time

## 1. Pre-requisites

== Installation
[source,bash]
. Setup Git configuraiton
----
* git config --global http.sslVerify false
* git https validate off
* git config --global user.name "Kildong Hong"
---- 
[source,bash]
. Generation Certification file
----
openssl genrsa -out rootCA.key 2048 
openssl req -x509 -new -nodes -key rootCA.key -subj "/CN=*.tonybai.com" -days 5000 -out rootCA.pem 
openssl genrsa -out cert.key 2048 
openssl req -new -key cert.key -subj "/CN=tonybai.com" -out cert.csr 
openssl x509 -req -in cert.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out cert.crt -days 5000
----

[source,bash]
. Download library using go packages
----
go get -u -v -insecure  gopkg.in/check.v1
go get -u -v -insecure  golang.org/x/sys
go get -u -v -insecure  golang.org/x/net
go get -v -u -insecure  gopkg.in/alecthomas/kingpin.v2
go get -v -u -insecure  gopkg.in/yaml.v2 v2.2.8
go get -v -u -insecure  github.com/pkg/errors v0.8.1
go get -v -u -insecure  github.com/prometheus/client_golang v1.0.0
go get -v -u -insecure  github.com/prometheus/common v0.7.0
----


