#!/bin/bash

echo "🔁 Sending test reviews..."

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"Отличный сервис!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"Ты дурак!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"nigger!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"gayakbfdkabk!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"abobaz!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"yooo nigga!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"Отличный сервис!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"Ты дурак!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"nigger!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"gayakbfdkabk!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"abobaz!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"content":"yooo nigga!"}'
echo ""

curl -s http://localhost:8080/moderation/history | jq
