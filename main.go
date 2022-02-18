package I2P

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func LinuxCommand(dir string) (*exec.Cmd, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	execPath := filepath.Join(dir, "I2P", "bin", "I2P")
	return exec.Command(execPath), nil
}

func RunLinuxCommand(dir string) error {
	cmd, err := LinuxCommand(dir)
	if err != nil {
		return err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func WindowsCommand(dir string) (*exec.Cmd, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	execPath := filepath.Join(dir, "I2P", "I2P.exe")
	return exec.Command(execPath), nil
}

func RunWindowsCommand(dir string) error {
	cmd, err := WindowsCommand(dir)
	if err != nil {
		return err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunCommand(dir string) error {
	switch runtime.GOOS {
	case "windows":
		return RunWindowsCommand(dir)
	default:
		return RunLinuxCommand(dir)
	}
}
