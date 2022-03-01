
jpackage:
	go generate

clean:
	git clean -fxd
	rm -rf i2p.firefox

version: release
	UPLOAD_OS=linux make upload
	UPLOAD_OS=windows make upload

USER=eyedeekay
REPO=go-I2P-jpackage
TITLE="Java I2P Jpackage(non-go components)"
UPLOAD_OS?=linux
VERSION=1.7.0

release:
	gothub release \
		--pre-release \
		--user $(USER) \
		--repo $(REPO) \
		--name "$(TITLE) - $(VERSION)" \
		--description `cat desc` \
		--tag v$(VERSION); true

upload:
	gothub upload \
		--replace \
		--user $(USER) \
		--repo $(REPO) \
		--tag v$(VERSION) \
		--name "build.$(UPLOAD_OS).I2P.tar.xz" \
		--label `sha256sum build.$(UPLOAD_OS).I2P.tar.xz` \
		--file "build.$(UPLOAD_OS).I2P.tar.xz"
