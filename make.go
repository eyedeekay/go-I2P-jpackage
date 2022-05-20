package I2P

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

//go:generate go run ./touch
//go:generate go run ./I2P -generate=true -dir=$GOPATH/src/github.com/eyedeekay/go-I2P-jpackage/
//go:generate go build -o $GOPATH/src/github.com/eyedeekay/go-I2P-jpackage/go-I2P-jpackage $GOPATH/src/github.com/eyedeekay/go-I2P-jpackage/I2P

func (d *Daemon) Generate() error {
	if err := d.gitCloneI2PFirefox(); err != nil {
		return fmt.Errorf("generate: gitCloneI2PFirefox failed %ss", err.Error())
	}
	if err := d.runI2PFirefoxBuildSh(); err != nil {
		return fmt.Errorf("generate: runI2PFirefoxBuildSh failed %s", err.Error())
	}
	if err := d.runI2PFirefoxExtensions(); err != nil {
		return fmt.Errorf("generate: runI2PFirefoxExtensions failed %s", err.Error())
	}
	if err := d.runI2PFirefoxMake(); err != nil {
		return fmt.Errorf("generate: runI2PFirefoxMake failed %s", err.Error())
	}
	if err := d.tarI2PdotFirefoxdotBuild(); err != nil {
		return fmt.Errorf("generate: tarI2PdotFirefoxdotBuild failed %s", err.Error())
	}
	return nil
}

func (d *Daemon) gitCloneI2PFirefox() error {
	dir := filepath.Join(d.Dir, "i2p.firefox")
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:           "https://i2pgit.org/i2p-hackers/i2p.firefox",
		Progress:      os.Stdout,
		SingleBranch:  true,
		ReferenceName: plumbing.NewBranchReferenceName("master"),
	})
	if err != nil {
		log.Printf("gitCloneI2PFirefox: git.PlainClone failed: %s", err.Error())
	}
	return nil
}

func (d *Daemon) runI2PFirefoxBuildSh() error {
	dir := filepath.Join(d.Dir, "i2p.firefox")
	fmt.Println("Running build.sh")
	switch runtime.GOOS {
	case "windows":
		gitbash, err := filepath.Abs(filepath.Join("/Program Files/", "/Git/", "git-bash.exe"))
		if err != nil {
			return err
		}
		cmd := exec.Command(gitbash, filepath.Join(dir, "build.sh"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		cmd := exec.Command(filepath.Join(dir, "build.sh"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}

func (d *Daemon) runI2PFirefoxExtensions() error {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Running wsl", "make", "extensions", "-C", "i2p.firefox")
		cmd := exec.Command("wsl", "make", "extensions", "-C", "i2p.firefox")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		fmt.Println("Running make", "extensions", "-C", "i2p.firefox")
		cmd := exec.Command("make", "extensions", "-C", filepath.Join(d.Dir, "i2p.firefox"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}

func (d *Daemon) runI2PFirefoxMake() error {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Running wsl", "make", "-C", "i2p.firefox")
		cmd := exec.Command("wsl", "make", "-C", "i2p.firefox")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		fmt.Println("Running make", "-C", "i2p.firefox")
		cmd := exec.Command("make", "-C", filepath.Join(d.Dir, "i2p.firefox"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}

func (d *Daemon) tarI2PdotFirefoxdotBuild() error {
	os.Remove(filepath.Join(d.Dir, "build."+runtime.GOOS+".I2P.tar.xz"))
	err := TarXzip(filepath.Join(d.Dir, "i2p.firefox", "build", "I2P"), filepath.Join(d.Dir, "build."+runtime.GOOS+".I2P.tar.xz"))
	if err != nil {
		return fmt.Errorf("tarI2PdotFirefoxdotBuild: Tar failed: %s", err.Error())
	}
	return nil
}
