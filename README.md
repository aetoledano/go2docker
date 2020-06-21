# go2docker

go2docker is a simple tool to dockerize a golang project. It produces a final tiny docker image based on `alpine` linux docker official image. 

Inspired on https://github.com/GoogleContainerTools/jib

## How to use it

Install: `go get -u github.com/aetoledano/go2docker`
Make sure $GOPATH/bin is in system path.

Run: `go2docker /path/to/golang/project`

## Custom go2docker configuration
Add a `go2docker.yml` file in project root. 
### Configuration file sample
```
app:
  name: my-awesome-app

go:
  version: 1.14

include-external-resources:
  - 'external.txt'
  - 'external-dir'
```

`app.name` is used for tagging the final docker image

`go.version` is the golang official image version used as builder

 `include-external-resources` is an array of files that should be included in the final image aside the main executable
 
 ## Troubleshooting
 There has been some issues between the docker sdk used and the docker instance running go2docker with a fix docker api version is a workaround:
 
 `DOCKER_API_VERSION=1.40 go2docker /path/to/awesome/project`
 
 As go2docker creates the docker client from env the following variables can be defined and will be used to build the docker client as of the docker sdk docs:
 
 Use DOCKER_HOST to set the url to the docker server.
 
 Use DOCKER_API_VERSION to set the version of the API to reach, leave empty for latest.
 
 Use DOCKER_CERT_PATH to load the tls certificates from.
 
 Use DOCKER_TLS_VERIFY to enable or disable TLS verification, off by default.
 
 