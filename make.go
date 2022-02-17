package I2P

//go:generate rm -rf i2p.firefox
//go:generate git clone https://i2pgit.org/i2p-hackers/i2p.firefox
//go:generate ./i2p.firefox/build.sh
//go:generate make -C i2p.firefox
//go:generate tar -C i2p.firefox/build/I2P -czf build.I2P.tar.gz

/* Express this Makefile segment with Go Generate.
jpackage: i2p.firefox i2p.firefox/I2P i2p.firefox/build/I2P build.I2P.tar.gz
	go build -o go-I2P-jpackage ./I2P

i2p.firefox:
	git clone https://i2pgit.org/i2p-hackers/i2p.firefox

i2p.firefox/I2P:
	cd i2p.firefox && \
		./build.sh

i2p.firefox/build/I2P:
	cd i2p.firefox && \
		make

build.I2P.tar.gz:
	cd i2p.firefox/build/I2P && \
		tar -czvf ../../../build.I2P.tar.gz .
*/

/*func gitCloneI2PFirefox(dir string) {
	//git clone https://i2pgit.org/i2p-hackers/i2p.firefox
}

func runI2PFirefoxBuildSh(dir string) {
	//cd i2p.firefox && ./build.sh
}

func runI2PFirefoxMake(dir string) {
	//cd i2p.firefox && make
}*/
