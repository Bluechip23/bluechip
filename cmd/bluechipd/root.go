package main

import (
	"bufio"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"

	dbm "github.com/cometbft/cometbft-db"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankcli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distcli "github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	slashingcli "github.com/cosmos/cosmos-sdk/x/slashing/client/cli"
	stakingcli "github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/BlueChip23/bluechip/app"
	"github.com/BlueChip23/bluechip/app/params"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	tmcfg "github.com/cometbft/cometbft/config"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	snapshottypes "github.com/cosmos/cosmos-sdk/snapshots/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
)

// initAppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
func initAppConfig() (string, interface{}) {
	// The following code snippet is just for reference.

	type CustomAppConfig struct {
		serverconfig.Config
		Wasm wasmtypes.WasmConfig `mapstructure:"wasm"`
	}

	// Optionally allow the chain developer to overwrite the SDK's default
	// server config.
	srvCfg := serverconfig.DefaultConfig()
	// The SDK's default minimum gas price is set to "" (empty value) inside
	// app.toml. If left empty by validators, the node will halt on startup.
	// However, the chain developer can set a default app.toml value for their
	// validators here.
	//
	// In summary:
	// - if you leave srvCfg.MinGasPrices = "", all validators MUST tweak their
	//   own app.toml config,
	// - if you set srvCfg.MinGasPrices non-empty, validators CAN tweak their
	//   own app.toml to override, or use this default value.
	//
	// In simapp, we set the min gas prices to 0.
	srvCfg.MinGasPrices = "0stake"
	// srvCfg.BaseConfig.IAVLDisableFastNode = true // disable fastnode by default
	customAppConfig := CustomAppConfig{
		Config: *srvCfg,
		Wasm:   wasmtypes.DefaultWasmConfig(),
	}
	customAppTemplate := serverconfig.DefaultConfigTemplate +
		wasmtypes.DefaultConfigTemplate()
	return customAppTemplate, customAppConfig
}

// initTendermintConfig helps to override default Tendermint Config values.
// return tmcfg.DefaultConfig if no custom configuration is required for the application.
func initTendermintConfig() *tmcfg.Config {
	cfg := tmcfg.DefaultConfig()
	return cfg
}

// NewRootCmd creates a new root command for bluechipd. It is called once in the
// main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	encodingConfig := app.MakeEncodingConfig()

	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(app.Bech32PrefixAccAddr, app.Bech32PrefixAccPub)
	cfg.SetBech32PrefixForValidator(app.Bech32PrefixValAddr, app.Bech32PrefixValPub)
	cfg.SetBech32PrefixForConsensusNode(app.Bech32PrefixConsAddr, app.Bech32PrefixConsPub)
	cfg.SetAddressVerifier(wasmtypes.VerifyAddressLen())
	cfg.Seal()

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithHomeDir(app.DefaultNodeHome).
		WithViper("")

	rootCmd := &cobra.Command{
		Use:   version.AppName,
		Short: "Bluechip Smart Contract Zone",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}
			customAppTemplate, customAppConfig := initAppConfig()
			customTMConfig := initTendermintConfig()

			// override the default timeout commit value
			customTMConfig.Consensus.TimeoutCommit = 3000 * time.Millisecond

			return server.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig, customTMConfig)
		},
	}

	initRootCmd(rootCmd, encodingConfig)

	return rootCmd, encodingConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig) {
	gentxModule := app.ModuleBasics[genutiltypes.ModuleName].(genutil.AppModuleBasic)
	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome, gentxModule.GenTxValidator),
		genutilcli.MigrateGenesisCmd(),
		genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		AddGenesisAccountCmd(app.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		// testnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}),
		debug.Cmd(),
		config.Cmd(),
	)

	ac := appCreator{
		encCfg: encodingConfig,
	}
	server.AddCommands(rootCmd, app.DefaultNodeHome, ac.newApp, ac.appExport, addModuleInitFlags)

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		rpc.StatusCommand(), // Query Node Status
		queryCommand(),
		txCommand(),
		keyCommand(),
		CatchingUpCommand(),
		NodeIDCommand(),
		AddEncryptedKeyCommand(app.DefaultNodeHome),
		keys.Commands(app.DefaultNodeHome),
	)
}

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
	wasm.AddModuleInitFlags(startCmd)
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
		bankcli.GetBalancesCmd(),                         // Query the total balance of an account or of a specific denomination.
		bankcli.GetCmdQueryTotalSupply(),                 // Query total supply of tokens that are held by accounts
		distcli.GetCmdQueryDelegatorRewards(),            // Query all rewards earned by a delegator, optionally restrict to rewards from a single validator
		distcli.GetCmdQueryValidatorSlashes(),            // Query all slashes of a validator for a given block range
		distcli.GetCmdQueryValidatorOutstandingRewards(), // Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations.
		distcli.GetCmdQueryValidatorCommission(),         // Query distribution validator commission
		slashingcli.GetCmdQuerySigningInfo(),             // Check if you are Jailed or Tombstoned
		stakingcli.GetCmdQueryValidator(),                // Query a validator by validator address
		stakingcli.GetCmdQueryValidators(),               // Query all validators
		stakingcli.GetCmdQueryValidatorDelegations(),     // Query all delegations to one validator
	)

	app.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		stakingcli.NewCreateValidatorCmd(), // Create Validator
		distcli.NewWithdrawRewardsCmd(),    // Withdraw Rewards from Your Validator Address
		stakingcli.NewDelegateCmd(),        // Staking Coins
		stakingcli.GetCmdQueryDelegation(), // Query a Delegation
		bankcli.NewSendTxCmd(),             // Send tokens to another address
		distcli.NewWithdrawAllRewardsCmd(), // Withdraw all rewards for a single delegator.
		slashingcli.NewUnjailTxCmd(),       // Unjail Validator
		stakingcli.NewEditValidatorCmd(),   // Edit an existing validator's settings, such as commission rate, name, etc.
		stakingcli.NewDelegateCmd(),        // Delegate tokens to a validator
		stakingcli.NewUnbondCmd(),          // Unbond tokens from a validator
		stakingcli.NewRedelegateCmd(),      // Redelegate some tokens to another validator
	)

	app.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func keyCommand() *cobra.Command {
	cmd := keys.Commands(app.DefaultNodeHome)
	cmd.AddCommand(
		keys.AddKeyCommand(),    // Create New Keys for Validator
		keys.ShowKeysCmd(),      // Query the Keystore for Your Public Address
		keys.DeleteKeyCommand(), // Delete the given key
	)
	return cmd
}

