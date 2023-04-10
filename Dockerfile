# build stage
FROM golang:1.20.2-alpine3.17 AS build-env
RUN apk --no-cache add git
ENV D=/go/src/github.com/fnproject/fn
ADD . $D
RUN cd $D/cmd/fnserver && go build -o fn-alpine && cp fn-alpine /tmp/

FROM docker:23.0.2-cli-alpine3.17
WORKDIR /app
COPY --from=build-env /tmp/fn-alpine /app/fnserver
CMD ["./fnserver"]
EXPOSE 8080
