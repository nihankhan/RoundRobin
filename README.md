# RoundRobin
This is a simple load balancing algorithm implemented with golang.

## Installation:
***

go get github.com/nihankhan/round-robin

Example:

rr, _ := roundrobin.New(
    &url.URL{Host: "192.168.33.10"},
    &url.URL{Host: "192.168.33.11"},
    &url.URL{Host: "192.168.33.12"},
    &url.URL{Host: "192.168.33.13"},
)

rr.Next() // {Host: "192.168.33.10"}
rr.Next() // {Host: "192.168.33.11"}
rr.Next() // {Host: "192.168.33.12"}
rr.Next() // {Host: "192.168.33.13"}
rr.Next() // {Host: "192.168.33.10"}
rr.Next() // {Host: "192.168.33.11"}
