package bot

import (
	"database/sql"

	"github.com/hortbot/hortbot/internal/dedupe"
)

const (
	DefaultPrefix = "!"
	DefaultBullet = "[HB]"
)

var isTesting = false

type Config struct {
	DB     *sql.DB
	Dedupe dedupe.Deduplicator
	Sender MessageSender

	Prefix string
	Bullet string
}

type Bot struct {
	db     *sql.DB
	dedupe dedupe.Deduplicator
	sender MessageSender

	prefix string
	bullet string
}

func New(config *Config) *Bot {
	b := &Bot{
		db:     config.DB,
		dedupe: config.Dedupe,
		sender: config.Sender,
		prefix: config.Prefix,
		bullet: config.Bullet,
	}

	// TODO: verify other dependencies

	if b.bullet == "" {
		b.bullet = DefaultBullet
	}

	if b.prefix == "" {
		b.prefix = DefaultPrefix
	}

	return b
}

//go:generate gobin -m -run github.com/maxbrunsfeld/counterfeiter/v6 . MessageSender

type MessageSender interface {
	SendMessage(origin, target, message string) error
}
