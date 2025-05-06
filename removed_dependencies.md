#Cosmos 50.X upgrade

##removed Dependencies
* github.com/cosmos/cosmos-sdk/x/crisis
- Deprecated https://github.com/cosmos/cosmos-sdk/tree/main/x#deprecated-modules
* cosmossdk.io/x/upgrade/client
- Removed from the base repo. No clue where this goes? look into auto client gens.
* github.com/cosmos/cosmos-sdk/x/slashing/client/cli
- Removed from the base repo. No clue where this goes? look into auto client gens.
* github.com/cosmos/ibc-go/v8/modules/core/02-client/client
- Removed from the base repo. No clue where this goes? look into auto client gens.
* github.com/cosmos/cosmos-sdk/snapshots
* Removed from the base repo. No clue where this goes? Think it's deprecated?
* cosmossdk.io/api
- Reverted from v0.8 to v0.7. Because they wanted to deprecate cosmos-sdk/module/crisis but it is still used by wasmd.
* x/pageinflation/keeper/msg_server_creator_pool_mint
- Deleted the truncation of int. Maybe this is not needed at all?
- Added missing app modules. see todos there. How do we fix it? (Note: Use the new ignite app to check how they define modules?)
* cmd/bluechipd/main.go
- "github.com/cosmos/cosmos-sdk/server" doesn't have error code anymore??
* /app/app.go
* GetEnabledProposals() enabledProposals for WASM was deprecated. what is this neeed for? https://github.com/CosmWasm/wasmd/blob/v0.55.0/x/wasm/types/proposal_legacy.go#L38
  * Also modified to add a route in app.go govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.wasmKeeper, enabledProposals))
* Address codec added to the new authkeeper.NewAccountKeeper. No clue what this is and what needs to actually be there.
* Eventservice is added to consensusparamkeeper.NewKeeper. Set as null right now
* stakingkeeper.NewKeeper has two new address codecs. What is this used for?
- "app.go"
* IBCKeeper needs an authority. No clue what this is since they just change everything for no reason.
* AddRoute(upgradetypes.RouterKey, upgradetypes.NewSoftwareUpgradeProposal((app.UpgradeKeeper))) Upgrade handler decided to deprecate and chagne to msg handlers. no clue what this is.
* ibctransferkeeper also added authority. No documentation in the code on why they like adding random parameters.
* icahostkeeper also added authority. No documentation in the code on why they like adding random parameters.
* icacontrollerkeeper also added authority. No documentation in the code on why they like adding random parameters.
* evidencekeeper added new address code and block info params. put null here cause who knows?
* AddRoute(wasm.ModuleName, wasm.NewIBCHandler(app.wasmKeeper, app.IBCKeeper.ChannelKeeper, app.IBCKeeper.ChannelKeeper)). Removed for now just so we can compile. Wasm changes paths for everything
* GovKeeper added a distribution keeper. SO need more keepers in the code.
* genutil.NewAppModule Changed the app.BaseApp.DeliverTX to a private method. have to figure out what to use now.
* newappmodule - added a new required parameter for an address codec
* slashing.newAppModule, added a new interface registry but i don't know what it is
* app init chainer doesn't exist anymore.
* BeginBlockers magically disappeared as well.
* app.UpgradeKeeper.SetUpgradeHandler type problem but when it's an anonymous function it's hard to tell
* RegisterNodeService added a new config to the parameters. no clue what it is
- "app/export.go"
* , tmproto.Header{Height: app.LastBlockHeight()} removed?? from new app context
* DistrKeeper - GetFeePool Disappeared!!
* Every address type changed to a byte array instead of string. So i updated types to convert to byte array
* KVStoreReversePrefixIterator doesn't exist in the sdk. where does this go now??
* applyAllowedAddrs removed but it will come back.
- "cmd/bluechip/root.go
* initRootcmd - Added address codec to collectGenTxsCmd??? added migrationMap???? added address codec to GenTxCmd??? config.CMD is gone????
* queryCommand -> AddCommand all of the authcmd,rpc,bankcli, distcli, stakingcli commands were removed??
* stakingcli.GetCmdQueryDelegation no longer exists??????
* txCommand -> AddCommand, Added nil to a lot of tx cmds.  Where do we get address codecs??
* newApp -> sdk.MultiStorePersistentCache removed ffrom types