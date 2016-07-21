package smtp

type Config struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	// Whether to skip TLS verify.
	NoVerify bool `toml:"no-verify"`
	// From address
	From string `toml:"from"`
	// Default To addresses
	To []string `toml:"to"`
	// Close connection to SMTP server after idle timeout has elapsed
}

func NewConfig() Config {
	return Config{
		Host:     "localhost",
		Port:     25,
		Username: "",
		Password: "",
	}
}
