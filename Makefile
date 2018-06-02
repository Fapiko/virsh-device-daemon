build-prep:
	rm -rf target
	mkdir target

build-linux64: build-prep
	GOOS=linux GOARCH=amd64 go build -o target/virsh-device-daemon github.com/fapiko/virsh-device-daemon

build: build-linux64
	GOOS=windows GOARCH=amd64 go build -o target/virsh-device-daemon-win64.exe github.com/fapiko/virsh-device-daemon

deb: build-linux64
	cp -r build/deb target/
	mkdir -p target/deb/usr/sbin
	chmod 0755 target/virsh-device-daemon
	mv target/virsh-device-daemon target/deb/usr/sbin
	find ./target/deb -type d | xargs chmod 0755
	fakeroot dpkg-deb --build target/deb
	mv target/deb.deb target/virsh-device-daemon_0.1.0_amd64.deb

test-deb: deb
	cp -r build/docker target/
	cp target/virsh-device-daemon_0.1.0_amd64.deb target/docker/
	docker build -t virsh-device-daemon target/docker

shell-test-deb:
	docker run -it virsh-device-daemon /bin/bash
