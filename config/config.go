package config

import (
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/hesusruiz/isbetmf/internal/sqlogger"
	"github.com/hesusruiz/isbetmf/tmfclient"
	"github.com/hesusruiz/isbetmf/types"
)

// Indicates the environment (SBX, DEV2, PRO, LCL) where the server is running.
// It is used to determine the default configuration profile if the user does not specify anything else.
type Environment int

const (
	DOME_PRO      Environment = 0
	DOME_DEV2     Environment = 1
	DOME_SBX      Environment = 2
	DOME_LCL      Environment = 3
	ISBE_EVIDENCE Environment = 4
	ISBE_MYCRED   Environment = 5
)

type Config struct {

	// The environment for the default configuration profile
	Environment Environment

	// Server operator information
	ServerOperatorOrganizationIdentifier string
	ServerOperatorDid                    string
	ServerOperatorName                   string
	ServerOperatorCountry                string
	ServerEmailAddress                   string

	// VerifierServer is the URL of the verifier server, which is used to verify the access tokens.
	VerifierServer string

	// The domain of the remote TMForum API server when we act as proxy
	RemoteTMFServer string

	// Dbname is the name of the database file where the TMForum cached data is stored
	// It is used to store the data in a local SQLite database, the best SQL database for this purpose.
	Dbname string

	// The power required by a caller to be considered LEAR
	LEARPower types.OnePower

	// PolicyFileName is the name of the file where the policies are stored.
	// It can specify a local file or a remote URL.
	PolicyFileName string

	// PDPAddress is the address of the PDP server.
	PDPAddress string

	// Debug mode, more logs and less caching
	Debug bool

	// TODO: this is temporary for testing
	FakeClaims bool

	// fixMode enables "smart" automatit fixing of objects so they comply with the DOME specs
	// There is no magic. however, and there are things that can not be done.
	fixMode bool

	// Enable synchronization with the remote server in background
	BackgroudSync bool

	// ClonePeriod is the period in which the reporting tool will clone the TMForum objects from the DOME instance,
	// to keep the local cache up to date.
	ClonePeriod time.Duration

	// LogHandler is the handler used to log messages.
	// It is a custom handler that uses the slog package to log messages both to the console and to a SQLite database.
	LogHandler *sqlogger.SQLogHandler

	// LogLevel is a slog.LevelVar that can be set to different log levels (e.g. Debug, Info, Warn, Error).
	LogLevel *slog.LevelVar

	// ProxyEnabled enables the TMF caching proxy functionality.
	ProxyEnabled bool

	// TMFClient is the configuration for the TMF client when calling the remote TMF server
	TMFClient *tmfclient.Config
}

// As this PDP is designed for DOME and ISBE environments, many config data items are hardcoded.
// This avoids many configuration errors and simplifies deployment, at the expense of some flexibility.
// However, this flexibility is not really needed in practice, as the DOME environments are well defined and stable.
// Minimizing errors is here much more important than the ability to configure these parameters.

var sbxConfig = &Config{
	Environment:  DOME_SBX,
	ProxyEnabled: true,

	ServerOperatorOrganizationIdentifier: "VATSB-12345678J",
	ServerOperatorDid:                    "did:elsi:VATSB-12345678J",
	ServerOperatorName:                   "DOME Foundation SBX",
	ServerOperatorCountry:                "ES",

	LEARPower: types.OnePower{
		Tmf_type:     "Domain",
		Tmf_domain:   "DOME",
		Tmf_function: "Onboarding",
		Tmf_action:   []string{"execute"},
	},

	PolicyFileName:  "auth_policies.star",
	RemoteTMFServer: "https://tmf.dome-marketplace-sbx.org",
	VerifierServer:  "https://verifier.dome-marketplace-sbx.org",
	Dbname:          "data/tmf.dome.sbx.db",
	ClonePeriod:     DefaultClonePeriod,
}

var isbeEvidenceConfig = &Config{
	Environment:  ISBE_EVIDENCE,
	ProxyEnabled: false,

	ServerOperatorOrganizationIdentifier: "VATES-11111111K",
	ServerOperatorDid:                    "did:elsi:VATES-11111111K",
	ServerOperatorName:                   "ISBE Foundation",
	ServerOperatorCountry:                "ES",

	LEARPower: types.OnePower{
		Tmf_type:     "Domain",
		Tmf_domain:   "DOME",
		Tmf_function: "Onboarding",
		Tmf_action:   []string{"execute"},
	},

	PolicyFileName: "auth_policies.star",
	VerifierServer: "https://verifier.dome-marketplace.eu",
	Dbname:         "data/isbetmf.db",
	ClonePeriod:    DefaultClonePeriod,
}

var isbeMycredentialConfig = &Config{
	Environment:  ISBE_MYCRED,
	ProxyEnabled: false,

	ServerOperatorOrganizationIdentifier: "VATES-11111111K",
	ServerOperatorDid:                    "did:elsi:VATES-11111111K",
	ServerOperatorName:                   "ISBE Foundation",
	ServerOperatorCountry:                "ES",

	LEARPower: types.OnePower{
		Tmf_type:     "Domain",
		Tmf_domain:   "ISBE",
		Tmf_function: "Onboarding",
		Tmf_action:   []string{"execute"},
	},

	PolicyFileName: "auth_policies.star",
	VerifierServer: "https://verifier.dome-marketplace.eu",
	Dbname:         "data/isbetmf.db",
	ClonePeriod:    DefaultClonePeriod,
}

