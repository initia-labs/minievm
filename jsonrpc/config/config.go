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
	DefaultEnableUnsafeCORS = false
	// DefaultMaxOpenConnections is the default maximum number of simultaneous connections
	// for the server listener.
	DefaultMaxOpenConnections = 100
	// DefaultLogsCap is the default max number of results can be returned from single `eth_getLogs` query.
	DefaultLogsCap int32 = 100
	// DefaultFilterCap is the default global cap for total number of filters that can be created.
	DefaultFilterCap int32 = 200
	// DefaultBlockRangeCap is the default max block range allowed for `eth_getLogs` query.
	DefaultBlockRangeCap int32 = 100
	// DefaultAddress defines the default HTTP server to listen on.
	DefaultAddress = "127.0.0.1:8545"
	// DefaultAddressWS defines the default WebSocket server address to bind to.
	DefaultAddressWS = "127.0.0.1:8546"
	// DefaultBatchRequestLimit is the default maximum number of requests in a batch
	DefaultBatchRequestLimit = 1000
	// DefaultBatchResponseMaxSize is the default maximum number of bytes returned from a batched call
	DefaultBatchResponseMaxSize = 25 * 1000 * 1000
	// DefaultQueuedTransactionCap is the default maximum number of queued transactions that can be in the transaction pool.
	DefaultQueuedTransactionCap = 1000
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
	flagJSONRPCLogsCap              = "json-rpc.logs-cap"
	flagJSONRPCFilterCap            = "json-rpc.filter-cap"
	flagJSONRPCBlockRangeCap        = "json-rpc.block-range-cap"
	flagJSONRPCHTTPTimeout          = "json-rpc.http-timeout"
	flagJSONRPCHTTPIdleTimeout      = "json-rpc.http-idle-timeout"
	flagJSONRPCMaxOpenConnections   = "json-rpc.max-open-connections"
	flagJSONRPCBatchRequestLimit    = "json-rpc.batch-request-limit"
	flagJSONRPCBatchResponseMaxSize = "json-rpc.batch-response-max-size"
	flagJSONRPCQueuedTransactionCap = "json-rpc.queued-transaction-cap"
	flagJSONRPCQueuedTransactionTTL = "json-rpc.queued-transaction-ttl"
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
	// FilterCap is the global cap for total number of filters that can be created.
	FilterCap int32 `mapstructure:"filter-cap"`
	// BlockRangeCap defines the max block range allowed for `eth_getLogs` query.
	BlockRangeCap int32 `mapstructure:"block-range-cap"`
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
	// QueuedTransactionCap is a maximum number of queued transactions that can be in the transaction pool.
	QueuedTransactionCap int `mapstructure:"queued-transaction-cap"`
}

// DefaultJSONRPCConfig returns a default configuration for the EVM RPC server.
func DefaultJSONRPCConfig() JSONRPCConfig {
	return JSONRPCConfig{
		Enable:           DefaultEnable,
		Address:          DefaultAddress,
		EnableWS:         DefaultEnable,
		AddressWS:        DefaultAddressWS,
		EnableUnsafeCORS: DefaultEnableUnsafeCORS,

		APIs: DefaultAPIs,

		FilterCap:     DefaultFilterCap,
		BlockRangeCap: DefaultBlockRangeCap,

		HTTPTimeout:        DefaultHTTPTimeout,
		HTTPIdleTimeout:    DefaultHTTPIdleTimeout,
		MaxOpenConnections: DefaultMaxOpenConnections,

		BatchRequestLimit:    DefaultBatchRequestLimit,
		BatchResponseMaxSize: DefaultBatchResponseMaxSize,

		QueuedTransactionCap: DefaultQueuedTransactionCap,
	}
}

// AddConfigFlags adds flags for a EVM RPC server to the StartCmd.
func AddConfigFlags(startCmd *cobra.Command) {
	startCmd.Flags().Bool(flagJSONRPCEnable, DefaultEnable, "Enable the EVM RPC server")
	startCmd.Flags().String(flagJSONRPCAddress, DefaultAddress, "Address to listen on for the EVM RPC server")
	startCmd.Flags().Bool(flagJSONRPCEnableWS, DefaultEnableWS, "Enable the WebSocket server")
	startCmd.Flags().String(flagJSONRPCAddressWS, DefaultAddressWS, "Address to listen on for the WebSocket server")
	startCmd.Flags().Bool(flagJSONRPCEnableUnsafeCORS, DefaultEnableUnsafeCORS, "Enable unsafe CORS")
	startCmd.Flags().StringSlice(flagJSONRPCAPIs, DefaultAPIs, "List of JSON-RPC namespaces that should be enabled")
	startCmd.Flags().Int32(flagJSONRPCFilterCap, DefaultFilterCap, "Sets the global cap for total number of filters that can be created")
	startCmd.Flags().Int32(flagJSONRPCBlockRangeCap, DefaultBlockRangeCap, "Max block range allowed for 'eth_getLogs' query")
	startCmd.Flags().Duration(flagJSONRPCHTTPTimeout, DefaultHTTPTimeout, "Read/write timeout of http json-rpc server")
	startCmd.Flags().Duration(flagJSONRPCHTTPIdleTimeout, DefaultHTTPIdleTimeout, "Idle timeout of http json-rpc server")
	startCmd.Flags().Int(flagJSONRPCMaxOpenConnections, DefaultMaxOpenConnections, "Maximum number of simultaneous connections for the server listener")
	startCmd.Flags().Int(flagJSONRPCBatchRequestLimit, DefaultBatchRequestLimit, "Maximum number of requests in a batch")
	startCmd.Flags().Int(flagJSONRPCBatchResponseMaxSize, DefaultBatchResponseMaxSize, "Maximum number of bytes returned from a batched call")
	startCmd.Flags().Int(flagJSONRPCQueuedTransactionCap, DefaultQueuedTransactionCap, "Maximum number of queued transactions that can be in the transaction pool")
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
		FilterCap:            cast.ToInt32(appOpts.Get(flagJSONRPCFilterCap)),
		BlockRangeCap:        cast.ToInt32(appOpts.Get(flagJSONRPCBlockRangeCap)),
		HTTPTimeout:          cast.ToDuration(appOpts.Get(flagJSONRPCHTTPTimeout)),
		HTTPIdleTimeout:      cast.ToDuration(appOpts.Get(flagJSONRPCHTTPIdleTimeout)),
		MaxOpenConnections:   cast.ToInt(appOpts.Get(flagJSONRPCMaxOpenConnections)),
		BatchRequestLimit:    cast.ToInt(appOpts.Get(flagJSONRPCBatchRequestLimit)),
		BatchResponseMaxSize: cast.ToInt(appOpts.Get(flagJSONRPCBatchResponseMaxSize)),
		QueuedTransactionCap: cast.ToInt(appOpts.Get(flagJSONRPCQueuedTransactionCap)),
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

# FilterCap is the global cap for total number of filters that can be created.
filter-cap = {{ .JSONRPCConfig.FilterCap }}

# BlockRangeCap defines the max block range allowed for 'eth_getLogs' query.
block-range-cap = {{ .JSONRPCConfig.BlockRangeCap }}

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

# QueuedTransactionCap is the maximum number of queued transactions that 
# can be in the transaction pool.
queued-transaction-cap = {{ .JSONRPCConfig.QueuedTransactionCap }}
`
