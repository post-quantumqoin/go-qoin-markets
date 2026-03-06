package discoveryimpl

import (
	"github.com/post-quantumqoin/go-qoin-markets/discovery"
)

func Multi(r discovery.PeerResolver) discovery.PeerResolver { // TODO: actually support multiple mechanisms
	return r
}
