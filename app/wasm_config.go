package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultBluechipInstanceCost is initially set the same as in wasmd
	DefaultBluechipInstanceCost uint64 = 60_000
	// DefaultBluechipCompileCost set to a large number for testing
	DefaultBluechipCompileCost uint64 = 3
)

// BluechipGasRegisterConfig is defaults plus a custom compile amount
func BluechipGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultBluechipInstanceCost
	gasConfig.CompileCost = DefaultBluechipCompileCost

	return gasConfig
}

func NewBluechipWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(BluechipGasRegisterConfig())
}
