package I2P

//import "github.com/walle/targz"
import (
	"compress/gzip"
	"fmt"

	"github.com/mholt/archiver"
)

func TarGzip(source, target string) error {
	tgz := archiver.NewTarGz()
	tgz.CompressionLevel = gzip.BestCompression
	err := tgz.Archive([]string{source}, target)
	if err != nil {
		return fmt.Errorf("TarGzip: TarGz() failed: %s", err.Error())
	}
	return nil
}
