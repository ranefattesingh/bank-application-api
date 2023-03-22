package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID                uuid.UUID      `json:"accountId" gorm:"primaryKey"`
	BankName          string         `json:"bankName"`
	BranchName        string         `json:"branchName"`
	AccoundHolderName string         `json:"accountHoldername"`
	IdenetiyID        string         `json:"identityName"`
	FirstName         string         `json:"firstName"`
	LastName          string         `json:"lastName"`
	Address           string         `json:"address"`
	CreatedAt         time.Time      `json:"-"`
	UpdatedAt         time.Time      `json:"-"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}
