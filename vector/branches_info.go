package vector

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/inter/pos"
)

// branchesInfo contains information about global branches of each validator
type branchesInfo struct {
	BranchIDLastSeq     []idx.Event       // branchID -> highest e.Seq in the branch
	BranchIDCreatorIdxs []idx.Validator   // branchID -> validator idx
	BranchIDCreators    []common.Address  // branchID -> validator addr
	BranchIDByCreators  [][]idx.Validator // validator idx -> list of branch IDs
}

// initBranchesInfo loads branchesInfo from store
func (vi *Index) initBranchesInfo() {
	if vi.bi == nil {
		// if not cached
		vi.bi = vi.getBranchesInfo()
		if vi.bi == nil {
			// first run
			vi.bi = newInitialBranchesInfo(vi.validators)
		}
	}
}

func newInitialBranchesInfo(validators pos.Validators) *branchesInfo {
	branchIDCreators := validators.SortedAddresses()
	branchIDCreatorIdxs := make([]idx.Validator, len(branchIDCreators))
	for i := range branchIDCreators {
		branchIDCreatorIdxs[i] = idx.Validator(i)
	}

	branchIDLastSeq := make([]idx.Event, len(branchIDCreatorIdxs))
	branchIDByCreators := make([][]idx.Validator, len(validators))
	for i := range branchIDByCreators {
		branchIDByCreators[i] = make([]idx.Validator, 1, len(validators)/2+1)
		branchIDByCreators[i][0] = idx.Validator(i)
	}
	return &branchesInfo{
		BranchIDLastSeq:     branchIDLastSeq,
		BranchIDCreatorIdxs: branchIDCreatorIdxs,
		BranchIDByCreators:  branchIDByCreators,
		BranchIDCreators:    branchIDCreators,
	}
}

func (vi *Index) atLeastOneFork() bool {
	return len(vi.bi.BranchIDCreators) > len(vi.validators)
}
