package nameservice

import (
	"github.com/azizyano/myapp/x/nameservice/keeper"
	"github.com/azizyano/myapp/x/nameservice/types"
)

const (
	// TODO: define constants that you would like exposed from your module

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	QueryParams       = types.QueryParams
	QuerierRoute      = types.QuerierRoute
	StoreKey          = types.StoreKey
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewMsgBuyName       = types.NewMsgBuyName
	NewMsgSetName       = types.NewMsgSetName
	NewMsgDeleteName    = types.NewMsgDeleteName
	NewWhois            = types.NewWhois
	ModuleCdc           = types.ModuleCdc
	RegisterCodec       = types.RegisterCodec
	// TODO: Fill out function aliases

	// variable aliases
	ModuleCdc = types.ModuleCdc
	// TODO: Fill out variable aliases
)

type (
	Keeper          = keeper.Keeper
	GenesisState    = types.GenesisState
	Params          = types.Params
	MsgSetName      = types.MsgSetName
	MsgBuyName      = types.MsgBuyName
	MsgDeleteName   = types.MsgDeleteName
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois

	// TODO: Fill out module types
)
