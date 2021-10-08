package async

import "net/http"

// Client stores a client and has two channels to aggregate
// responses and errors
type Client struct {
	*http.Client
	Resp chan *http.Response
	Err  chan error
}

func (c *Client) AsyncGet(url string) {
	resp,err:=c.Get(url)
	if err != nil {
		c.Err<-err
		return
	}
	c.Resp<-resp
}

func NewClient(client *http.Client,bufferSize int) *Client {
	respch:=make(chan *http.Response,bufferSize)
	errch:=make(chan error,bufferSize)
	return &Client{
		client,
		respch,
		errch,
	}
}