// Check if Node is Catching Up
func CatchingUpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "catching-up",
		Short: "Check if the node is catching up with the network",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			status, err := clientCtx.Client.Status(context.Background())
			if err != nil {
				return err
			}
			cmd.Println("Is Catching Up:", status.SyncInfo.CatchingUp)
			return nil
		},
	}
}

// Get Node ID
func NodeIDCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "node-id",
		Short: "Get the node's ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			status, err := clientCtx.Client.Status(context.Background())
			if err != nil {
				return err
			}
			cmd.Println("Node ID:", status.NodeInfo.ID)
			return nil
		},
	}
}

// Add an encrypted private key, encrypt it, and save to disk.
func AddEncryptedKeyCommand(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-encrypted-key [name]",
		Short: "Add an encrypted private key",
		Long:  `Create a new encrypted private key, encrypt it with a passphrase, and save it to disk.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec
			inBuf := bufio.NewReader(cmd.InOrStdin())

			name := args[0]
			kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendFile, defaultNodeHome, inBuf, cdc)
			if err != nil {
				return err
			}

			// Get the passphrase
			passphrase, err := input.GetPassword("Enter passphrase to encrypt the private key:", inBuf)
			if err != nil {
				return err
			}

			// Add the key
			info, err := kb.NewAccount(name, passphrase, "", hd.CreateHDPath(118, 0, 0).String(), hd.Secp256k1)
			if err != nil {
				return err
			}

			addr, err := info.GetAddress()
			if err != nil {
				return err
			}

			pubkey, err := info.GetPubKey()
			if err != nil {
				return err
			}
			// Output the address and pubkey
			cmd.Println("Address:", addr.String())
			cmd.Println("PubKey:", pubkey.String())

			return nil
		},
	}

	return cmd
}

type appCreator struct {
	encCfg params.EncodingConfig
}

func (ac appCreator) newApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	var cache sdk.MultiStorePersistentCache

	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := dbm.NewDB("metadata", dbm.GoLevelDBBackend, snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	snapshotOptions := snapshottypes.NewSnapshotOptions(
		cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval)),
		cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent)),
	)

	var wasmOpts []wasm.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	return app.New(logger, db, traceStore, true, skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		ac.encCfg,
		app.GetEnabledProposals(),
		appOpts,
		wasmOpts,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshot(snapshotStore, snapshotOptions),
	)
}

func (ac appCreator) appExport(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
	modulesToExport []string,
) (servertypes.ExportedApp, error) {
	var bluechipApp *app.App
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home is not set")
	}

	loadLatest := height == -1
	var emptyWasmOpts []wasm.Option
	bluechipApp = app.New(
		logger,
		db,
		traceStore,
		loadLatest,
		map[int64]bool{},
		homePath,
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		ac.encCfg,
		app.GetEnabledProposals(),
		appOpts,
		emptyWasmOpts,
	)

	if height != -1 {
		if err := bluechipApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return bluechipApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs, modulesToExport)
}
