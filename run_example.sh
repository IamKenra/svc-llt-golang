#!/bin/bash

# Untuk macOS
kill -9 $(lsof -ti:3000) 2>/dev/null || true

# Tunggu sebentar
sleep 1

# Jalankan server
go run ./cmd/server/main.go