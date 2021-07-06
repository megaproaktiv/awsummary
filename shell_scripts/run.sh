#!/bin/sh

DIR="${BASH_SOURCE%/*}"
if [[ ! -d "$DIR" ]]; then
  DIR="$PWD"
fi
source "$DIR/setenv.sh"

if [[ ! -d "$AO_EXECUTABLE_PATH" ]]; then
  echo "Executable path is not accecible: $AO_EXECUTABLE_PATH"
  exit 5
fi

echo "Starting..."
$AO_EXECUTABLE_PATH/aws-overview -log-file=$AO_LOG_PATH/aws.log -daemon=$AO_REPEAT_TIME
