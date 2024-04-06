package model

import (
	"gorm.io/gorm"
	"time"
)

/* --- BASE MODEL CONFIGURATION --- */

type RabbitMQBaseModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:createdAt;type:timestamp"`
	UpdatedAt time.Time      `gorm:"column:updatedAt;type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt;index"`
}

func (m *RabbitMQBaseModel) BeforeCreate(tx *gorm.DB) error {
	if m.CreatedAt == (time.Time{}) {
		m.CreatedAt = time.Now()
	}

	if m.UpdatedAt == (time.Time{}) {
		m.UpdatedAt = time.Now()
	}

	return nil
}

func (m *RabbitMQBaseModel) BeforeSave(tx *gorm.DB) error {
	if m.CreatedAt == (time.Time{}) {
		m.CreatedAt = time.Now()
	}

	if m.UpdatedAt == (time.Time{}) {
		m.UpdatedAt = time.Now()
	}

	return nil
}

func (m *RabbitMQBaseModel) BeforeUpdate(tx *gorm.DB) error {
	if m.UpdatedAt == (time.Time{}) {
		m.UpdatedAt = time.Now()
	}

	return nil
}
