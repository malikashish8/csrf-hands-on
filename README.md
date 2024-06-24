# CSRF Hand-on

Learn CSRF practically through various use cases. Use cases are listed in [web](/web/) folder.

Victim url is https://localhost:8000/
Attacker url is https://localhost:8001/

## Run

- Open 2 bash tabs
- in 1 tab run

  `go run victim.go`

- in 2nd run

  `go run attacker.go`

- open <https://localhost:8001> in a private/incognito browser window and visit the init link to initiate a session with the victim domain.

## Prerequisites

- Go 1.11
