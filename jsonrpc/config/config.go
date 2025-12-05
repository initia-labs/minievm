package config

import (
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

const (
	// DefaultEnable defines the default value for enabling the EVM RPC server.
	DefaultEnable = true
	// DefaultHTTPTimeout is the default read/write timeout of http json-rpc server.
	DefaultHTTPTimeout = 10 * time.Second
	// DefaultEnableWS defines the default value for enabling the WebSocket server.
	DefaultEnableWS = true
	// DefaultHTTPIdleTimeout is the default idle timeout of http json-rpc server.
	DefaultHTTPIdleTimeout = 120 * time.Second
	// DefaultEnableUnsafeCORS defines the default value for enabling unsafe CORS.
	DefaultEnableUnsafeCORS = true
	// DefaultMaxOpenConnections is the default maximum number of simultaneous connections
	// for the server listener.
	DefaultMaxOpenConnections = 1000
	// DefaultAddress defines the default HTTP server to listen on.
	DefaultAddress = "127.0.0.1:8545"
	// DefaultAddressWS defines the default WebSocket server address to bind to.
	DefaultAddressWS = "127.0.0.1:8546"
	// DefaultBatchRequestLimit is the default maximum number of requests in a batch
	DefaultBatchRequestLimit = 1000
	// DefaultBatchResponseMaxSize is the default maximum number of bytes returned from a batched call
	DefaultBatchResponseMaxSize = 25 * 1000 * 1000
	// DefaultFeeHistoryMaxHeaders is the default maximum number of headers, which can be used to lookup the fee history.
	DefaultFeeHistoryMaxHeaders = 1024
	// DefaultFeeHistoryMaxBlocks is the default maximum number of blocks, which can be used to lookup the fee history.
	DefaultFeeHistoryMaxBlocks = 1024
	// DefaultFilterTimeout is the default filter timeout, how long filters stay active.
	DefaultFilterTimeout = 5 * time.Minute
	// DefaultFilterMaxBlockRange is the default maximum number of blocks that can be queried in a filter.
	DefaultFilterMaxBlockRange = 1_000_000
	// DefaultFilterMaxAddresses is the default maximum number of addresses that can be used in a log filter.
	DefaultFilterMaxAddresses = 100
	// DefaultLogCacheSize is the maximum number of cached blocks.
	DefaultLogCacheSize = 32
	// DefaultGasMultiplier is the default gas multiplier for the EVM state transition.
	DefaultGasMultiplier = "1.4"
	// DefaultTracerTimeout is the default tracer timeout.
	DefaultTracerTimeout = 10 * time.Second
)

var (
	// DefaultAPIs defines the default list of JSON-RPC namespaces that should be enabled.
	DefaultAPIs = []string{"eth", "net", "txpool", "web3", "personal", "debug", "miner", "cosmos"}
)

const (
	flagJSONRPCEnable               = "json-rpc.enable"
	flagJSONRPCAddress              = "json-rpc.address"
	flagJSONRPCEnableWS             = "json-rpc.enable-ws"
	flagJSONRPCAddressWS            = "json-rpc.address-ws"
	flagJSONRPCEnableUnsafeCORS     = "json-rpc.enable-unsafe-cors"
	flagJSONRPCAPIs                 = "json-rpc.apis"
	flagJSONRPCHTTPTimeout          = "json-rpc.http-timeout"
	flagJSONRPCHTTPIdleTimeout      = "json-rpc.http-idle-timeout"
	flagJSONRPCMaxOpenConnections   = "json-rpc.max-open-connections"
	flagJSONRPCBatchRequestLimit    = "json-rpc.batch-request-limit"
	flagJSONRPCBatchResponseMaxSize = "json-rpc.batch-response-max-size"
	flagJSONRPCFeeHistoryMaxHeaders = "json-rpc.fee-history-max-headers"
	flagJSONRPCFeeHistoryMaxBlocks  = "json-rpc.fee-history-max-blocks"
	flagJSONRPCFilterTimeout        = "json-rpc.filter-timeout"
	flagJSONRPCLogCacheSize         = "json-rpc.log-cache-size"
	flagJSONRPCGasMultiplier        = "json-rpc.gas-multiplier"
	flagJSONRPCFilterMaxBlockRange  = "json-rpc.filter-max-block-range"
	flagJSONRPCFilterMaxAddresses   = "json-rpc.filter-max-addresses"
	flagJSONRPCTracerTimeout        = "json-rpc.tracer-timeout"
)

// JSONRPCConfig defines configuration for the EVM RPC server.
type JSONRPCConfig struct {
	// Enable defines if the EVM RPC server should be enabled.
	Enable bool `mapstructure:"enable"`
	// Address defines the HTTP server to listen on
	Address string `mapstructure:"address"`
	// EnableWS defines if the WebSocket server should be enabled.
	EnableWS bool `mapstructure:"enable-ws"`
	// AddressWS defines the WebSocket server address to bind to.
	AddressWS string `mapstructure:"address-ws"`
	// EnableUnsafeCORS defines if the EVM RPC server should enable unsafe CORS.
	EnableUnsafeCORS bool `mapstructure:"enable-unsafe-cors"`
	// API defines a list of JSON-RPC namespaces that should be enabled
	APIs []string `mapstructure:"apis"`
	// HTTPTimeout is the read/write timeout of http json-rpc server.
	HTTPTimeout time.Duration `mapstructure:"http-timeout"`
	// HTTPIdleTimeout is the idle timeout of http json-rpc server.
	HTTPIdleTimeout time.Duration `mapstructure:"http-idle-timeout"`
	// MaxOpenConnections sets the maximum number of simultaneous connections
	// for the server listener.
	MaxOpenConnections int `mapstructure:"max-open-connections"`
	// Maximum number of requests in a batch
	BatchRequestLimit int `mapstructure:"batch-request-limit"`
	// Maximum number of bytes returned from a batched call
	BatchResponseMaxSize int `mapstructure:"batch-response-max-size"`
	// FeeHistoryMaxHeaders is the maximum number of headers, which can be used to lookup the fee history.
	FeeHistoryMaxHeaders int `mapstructure:"fee-history-max-headers"`
	// FeeHistoryMaxBlocks is the maximum number of blocks, which can be used to lookup the fee history.
	FeeHistoryMaxBlocks int `mapstructure:"fee-history-max-blocks"`
	// FilterTimeout is a duration how long filters stay active (default: 5min)
	FilterTimeout time.Duration `mapstructure:"filter-timeout"`
	// FilterMaxBlockRange is the maximum number of blocks that can be queried in a filter.
	FilterMaxBlockRange int `mapstructure:"filter-max-block-range"`
	// LogCacheSize is the maximum number of cached blocks.
	LogCacheSize int `mapstructure:"log-cache-size"`
	// GasMultiplier is the gas multiplier for the EVM state transition.
	GasMultiplier string `mapstructure:"gas-multiplier"`
	// FilterMaxAddresses is the maximum number of addresses that can be used in a log filter.
	FilterMaxAddresses int `mapstructure:"filter-max-addresses"`
	// TracerTimeout is the timeout for the tracer.
	TracerTimeout time.Duration `mapstructure:"tracer-timeout"`
}

// DefaultJSONRPCConfig returns a default configuration for the EVM RPC server.
func DefaultJSONRPCConfig() JSONRPCConfig {
	return JSONRPCConfig{
		Enable:           DefaultEnable,
		Address:          DefaultAddress,
		EnableWS:         DefaultEnableWS,
		AddressWS:        DefaultAddressWS,
		EnableUnsafeCORS: DefaultEnableUnsafeCORS,

		APIs: DefaultAPIs,

		HTTPTimeout:        DefaultHTTPTimeout,
		HTTPIdleTimeout:    DefaultHTTPIdleTimeout,
		MaxOpenConnections: DefaultMaxOpenConnections,

		BatchRequestLimit:    DefaultBatchRequestLimit,
		BatchResponseMaxSize: DefaultBatchResponseMaxSize,

		FeeHistoryMaxHeaders: DefaultFeeHistoryMaxHeaders,
		FeeHistoryMaxBlocks:  DefaultFeeHistoryMaxBlocks,

		FilterTimeout:       DefaultFilterTimeout,
		FilterMaxBlockRange: DefaultFilterMaxBlockRange,
		FilterMaxAddresses:  DefaultFilterMaxAddresses,
		LogCacheSize:        DefaultLogCacheSize,

		GasMultiplier: DefaultGasMultiplier,

		TracerTimeout: DefaultTracerTimeout,
	}
}

// AddConfigFlags adds flags for a EVM RPC server to the StartCmd.
func AddConfigFlags(startCmd *cobra.Command) {
	startCmd.Flags().Bool(flagJSONRPCEnable, DefaultEnable, "Enable the EVM RPC server")
	startCmd.Flags().String(flagJSONRPCAddress, DefaultAddress, "Address to listen on for the EVM RPC server")
	startCmd.Flags().Bool(flagJSONRPCEnableWS, DefaultEnableWS, "Enable the WebSocket server")
	startCmd.Flags().String(flagJSONRPCAddressWS, DefaultAddressWS, "Address to listen on for the WebSocket server")
	startCmd.Flags().Bool(flagJSONRPCEnableUnsafeCORS, DefaultEnableUnsafeCORS, "Enable unsafe CORS")
	startCmd.Flags().String(flagJSONRPCAPIs, strings.Join(DefaultAPIs, ","), "List of JSON-RPC namespaces that should be enabled")
	startCmd.Flags().Duration(flagJSONRPCHTTPTimeout, DefaultHTTPTimeout, "Read/write timeout of http json-rpc server")
	startCmd.Flags().Duration(flagJSONRPCHTTPIdleTimeout, DefaultHTTPIdleTimeout, "Idle timeout of http json-rpc server")
	startCmd.Flags().Int(flagJSONRPCMaxOpenConnections, DefaultMaxOpenConnections, "Maximum number of simultaneous connections for the server listener")
	startCmd.Flags().Int(flagJSONRPCBatchRequestLimit, DefaultBatchRequestLimit, "Maximum number of requests in a batch")
	startCmd.Flags().Int(flagJSONRPCBatchResponseMaxSize, DefaultBatchResponseMaxSize, "Maximum number of bytes returned from a batched call")
	startCmd.Flags().Int(flagJSONRPCFeeHistoryMaxHeaders, DefaultFeeHistoryMaxHeaders, "Maximum number of headers used to lookup the fee history")
	startCmd.Flags().Int(flagJSONRPCFeeHistoryMaxBlocks, DefaultFeeHistoryMaxBlocks, "Maximum number of blocks used to lookup the fee history")
	startCmd.Flags().Duration(flagJSONRPCFilterTimeout, DefaultFilterTimeout, "Duration how long filters stay active")
	startCmd.Flags().Int(flagJSONRPCFilterMaxBlockRange, DefaultFilterMaxBlockRange, "Maximum number of blocks that can be queried in a filter")
	startCmd.Flags().Int(flagJSONRPCFilterMaxAddresses, DefaultFilterMaxAddresses, "Maximum number of addresses that can be used in a log filter")
	startCmd.Flags().Int(flagJSONRPCLogCacheSize, DefaultLogCacheSize, "Maximum number of cached blocks for the log filter")
	startCmd.Flags().String(flagJSONRPCGasMultiplier, DefaultGasMultiplier, "Gas multiplier for the EVM state transition")
	startCmd.Flags().Duration(flagJSONRPCTracerTimeout, DefaultTracerTimeout, "Timeout for the tracer")
}

// GetConfig load config values from the app options
func GetConfig(appOpts servertypes.AppOptions) JSONRPCConfig {
	return JSONRPCConfig{
		Enable:               cast.ToBool(appOpts.Get(flagJSONRPCEnable)),
		Address:              cast.ToString(appOpts.Get(flagJSONRPCAddress)),
		EnableWS:             cast.ToBool(appOpts.Get(flagJSONRPCEnableWS)),
		AddressWS:            cast.ToString(appOpts.Get(flagJSONRPCAddressWS)),
		EnableUnsafeCORS:     cast.ToBool(appOpts.Get(flagJSONRPCEnableUnsafeCORS)),
		APIs:                 strings.Split(cast.ToString(appOpts.Get(flagJSONRPCAPIs)), ","),
		HTTPTimeout:          cast.ToDuration(appOpts.Get(flagJSONRPCHTTPTimeout)),
		HTTPIdleTimeout:      cast.ToDuration(appOpts.Get(flagJSONRPCHTTPIdleTimeout)),
		MaxOpenConnections:   cast.ToInt(appOpts.Get(flagJSONRPCMaxOpenConnections)),
		BatchRequestLimit:    cast.ToInt(appOpts.Get(flagJSONRPCBatchRequestLimit)),
		BatchResponseMaxSize: cast.ToInt(appOpts.Get(flagJSONRPCBatchResponseMaxSize)),
		FeeHistoryMaxHeaders: cast.ToInt(appOpts.Get(flagJSONRPCFeeHistoryMaxHeaders)),
		FeeHistoryMaxBlocks:  cast.ToInt(appOpts.Get(flagJSONRPCFeeHistoryMaxBlocks)),
		FilterTimeout:        cast.ToDuration(appOpts.Get(flagJSONRPCFilterTimeout)),
		FilterMaxBlockRange:  cast.ToInt(appOpts.Get(flagJSONRPCFilterMaxBlockRange)),
		FilterMaxAddresses:   cast.ToInt(appOpts.Get(flagJSONRPCFilterMaxAddresses)),
		LogCacheSize:         cast.ToInt(appOpts.Get(flagJSONRPCLogCacheSize)),
		GasMultiplier:        cast.ToString(appOpts.Get(flagJSONRPCGasMultiplier)),
		TracerTimeout:        cast.ToDuration(appOpts.Get(flagJSONRPCTracerTimeout)),
	}
}

// DefaultConfigTemplate defines the configuration template for the EVM RPC configuration
const DefaultConfigTemplate = `
###############################################################################
###                           JSON RPC Configuration                        ###
###############################################################################

[json-rpc]

# Enable defines if the EVM RPC server should be enabled.
enable = {{ .JSONRPCConfig.Enable }}

# Address defines the EVM RPC HTTP server address to bind to.
address = "{{ .JSONRPCConfig.Address }}"

# Enable defines if the gRPC websocket server should be enabled.
enable-ws = {{ .JSONRPCConfig.EnableWS }}

# WsAddress defines the EVM RPC WebSocket server address to bind to.
address-ws = "{{ .JSONRPCConfig.AddressWS }}"

# EnableUnsafeCORS defines if the EVM RPC server should enable unsafe CORS.
enable-unsafe-cors = {{ .JSONRPCConfig.EnableUnsafeCORS }}

# API defines a list of JSON-RPC namespaces that should be enabled
# Example: "eth,txpool,personal,net,debug,web3"
apis = "{{range $index, $elmt := .JSONRPCConfig.APIs}}{{if $index}},{{$elmt}}{{else}}{{$elmt}}{{end}}{{end}}"

# HTTPTimeout is the read/write timeout of http json-rpc server.
http-timeout = "{{ .JSONRPCConfig.HTTPTimeout }}"

# HTTPIdleTimeout is the idle timeout of http json-rpc server.
http-idle-timeout = "{{ .JSONRPCConfig.HTTPIdleTimeout }}"

# MaxOpenConnections sets the maximum number of simultaneous connections
# for the server listener.
max-open-connections = {{ .JSONRPCConfig.MaxOpenConnections }}

# Maximum number of requests in a batch
batch-request-limit = {{ .JSONRPCConfig.BatchRequestLimit }}

# Maximum number of bytes returned from a batched call
batch-response-max-size = {{ .JSONRPCConfig.BatchResponseMaxSize }}

# FeeHistoryMaxHeaders is the maximum number of headers, which can be used to lookup the fee history.
fee-history-max-headers = {{ .JSONRPCConfig.FeeHistoryMaxHeaders }}

# FeeHistoryMaxBlocks is the maximum number of blocks, which can be used to lookup the fee history.
fee-history-max-blocks = {{ .JSONRPCConfig.FeeHistoryMaxBlocks }}

# FilterTimeout is a duration how long filters stay active (default: 5min)
filter-timeout = "{{ .JSONRPCConfig.FilterTimeout }}"

# FilterMaxBlockRange is the maximum number of blocks that can be queried in a filter.
filter-max-block-range = {{ .JSONRPCConfig.FilterMaxBlockRange }}

# FilterMaxAddresses is the maximum number of addresses that can be used in a log filter.
filter-max-addresses = {{ .JSONRPCConfig.FilterMaxAddresses }}

# LogCacheSize is the maximum number of cached blocks for the log filter.
log-cache-size = {{ .JSONRPCConfig.LogCacheSize }}

# GasMultiplier is the gas multiplier for the EVM state transition.
gas-multiplier = "{{ .JSONRPCConfig.GasMultiplier }}"

# TracerTimeout is the timeout for the tracer.
tracer-timeout = "{{ .JSONRPCConfig.TracerTimeout }}"
`
