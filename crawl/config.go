package crawl

type Config struct {
	URL      string `toml:"url"`
	Interval int    `toml:"intreval"` //s
}

func NewConfig() Config {

	return Config{}
}
