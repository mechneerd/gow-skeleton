#!/usr/bin/env bash
echo "Running post-install for {{ .AppName }}..."

go mod tidy

if [ -f ".env.example" ] && [ ! -f ".env" ]; then
    cp .env.example .env
    echo ".env file created from .env.example"
fi

echo "Post-install completed."
echo "Run: gow serve"
