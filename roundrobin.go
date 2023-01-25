package roundrobin

import (
	"errors"
	"net/url"
	"sync/atomic"
)

var ServerNotExist = errors.New("Server Does Not Exist!!!")

type RoundRobin interface {
	Next() *url.URL
}

type roundrobin struct {
	urls []*url.URL
	next uint32
}

func New(urls ...*url.URL) (RoundRobin, error) {
	if len(urls) == 0 {
		return nil, ServerNotExist
	}

	return &roundrobin{
		urls: urls,
	}, nil

}

func (r *roundrobin) Next() *url.URL {
	n := atomic.AddUint32(&r.next, 1)
	return r.urls[(int(n)-1)%len(r.urls)]
}
