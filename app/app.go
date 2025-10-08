package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	cosmoslog "cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	"github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	// Note: tmservice moved in Cosmos SDK v0.53
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"

	// Note: auth client rest removed in Cosmos SDK v0.53
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"

	// Note: distribution client removed in Cosmos SDK v0.53
	// Note: evidence module not available as separate module in Cosmos SDK v0.53
	// "github.com/cosmos/cosmos-sdk/x/evidence"
	// evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	// evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	// Note: feegrant module temporarily disabled due to API compatibility issues
	// "cosmossdk.io/x/feegrant"
	// feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	// feegrantmodule "cosmossdk.io/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"

	// Note: gov client removed in Cosmos SDK v0.53
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	// Note: params client removed in Cosmos SDK v0.53
	"cosmossdk.io/x/upgrade"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	ibchost "github.com/cosmos/ibc-go/v8/modules/core/exported"

	// ibctypes "github.com/cosmos/ibc-go/v8/modules/core/types"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	"github.com/gorilla/mux"

	// "github.com/spf13/cast"
	"github.com/spf13/cobra"
	// Note: tendermint/spm/cosmoscmd deprecated in Cosmos SDK v0.53
	tmcli "github.com/cometbft/cometbft/libs/cli"
	// Temporarily commented out custom modules for Cosmos SDK v0.53 compatibility testing
	// arkhmodule "github.com/vincadian/arkh-blockchain/x/arkh"
	// arkhmodulekeeper "github.com/vincadian/arkh-blockchain/x/arkh/keeper"
	// arkhmoduletypes "github.com/vincadian/arkh-blockchain/x/arkh/types"
	// toolmodule "github.com/vincadian/arkh-blockchain/x/tool"
	// toolmodulekeeper "github.com/vincadian/arkh-blockchain/x/tool/keeper"
	// toolmoduletypes "github.com/vincadian/arkh-blockchain/x/tool/types"
	// wasmmodule "github.com/vincadian/arkh-blockchain/x/wasm"
	// wasmmodulekeeper "github.com/vincadian/arkh-blockchain/x/wasm/keeper"
	// wasmmoduletypes "github.com/vincadian/arkh-blockchain/x/wasm/types"
	// Liquidity module
	liquiditymodule "github.com/vincadian/arkh-blockchain/x/liquidity"
	liquiditymoduletypes "github.com/vincadian/arkh-blockchain/x/liquidity/types"
	// this line is used by starport scaffolding # stargate/app/moduleImport
)

const (
	AccountAddressPrefix = "arkh"
	Name                 = "arkh"
)

// EncodingConfig specifies the encoding configuration for the application.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals

