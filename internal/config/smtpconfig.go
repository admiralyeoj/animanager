package config

// SMTPConfig holds the SMTP-related configuration settings.
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Sender   string
}
