package commands

import (
	"github.com/admiralyeoj/anime-announcements/internal/config"
)

type Command struct {
	Name        string
	Description string
	Handler     func(repos *config.Repositories, args []string) error
}
