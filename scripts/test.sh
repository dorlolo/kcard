#!/usr/bin/env bash
set -euo pipefail
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
( cd "$ROOT_DIR/backend" && go test ./... )
( cd "$ROOT_DIR/frontend" && npm run lint && npm run typecheck && npm run test && npm run e2e )
