package models

import "sync/atomic"

// DB Interface is our storage
// layer
type DB interface {
	GetScore() (int64, error)
	SetScore(int64) error
}

// NewDB returns our db struct that
// satisfies DB interface
func NewDB() DB {
	return &db{0}
}

type db struct {
	score int64
}

// GetScore returns the score atomically
func (d *db) GetScore() (int64, error) {
	return atomic.LoadInt64(&d.score), nil
}

// SetScore stores a new value atomically
func (d *db) SetScore(score int64) error {
	atomic.StoreInt64(&d.score, score)
	return nil
}
