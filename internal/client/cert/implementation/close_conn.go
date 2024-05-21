package certclientimpl

func (c *client) CloseConn() error {
	return c.conn.Close()
}
