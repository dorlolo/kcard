package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDotEnvSetsTemporaryEnvironmentVariables(t *testing.T) {
	dir := t.TempDir()
	oldwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(oldwd) })
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".env"), []byte("DOTENV_TEST_VALUE='from file'\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("DOTENV_TEST_VALUE", "")
	_ = os.Unsetenv("DOTENV_TEST_VALUE")

	if err := loadDotEnv(); err != nil {
		t.Fatal(err)
	}
	if got := os.Getenv("DOTENV_TEST_VALUE"); got != "from file" {
		t.Fatalf("DOTENV_TEST_VALUE=%q", got)
	}
}

func TestLoadDotEnvDoesNotOverrideExistingEnvironment(t *testing.T) {
	dir := t.TempDir()
	oldwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(oldwd) })
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".env"), []byte("DOTENV_EXISTING=from-file\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("DOTENV_EXISTING", "from-env")

	if err := loadDotEnv(); err != nil {
		t.Fatal(err)
	}
	if got := os.Getenv("DOTENV_EXISTING"); got != "from-env" {
		t.Fatalf("DOTENV_EXISTING=%q", got)
	}
}
