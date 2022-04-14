# use golang 1.17 version as base layer
FROM golang:1.17-alpine as base
# create working dir
RUN mkdir /my-app
# add from root to workdir
ADD . /my-app
# define it as workding
WORKDIR /my-app
# build the app with no dependences
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url-requestor cmd/main.go
# using second layer to make image smaller
FROM scratch
# copy from base layer to current one
COPY --from=base /my-app/url-requestor /usr/bin/url-requestor
# export port, should be same as in config.yaml
EXPOSE 8080
# define entry point of container
ENTRYPOINT ["/usr/bin/url-requestor"]