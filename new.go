package I2P

import (
	"fmt"

	"github.com/mholt/archiver"
)

func NewDaemon(dir string, generate bool) (*Daemon, error) {
	I2Pdaemon := &Daemon{
		Dir: dir,
	}
	if generate {
		err := I2Pdaemon.Generate()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Generate() called for code generation hack. Not continuing to start I2P router")
	}
	err := I2Pdaemon.Unpack()
	if err != nil {
		return nil, err
	}
	return I2Pdaemon, nil
}

func UnTarXzip(source, target string) error {
	txz := archiver.NewTarXz()
	txz.Tar.OverwriteExisting = true
	txz.Tar.ContinueOnError = true
	err := txz.Unarchive(source, target)
	if err != nil {
		return fmt.Errorf("TarGzip: Unarchive() failed: %s", err.Error())
	}
	return nil
}
