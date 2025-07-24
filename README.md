# WASAPhoto

**A full-stack project in Go and Vue.js**

![YAML](https://img.shields.io/badge/YAML-85EA2D?style=for-the-badge&logo=YAML&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)
![Vue.js](https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D)
![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)

WASAPhoto is a social network where users can post photos, leave likes, comments and
also ban other users, with all the implications about information hiding.

It consists of:

* Documented REST API (OpenAPI 3.0) with all the endpoints described.
  You can find the specification [here](doc/api.yaml)
* Golang backend which implements the REST API. According to the given project
  specification, an authentication mechanism is not provided. Instead,
  the User ID is sent as an Authorization Bearer header, as it was a token in some way.
* Vue.js frontend app, which of course interfaces with the implemented REST API.
* All distributed using a Docker image

## Project structure

* `cmd/` contains all executables; Go programs here should only do "executable-stuff", like reading options from the CLI/env, etc.
	* `cmd/healthcheck` is an example of a daemon for checking the health of servers daemons; useful when the hypervisor is not providing HTTP readiness/liveness probes (e.g., Docker engine)
	* `cmd/webapi` contains an example of a web API server daemon
* `demo/` contains a demo config file
* `doc/` contains the documentation (usually, for APIs, this means an OpenAPI file)
* `service/` has all packages for implementing project-specific functionalities
	* `service/api` contains an example of an API server
	* `service/globaltime` contains a wrapper package for `time.Time` (useful in unit testing)
* `vendor/` is managed by Go, and contains a copy of all dependencies
* `webui/` is an example of a web frontend in Vue.js; it includes:
	* Bootstrap JavaScript framework
	* a customized version of "Bootstrap dashboard" template
	* feather icons as SVG
	* Go code for release embedding

Other project files include:
* `open-npm.sh` starts a new (temporary) container using `node:lts` image for safe web frontend development (you don't want to use `npm` in your system, do you?)


## How to build

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:

```shell
go build ./cmd/webapi/
```

If you're using the WebUI and you want to embed it into the final executable:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

## How to run (in development mode)

You can launch the backend only using:

```shell
go run ./cmd/webapi/
```

If you want to launch the WebUI, open a new tab and launch:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run dev
```

## How to run (Docker)

**Backend**

```shell
docker build -f Dockerfile.backend -t backend:latest .
docker run -it --rm -p 3000:3000 backend:latest
```

**Frontend**

```shell
docker build -f Dockerfile.frontend -t frontend:latest .
docker run -it --rm -p 8080:80 frontend:latest
```

## License

See [LICENSE](LICENSE).
