# Style and Conventions
- Standard Go style: gofmt formatting, short receiver names, exported types for domain entities/usecases, unexported struct fields where possible.
- Clean Architecture layering: domain entities free of infra deps; usecases depend on domain repository interfaces; interface/http packages convert transport payloads.
- Error handling: return Go errors, map to HTTP status codes in handlers; prefer sentinel errors (`ErrInvalidDates`, etc.).
- Repositories expose interfaces for Plan and Reservation storage; concrete adapters live under infrastructure (MySQL, in-memory).
- Testing not present yet; in-memory repos can support unit tests without DB.
- Comments primarily Japanese/English mix, keep brief purposeful notes when logic not obvious.