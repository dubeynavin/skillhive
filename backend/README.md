# Backend (Go API)

This folder contains the backend Go API for the SkillHive project.

Prerequisites
- Go 1.23+
- MongoDB (Atlas or local)

Setup
1. Copy the example env file:

   cp .env.example .env

2. Update environment variables in `.env` (see `.env.example`).

3. Install dependencies (Go modules automatically used):

   go mod download

Run locally

   go run ./

Build

   go build

Notes
- The main package expects `PORT`, `MONGODB_URI`, and `DB_NAME` environment variables to be set.
- Example environment variables are in `.env.example`.