var proConfig = &Config{
	Environment:  DOME_PRO,
	ProxyEnabled: true,

	LEARPower: types.OnePower{
		Tmf_type:     "Domain",
		Tmf_domain:   "DOME",
		Tmf_function: "Onboarding",
		Tmf_action:   []string{"execute"},
	},

	PolicyFileName:  "auth_policies.star",
	RemoteTMFServer: "https://tmf.dome-marketplace.eu",
	VerifierServer:  "https://verifier.dome-marketplace.eu",
	Dbname:          "data/tmf.dome.pro.db",
	ClonePeriod:     DefaultClonePeriod,
}

var dev2Config = &Config{
	Environment:  DOME_DEV2,
	ProxyEnabled: true,

	LEARPower: types.OnePower{
		Tmf_type:     "Domain",
		Tmf_domain:   "DOME",
		Tmf_function: "Onboarding",
		Tmf_action:   []string{"execute"},
	},

	PolicyFileName:  "auth_policies.star",
	RemoteTMFServer: "https://tmf.dome-marketplace-dev2.org",
	VerifierServer:  "https://verifier.dome-marketplace-dev2.org",
	Dbname:          "data/tmf.dome.dev2.db",
	ClonePeriod:     DefaultClonePeriod,
}

var lclConfig = &Config{
	Environment:  DOME_LCL,
	ProxyEnabled: true,

	LEARPower: types.OnePower{
		Tmf_type:     "Domain",
		Tmf_domain:   "DOME",
		Tmf_function: "Onboarding",
		Tmf_action:   []string{"execute"},
	},

	PolicyFileName:  "auth_policies.star",
	RemoteTMFServer: "https://tmf.dome-marketplace-lcl.org",
	VerifierServer:  "https://verifier.dome-marketplace-lcl.org",
	Dbname:          "data/tmf.dome.lcl.db",
	ClonePeriod:     DefaultClonePeriod,
}

// LoadConfig initializes and returns a Config struct based on the provided parameters.
// It sets up logging, selects the appropriate environment, and applies configuration options.
//
// Parameters:
//   - envir:        The environment to use ("pro", "dev2", "sbx", "lcl").
//   - debug:        Enables debug logging if true.
//
// Returns:
//   - *Config: The initialized configuration struct.
//   - error:   An error if configuration or logger setup fails.
func LoadConfig(
	envir string,
	debug bool,
) (*Config, error) {
	var conf *Config

	// Normalize to lowercase for comparisons
	envir = strings.ToLower(envir)

	// Configure the slog logger
	var logLevel slog.Level
	if debug {
		logLevel = slog.LevelDebug
	} else {
		logLevel = slog.LevelInfo
	}

	// Initialize the custom SQLogHandler
	logOptions := &sqlogger.Options{
		Level: &logLevel,
	}

	// Check if the logs should be colored
	if os.Getenv("ISBETMF_LOGS_NOCOLOR") == "true" {
		logOptions.NoColor = true
	}

	// Initialize the logging system
	sqlog, err := sqlogger.NewSQLogHandler(logOptions)
	if err != nil {
		slog.Error("failed to initialize SQLogHandler, exiting", slog.Any("error", err))
		os.Exit(1)
	}
	defer sqlog.Close()

	// And set the default logging system for all components
	slog.SetDefault(slog.New(sqlog))

	// Choose the profile from the environment passed
	switch envir {
	case "pro":
		conf = proConfig
		slog.Info("Using the PRODUCTION environment")
	case "dev2":
		conf = dev2Config
		slog.Info("Using the DEV2 environment")
	case "sbx":
		conf = sbxConfig
		slog.Info("Using the SBX environment")
	case "lcl":
		conf = lclConfig
		slog.Info("Using the LCL environment")
	case "evidenceledger":
		conf = isbeEvidenceConfig
		slog.Info("Using the ISBE evidence environment")
	case "mycredential":
		conf = isbeMycredentialConfig
		slog.Info("Using the ISBE mycredential environment")
	default:
		conf = isbeMycredentialConfig
		slog.Info("Using the default (ISBE mycredential) environment")
	}

	conf.Debug = debug

	// Check for overrides with environment variables

	verifierServer := os.Getenv("ISBETMF_VERIFIER")
	if verifierServer != "" {
		conf.VerifierServer = verifierServer
	}
	slog.Info("Verifier", slog.String("url", conf.VerifierServer))

	remoteTMFServer := os.Getenv("ISBETMF_REMOTE_SERVER")
	if remoteTMFServer != "" {
		conf.RemoteTMFServer = remoteTMFServer
	}
	slog.Info("RemoteTMFServer", slog.String("url", conf.RemoteTMFServer))

	proxyEnabled := os.Getenv("ISBETMF_PROXY_ENABLED")
	switch proxyEnabled {
	case "true":
		conf.ProxyEnabled = true
	case "false":
		conf.ProxyEnabled = false
	}
	slog.Info("Proxy", slog.Bool("enabled", conf.ProxyEnabled))

	return conf, nil

}
