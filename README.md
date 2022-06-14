# Spectrum

[![Go Report Card](https://goreportcard.com/badge/github.com/maxthom/spectrum-go)](https://goreportcard.com/report/github.com/maxthom/spectrum-go)

https://github.com/rpi-ws281x/rpi-ws281x-go

[LED Colors](https://www.springtree.net/audio-visual-blog/rgb-led-color-mixing/)

## Deployment on prod devices
```
curl https://codeload.github.com/maxthom/spectrum-go/tar.gz/main | tar -xz --strip=2 spectrum-go-main/build/docker-compose
rn ledstrip_default.json ledstrip.json
chmod +x docker_installer.sh
./docker_installer.sh
docker-compose up
```

## File structure

```
/cmd/<app_name>: main applications for this project.
/internal: source code that are private to this module.
/pkg: source code that can be shared.
/vendor: a pplication dependencies.
/api: openAPI/Swagger specs, JSON schema files, protocol definition files.
/web: web application specific components: static web assets, server side templates and SPAs.
/configs: configuration file templates or default configs.
/init: system init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.
/scripts: scripts to perform various build, install, analysis, etc operations.
/build: packaging and Continuous Integration. Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the /build/package directory.
/deployments: IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).
/test: additional external test apps and test data.
/docs: Design and user documents (in addition to your godoc generated documentation).
/tools: Supporting tools for this project. Note that these tools can import code from the /pkg and /internal directories.
/examples: Examples for your applications and/or public libraries.
/third_party: External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).
/assets: Other assets to go along with your repository (images, logos, etc).
/website: This is the place to put your project's website data if you are not using GitHub pages.
/githooks: Git Hooks.
```
## Docker

```sh
docker build --tag spectrum:latest .
```

```sh
docker run --rm spectrum:latest
```

## Development 🧑‍💻

1.  Install [rpi-ws281x](https://github.com/jgarff/rpi_ws281x) C variant

```sh
git clone https://github.com/jgarff/rpi_ws281x.git
cd rpi_ws281x
mkdir build
cd build
cmake -D BUILD_SHARED=OFF -D BUILD_TEST=OFF ..
cmake --build .
sudo make install
```
2. Create go project, install [rpi-ws281x](https://github.com/rpi-ws281x/rpi-ws281x-go) Go variant
```sh
go mod init github.com/maxthom/spectrum-go
go get github.com/rpi-ws281x/rpi-ws281x-go
go mod tidy
go mod vendor
```
3. Run and test project
```sh
go build -o swiss
sudo ./swiss
```


docker run --rm -v "$PWD":/usr/src/$APP -w /usr/src/$APP ws2811-builder:latest go build -o "$swiss-docker" -v

## Install docker
```sh
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker ${USER}
groups ${USER}
‍sudo su - ${USER}
docker version
docker info
docker run hello-world
docker image rm hello-world
```

## Install docker-compose
```sh
sudo apt-get install libffi-dev libssl-dev
sudo apt install python3-dev
sudo apt-get install -y python3 python3-pip
sudo pip3 install docker-compose
‍sudo systemctl enable docker
docker-compose version
```

## Install go
```sh
chmod +x scripts/go_installer.sh
./go_installer.sh
```