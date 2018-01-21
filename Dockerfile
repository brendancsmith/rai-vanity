FROM golang:1.9 AS gobuild

WORKDIR /go/src/github.com/brendancsmith/rai-vanity
RUN go get github.com/spf13/cobra \
  github.com/frankh/crypto/ed25519 \
  github.com/golang/crypto/blake2b \
  github.com/frankh/rai \
  github.com/a-h/round
ADD . ./

RUN go build -o rai-vanity .

FROM alpine

COPY --from=gobuild /go/src/github.com/brendancsmith/rai-vanity/rai-vanity /rai-vanity

ENTRYPOINT ["/rai-vanity"]
