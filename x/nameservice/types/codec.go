package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "nameservice/BuyName", nil)
	cdc.RegisterConcrete(MsgAddNewCoin{}, "nameservice/AddNewCoin", nil)
	cdc.RegisterConcrete(MsgAddCoin{}, "nameservice/AddCoin", nil)
	cdc.RegisterConcrete(MsgBurnCoin{}, "nameservice/BurnCoin", nil)
}
