#!/bin/bash

set -e

# Change to directory if it exists (for container environments)
if [ -d /app ]; then
  cd /app
fi

# Start the python Server
uv run -m app
