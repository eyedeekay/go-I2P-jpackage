package I2P

import "testing"

func TestPrintLicense(t *testing.T) {
	s, err := GetLicenses()
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}
