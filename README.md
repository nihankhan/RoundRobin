# round-robin
round-robin is balancing algorithm written in golang

## Installation

```shell
go get github.com/nihankhan/round-robin
```

## Example

```go
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
```
## Author
[nihankhan](https://github.com/nihankhan)

## LICENSE
round-robin released under GNU license, refer [LICENSE](https://github.com/nihankhan/round-robin/blob/main/LICENSE) file.
