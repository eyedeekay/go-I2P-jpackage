package I2P

import "fmt"

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
