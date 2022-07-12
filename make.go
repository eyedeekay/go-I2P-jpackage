package I2P

import (
	"fmt"
	"io/ioutil"
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
	if err := d.gitPullI2PFirefox(); err != nil {
		return fmt.Errorf("generate: gitPullI2PFirefox failed %ss", err.Error())
	}
	if err := d.setMasterOveride(); err != nil {
		return fmt.Errorf("generate: setMasterOveride failed %ss", err.Error())
	}
	if err := d.removeI2PJpackageDir(); err != nil {
		return fmt.Errorf("generate: runI2PFirefoxCleanSh failed %ss", err.Error())
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

func (d *Daemon) gitPullI2PFirefox() error {
	dir := filepath.Join(d.Dir, "i2p.firefox")
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	err = w.Pull(&git.PullOptions{RemoteName: "origin", ReferenceName: plumbing.NewBranchReferenceName("master")})
	if err != nil {
		if err.Error() == "already up-to-date" {
			log.Printf("gitPullI2PFirefox: w.Pull failed %ss", err.Error())
			err = nil
		}
		return err
	}
	return nil
}

func (d *Daemon) removeI2PJpackageDir() error {
	dir := filepath.Join(d.Dir, "i2p.firefox-jpackage-build")
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}

func (d *Daemon) runI2PFirefoxBuildSh() error {
	dir := filepath.Join(d.Dir, "i2p.firefox")
	fmt.Println("Running build.sh")
	args := []string{"--login", "--interactive", filepath.Join(dir, "build.sh")}
	switch runtime.GOOS {
	case "windows":
		gitbash, err := filepath.Abs(filepath.Join("/Program Files/", "/Git/", "git-bash.exe"))
		if err != nil {
			return err
		}
		cmd := exec.Command(gitbash, args...)
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
		fmt.Println("Running wsl", "make", "version", "prep", "-C", "i2p.firefox")
		cmd := exec.Command("wsl", "make", "version", "prep", "-C", "i2p.firefox")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		fmt.Println("Running make", "version", "prep", "-C", "i2p.firefox")
		cmd := exec.Command("make", "version", "prep", "-C", filepath.Join(d.Dir, "i2p.firefox"))
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

func (d *Daemon) setMasterOveride() error {
	override := `
	DATE=$(($(date +%s) / 60 / 60 / 24))
	I2P_VERSION=1.8.0.${DATE}
	export I2P_VERSION=1.8.0.${DATE}
	VERSION=master
	export VERSION="$VERSION"
	`
	err := ioutil.WriteFile(filepath.Join(d.Dir, "i2p.firefox", "i2pversion_override"), []byte(override), 0644)
	if err != nil {
		return fmt.Errorf("setMasterOveride: WriteFile failed: %s", err.Error())
	}
	return nil
}
