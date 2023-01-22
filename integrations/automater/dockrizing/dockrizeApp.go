package main

import (
	"os"
	"text/template"
)

/*
This script does the following:

1. It imports the necessary Go packages for working with files and templates.
2. It defines two constant strings, dockerfile and composefile, that contain the contents of the Dockerfile and docker-compose.yml files respectively.
3. In the main function it creates two template, one for Dockerfile and one for docker-compose.yml
4. It creates two files Dockerfile and docker-compose.yml and writes the contents of the templates to them.

The Dockerfile in this example uses the official Go image as the base image, sets the working directory to /go/src/app, copies the application files to the container, runs go get and go install, exposes port 8080 and runs the command app when the container starts.

The docker-compose.yml in this example defines a service called app that build the current directory and maps port 8080 on the host to port 8080 in the container.


*/

const dockerfile = `
FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["app"]
`

const composefile = `
version: '3'
services:
  app:
      build: .
	      ports:
		        - "8080:8080"
				`

func main() {
	// create the Dockerfile
	dockerfileTemplate, err := template.New("dockerfile").Parse(dockerfile)
	if err != nil {
		panic(err)

	}
	dockerfileFile, err := os.Create("Dockerfile")
	if err != nil {
		panic(err)

	}
	if err := dockerfileTemplate.Execute(dockerfileFile, nil); err != nil {
		panic(err)

	}
	dockerfileFile.Close()

	// create the docker-compose.yml
	composefileTemplate, err := template.New("composefile").Parse(composefile)
	if err != nil {
		panic(err)

	}
	composefileFile, err := os.Create("docker-compose.yml")
	if err != nil {
		panic(err)

	}
	if err := composefileTemplate.Execute(composefileFile, nil); err != nil {
		panic(err)

	}
	composefileFile.Close()

}
