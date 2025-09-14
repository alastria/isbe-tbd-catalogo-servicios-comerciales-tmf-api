package tmfclient

// Config holds the configuration for the tmfclient package
type Config struct {
	// BaseURL of the remote TMForum server
	BaseURL string `json:"base_url" yaml:"base_url"`

	// Timeout in seconds for HTTP requests
	Timeout int `json:"timeout" yaml:"timeout"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		BaseURL: "https://tmf.dome-marketplace-sbx.org",
		Timeout: 30,
	}
}
