# BookingApp Overview
- Purpose: Clean Architecture sample for hotel reservation bookings with RESTful HTTP API backed by MySQL via GORM.
- Tech stack: Go 1.24+ modules, net/http, GORM (mysql driver), optional in-memory repos for tests.
- Structure: `cmd/api` bootstrap, `internal/domain` entities + repository contracts, `internal/usecase` orchestrates, `internal/interface/http` HTTP handlers, `internal/infrastructure/db` GORM models + migrations, `internal/infrastructure/memory` in-memory repos.
- Key patterns: layered Clean Architecture with repositories injected into usecases, HTTP handlers consuming usecases, optional DB seeding on startup.
- Running env: expects MySQL reachable via env vars `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`; seeds sample plans if table empty.