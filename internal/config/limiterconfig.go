package config

// LimiterConfig holds the rate limiter configuration settings.
type LimiterConfig struct {
	RPS     float64
	Burst   int
	Enabled bool
}
