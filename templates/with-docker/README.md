# Docker Support for GoW Projects

This kit adds Docker and docker-compose support to any GoW starter.

## Files Included

- `Dockerfile`
- `docker-compose.yml`
- `.dockerignore`

## Usage

```bash
docker-compose up --build
```

## Status
**In Progress** - Basic Docker files are being added.

## Recommended Combination
Use with `full`, `web-auth`, or `admin-panel` for a production-like local development environment.

## New Features Available
All kits now include:
- Beautiful landing page at `/`
- GoW Livewire for reactive UI
- `gow make:view` command (creates `.goblade` files by default)

These work out of the box when you run the container.
