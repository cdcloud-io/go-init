#!/usr/bin/env bash

# ==============================================================================
# Title          : Go Init Script
# Description    : This script sets up a Go project environment and should be sourced.
# Company        : cdcloud-io
# Author         : cd-stephen
# References     : [URLs or other references]
# Last Modified  : 7/2/2024
# Version        : 1.0
# Usage          : source go_init.sh
# Notes          : Ensure this script is sourced to run the function in the current shell.
# ==============================================================================

clear
echo '🟨 sourcing go_init.sh'
sleep 1

# Function to set up a Go project
function go_init() {
    # Check if the script is being run in a subdirectory of the user's home directory
    if [ "$PWD" == "$HOME" ] || [[ "$PWD" != "$HOME/"* ]]; then
        echo ''
        echo "🟥 Error: Script must be run in a subdirectory of ${HOME}. Exiting..."
        return 1
    fi

    # Check if the directory contains only .git/ and README.md
    for file in * .*; do
        if [ "$file" != "." ] && [ "$file" != ".." ] && [ "$file" != "*" ] && [ "$file" != ".git" ] && [ "$file" != "README.md" ]; then
            echo ''
            echo "🟥 Error: Not an empty project. Only .git/ and README.md are allowed. Exiting..."
            return 1
        fi
    done

    # Download the Makefile and .gitignore
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/Makefile -O Makefile
    wget -q https://raw.githubusercontent.com/cdcloud-io/go_init/main/.gitignore -O .gitignore
    
    # Extract the MODULE_NAME from the current directory name
    MODULE_NAME=$(basename "$(pwd)")
    
    # Extract the URL_PATH from the Git configuration if a .git directory exists
    if [ -d ".git" ]; then
        GIT_URL=$(git config --get remote.origin.url)
        echo "GIT_URL: $GIT_URL"
        URL_PATH=$(echo "$GIT_URL" | sed -E "s|git@([^:]+):([^/]+/[^/]+)\.git$|\\1/\\2|")
        echo "GO MOD: $URL_PATH"
    else
        URL_PATH=""
    fi
    
    # Use sed to replace the placeholders in the Makefile
    sed -i "s|^MODULE_NAME :=.*|MODULE_NAME := $MODULE_NAME|" Makefile
    sed -i "s|^URL_PATH :=.*|URL_PATH := $URL_PATH|" Makefile

    echo ''
    echo "🟩 Makefile has been set up with MODULE_NAME: $MODULE_NAME and URL_PATH: $URL_PATH"
    echo ''

    if [ -z "${URL_PATH}" ]; then
        echo "Initializing Go module..."
        go mod init "${MODULE_NAME}"
    else
        echo "Initializing Go module with URL path..."
        go mod init "${URL_PATH}"
    fi

    # Create necessary directories
    mkdir -p api > /dev/null 2>&1                            ## openapi spec
    mkdir -p bin > /dev/null 2>&1                            ## compilation bin destination
    mkdir -p build/{docker,k8s/kustomize} > /dev/null 2>&1   ## scripts for build, run, deploy
    mkdir -p cmd/${MODULE_NAME} > /dev/null 2>&1             ## application entry point. main.go
    mkdir -p config > /dev/null 2>&1                         ## config.yaml used by internal/config/config.go
    mkdir -p docs/img > /dev/null 2>&1                       ## module/app documentation and images
    mkdir -p example > /dev/null 2>&1                        ## optional use for app/code usage examples
    mkdir -p internal/{config,endpoint/user,middleware/auth,middleware/logging,middleware/trace,server} > /dev/null 2>&1 ## module/app internal packages
    mkdir -p test > /dev/null 2>&1                           ## unit/integration tests

    # Create README.md if it does not exist
    if [ ! -f README.md ]; then
        printf "# %s" "${MODULE_NAME}" > README.md
        echo ''
        echo '🟩 INFO: Go module has been initialized'
        echo ''
    else
        rm -f README.md
        printf "# %s" "${MODULE_NAME}" > README.md
        echo ''
        echo '🟨 WARN: README.md has been modified'
        echo '🟩 INFO: Go module has been initialized'
        echo ''
    fi
}

clear
echo '🟩 sourcing go_init.sh'

# Make sure to source this file in .bashrc
# source /path/to/this/file/go_init.sh or source $HOME/go_init.sh
