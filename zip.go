package I2P

//import "github.com/walle/tarxz"
import (
	"fmt"

	"github.com/mholt/archiver"
)

func TarXzip(source, target string) error {
	txz := archiver.NewTarXz()
	//	txz.CompressionLevel = 12
	err := txz.Archive([]string{source}, target)
	if err != nil {
		return fmt.Errorf("TarGzip: TarGz() failed: %s", err.Error())
	}
	return nil
}
