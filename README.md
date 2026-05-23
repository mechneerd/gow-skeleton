# GoW Skeleton

This repository contains the official project skeletons for the **GoW Framework**.

When a user runs `gow new myapp`, the GoW CLI clones this repository and copies one of the templates into the new project.

## Available Starter Kits

| Starter Kit   | Command                     | Description                              |
|---------------|-----------------------------|------------------------------------------|
| minimal       | `gow new myapp --minimal`   | Basic routing + views                    |
| api           | `gow new myapp --api`       | REST API with Sanctum                    |
| web           | `gow new myapp`             | Full web app with Blade templating       |
| web-auth      | `gow new myapp --auth`      | Web app + complete authentication system |

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
