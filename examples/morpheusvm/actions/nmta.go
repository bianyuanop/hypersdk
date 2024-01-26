package actions

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/consts"
	mconsts "github.com/ava-labs/hypersdk/examples/morpheusvm/consts"
	"github.com/ava-labs/hypersdk/examples/morpheusvm/storage"
	"github.com/ava-labs/hypersdk/state"
)

var _ chain.Action = (*NMTTestAction)(nil)

type NMTTestAction struct {
	ChainID []byte `json:"chainID"`
	// Amount are transferred to [To].
	Value uint64 `json:"value"`
}

func (*NMTTestAction) GetTypeID() uint8 {
	return mconsts.NMTTestActionID
}

func (t *NMTTestAction) StateKeys(actor codec.Address, _ ids.ID) []string {
	return []string{
		string(storage.BalanceKey(actor)),
	}
}

func (*NMTTestAction) StateKeysMaxChunks() []uint16 {
	return []uint16{storage.BalanceChunks, storage.BalanceChunks}
}

func (*NMTTestAction) OutputsWarpMessage() bool {
	return false
}

func (t *NMTTestAction) Execute(
	ctx context.Context,
	_ chain.Rules,
	mu state.Mutable,
	_ int64,
	actor codec.Address,
	_ ids.ID,
	_ bool,
) (bool, uint64, []byte, *warp.UnsignedMessage, error) {
	return true, 1, nil, nil, nil
}

func (*NMTTestAction) MaxComputeUnits(chain.Rules) uint64 {
	return TransferComputeUnits
}

func (*NMTTestAction) Size() int {
	return codec.AddressLen + consts.Uint64Len
}

func (t *NMTTestAction) Marshal(p *codec.Packer) {
	p.PackUint64(t.Value)
}

func UnmarshalNMTTestAction(p *codec.Packer, _ *warp.Message) (chain.Action, error) {
	var transfer NMTTestAction
	transfer.Value = p.UnpackUint64(true)
	if err := p.Err(); err != nil {
		return nil, err
	}
	return &transfer, nil
}

func (*NMTTestAction) ValidRange(chain.Rules) (int64, int64) {
	// Returning -1, -1 means that the action is always valid.
	return -1, -1
}

func (na *NMTTestAction) NMTNamespace() []byte {
	// do calculation here to get namespace id
	// go-eth use chainID as uint64 and it cannot be 0
	// hence we can direclty using chain id as namespace id (8 bytes)

	return na.ChainID
}
