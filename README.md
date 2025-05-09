# 🛡️ ReviewGuard

**ReviewGuard** — микросервис для фильтрации и модерации пользовательских отзывов на лету.  
Автоматически помечает отзывы для модерации, если они содержат спам, ссылки или запрещённые слова.

---

## Возможности

- Фильтрация запрещённых слов и URL
- Rate limiting по IP с использованием Redis
- Очередь модерации через Redis Streams
- Долговременное хранение и отчёты в PostgreSQL
- REST API:  
  - `POST /review` — принять отзыв  
  - `GET /report/spam` — отзывы со статусом "moderation"  
  - `GET /report/by-ip?ip=...` — все отзывы по IP  
  - `GET /report/stats?period=hour|day` — статистика по IP
- Асинхронная обработка отзывов
- Статус `"pending"` обрабатывается фоновым воркером и становится `"approved"` или `"moderation"`

---

## Стек

| Компонент       | Используется           |
|-----------------|------------------------|
| Язык            | Go 1.22                |
| Web Framework   | Gin                    |
| Очередь / Cache | Redis 7 (Streams, TTL) |
| БД              | PostgreSQL 13          |
| Инфра           | Docker, Docker Compose |

---

## Быстрый старт

### 📁 Клонирование

```bash
git clone https://github.com/maksroxx/ReviewGuard.git
cd ReviewGuard
