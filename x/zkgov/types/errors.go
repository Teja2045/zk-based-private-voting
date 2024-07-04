package types

import (
	"cosmossdk.io/errors"
)

// x/zk-gov module sentinel errors
var (
	EmptyCommitment = errors.Register(ModuleName, 1, "Commitment is empty")
)
