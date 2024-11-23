# go_prometheus_metrics
A simple Go API which exports metrics with prometheus client


## metric for counting incoming http requests

$ curl -s localhost:8080/users
{"users":[{"name":"u1","Role":"dev"},{"name":"u2","Role":"qa"}]}

$ curl -s localhost:8080/health
{"status":"running"}

$ curl -s localhost:8080/metrics | grep -i http_total_requests
# HELP http_total_requests total number of http requests received
# TYPE http_total_requests counter
http_total_requests{path="/health"} 1
http_total_requests{path="/metrics"} 1
http_total_requests{path="/users"} 1
http_total_requests{path="all"} 3


$ curl -s localhost:8080/users
{"users":[{"name":"u1","Role":"dev"},{"name":"u2","Role":"qa"}]}

$ curl -s localhost:8080/metrics | grep -i http_total_requests
# HELP http_total_requests total number of http requests received
# TYPE http_total_requests counter
http_total_requests{path="/health"} 1
http_total_requests{path="/metrics"} 2
http_total_requests{path="/users"} 2
http_total_requests{path="all"} 5



To do:
1. need a metric that returns number of users

