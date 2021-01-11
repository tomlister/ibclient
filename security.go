package ib

// Security stores information about a security
type Security struct {
	Broker BrokerAccount
	Conid  int
}

// Security creates a security object from a position
func (ba BrokerAccount) Security(p Position) Security {
	security := Security{
		Broker: ba,
		Conid:  p.Conid,
	}
	return security
}