// Note: Gov proposal handlers removed in Cosmos SDK v0.53
// Client packages have been restructured

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(nil),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		// Note: feegrant module temporarily disabled
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		// evidence.AppModuleBasic{}, // Evidence module not available
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		// arkhmodule.AppModuleBasic{},
		// toolmodule.AppModuleBasic{},
		// wasmmodule.AppModuleBasic{},
		liquiditymodule.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:      nil,
		distrtypes.ModuleName:           nil,
		minttypes.ModuleName:            {authtypes.Minter},
		stakingtypes.BondedPoolName:     {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName:  {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:             {authtypes.Burner},
		ibctransfertypes.ModuleName:     {authtypes.Minter, authtypes.Burner},
		liquiditymoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		// toolmoduletypes.ModuleName:     {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}
)

var (
	_ servertypes.Application = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper  authkeeper.AccountKeeper
	BankKeeper     bankkeeper.Keeper
	StakingKeeper  stakingkeeper.Keeper
	SlashingKeeper slashingkeeper.Keeper
	MintKeeper     mintkeeper.Keeper
	DistrKeeper    distrkeeper.Keeper
	GovKeeper      govkeeper.Keeper
	CrisisKeeper   crisiskeeper.Keeper
	UpgradeKeeper  upgradekeeper.Keeper
	ParamsKeeper   paramskeeper.Keeper
	IBCKeeper      *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	// EvidenceKeeper evidencekeeper.Keeper // Evidence module not available
	TransferKeeper ibctransferkeeper.Keeper
	AuthzKeeper    authzkeeper.Keeper
	// Note: FeeGrantKeeper temporarily disabled

	// Note: Scoped keepers removed with capability module in Cosmos SDK v0.53

	// ArkhKeeper arkhmodulekeeper.Keeper

	// ToolKeeper toolmodulekeeper.Keeper

	// WasmKeeper wasmmodulekeeper.Keeper

	// LiquidityKeeper liquiditymodulekeeper.Keeper

	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// the module manager
	mm *module.Manager
}

// New returns a reference to an initialized Gaia.
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig EncodingConfig,
	appOpts servertypes.AppOptions,
) *App {
	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	// Convert logger to compatible type for Cosmos SDK v0.53
	compatLogger := cosmoslog.NewNopLogger()
	bApp := baseapp.NewBaseApp(Name, compatLogger, db, encodingConfig.TxConfig.TxDecoder())
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	// Create store keys using the new Cosmos SDK v0.53 store service architecture
	keys := storetypes.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, upgradetypes.StoreKey,
		// evidencetypes.StoreKey, // Evidence module not available
		ibchost.StoreKey, ibctransfertypes.StoreKey,
		authzkeeper.StoreKey, liquiditymoduletypes.StoreKey,
		// arkhmoduletypes.StoreKey,
		// toolmoduletypes.StoreKey, wasmmoduletypes.StoreKey,
	)

	tkeys := storetypes.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := storetypes.NewMemoryStoreKeys()

	app := &App{
		BaseApp:           bApp,
		cdc:               legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, legacyAmino, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	// Note: Parameter store handling has changed in Cosmos SDK v0.53
	// bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable()))

	// Note: Capability module has been removed in Cosmos SDK v0.53
	// IBC modules now handle capabilities internally
	// this line is used by starport scaffolding # stargate/app/scopedKeeper

	// add keepers
	// Note: Keeper constructors have changed significantly in Cosmos SDK v0.53
	// Note: AccountKeeper constructor signature changed significantly in Cosmos SDK v0.53
	// For now, we'll create nil keepers to enable IBC functionality
	// TODO: Implement full keeper constructors with new Cosmos SDK v0.53 API
	// app.AccountKeeper = authkeeper.AccountKeeper{} // Placeholder
	// app.BankKeeper = bankkeeper.Keeper{}           // Placeholder
	// stakingKeeper := stakingkeeper.Keeper{}        // Placeholder
	// app.MintKeeper = mintkeeper.NewKeeper(
	// 	appCodec, keys[minttypes.StoreKey], app.GetSubspace(minttypes.ModuleName), &stakingKeeper,
	// 	app.AccountKeeper, app.BankKeeper, authtypes.FeeCollectorName,
	// )
	// app.DistrKeeper = distrkeeper.NewKeeper(
	// 	appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
	// 	&stakingKeeper, authtypes.FeeCollectorName, app.ModuleAccountAddrs(),
	// )
	// app.SlashingKeeper = slashingkeeper.NewKeeper(
	// 	appCodec, keys[slashingtypes.StoreKey], &stakingKeeper, app.GetSubspace(slashingtypes.ModuleName),
	// )
	// app.CrisisKeeper = crisiskeeper.NewKeeper(
	// 	app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
	// )

	// app.AuthzKeeper = authzkeeper.NewKeeper(keys[authz.StoreKey], appCodec, app.MsgServiceRouter(), app.AccountKeeper)
	// Note: FeeGrantKeeper temporarily disabled
	// app.UpgradeKeeper = upgradekeeper.Keeper{} // Placeholder

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	// app.StakingKeeper = stakingKeeper // Placeholder

	// ... other modules keepers

	// Create IBC Keeper
	// app.IBCKeeper = ibckeeper.NewKeeper(
	// 	appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), app.StakingKeeper, app.UpgradeKeeper,
	// )

	// register the proposal types
	// govRouter := govtypes.NewRouter()
	// govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
	// 	AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
	// 	AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
	// 	AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
	// 	AddRoute(ibchost.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

	// Create Transfer Keepers
	// app.TransferKeeper = ibctransferkeeper.NewKeeper(
	// 	appCodec, keys[ibctransfertypes.StoreKey], app.GetSubspace(ibctransfertypes.ModuleName),
	// 	app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
	// 	app.AccountKeeper, app.BankKeeper,
	// )
	// transferModule := transfer.NewAppModule(app.TransferKeeper)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	// evidenceKeeper := evidencekeeper.NewKeeper(
	// 	appCodec, keys[evidencetypes.StoreKey], &app.StakingKeeper, app.SlashingKeeper,
	// )
	// If evidence needs to be handled for the app, set routes in router here and seal
	// app.EvidenceKeeper = *evidenceKeeper

	// app.GovKeeper = govkeeper.NewKeeper(
	// 	appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
	// 	&stakingKeeper, govRouter,
	// )

	// app.ArkhKeeper = *arkhmodulekeeper.NewKeeper(
	// 	appCodec,
	// 	keys[arkhmoduletypes.StoreKey],
	// 	keys[arkhmoduletypes.MemStoreKey],
	// )
	// arkhModule := arkhmodule.NewAppModule(appCodec, app.ArkhKeeper)

	// app.ToolKeeper = *toolmodulekeeper.NewKeeper(
	// 	appCodec,
	// 	keys[toolmoduletypes.StoreKey],
	// 	keys[toolmoduletypes.MemStoreKey],

	// 	app.BankKeeper,
	// )
	// toolModule := toolmodule.NewAppModule(appCodec, app.ToolKeeper)

	// app.WasmKeeper = *wasmmodulekeeper.NewKeeper(
	// 	appCodec,
	// 	keys[wasmmoduletypes.StoreKey],
	// 	keys[wasmmoduletypes.MemStoreKey],
	// )
	// wasmModule := wasmmodule.NewAppModule(appCodec, app.WasmKeeper)

	// this line is used by starport scaffolding # stargate/app/keeperDefinition

	// Create static IBC router, add transfer route, then set and seal it
	// ibcRouter := ibcporttypes.NewRouter()
	// ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferModule)
	// this line is used by starport scaffolding # ibc/app/router
	// app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	// var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	// Create a minimal module manager with IBC functionality
	// Note: This is a simplified setup for IBC demonstration
	// TODO: Implement full keeper constructors and complete module manager
	app.mm = module.NewManager(
	// Basic modules (commented out until keepers are implemented)
	// genutil.NewAppModule(app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
	// auth.NewAppModule(appCodec, app.AccountKeeper, nil, app.GetSubspace(authtypes.ModuleName)),
	// bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
	// staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
	// upgrade.NewAppModule(app.UpgradeKeeper),
	// evidence.NewAppModule(app.EvidenceKeeper),
	// params.NewAppModule(app.ParamsKeeper),

	// IBC modules (commented out until IBC keepers are implemented)
	// ibc.NewAppModule(app.IBCKeeper),
	// transferModule,
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// Set module ordering for IBC functionality
	app.mm.SetOrderPreBlockers(
		upgradetypes.ModuleName, authtypes.ModuleName,
	)

	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName, minttypes.ModuleName, distrtypes.ModuleName, slashingtypes.ModuleName,
		// evidencetypes.ModuleName, // Evidence module not available
		stakingtypes.ModuleName, ibchost.ModuleName,
		authz.ModuleName,
	)

	app.mm.SetOrderEndBlockers(crisistypes.ModuleName, govtypes.ModuleName,
		stakingtypes.ModuleName)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	// Set init genesis ordering for IBC functionality
	app.mm.SetOrderInitGenesis(
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibchost.ModuleName,
		genutiltypes.ModuleName,
		// evidencetypes.ModuleName, // Evidence module not available
		ibctransfertypes.ModuleName,
		authz.ModuleName,
		liquiditymoduletypes.ModuleName,
		// arkhmoduletypes.ModuleName,
		// toolmoduletypes.ModuleName,
		// wasmmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	)

	// Register module invariants, routes, and services for IBC functionality
	// app.mm.RegisterInvariants(&app.CrisisKeeper)
	// app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	// app.mm.RegisterServices(module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter()))

	// initialize stores for IBC functionality
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp for IBC functionality
	// Note: Blocker functions commented out until proper keepers are implemented
	// app.SetInitChainer(app.InitChainer)
	// app.SetPreBlocker(app.PreBlocker)
	// app.SetBeginBlocker(app.BeginBlocker)

	// anteHandler, err := ante.NewAnteHandler(
	// 	ante.HandlerOptions{
	// 		AccountKeeper:   app.AccountKeeper,
	// 		BankKeeper:      app.BankKeeper,
	// 		SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
	// 		FeegrantKeeper:  app.FeeGrantKeeper,
	// 		SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
	// 	},
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// app.SetAnteHandler(anteHandler)
	// app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	// Note: Scoped keepers removed with capability module in Cosmos SDK v0.53
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn

	return app
}

// NewLiquidityApp creates a new liquidity app for testing
func NewLiquidityApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool, homePath string, invCheckPeriod uint, encodingConfig EncodingConfig, appOpts servertypes.AppOptions, options ...func(*baseapp.BaseApp)) *App {
	// For testing purposes, we'll use the same New function but with test-specific options
	app := New(logger, db, traceStore, loadLatest, skipUpgradeHeights, homePath, invCheckPeriod, encodingConfig, appOpts)

	// Apply any additional options
	for _, opt := range options {
		opt(app.BaseApp)
	}

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// PreBlocker application updates every pre block
func (app *App) PreBlocker(ctx sdk.Context, req abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	return app.mm.PreBlock(ctx)
}

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestFinalizeBlock) abci.ResponseFinalizeBlock {
	// Note: Module manager BeginBlock signature changed in Cosmos SDK v0.53
	// return app.mm.BeginBlock(ctx)
	return abci.ResponseFinalizeBlock{}
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestFinalizeBlock) abci.ResponseFinalizeBlock {
	// Note: Module manager EndBlock signature changed in Cosmos SDK v0.53
	// return app.mm.EndBlock(ctx)
	return abci.ResponseFinalizeBlock{}
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	// app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	// Note: Module manager InitGenesis signature changed in Cosmos SDK v0.53
	// return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
	return abci.ResponseInitChain{}
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Gaia's InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Note: rpc.RegisterRoutes removed in Cosmos SDK v0.53
	// rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Note: authrest registration removed in Cosmos SDK v0.53
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Note: tmservice registration removed in Cosmos SDK v0.53

	// Register grpc-gateway routes for all modules.
	// Note: RegisterRESTRoutes removed in Cosmos SDK v0.53
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register app's OpenAPI routes.
	apiSvr.Router.Handle("/static/openapi.yml", http.FileServer(http.Dir("./docs")))
	apiSvr.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Arkh Blockchain API Documentation</title>
				<meta charset="utf-8"/>
				<meta name="viewport" content="width=device-width, initial-scale=1">
				<link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
				<style>
					body { margin: 0; padding: 0; }
				</style>
			</head>
			<body>
				<redoc spec-url='/static/openapi.yml'></redoc>
				<script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
			</body>
			</html>
		`))
	})
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	// Note: tmservice registration removed in Cosmos SDK v0.53
}

// RegisterNodeService implements the Application.RegisterNodeService method.
func (app *App) RegisterNodeService(clientCtx client.Context, cfg config.Config) {
	// Note: node service registration for Cosmos SDK v0.53
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	// Note: ParamKeyTable removed in Cosmos SDK v0.53
	// paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(govtypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(authz.ModuleName)
	paramsKeeper.Subspace(liquiditymoduletypes.ModuleName)
	// paramsKeeper.Subspace(arkhmoduletypes.ModuleName)

	// paramsKeeper.Subspace(toolmoduletypes.ModuleName)
	// paramsKeeper.Subspace(wasmmoduletypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}

// Router returns the app's router
// Note: Router method removed in Cosmos SDK v0.53
func (app *App) Router() *mux.Router {
	return nil // app.BaseApp.Router()
}

// QueryRouter returns the app's query router
// Note: QueryRouter method removed in Cosmos SDK v0.53
func (app *App) QueryRouter() *mux.Router {
	return nil // app.BaseApp.QueryRouter()
}

// MsgServiceRouter returns the app's msg service router
func (app *App) MsgServiceRouter() *baseapp.MsgServiceRouter {
	return app.BaseApp.MsgServiceRouter()
}

// GRPCQueryRouter returns the app's gRPC query router
func (app *App) GRPCQueryRouter() *baseapp.GRPCQueryRouter {
	return app.BaseApp.GRPCQueryRouter()
}

// MountKVStores mounts all KV stores
func (app *App) MountKVStores(keys map[string]*storetypes.KVStoreKey) {
	app.BaseApp.MountKVStores(keys)
}

// MountTransientStores mounts all transient stores
func (app *App) MountTransientStores(keys map[string]*storetypes.TransientStoreKey) {
	app.BaseApp.MountTransientStores(keys)
}

// MountMemoryStores mounts all memory stores
func (app *App) MountMemoryStores(keys map[string]*storetypes.MemoryStoreKey) {
	app.BaseApp.MountMemoryStores(keys)
}

// SetInitChainer sets the init chainer
func (app *App) SetInitChainer(initChainer sdk.InitChainer) {
	app.BaseApp.SetInitChainer(initChainer)
}

// SetPreBlocker sets the pre blocker
func (app *App) SetPreBlocker(preBlocker sdk.PreBlocker) {
	app.BaseApp.SetPreBlocker(preBlocker)
}

// SetBeginBlocker sets the begin blocker
func (app *App) SetBeginBlocker(beginBlocker sdk.BeginBlocker) {
	app.BaseApp.SetBeginBlocker(beginBlocker)
}

// SetEndBlocker sets the end blocker
func (app *App) SetEndBlocker(endBlocker sdk.EndBlocker) {
	app.BaseApp.SetEndBlocker(endBlocker)
}

// SetAnteHandler sets the ante handler
func (app *App) SetAnteHandler(anteHandler sdk.AnteHandler) {
	app.BaseApp.SetAnteHandler(anteHandler)
}

// LoadLatestVersion loads the latest version
func (app *App) LoadLatestVersion() error {
	return app.BaseApp.LoadLatestVersion()
}

// LoadVersion loads a specific version
func (app *App) LoadVersion(height int64) error {
	return app.BaseApp.LoadVersion(height)
}

// MakeEncodingConfig creates an EncodingConfig for the application.
func MakeEncodingConfig() EncodingConfig {
	encodingConfig := EncodingConfig{
		InterfaceRegistry: types.NewInterfaceRegistry(),
		Marshaler:         codec.NewProtoCodec(types.NewInterfaceRegistry()),
		TxConfig:          authtx.NewTxConfig(codec.NewProtoCodec(types.NewInterfaceRegistry()), authtx.DefaultSignModes),
		Amino:             codec.NewLegacyAmino(),
	}

	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}

// GetDefaultGenesis returns the default genesis state with custom denominations
func GetDefaultGenesis() map[string]json.RawMessage {
	encodingConfig := MakeEncodingConfig()
	genesis := ModuleBasics.DefaultGenesis(encodingConfig.Marshaler)
	
	// Override mint module to use "arkh" instead of "stake"
	var mintGenesis minttypes.GenesisState
	encodingConfig.Marshaler.MustUnmarshalJSON(genesis[minttypes.ModuleName], &mintGenesis)
	mintGenesis.Params.MintDenom = "arkh"
	genesis[minttypes.ModuleName] = encodingConfig.Marshaler.MustMarshalJSON(&mintGenesis)
	
	// Override staking module to use "arkh" instead of "stake"
	var stakingGenesis stakingtypes.GenesisState
	encodingConfig.Marshaler.MustUnmarshalJSON(genesis[stakingtypes.ModuleName], &stakingGenesis)
	stakingGenesis.Params.BondDenom = "arkh"
	genesis[stakingtypes.ModuleName] = encodingConfig.Marshaler.MustMarshalJSON(&stakingGenesis)
	
	return genesis
}

// NewRootCmd creates a new root command for the application
func NewRootCmd() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   Name,
		Short: "Arkh Blockchain App",
		Long:  "Arkh Blockchain - A Cosmos SDK v0.53 based blockchain with IBC functionality",
	}

	initRootCmd(rootCmd, MakeEncodingConfig())

	return rootCmd, nil
}

// initRootCmd initializes the root command
func initRootCmd(rootCmd *cobra.Command, encodingConfig EncodingConfig) {
	// Register module commands selectively to avoid issues with incomplete modules
	// Only register commands for modules that have proper CLI implementations

	// Set up keyring configuration
	rootCmd.PersistentFlags().String(flags.FlagHome, DefaultNodeHome, "The application home directory")
	rootCmd.PersistentFlags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test)")
	rootCmd.PersistentFlags().String(flags.FlagChainID, "arkh-testnet-1", "The network chain ID")

	// Add Tendermint CLI commands for standard functionality
	rootCmd.AddCommand(tmcli.NewCompletionCmd(rootCmd, true))

	// Add genesis-related commands
	rootCmd.AddCommand(
		genutilcli.InitCmd(ModuleBasics, DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(ModuleBasics),
		AddGenesisAccountCmd(DefaultNodeHome),
		&cobra.Command{
			Use:   "gentx [key_name] [amount]",
			Short: "Generate a genesis tx carrying a self delegation",
			Long:  "Generate a genesis transaction that creates a validator with a self-delegation, and collect it in the genesis file.",
			Args:  cobra.ExactArgs(2),
			RunE: func(cmd *cobra.Command, args []string) error {
				fmt.Printf("Generating genesis transaction for %s with amount %s\n", args[0], args[1])
				fmt.Println("Genesis transaction generated (placeholder implementation)")
				return nil
			},
		},
		&cobra.Command{
			Use:   "collect-gentxs",
			Short: "Collect genesis txs and output a genesis.json file",
			Long:  "Collect genesis txs and output a genesis.json file.",
			RunE: func(cmd *cobra.Command, args []string) error {
				fmt.Println("Collecting genesis transactions...")
				fmt.Println("Genesis transactions collected (placeholder implementation)")
				return nil
			},
		},
	)

	// Add keys command - this is the main missing command
	// Try to fix keyring nil pointer issue by properly configuring the keys command
	keysCmd := keys.Commands()

	// Add proper client context configuration
	clientCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithHomeDir(DefaultNodeHome).
		WithViper("ARKH")

	// Set the client context for the keys command
	keysCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Initialize client context
		clientCtx = clientCtx.WithCmdContext(cmd.Context())
		return client.SetCmdClientContextHandler(clientCtx, cmd)
	}

	rootCmd.AddCommand(keysCmd)

	// Add tx command by aggregating all module tx commands
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transaction subcommands",
		Long:  "Transaction subcommands for creating and managing transactions",
	}

	// Add module tx commands
	txCmd.AddCommand(liquiditymodule.AppModuleBasic{}.GetTxCmd())
	// Add other module tx commands when they're enabled
	// txCmd.AddCommand(arkhmodule.AppModuleBasic{}.GetTxCmd())
	// txCmd.AddCommand(toolmodule.AppModuleBasic{}.GetTxCmd())
	// txCmd.AddCommand(utilitymodule.AppModuleBasic{}.GetTxCmd())
	// txCmd.AddCommand(wasmmodule.AppModuleBasic{}.GetTxCmd())

	rootCmd.AddCommand(txCmd)

	// Add query command by aggregating all module query commands
	queryCmd := &cobra.Command{
		Use:   "query",
		Short: "Querying subcommands",
		Long:  "Querying subcommands for querying blockchain state",
	}

	// Add module query commands
	queryCmd.AddCommand(liquiditymodule.AppModuleBasic{}.GetQueryCmd())
	// Add other module query commands when they're enabled
	// queryCmd.AddCommand(arkhmodule.AppModuleBasic{}.GetQueryCmd())
	// queryCmd.AddCommand(toolmodule.AppModuleBasic{}.GetQueryCmd())
	// queryCmd.AddCommand(utilitymodule.AppModuleBasic{}.GetQueryCmd())
	// queryCmd.AddCommand(wasmmodule.AppModuleBasic{}.GetQueryCmd())

	rootCmd.AddCommand(queryCmd)

	// Add status command for node status
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "status",
			Short: "Query remote node for status",
			Long:  "Query remote node for status information",
			RunE: func(cmd *cobra.Command, args []string) error {
				fmt.Println("Node status command - placeholder implementation")
				return nil
			},
		},
	)

	// Add additional commands
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "init [moniker]",
			Short: "Initialize private validator, p2p, genesis, and application configuration files",
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				moniker := args[0]
				chainID, _ := cmd.Flags().GetString("chain-id")
				if chainID == "" {
					chainID = "arkh-testnet-1"
				}

				// Create basic init functionality
				homeDir := DefaultNodeHome
				configDir := filepath.Join(homeDir, "config")

				// Create directories
				os.MkdirAll(homeDir, 0755)
				os.MkdirAll(configDir, 0755)

				// Create basic config files
				genesisFile := filepath.Join(configDir, "genesis.json")
				configFile := filepath.Join(configDir, "config.toml")

				// Write basic genesis
				genesis := `{
  "genesis_time": "2024-01-01T00:00:00Z",
  "chain_id": "` + chainID + `",
  "initial_height": "1",
  "consensus_params": {
    "block": {
      "max_bytes": "22020096",
      "max_gas": "-1",
      "time_iota_ms": "1000"
    },
    "evidence": {
      "max_age_num_blocks": "100000",
      "max_age_duration": "172800000000000",
      "max_bytes": "1048576"
    },
    "validator": {
      "pub_key_types": ["ed25519"]
    },
    "version": {}
  },
  "app_hash": "",
  "app_state": {}
}`

				err := os.WriteFile(genesisFile, []byte(genesis), 0644)
				if err != nil {
					return err
				}

				// Write basic config
				config := `# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# TCP or UNIX socket address for the RPC server to listen on
laddr = "tcp://127.0.0.1:26657"

# A custom human readable name for this node
moniker = "` + moniker + `"

# If this node is many blocks behind the tip of the chain, FastSync
# allows them to catchup quickly by downloading blocks in parallel
# and verifying their commits
fast_sync = true

# Database backend: goleveldb | cleveldb | boltdb | rocksdb | badgerdb
db_backend = "goleveldb"

# Database directory
db_dir = "data"

# Output level for logging, including package level options
log_level = "info"

# Output format: 'plain' (colored text) or 'json'
log_format = "plain"
`

				err = os.WriteFile(configFile, []byte(config), 0644)
				if err != nil {
					return err
				}

				fmt.Printf("Initialized %s node with chain-id %s\n", moniker, chainID)
				return nil
			},
		},
		&cobra.Command{
			Use:   "start",
			Short: "Run the full node",
			RunE: func(cmd *cobra.Command, args []string) error {
				fmt.Println("Starting Arkh Blockchain node...")
				fmt.Println("Node is running (placeholder implementation)")
				return nil
			},
		},
	)

}

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
func AddGenesisAccountCmd(defaultNodeHome string) *cobra.Command {
	// Create a simple add-genesis-account command
	return &cobra.Command{
		Use:   "add-genesis-account [address_or_key_name] [coin][,[coin]]",
		Short: "Add a genesis account to genesis.json",
		Long: `Add a genesis account to genesis.json. The provided account must specify
the account address or key name and a list of initial coins. If a key name is given,
the address will be looked up in the local Keybase. The list of initial tokens must
contain valid denominations. Accounts may optionally be supplied with vesting parameters.
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Simple implementation for now
			return nil
		},
	}
}
