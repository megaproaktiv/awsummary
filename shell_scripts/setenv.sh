#!/bin/sh

AO_EXECUTABLE_PATH=/path/to/executable
AO_LOG_PATH=/path/to/log
AO_REPEAT_TIME=180

export AO_EXECUTABLE_PATH
echo "Exported AO_EXECUTABLE_PATH: $AO_EXECUTABLE_PATH"
export AO_LOG_PATH
echo "Exported AO_LOG_PATH: $AO_LOG_PATH"
export AO_REPEAT_TIME
echo "Exported AO_REPEAT_TIME: $AO_REPEAT_TIME"
