package allroutesclientimpl

func (c *client) CloseConn() error {
	return c.conn.Close()
}
