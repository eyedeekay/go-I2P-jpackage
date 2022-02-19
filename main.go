package I2P

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Daemon struct {
	Dir     string
	Command *exec.Cmd
}

func (d *Daemon) LinuxCommand() (*exec.Cmd, error) {
	if err := SetEnv(d.Dir); err != nil {
		return nil, err
	}
	dir, err := filepath.Abs(d.Dir)
	if err != nil {
		return nil, err
	}
	execPath := filepath.Join(dir, "I2P", "bin", "I2P")
	d.Command = exec.Command(execPath)
	return d.Command, nil
}

func (d *Daemon) RunLinuxCommand() error {
	if err := SetEnv(d.Dir); err != nil {
		return err
	}
	var err error
	d.Command, err = d.LinuxCommand()
	if err != nil {
		return err
	}
	d.Command.Stdout = os.Stdout
	d.Command.Stderr = os.Stderr
	return d.Command.Run()
}

func (d *Daemon) WindowsCommand() (*exec.Cmd, error) {
	if err := SetEnv(d.Dir); err != nil {
		return nil, err
	}
	dir, err := filepath.Abs(d.Dir)
	if err != nil {
		return nil, err
	}
	execPath := filepath.Join(dir, "I2P", "I2P.exe")
	d.Command = exec.Command(execPath)
	return d.Command, nil
}

func (d *Daemon) RunWindowsCommand() error {
	if err := SetEnv(d.Dir); err != nil {
		return err
	}
	var err error
	d.Command, err = d.WindowsCommand()
	if err != nil {
		return err
	}
	d.Command.Stdout = os.Stdout
	d.Command.Stderr = os.Stderr
	return d.Command.Run()
}

func (d *Daemon) RunCommand() error {
	if err := SetEnv(d.Dir); err != nil {
		return err
	}
	switch runtime.GOOS {
	case "windows":
		return d.RunWindowsCommand()
	default:
		return d.RunLinuxCommand()
	}
}

func SetEnv(dir string) error {
	err := os.Setenv("I2P", filepath.Join(dir, "I2P"))
	if err != nil {
		return fmt.Errorf("SetEnv: os.Setenv failed: %s", err.Error())
	}
	err = os.Setenv("I2P_CONFIG", filepath.Join(dir, "I2P", "workdir"))
	if err != nil {
		return fmt.Errorf("SetEnv: os.Setenv failed: %s", err.Error())
	}
	return nil
}
