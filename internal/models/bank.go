package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bank struct {
	ID         uuid.UUID      `json:"bankId" gorm:"primaryKey"`
	Name       string         `json:"bankName"`
	IFSCCode   string         `json:"ifscCode"`
	BranchName string         `json:"branchName"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
