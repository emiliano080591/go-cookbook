package async

func FetchAll(urls []string,c *Client)  {
	for _,url:=range urls{
		go c.AsyncGet(url)
	}
}
