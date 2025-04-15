#!/usr/bin/env bash
#### build-dockerfile.sh ####

# global variables
_FILENAME="build-dockerfile.sh"
_REQUIRED_GO_VERSION="1.22.0"
_INVOKED_SCRIPT_NAME=$(basename "$0")
_INVOKED_SCRIPT_PATH=$(dirname "$0")
_SCRIPT_PATH=$(basename "$0")
_ENV="local"
_REGISTRY_URL_PORT="localhost:5000"

# Check if the script is executed from the project root
executionPathCheck() {
  if [ "$_INVOKED_SCRIPT_NAME" = "$_FILENAME" ] && [ "$_INVOKED_SCRIPT_PATH" = "./build" ]; then
    echo "🟩 INVOKED_SCRIPT_NAME is: $_INVOKED_SCRIPT_NAME"
    echo "🟩 INVOKED_SCRIPT_PATH is: $_INVOKED_SCRIPT_PATH"

    echo "🟩 script was executed as ./build/${_FILENAME}"
  else
    echo "🟥 ERROR: script must be executed from project root:"
    echo "CMD: ./builds/${_FILENAME}"
    echo "INVOKED_SCRIPT_NAME is: $_INVOKED_SCRIPT_NAME"
    echo "INVOKED_SCRIPT_PATH is: $_INVOKED_SCRIPT_PATH"
    exit 1
  fi
}
executionPathCheck
sleep 1

# Set application name from go.mod.
if [[ -f "./go.mod" ]]; then
    _GO_MOD=$(awk '/^module /{print $2}' ./go.mod | xargs 2>/dev/null)
    if [[ -z "$_GO_MOD" ]]; then
        echo "🟥 ERROR: Could not find module name in go.mod file."
        echo 'Exiting...'
        exit 1
    fi
else
    echo "🟥 ERROR: go.mod file does not exist. Exiting..."
    exit 1
fi

# function declarations
pause() {
    read -n1 -r -p "💻 Press any key to continue... 💻" key
}

goclean() {
    go mod tidy
    go mod download
    go fmt ./...
}

# Check Go version
_GO_VERSION=$(go version | awk '{print $3}' | cut -d "o" -f 2)

if [[ $(printf '%s\n' "$_REQUIRED_GO_VERSION" "$_GO_VERSION" | sort -V | head -n1) != "$_REQUIRED_GO_VERSION" ]]; then
    echo "🟥 ERROR: Go version must be greater than $_REQUIRED_GO_VERSION. Current version: $_GO_VERSION"
    exit 1
fi

# Check if config.yaml file exists
if [[ -f "./config/config.yaml" ]]; then
    # Read the app version from the config file using awk and trim whitespace
    _APP_VERSION=$(awk -F': ' '/version:/ {print $2}' ./config/config.yaml | tr -d '"' | xargs 2>/dev/null)
    if [[ -z "$_APP_VERSION" ]]; then
        echo "🟥 ERROR: Could not find version in config file. Exiting..."
        exit 1
    fi
else
    echo "🟥 ERROR: config.yaml file does not exist. Exiting..."
    exit 1
fi

# Check if the version starts with 'v'
if [[ $_APP_VERSION != v* ]]; then
    _APP_VERSION="v$_APP_VERSION"
fi

# Export ENV to current shell for application startup
export _REGISTRY_IMAGE_NAME=${_REGISTRY_URL_PORT}/${_GO_MOD}
export _APP_ENV=${_ENV}
export _APP_NAME="$_GO_MOD"
export _APP_VERSION

# Get the short commit SHA, default to "abcd1234" if null
_APP_COMMIT_SHA=$(git rev-parse --short HEAD 2>/dev/null)
if [[ -z "$_APP_COMMIT_SHA" ]]; then
    _APP_COMMIT_SHA="abcd1234"
fi

export _APP_COMMIT_SHA
export _APP_BUILD_ID="localbuild"
export _APP_BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

goclean

clear
echo '🟦 Building Dockerfile image with the following variables 🟦'
echo ''
echo '################################################################'
echo ''
echo "_APP_NAME: ................. $_APP_NAME"
echo "_APP_VERSION: .............. $_APP_VERSION"
echo "_APP_COMMIT_SHA: ........... $_APP_COMMIT_SHA"
echo "_APP_BUILD_ID: ............. $_APP_BUILD_ID"
echo "_APP_BUILD_DATE: ........... $_APP_BUILD_DATE"
echo "_REGISTRY_IMAGE_NAME: ...... $_REGISTRY_IMAGE_NAME"
echo ''
echo '################################################################'

pause

echo '⚙️ envoking docker build ⚙️'
docker build \
--build-arg=_APP_NAME=${_APP_NAME} \
--build-arg=_APP_VERSION=${_APP_VERSION} \
--build-arg=_APP_COMMIT_SHA=${_APP_COMMIT_SHA} \
--build-arg=_APP_ORIGIN=${_APP_ORIGIN} \
--build-arg=_APP_BUILD_ID=${_APP_BUILD_ID} \
--build-arg=_APP_BUILD_DATE=${_APP_BUILD_DATE} \
--tag=${_REGISTRY_IMAGE_NAME}:${_APP_COMMIT_SHA} \
--tag=${_REGISTRY_IMAGE_NAME}:${_APP_VERSION} \
--tag=${_REGISTRY_IMAGE_NAME}:${_APP_VERSION}.${_APP_COMMIT_SHA} \
--tag=${_REGISTRY_IMAGE_NAME}:latest \
--file=./build/docker/Dockerfile .
