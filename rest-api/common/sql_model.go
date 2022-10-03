package common

import (
	"time"
)

type SQLModel struct {
	ID        int        `json:"-" gorm:"column:id"`
	FakeID    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (s *SQLModel) GenerateUID(dbType int) {
	uid := NewUID(uint32(s.ID), dbType, 1)
	s.FakeID = uid
}
