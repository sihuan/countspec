package models

import "gorm.io/gorm"

type Tarball struct {
	gorm.Model
	Filename    string
	UUID        string
	Description string
	Size        int64
	ConfigPath  string

	Benchmarks Strs `gorm:"type:text[]"`
	QemuTasks  []QemuTask

	// Tasks   []Task
	Deleted bool
}

func (t *Tarball) RealPath() string {
	// 查询 tarball uuid
	return "data/upload/tarball/" + t.UUID
}

func (t *Tarball) HasBenchmark(benchmark string) bool {
	for _, b := range t.Benchmarks {
		if b == benchmark {
			return true
		}
	}
	return false
}
