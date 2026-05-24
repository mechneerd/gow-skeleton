# GoW Skeleton

This repository contains the official project skeletons for the **GoW Framework**.

When a user runs `gow new myapp`, the GoW CLI clones this repository and copies one of the templates into the new project.

## Available Starter Kits

| Starter Kit       | Command                              | Description                                      | Status      |
|-------------------|--------------------------------------|--------------------------------------------------|-------------|
| minimal           | `gow new ... --minimal`              | Basic routing + views                            | Ready       |
| minimal-api       | `gow new ... --skeleton=...`         | Ultra-light API only (with basic auth)           | Ready       |
| api               | `gow new ... --api`                  | REST API with Sanctum                            | Ready       |
| web               | `gow new ...`                        | Full web app with Blade                        | Ready       |
| web-auth          | `gow new ... --auth`                 | Web + complete authentication + RBAC             | Ready       |
| full              | `gow new ... --skeleton=...`         | Web + API + Auth + RBAC + advanced structure     | Ready       |
| admin-panel       | -                                    | Dashboard / Internal tools                       | Planned     |
| with-docker       | -                                    | Includes Docker + docker-compose                 | Planned     |
| inertia-react     | -                                    | Inertia.js + React                               | Planned     |
| inertia-vue       | -                                    | Inertia.js + Vue 3                               | Planned     |
| livewire          | -                                    | Livewire-style reactive components               | Planned     |

## How It Works

1. `gow` CLI clones this repo (shallow clone).
2. It selects the correct template based on flags or interactive prompts.
3. It replaces placeholders (`{{ .AppName }}`, `{{ .ModulePath }}`, etc.).
4. It runs post-install steps (`go mod tidy`, `.env` creation, etc.).
5. The user gets a clean, ready-to-use project with all files in their own directory.

## Contributing

See `STRUCTURE.md` for how the repository is organized and how to add new starter kits.

## License

MIT
