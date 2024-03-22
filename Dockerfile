FROM docker.io/golang:alpine AS builder
RUN mkdir /app
RUN apk add git make --no-cache
COPY . /app
WORKDIR /app
RUN make test
RUN make ipcalc

FROM scratch AS binaries
COPY --from=builder /app/release/* /


