package I2P

//import "github.com/walle/tarlz4"
import (
	"fmt"

	"github.com/mholt/archiver"
)

func TarGzip(source, target string) error {
	tlz4 := archiver.NewTarLz4()
	err := tlz4.Archive([]string{source}, target)
	if err != nil {
		return fmt.Errorf("TarGzip: TarGz() failed: %s", err.Error())
	}
	return nil
}
