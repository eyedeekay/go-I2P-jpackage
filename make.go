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

//go:generate ./touch/touch
//go:generate go run ./I2P -generate=true -dir=$GOPATH/src/github.com/eyedeekay/go-I2P-jpackage/
//go:generate go build -o $GOPATH/src/github.com/eyedeekay/go-I2P-jpackage/go-I2P-jpackage ./I2P

func Generate(dir string) error {
	if err := gitCloneI2PFirefox(dir); err != nil {
		return fmt.Errorf("generate: gitCloneI2PFirefox failed %ss", err.Error())
	}
	if err := runI2PFirefoxBuildSh(dir); err != nil {
		return fmt.Errorf("generate: runI2PFirefoxBuildSh failed %s", err.Error())
	}
	if err := runI2PFirefoxMake(dir); err != nil {
		return fmt.Errorf("generate: runI2PFirefoxMake failed %s", err.Error())
	}
	if err := tarI2PdotFirefoxdotBuild(dir); err != nil {
		return fmt.Errorf("generate: tarI2PdotFirefoxdotBuild failed %s", err.Error())
	}
	return nil
}

func gitCloneI2PFirefox(dir string) error {
	dir = filepath.Join(dir, "i2p.firefox")
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:           "https://i2pgit.org/i2p-hackers/i2p.firefox",
		Progress:      os.Stdout,
		SingleBranch:  true,
		ReferenceName: plumbing.NewBranchReferenceName("settable-paths"),
	})
	if err != nil {
		log.Printf("gitCloneI2PFirefox: git.PlainClone failed: %s", err.Error())
	}
	return nil
}

func runI2PFirefoxBuildSh(dir string) error {
	dir = filepath.Join(dir, "i2p.firefox")
	fmt.Println("Running build.sh")
	cmd := exec.Command(filepath.Join(dir, "build.sh"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runI2PFirefoxMake(dir string) error {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Running wsl make")
		cmd := exec.Command("wsl", "make", "-C", filepath.Join(dir, "i2p.firefox"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		fmt.Println("Running make")
		cmd := exec.Command("make", "-C", filepath.Join(dir, "i2p.firefox"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}

func tarI2PdotFirefoxdotBuild(dir string) error {
	os.Remove(filepath.Join(dir, "build.I2P.tar.gz"))
	err := TarGzip(filepath.Join(dir, "i2p.firefox", "build", "I2P"), filepath.Join(dir, "build.I2P.tar.gz"))
	if err != nil {
		return fmt.Errorf("tarI2PdotFirefoxdotBuild: Tar failed: %s", err.Error())
	}
	return nil
}
