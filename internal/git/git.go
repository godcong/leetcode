package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// Config holds git configuration
type Config struct {
	Enabled bool
}

// NewConfig creates a new git config
func NewConfig(enabled bool) *Config {
	return &Config{
		Enabled: enabled,
	}
}

// AddAndCommit adds file to git and commits with message
func AddAndCommit(path string, name string, enabled bool) error {
	if !enabled {
		return nil
	}

	command := exec.Command("git", "add", path)
	_, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("add to git error: %v", err)
	}

	command = exec.Command("git", "commit", "-a", "-m", fmt.Sprintf("feat(code): add new %s", name))
	_, err = command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("commit to git error: %v", err)
	}

	return nil
}

// Check verifies if git is available
func Check() error {
	command := exec.Command("git", "version")
	ret, err := command.CombinedOutput()
	if err != nil {
		return err
	}

	if !strings.Contains(string(ret), "git version") {
		//skip if git is not exist
		return err
	}
	return nil
}
