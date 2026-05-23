Write-Host "Running post-install for {{ .AppName }}..." -ForegroundColor Cyan

go mod tidy

if (Test-Path ".env.example" -and -not (Test-Path ".env")) {
    Copy-Item ".env.example" ".env"
    Write-Host ".env file created from .env.example" -ForegroundColor Green
}

Write-Host "Post-install completed." -ForegroundColor Green
Write-Host "Run: gow serve" -ForegroundColor Yellow
