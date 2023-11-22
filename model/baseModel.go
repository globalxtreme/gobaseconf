package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/* --- BASE MODEL CONFIGURATION --- */

type BaseModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:createdAt;type:timestamp"`
	UpdatedAt time.Time      `gorm:"column:updatedAt;type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt;index"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

/* --- COLUMN TYPE CONFIGURATION: OBJECT / MAP IN ARRAY --- */

type ArrayMapInterfaceColumn []map[string]interface{}

func (j *ArrayMapInterfaceColumn) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var result []map[string]interface{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

func (j ArrayMapInterfaceColumn) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.Marshal(j)
}

type MapInterfaceColumn map[string]interface{}

func (j *MapInterfaceColumn) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var result map[string]interface{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

func (j MapInterfaceColumn) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.Marshal(j)
}

type MapBoolColumn map[string]bool

func (j *MapBoolColumn) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var result map[string]bool
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

func (j MapBoolColumn) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.Marshal(j)
}

type ArrayStringColumn []string

func (j *ArrayStringColumn) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var result []string
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

func (j ArrayStringColumn) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.Marshal(j)
}

type ArrayIntColumn []int

func (j *ArrayIntColumn) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var result []int
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

func (j ArrayIntColumn) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.Marshal(j)
}

type ArrayUintColumn []uint

func (j *ArrayUintColumn) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var result []uint
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

func (j ArrayUintColumn) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.Marshal(j)
}

type ArrayBoolColumn []bool

func (j *ArrayBoolColumn) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var result []bool
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

func (j ArrayBoolColumn) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.Marshal(j)
}
