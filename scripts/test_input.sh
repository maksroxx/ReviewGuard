#!/bin/bash

echo "🔁 Sending test reviews..."

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r1", "user_ip":"1.1.1.1", "content":"Отличный сервис!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r2", "user_ip":"1.1.1.1", "content":"Ты дурак!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r3", "user_ip":"1.1.1.1", "content":"nigger!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r4", "user_ip":"1.1.1.1", "content":"gayakbfdkabk!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r5", "user_ip":"1.1.1.1", "content":"abobaz!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r6", "user_ip":"1.1.1.1", "content":"yooo nigga!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r1", "user_ip":"2.2.2.2", "content":"Отличный сервис!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r2", "user_ip":"2.2.2.2", "content":"Ты дурак!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r3", "user_ip":"2.2.2.2", "content":"nigger!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r4", "user_ip":"2.2.2.2", "content":"gayakbfdkabk!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r5", "user_ip":"2.2.2.2", "content":"abobaz!"}'
echo ""

curl -s -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r6", "user_ip":"2.2.2.2", "content":"yooo nigga!"}'
echo ""

curl -s http://localhost:8080/moderation/history | jq
curl -s http://localhost:8080/moderation/history | jq
