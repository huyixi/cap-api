package service

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Memo struct {
	ID        int64  `db:"id" json:"id"`
	Content   string `db:"content" json:"content"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

type MemoService struct {
	db *sqlx.DB
}

func NewMemoService(db *sqlx.DB) *MemoService {
	return &MemoService{db: db}
}

func (s *MemoService) Create(ctx context.Context, content string) (*Memo, error) {
	res, err := s.db.ExecContext(ctx,
		`INSERT INTO memos(content) VALUES (?)`,
		content,
	)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	return s.GetByID(ctx, id)
}

func (s *MemoService) GetByID(ctx context.Context, id int64) (*Memo, error) {
	var m Memo
	err := s.db.GetContext(ctx, &m,
		`SELECT id, content, created_at, updated_at FROM memos WHERE id = ?`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *MemoService) Update(ctx context.Context, id int64, content string) (*Memo, error) {
	_, err := s.db.ExecContext(ctx,
		`UPDATE memos SET content=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`,
		content, id,
	)
	if err != nil {
		return nil, err
	}
	return s.GetByID(ctx, id)
}

func (s *MemoService) Delete(ctx context.Context, id int64) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM memos WHERE id=?`, id)
	return err
}

func (s *MemoService) List(ctx context.Context, limit, offset int, q string) ([]Memo, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	memos := []Memo{}

	if q == "" {
		err := s.db.SelectContext(ctx, &memos,
			`SELECT id, content, created_at, updated_at
			 FROM memos
			 ORDER BY updated_at DESC
			 LIMIT ? OFFSET ?`,
			limit, offset,
		)
		return memos, err
	}

	like := "%" + q + "%"
	err := s.db.SelectContext(ctx, &memos,
		`SELECT id, content, created_at, updated_at
		 FROM memos
		 WHERE content LIKE ?
		 ORDER BY updated_at DESC
		 LIMIT ? OFFSET ?`,
		like, limit, offset,
	)
	return memos, err
}
