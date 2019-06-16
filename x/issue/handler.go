package issue

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler creates a new handler for {{ .Name }} module
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		errMsg := fmt.Sprintf("Unrecognized issue message type: %T", msg)
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
}