#!/usr/bin/env bash
set -euo pipefail

PORT="${PORT:-3000}"

echo "Ensuring port ${PORT} is free..."

kill_by_pids() {
  local pids="$1"
  [[ -z "${pids// /}" ]] && return 0
  echo "Killing PIDs on port ${PORT}: ${pids}"
  kill ${pids} 2>/dev/null || true
  sleep 0.3
  local still=""
  for pid in ${pids}; do
    if kill -0 "${pid}" 2>/dev/null; then
      still+=" ${pid}"
    fi
  done
  if [[ -n "${still// /}" ]]; then
    echo "Force killing PIDs:${still}"
    kill -9 ${still} 2>/dev/null || true
  fi
}

freed=0

if command -v lsof >/dev/null 2>&1; then
  pids=$(lsof -t -i TCP:"${PORT}" -sTCP:LISTEN 2>/dev/null || true)
  [[ -n "${pids// /}" ]] && kill_by_pids "${pids}" && freed=1
fi

if [[ "${freed}" -eq 0 ]] && command -v fuser >/dev/null 2>&1; then
  # fuser kills directly
  fuser -k "${PORT}/tcp" 2>/dev/null || true
  freed=1
fi

if [[ "${freed}" -eq 0 ]] && command -v ss >/dev/null 2>&1; then
  pids=$(ss -ltnp 2>/dev/null | awk -v port=":${PORT}" '$4 ~ port { match($6, /pid=([0-9]+)/, m); if (m[1] != "") print m[1]; }' | xargs -r echo)
  [[ -n "${pids// /}" ]] && kill_by_pids "${pids}" && freed=1
fi

echo "Starting local-main.go on port ${PORT}..."
exec go run ./local-main.go

