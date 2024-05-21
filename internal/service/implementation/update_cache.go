package serviceimpl

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

func (s *service) updateCache(c *cache.LinearCache, resource map[string]types.Resource) (*cache.LinearCache, error) {
	switch {
	case c.NumResources() == 0:
		c.SetResources(resource)
	default:
		if err := c.UpdateResources(resource, nil); err != nil {
			return nil, err
		}
	}
	return c, nil
}
