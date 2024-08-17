#!/bin/bash

PROJECT_PATH="$1"
CUSTOM_PATH="$2"

if [[ -z $CUSTOM_PATH ]]; then
    echo "CUSTOM_PATH was not provided. Trying to install into GOBIN..."
else
    go build -o "$CUSTOM_PATH/passman" "$PROJECT_PATH"
    exit
fi

if [[ -z $GOBIN ]]; then
    echo "GOBIN is empty... installing into /usr/local/bin/"
else
    go build -o "$GOBIN/passman" "$PROJECT_PATH"
    exit
fi

sudo go build -o /usr/local/bin/passman "$PROJECT_PATH"
