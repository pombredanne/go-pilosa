package pilosa

// Cluster is a simple ICluster implementation
type Cluster struct {
	hosts     []*URI
	nextIndex int
}

// NewCluster creates a Cluster with no addresses
func NewCluster() *Cluster {
	return &Cluster{
		hosts: make([]*URI, 0),
	}
}

// NewClusterWithHost creates a Cluster with the given address
func NewClusterWithHost(host *URI) *Cluster {
	cluster := NewCluster()
	cluster.AddHost(host)
	return cluster
}

// AddHost adds an address to the cluster
func (c *Cluster) AddHost(address *URI) {
	c.hosts = append(c.hosts, address)
}

// GetHost returns the next address in the cluster
func (c *Cluster) GetHost() *URI {
	if len(c.hosts) == 0 {
		return nil
	}
	// Return the transport, e.g., http from http+protobuf
	uri := c.hosts[c.nextIndex%len(c.hosts)]
	c.nextIndex = (c.nextIndex + 1) % len(c.hosts)
	return uri
}

// RemoveHost removes an address from the cluster
func (c *Cluster) RemoveHost(address *URI) {
	// TODO:
}

// GetHosts returns all addresses in this cluster
func (c *Cluster) GetHosts() []URI {
	arr := make([]URI, 0, len(c.hosts))
	for _, u := range c.hosts {
		arr = append(arr, *u)
	}
	return arr
}