package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	Cfg    = Load()
)

type Config struct {
	API        APIConfig    `env:",prefix=API_"`
	Environemt string       `env:"ENV"`
	Databases  DBConfig     `env:",prefix=DB_"`
	Token      TokenConfig  `env:",prefix=TOKEN_"`
	Email      EmailAddress `env:",prefix=EMAIL_"`
	Frontend   string       `env:"FRONTEND_URL, default=http://localhost:3000"`
	OAuth      OAuthConfig  `env:",prefix=OAUTH_"`
}

type APIConfig struct {
	Host     string `env:"SERVER_HOST, default=0.0.0.0"`
	Port     int    `env:"SERVER_PORT, default=8000"`
	LogLevel string `env:"LOG_LEVEL, default=WARN"`
}

type DBConfig struct {
	UseReplication  bool          `env:"USE_REPLICATION, default=false"`
	Driver          string        `env:"DRIVER, default=postgres"`
	Source          DBAttributes  `env:",prefix=SOURCE_"`
	Replica         DBAttributes  `env:",prefix=REPLICA_"`
	LogLevel        string        `env:"LOG_LEVEL, default=WARN"`
	ConnMaxIdleTime time.Duration `env:"CONN_MAX_IDLE_TIME, default=300s"`
	ConnMaxLifeTime time.Duration `env:"CONN_MAX_LIFE_TIME, default=300s"`
	MaxIdleConns    int           `env:"MAX_IDLE_CONNS, default=5"`
	MaxOpenConns    int           `env:"MAX_OPEN_CONNS, default=10"`
}

type DBAttributes struct {
	Dsn      string `env:"DSN"`
	Host     string `env:"HOST, default=127.0.0.1"`
	Port     int    `env:"PORT, default=5432"`
	Database string `env:"DATABASE, default=lusca"`
	Username string `env:"USERNAME, default=postgres_user"`
	Password string `env:"PASSWORD, default=postgres_pw"`
	SSLMode  string `env:"SSL_MODE, default=disable"`
}

type TokenConfig struct {
	SecretKey  string        `env:"SECRET_KEY, default=secret"`
	Secure     bool          `env:"SECURE, default=true"`
	AccessKey  string        `env:"ACCESS_KEY, default=access_key"`
	AccessTTL  time.Duration `env:"ACCESS_TTL, default=5m"`
	RefreshKey string        `env:"REFRESH_KEY, default=mysecretkey32byteslongforpaseto12"`
	RefreshTTL time.Duration `env:"REFRESH_TTL, default=24h"`
	TestToken  string        `env:"TEST, default=nothing_here"`
}

type EmailAddress struct {
	SMTPHost     string        `env:"SMTP_HOST, default=smtp.example.com"`
	SMTPPort     string        `env:"SMTP_PORT, default=587"`
	SMTPUsername string        `env:"SMTP_USERNAME, default=example_user"`
	SMTPPassword string        `env:"SMTP_PASSWORD, default=example_password"`
	EmailAdmin   string        `env:"ADMIN, default=example_email"`
	DurationTTL  time.Duration `env:"DURATION_TTL, default=15m"`
}

type OAuthConfig struct {
	GoogleClientID            string `env:"GOOGLE_CLIENT_ID, default=your_google_client_id"`
	GoogleClientSecret        string `env:"GOOGLE_CLIENT_SECRET, default=your_google_client_secret"`
	GoogleRedirectURL         string `env:"GOOGLE_REDIRECT_URL, default=http://localhost:8000"`
	GoogleFrontendRedirectURL string `env:"GOOGLE_FRONTEND_REDIRECT_URL, default=http://localhost:5000"`
}

func (c DBAttributes) GetPostgresDSN() string {
	if c.Dsn != "" {
		return c.Dsn
	}
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", c.Username, c.Password, c.Database, c.Host, c.Port, c.SSLMode)
}

func Load() Config {
	ctx := context.Background()

	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		logger.ErrorContext(ctx, "Failed to load config", slog.Any("error", err))
	}
	return c
}

func init() {
	if err := godotenv.Load(); err != nil {
		logger.Info("get .env failed")
	}
	Cfg = Load()
}
