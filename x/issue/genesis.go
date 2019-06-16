package issue

import (
	"github.com/tichex-project/hackatom-cosmos-tichex/x/issue/config"
	"github.com/tichex-project/hackatom-cosmos-tichex/x/issue/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState defines genesis data for the module
type GenesisState struct {
	StartingIssueId uint64          `json:"starting_issue_id"`
	Issues          []CoinIssueInfo `json:"issues"`
	Params          config.Params   `json:"params"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(startingIssueId uint64) GenesisState {
	return GenesisState{StartingIssueId: startingIssueId}
}


// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState(types.CoinIssueMinId)
}

// InitGenesis initializes story state from genesis file
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) {
	err := keeper.SetInitialIssueStartingIssueId(ctx, data.StartingIssueId)
	if err != nil {
		panic(err)
	}

	keeper.SetParams(ctx, data.Params)

	for _, issue := range data.Issues {
		keeper.AddIssue(ctx, &issue)
	}
}

// ExportGenesis exports the genesis state
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {
	genesisState := GenesisState{}

	startingIssueId, _ := keeper.PeekCurrentIssueID(ctx)
	genesisState.StartingIssueId = startingIssueId

	genesisState.Params = keeper.GetParams(ctx)
	genesisState.Issues = keeper.ListAll(ctx)

	return genesisState

}	

// ValidateGenesis validates the genesis state data
func ValidateGenesis(data GenesisState) error {
	return nil
}