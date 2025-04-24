name: CI
on:
push:
branches: [main]
pull_request:

jobs:
test:
runs-on: ubuntu-latest
services:
postgres:
image: postgres:16-alpine
env:
POSTGRES_USER: myapp
POSTGRES_PASSWORD: myapp
POSTGRES_DB: myapp_test
ports: ['5432:5432']
options: >-
--health-cmd="pg_isready -U myapp"
--health-interval=5s
--health-timeout=5s
--health-retries=5
steps:
- uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod         # ← go.mod の version を採用
          cache: true                     # modules + build 複合キャッシュ有効化
                                          # :contentReference[oaicite:0]{index=0}

      - name: Install migrate
        run: |
          curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.2/migrate.linux-amd64.tar.gz \
          | tar xz -C /usr/local/bin migrate
      - name: Apply migrations
        run: migrate -path migrations -database "postgres://myapp:myapp@localhost:5432/myapp_test?sslmode=disable" up

      - name: Run vet/lint
        run: |
          go vet ./...
          # staticcheck や revive を足すならここ

      - name: Run tests
        run: go test ./... -count=1 -race -coverprofile=coverage.out

      - name: Upload coverage to summary
        run: go tool cover -func=coverage.out | grep total >> $GITHUB_STEP_SUMMARY
