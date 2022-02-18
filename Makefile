
jpackage: i2p.firefox i2p.firefox/I2P i2p.firefox/build/I2P build.I2P.tar.gz
	go build -o go-I2P-jpackage ./I2P

i2p.firefox:
	git clone https://i2pgit.org/i2p-hackers/i2p.firefox

i2p.firefox/I2P:
	./i2p.firefox/build.sh
		
i2p.firefox/build/I2P:
	make -C i2p.firefox

build.I2P.tar.gz:
	tar -C i2p.firefox/build/I2P -czf build.I2P.tar.gz .

tar: build.I2P.tar.gz

clean:
	git clean -fxd
