FROM golang as build
WORKDIR /go/src/quantstopterminal
COPY . .

RUN GO111MODULE=on go mod vendor

RUN GOARCH=386 GOOS=linux go run ./cmd/genassets/main.go
RUN go generate
RUN GOARCH=386 GOOS=linux go build -ldflags="-X=main.BuildFlagVersion=0.1.0 -X=main.BuildFlagIsRelease=true" -tags=prod -o qst ./cmd/qst/main_linux.go
RUN GOARCH=386 GOOS=linux go build ./cmd/qstcli/

RUN mv qst /go/bin/qst
RUN mv qstcli /go/bin/qstcli

FROM alpine:latest

# Set up data directory
#VOLUME /root/.qstrader
#COPY /strategy/examples /root/.qstrader/scripts/examples
RUN apk update && apk add --no-cache ca-certificates bash go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

COPY --from=build /go/bin/qst /go/bin
COPY --from=build /go/bin/qstcli /go/bin

#COPY --from=build /go/src/github.com/thrasher-corp/gocryptotrader/config.json /root/.gocryptotrader/
ENTRYPOINT [ "/go/bin/qst" ]