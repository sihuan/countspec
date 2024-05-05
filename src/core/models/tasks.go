package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string
	Description string
	Completed   bool
	// TarballID   uint

	Benchmarks Strs `gorm:"type:text[]"`

	// QemuTasks []QemuTask

	// Type: test, build
	Type string
	// Status: pending, running, completed, failed
	Status string
}
