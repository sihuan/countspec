package models

import "gorm.io/gorm"

type QemuTask struct {
	gorm.Model
	Name string
	// TaskID    uint
	TarballID uint
	Benchmark string
	Type      string

	Error string

	Status   string
	Inscount int
	Path     string
	Cmd      string
	// Args     []string
	Stdin  string
	Stdout string
	Stderr string
}
