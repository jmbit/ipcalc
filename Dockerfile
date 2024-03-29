# Start from the latest golang base image
FROM golang:alpine as builder
RUN apk add --no-cache git

# Add Maintainer Info
LABEL maintainer="Johannes Bülow <johannes.buelow@jmbit.de>"

# Set the Current Working Directory inside the container
WORKDIR /ipcalc

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go ipcalc for different platforms
RUN GOOS=linux GOARCH=amd64 go build --buildvcs=true -o /go/bin/ipcalc-linux-amd64
RUN GOOS=linux GOARCH=arm64 go build --buildvcs=true -o /go/bin/ipcalc-linux-arm64
RUN GOOS=darwin GOARCH=amd64 go build --buildvcs=true -o /go/bin/ipcalc-mac-amd64
RUN GOOS=darwin GOARCH=arm64 go build --buildvcs=true -o /go/bin/ipcalc-mac-arm64
RUN GOOS=windows GOARCH=amd64 go build --buildvcs=true -o /go/bin/ipcalc-windows-amd64.exe
RUN GOOS=windows GOARCH=arm64 go build --buildvcs=true -o /go/bin/ipcalc-windows-arm64.exe

# Final stage
FROM scratch

# Copy the pre-built binaries from the previous stage
COPY --from=builder /go/bin/ipcalc-linux-amd64 ipcalc-linux-amd64
COPY --from=builder /go/bin/ipcalc-linux-arm64 ipcalc-linux-arm64
COPY --from=builder /go/bin/ipcalc-mac-amd64 ipcalc-mac-amd64
COPY --from=builder /go/bin/ipcalc-mac-arm64 ipcalc-mac-arm64
COPY --from=builder /go/bin/ipcalc-windows-amd64.exe ipcalc-windows-amd64.exe
COPY --from=builder /go/bin/ipcalc-windows-arm64.exe ipcalc-windows-arm64.exe
