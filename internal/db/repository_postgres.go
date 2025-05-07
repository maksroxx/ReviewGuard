package db

import (
	"context"

	"github.com/maksroxx/ReviewGuard/internal/models"
)

type PostgresRepository struct {
	db PostgreDB
}

func NewPostgresRepository(db PostgreDB) *PostgresRepository {
	return &PostgresRepository{db}
}

func (p *PostgresRepository) SaveReview(ctx context.Context, r *models.Review) error {
	query := `INSERT INTO reviews (id, user_ip, content, status, created_at)
              VALUES ($1, $2, $3, $4, $5)`
	_, err := p.db.Database.ExecContext(ctx, query, r.ID, r.UserIP, r.Content, r.Status, r.CreatedAt)
	return err
}

func (p *PostgresRepository) GetReviewsByIP(ctx context.Context, ip string) ([]models.Review, error) {
	query := `SELECT id, user_ip, content, status, created_at
              FROM reviews
              WHERE user_ip = $1
              ORDER BY created_at DESC`
	rows, err := p.db.Database.QueryContext(ctx, query, ip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		if err := rows.Scan(&r.ID, &r.UserIP, &r.Content, &r.Status, &r.CreatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	return reviews, nil
}

func (p *PostgresRepository) GetIPStats(ctx context.Context, period string) ([]models.IPStats, error) {
	var interval string
	switch period {
	case "hour":
		interval = "1 hour"
	case "day":
		interval = "1 day"
	default:
		interval = "1 hour"
	}

	query := `
		SELECT user_ip, COUNT(*) as count
		FROM reviews
		WHERE created_at >= NOW() - INTERVAL '` + interval + `'
		GROUP BY user_ip
		ORDER BY count DESC
	`

	rows, err := p.db.Database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []models.IPStats
	for rows.Next() {
		var stat models.IPStats
		if err := rows.Scan(&stat.IP, &stat.Count); err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}
	return stats, nil
}

func (p *PostgresRepository) GetReviewsByStatus(ctx context.Context, status string) ([]models.Review, error) {
	rows, err := p.db.Database.QueryContext(ctx, `
		SELECT id, content, user_ip, status, created_at
		FROM reviews
		WHERE status = $1
	`, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		err := rows.Scan(&r.ID, &r.Content, &r.UserIP, &r.Status, &r.CreatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	return reviews, nil
}

func (p *PostgresRepository) UpdateReviewStatus(ctx context.Context, id, status string) error {
	query := `UPDATE reviews SET status = $1 WHERE id = $2`
	_, err := p.db.Database.ExecContext(ctx, query, status, id)
	return err
}
