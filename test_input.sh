curl -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r1", "user_ip":"192.168.1.100", "content":"Это отличный сервис!"}'

curl -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r2", "user_ip":"192.168.1.100", "content":"Ты дурак!"}'

curl -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r3", "user_ip":"192.168.1.100", "content":"gay!"}'

curl -X POST http://localhost:8080/review \
  -H "Content-Type: application/json" \
  -d '{"id":"r4", "user_ip":"192.168.1.100", "content":"niggersaidhfsafh!"}'
