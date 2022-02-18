package I2P

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func LinuxCommand(dir string) (*exec.Cmd, error) {
	if err := SetEnv(dir); err != nil {
		return nil, err
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	execPath := filepath.Join(dir, "I2P", "bin", "I2P")
	return exec.Command(execPath), nil
}

func RunLinuxCommand(dir string) error {
	if err := SetEnv(dir); err != nil {
		return err
	}
	cmd, err := LinuxCommand(dir)
	if err != nil {
		return err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func WindowsCommand(dir string) (*exec.Cmd, error) {
	if err := SetEnv(dir); err != nil {
		return nil, err
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	execPath := filepath.Join(dir, "I2P", "I2P.exe")
	return exec.Command(execPath), nil
}

func RunWindowsCommand(dir string) error {
	if err := SetEnv(dir); err != nil {
		return err
	}
	cmd, err := WindowsCommand(dir)
	if err != nil {
		return err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunCommand(dir string) error {
	if err := SetEnv(dir); err != nil {
		return err
	}
	switch runtime.GOOS {
	case "windows":
		return RunWindowsCommand(dir)
	default:
		return RunLinuxCommand(dir)
	}
}

func SetEnv(dir string) error {
	err := os.Setenv("I2P", filepath.Join(dir, "I2P"))
	if err != nil {
		return fmt.Errorf("SetEnv: os.Setenv failed: %s", err.Error())
	}
	err = os.Setenv("I2P_CONFIG", filepath.Join(dir, "I2P", "config"))
	if err != nil {
		return fmt.Errorf("SetEnv: os.Setenv failed: %s", err.Error())
	}
	return nil
}
