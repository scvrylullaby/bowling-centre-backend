#!/usr/bin/env bash

echo "Starting dev mode..."

if [ -f ".env" ]; then
  export $(grep -v '^#' .env | xargs)
fi

air
