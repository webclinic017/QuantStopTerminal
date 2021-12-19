.DEFAULT_GOAL := run
FALSE=false
TRUE=true

# Project settings
PROJECTNAME=QuantstopTerminal
MAJORVERSION=0
MINORVERSION=1
ISRELEASE=false

# Build system settings
GCVARS=GOARCH=amd64 CGO_ENABLED=0

# Production build command
PRODBUILDFLAGS=-X 'main.MajorVersion=${MAJORVERSION}' -X 'main.MinorVersion=${MINORVERSION}' -X 'main.IsRelease=${ISRELEASE}' -X 'main.IsDevelopment=${FALSE}'
PRODBUILD=${GCVARS} go build -ldflags="${PRODBUILDFLAGS}" -tags=prod

# Development build command
DEVBUILDFLAGS=-X 'main.MajorVersion=${MAJORVERSION}' -X 'main.MinorVersion=${MINORVERSION}' -X 'main.IsRelease=${ISRELEASE}' -X 'main.IsDevelopment=${TRUE}'
DEVBUILD=go run -ldflags="${DEVBUILDFLAGS}" -tags=dev .

# Run in development mode
run: npm-install
	cd cmd/quantstopterminal && ${DEVBUILD}

# Build production executables for all supported operating systems
build-all: build-windows
	build-mac && build-linux

# Build production executable for windows
build-windows: generate-compiled-assets
	cd cmd/quantstopterminal && GOOS=windows ${PRODBUILD} -o ../../builds/windows/${PROJECTNAME}.exe

# Build production executable for mac
build-mac: generate-compiled-assets
	cd cmd/quantstopterminal && GOOS=darwin ${PRODBUILD} -o ${PROJECTNAME}

# Build production executable for linux
build-linux: generate-compiled-assets
	cd cmd/quantstopterminal && GOOS=linux ${PRODBUILD} -o ${PROJECTNAME}

# Helper to build and compile the frontend
generate-compiled-assets: build-vue-app
	go run cmd/genassets/main.go
	go generate

# Build the vue app with npm
build-vue-app: npm-install
	cd web && npm run build

# Install npm packages
npm-install:
	cd web && npm install