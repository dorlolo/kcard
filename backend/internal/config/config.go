package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	HTTPAddr               string
	DatabaseURL            string
	RedisURL               string
	StoragePath            string
	AnthropicAPIKey        string
	AnthropicModel         string
	AllowedFrontendOrigin  string
	JobVisibilityTimeout   time.Duration
	DefaultWorkspaceName   string
	RequireAnthropicAPIKey bool
}

func Load() (Config, error) {
	if err := loadDotEnv(); err != nil {
		return Config{}, err
	}

	cfg := Config{
		HTTPAddr:              env("HTTP_ADDR", ":8080"),
		DatabaseURL:           env("DATABASE_URL", "postgres://kcard:kcard@localhost:5432/kcard?sslmode=disable"),
		RedisURL:              env("REDIS_URL", "redis://localhost:6379/0"),
		StoragePath:           env("STORAGE_PATH", ".local/storage"),
		AnthropicAPIKey:       os.Getenv("ANTHROPIC_API_KEY"),
		AnthropicModel:        env("ANTHROPIC_MODEL", "claude-opus-4-8"),
		AllowedFrontendOrigin: env("FRONTEND_ORIGIN", "http://localhost:5173"),
		DefaultWorkspaceName:  env("DEFAULT_WORKSPACE_NAME", "My Study Workspace"),
	}
	seconds, err := strconv.Atoi(env("JOB_VISIBILITY_TIMEOUT_SECONDS", "120"))
	if err != nil || seconds <= 0 {
		return Config{}, fmt.Errorf("JOB_VISIBILITY_TIMEOUT_SECONDS must be a positive integer")
	}
	cfg.JobVisibilityTimeout = time.Duration(seconds) * time.Second
	cfg.RequireAnthropicAPIKey = strings.EqualFold(env("REQUIRE_ANTHROPIC_API_KEY", "false"), "true")
	return cfg, cfg.Validate()
}

func (c Config) Validate() error {
	var missing []string
	if c.DatabaseURL == "" {
		missing = append(missing, "DATABASE_URL")
	}
	if c.RedisURL == "" {
		missing = append(missing, "REDIS_URL")
	}
	if c.StoragePath == "" {
		missing = append(missing, "STORAGE_PATH")
	}
	if c.AllowedFrontendOrigin == "" {
		missing = append(missing, "FRONTEND_ORIGIN")
	}
	if c.RequireAnthropicAPIKey && c.AnthropicAPIKey == "" {
		missing = append(missing, "ANTHROPIC_API_KEY")
	}
	if len(missing) > 0 {
		return errors.New("missing required configuration: " + strings.Join(missing, ", "))
	}
	return nil
}

func loadDotEnv() error {
	path, ok := findDotEnv()
	if !ok {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open .env: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimPrefix(line, "export ")
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			return fmt.Errorf("parse .env line %d: expected KEY=VALUE", lineNumber)
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if key == "" {
			return fmt.Errorf("parse .env line %d: empty key", lineNumber)
		}
		if _, exists := os.LookupEnv(key); exists {
			continue
		}
		if err := os.Setenv(key, unquote(value)); err != nil {
			return fmt.Errorf("set .env variable %s: %w", key, err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read .env: %w", err)
	}
	return nil
}

func findDotEnv() (string, bool) {
	candidates := []string{".env", filepath.Join("..", ".env")}
	for _, candidate := range candidates {
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
			return candidate, true
		}
	}
	return "", false
}

func unquote(value string) string {
	if len(value) < 2 {
		return value
	}
	if (value[0] == '\'' && value[len(value)-1] == '\'') || (value[0] == '"' && value[len(value)-1] == '"') {
		return value[1 : len(value)-1]
	}
	return value
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
