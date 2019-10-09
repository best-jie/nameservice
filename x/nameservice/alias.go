package nameservice

import "github.com/best-jie/nameservice/x/nameservice/types"

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgBuyName      = types.NewMsgBuyName
	NewMsgSetName      = types.NewMsgSetName
	NewWhois           = types.NewWhois
	ModuleCdc          = types.ModuleCdc
	RegisterCodec      = types.RegisterCodec
	NewMsgAddNewCoin   = types.NewMsgAddNewCoin
	NewMsgAddCoin      = types.NewMsgAddCoin
	NewMsgBurnCoin     = types.NewMsgBurnCoin

)

type (
	MsgSetName      = types.MsgSetName
	MsgBuyName      = types.MsgBuyName
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois
	MsgAddNewCoin   = types.MsgAddNewCoin
	MsgAddCoin      = types.MsgAddCoin
	MsgBurnCoin     = types.MsgBurnCoin
)

