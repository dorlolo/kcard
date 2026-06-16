// Package config 提供应用程序配置的加载与验证功能。
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

// Config 表示应用程序配置，包含 HTTP、数据库、Redis、AI 及存储等各项参数。
type Config struct {
	HTTPAddr              string
	DatabaseURL           string
	RedisURL              string
	StoragePath           string
	AIProvider            string
	ArkAPIKey             string
	ArkModel              string
	ArkBaseURL            string
	AllowedFrontendOrigin string
	JobVisibilityTimeout  time.Duration
	DefaultWorkspaceName  string
	RequireAIAPIKey       bool
}

// Load 从环境变量加载配置并执行验证，返回有效的 Config 实例。
func Load() (Config, error) {
	if err := loadDotEnv(); err != nil {
		return Config{}, err
	}

	cfg := Config{
		HTTPAddr:              env("HTTP_ADDR", ":8080"),
		DatabaseURL:           env("DATABASE_URL", "postgres://kcard:kcard@localhost:5432/kcard?sslmode=disable"),
		RedisURL:              env("REDIS_URL", "redis://localhost:6379/0"),
		StoragePath:           env("STORAGE_PATH", ".local/storage"),
		AIProvider:            env("AI_PROVIDER", "ark"),
		ArkAPIKey:             os.Getenv("ARK_API_KEY"),
		ArkModel:              env("ARK_MODEL", ""),
		ArkBaseURL:            env("ARK_BASE_URL", "https://ark.cn-beijing.volces.com/api/v3"),
		AllowedFrontendOrigin: env("FRONTEND_ORIGIN", "http://localhost:5173"),
		DefaultWorkspaceName:  env("DEFAULT_WORKSPACE_NAME", "My Study Workspace"),
	}
	seconds, err := strconv.Atoi(env("JOB_VISIBILITY_TIMEOUT_SECONDS", "120"))
	if err != nil || seconds <= 0 {
		return Config{}, fmt.Errorf("JOB_VISIBILITY_TIMEOUT_SECONDS must be a positive integer")
	}
	cfg.JobVisibilityTimeout = time.Duration(seconds) * time.Second
	cfg.RequireAIAPIKey = strings.EqualFold(env("REQUIRE_AI_API_KEY", "false"), "true")
	return cfg, cfg.Validate()
}

// Validate 检查配置中必需字段是否已设置，返回缺失项的错误描述。
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
	if c.RequireAIAPIKey && c.ArkAPIKey == "" {
		missing = append(missing, "ARK_API_KEY")
	}
	if c.ArkAPIKey != "" && c.ArkModel == "" {
		missing = append(missing, "ARK_MODEL")
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
