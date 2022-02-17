package I2P

import (
	"os/exec"
	"path/filepath"
)

func LinuxCommand(dir string) (*exec.Cmd, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	execPath := filepath.Join(dir, "I2P", "bin", "I2P")
	return exec.Command(execPath), nil
}

func WindowsCommand(dir string) (*exec.Cmd, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	execPath := filepath.Join(dir, "I2P", "I2P.exe")
	return exec.Command(execPath), nil
}
