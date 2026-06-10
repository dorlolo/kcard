#!/usr/bin/env bash
set -euo pipefail
printf 'Start PostgreSQL and Redis locally, then run backend API, worker, and frontend in separate shells.\n'
printf 'Backend API: (cd backend && go run ./cmd/api)\n'
printf 'Worker:      (cd backend && go run ./cmd/worker)\n'
printf 'Frontend:    (cd frontend && npm install && npm run dev)\n'
