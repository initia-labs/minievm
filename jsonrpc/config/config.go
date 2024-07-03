package config

import (
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

const (
	// DefaultEnable defines the default value for enabling the EVM RPC server.
	DefaultEnable = true
	// DefaultEnableUnsafeCORS defines the default value for enabling unsafe CORS.
	DefaultEnableUnsafeCORS = false
	// DefaultHTTPTimeout is the default read/write timeout of http json-rpc server.
	DefaultHTTPTimeout = 10 * time.Second
	// DefaultHTTPIdleTimeout is the default idle timeout of http json-rpc server.
	DefaultHTTPIdleTimeout = 120 * time.Second
	// DefaultMaxOpenConnections is the default maximum number of simultaneous connections
	// for the server listener.
	DefaultMaxOpenConnections = 100
	// DefaultLogsCap is the default max number of results can be returned from single `eth_getLogs` query.
	DefaultLogsCap = 100
	// DefaultBlockRangeCap is the default max block range allowed for `eth_getLogs` query.
	DefaultBlockRangeCap = 100
	// DefaultMetricsAddress defines the default EVM Metrics server address to bind to.
	DefaultMetricsAddress = "127.0.0.1:6065"
	// DefaultAddress defines the default HTTP server to listen on.
	DefaultAddress = "127.0.0.1:8545"
	// DefaultFilterCap
	DefaultFilterCap int32 = 200
)

var (
	// DefaultAPIs defines the default list of JSON-RPC namespaces that should be enabled.
	DefaultAPIs = []string{"eth" /*"txpool", "personal", "net", "debug", "web3"*/}
)

const (
	flagJSONRPCEnable             = "json-rpc.enable"
	flagJSONRPCEnableUnsafeCORS   = "json-rpc.enable-unsafe-cors"
	flagJSONRPCAddress            = "json-rpc.address"
	flagJSONRPCAPIs               = "json-rpc.apis"
	flagJSONRPCLogsCap            = "json-rpc.logs-cap"
	flagJSONRPCFilterCap          = "json-rpc.filter-cap"
	flagJSONRPCBlockRangeCap      = "json-rpc.block-range-cap"
	flagJSONRPCHTTPTimeout        = "json-rpc.http-timeout"
	flagJSONRPCHTTPIdleTimeout    = "json-rpc.http-idle-timeout"
	flagJSONRPCMaxOpenConnections = "json-rpc.max-open-connections"
	flagJSONRPCMetricsAddress     = "json-rpc.metrics-address"
)

// JSONRPCConfig defines configuration for the EVM RPC server.
type JSONRPCConfig struct {
	// Enable defines if the EVM RPC server should be enabled.
	Enable bool `mapstructure:"enable"`
	// EnableUnsafeCORS defines if the EVM RPC server should enable unsafe CORS.
	EnableUnsafeCORS bool `mapstructure:"enable-unsafe-cors"`
	// Address defines the HTTP server to listen on
	Address string `mapstructure:"address"`
	// API defines a list of JSON-RPC namespaces that should be enabled
	APIs []string `mapstructure:"apis"`
	// FilterCap is the global cap for total number of filters that can be created.
	FilterCap int32 `mapstructure:"filter-cap"`
	// LogsCap defines the max number of results can be returned from single `eth_getLogs` query.
	LogsCap int32 `mapstructure:"logs-cap"`
	// BlockRangeCap defines the max block range allowed for `eth_getLogs` query.
	BlockRangeCap int32 `mapstructure:"block-range-cap"`
	// HTTPTimeout is the read/write timeout of http json-rpc server.
	HTTPTimeout time.Duration `mapstructure:"http-timeout"`
	// HTTPIdleTimeout is the idle timeout of http json-rpc server.
	HTTPIdleTimeout time.Duration `mapstructure:"http-idle-timeout"`
	// MaxOpenConnections sets the maximum number of simultaneous connections
	// for the server listener.
	MaxOpenConnections int `mapstructure:"max-open-connections"`
	// MetricsAddress defines the metrics server to listen on
	MetricsAddress string `mapstructure:"metrics-address"`
}

// DefaultJSONRPCConfig returns a default configuration for the EVM RPC server.
func DefaultJSONRPCConfig() JSONRPCConfig {
	return JSONRPCConfig{
		Enable:             DefaultEnable,
		EnableUnsafeCORS:   DefaultEnableUnsafeCORS,
		Address:            DefaultAddress,
		APIs:               DefaultAPIs,
		FilterCap:          DefaultFilterCap,
		LogsCap:            DefaultLogsCap,
		BlockRangeCap:      DefaultBlockRangeCap,
		HTTPTimeout:        DefaultHTTPTimeout,
		HTTPIdleTimeout:    DefaultHTTPIdleTimeout,
		MaxOpenConnections: DefaultMaxOpenConnections,
		MetricsAddress:     DefaultMetricsAddress,
	}
}

// AddConfigFlags adds flags for a EVM RPC server to the StartCmd.
func AddConfigFlags(startCmd *cobra.Command) {
	startCmd.Flags().Bool(flagJSONRPCEnable, DefaultEnable, "Enable the EVM RPC server")
	startCmd.Flags().Bool(flagJSONRPCEnableUnsafeCORS, DefaultEnableUnsafeCORS, "Enable unsafe CORS")
	startCmd.Flags().String(flagJSONRPCAddress, DefaultAddress, "Address to listen on for the EVM RPC server")
	startCmd.Flags().StringSlice(flagJSONRPCAPIs, DefaultAPIs, "List of JSON-RPC namespaces that should be enabled")
	startCmd.Flags().Int32(flagJSONRPCLogsCap, DefaultLogsCap, "Max number of results can be returned from single 'eth_getLogs' query")
	startCmd.Flags().Int32(flagJSONRPCFilterCap, DefaultFilterCap, "Sets the global cap for total number of filters that can be created")
	startCmd.Flags().Int32(flagJSONRPCBlockRangeCap, DefaultBlockRangeCap, "Max block range allowed for 'eth_getLogs' query")
	startCmd.Flags().Duration(flagJSONRPCHTTPTimeout, DefaultHTTPTimeout, "Read/write timeout of http json-rpc server")
	startCmd.Flags().Duration(flagJSONRPCHTTPIdleTimeout, DefaultHTTPIdleTimeout, "Idle timeout of http json-rpc server")
	startCmd.Flags().Int(flagJSONRPCMaxOpenConnections, DefaultMaxOpenConnections, "Maximum number of simultaneous connections for the server listener")
	startCmd.Flags().String(flagJSONRPCMetricsAddress, DefaultMetricsAddress, "Address to listen on for the EVM Metrics server")
}

// GetConfig load config values from the app options
func GetConfig(appOpts servertypes.AppOptions) JSONRPCConfig {
	return JSONRPCConfig{
		Enable:             cast.ToBool(appOpts.Get(flagJSONRPCEnable)),
		EnableUnsafeCORS:   cast.ToBool(appOpts.Get(flagJSONRPCEnableUnsafeCORS)),
		Address:            cast.ToString(appOpts.Get(flagJSONRPCAddress)),
		APIs:               cast.ToStringSlice(appOpts.Get(flagJSONRPCAPIs)),
		LogsCap:            cast.ToInt32(appOpts.Get(flagJSONRPCLogsCap)),
		FilterCap:          cast.ToInt32(appOpts.Get(flagJSONRPCFilterCap)),
		BlockRangeCap:      cast.ToInt32(appOpts.Get(flagJSONRPCBlockRangeCap)),
		HTTPTimeout:        cast.ToDuration(appOpts.Get(flagJSONRPCHTTPTimeout)),
		HTTPIdleTimeout:    cast.ToDuration(appOpts.Get(flagJSONRPCHTTPIdleTimeout)),
		MaxOpenConnections: cast.ToInt(appOpts.Get(flagJSONRPCMaxOpenConnections)),
		MetricsAddress:     cast.ToString(appOpts.Get(flagJSONRPCMetricsAddress)),
	}
}

// DefaultConfigTemplate defines the configuration template for the EVM RPC configuration
const DefaultConfigTemplate = `
###############################################################################
###                           JSON RPC Configuration                        ###
###############################################################################

[json-rpc]

# Enable defines if the gRPC server should be enabled.
enable = {{ .JSONRPCConfig.Enable }}

# Address defines the EVM RPC HTTP server address to bind to.
address = "{{ .JSONRPCConfig.Address }}"

# EnableUnsafeCORS defines if the EVM RPC server should enable unsafe CORS.
enable-unsafe-cors = {{ .JSONRPCConfig.EnableUnsafeCORS }}

# API defines a list of JSON-RPC namespaces that should be enabled
# Example: "eth,txpool,personal,net,debug,web3"
apis = "{{range $index, $elmt := .JSONRPCConfig.APIs}}{{if $index}},{{$elmt}}{{else}}{{$elmt}}{{end}}{{end}}"

# FilterCap is the global cap for total number of filters that can be created.
filter-cap = {{ .JSONRPCConfig.FilterCap }}

# LogsCap defines the max number of results can be returned from single 'eth_getLogs' query.
logs-cap = {{ .JSONRPCConfig.LogsCap }}

# BlockRangeCap defines the max block range allowed for 'eth_getLogs' query.
block-range-cap = {{ .JSONRPCConfig.BlockRangeCap }}

# HTTPTimeout is the read/write timeout of http json-rpc server.
http-timeout = "{{ .JSONRPCConfig.HTTPTimeout }}"

# HTTPIdleTimeout is the idle timeout of http json-rpc server.
http-idle-timeout = "{{ .JSONRPCConfig.HTTPIdleTimeout }}"

# MaxOpenConnections sets the maximum number of simultaneous connections
# for the server listener.
max-open-connections = {{ .JSONRPCConfig.MaxOpenConnections }}

# MetricsAddress defines the EVM Metrics server address to bind to. Pass --metrics in CLI to enable
# Prometheus metrics path: /debug/metrics/prometheus
metrics-address = "{{ .JSONRPCConfig.MetricsAddress }}"
`
