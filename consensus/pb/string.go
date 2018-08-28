package consensuspb

import (
	"fmt"

	"github.com/pepperdb/pepperdb-core/common/util/byteutils"
)

// ToString return a string of consensus root
func (m *ConsensusRoot) ToString() string {
	return fmt.Sprintf(`{"proposer": %s, "timestamp": "%d", "dynasty": "%s"}`,
		byteutils.Hex(m.Proposer),
		m.Timestamp,
		byteutils.Hex(m.DynastyRoot),
	)
}
