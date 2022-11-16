package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultbluechipInstanceCost is initially set the same as in wasmd
	DefaultbluechipInstanceCost uint64 = 60_000
	// DefaultbluechipCompileCost set to a large number for testing
	DefaultbluechipCompileCost uint64 = 3
)

// bluechipGasRegisterConfig is defaults plus a custom compile amount
func bluechipGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultbluechipInstanceCost
	gasConfig.CompileCost = DefaultbluechipCompileCost

	return gasConfig
}

func NewbluechipWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(bluechipGasRegisterConfig())
}
