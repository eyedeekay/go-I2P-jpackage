
jpackage:
	/usr/bin/go generate

clean:
	git clean -fd
	rm -rf i2p.firefox i2p.i2p-jpackage-build

version: release
	UPLOAD_OS=linux make upload
	UPLOAD_OS=windows make upload

USER=eyedeekay
REPO=go-I2P-jpackage
TITLE=Java I2P Jpackage\(non-go components\)
UPLOAD_OS?=linux
VERSION=1.7.0-`date +%Y%m%d`

release:
	gothub release \
		--pre-release \
		--user $(USER) \
		--repo $(REPO) \
		--name "$(TITLE) - $(VERSION)" \
		--description "$(cat desc)" \
		--tag v$(VERSION); true

SUM=`sha256sum build.$(UPLOAD_OS).I2P.tar.xz`

upload:
	gothub upload \
		--replace \
		--user $(USER) \
		--repo $(REPO) \
		--tag v$(VERSION) \
		--name "build.$(UPLOAD_OS).I2P.tar.xz" \
		--label "$(SUM)" \
		--file "build.$(UPLOAD_OS).I2P.tar.xz"
