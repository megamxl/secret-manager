package persistence

// ManagedFiles represents a request with its associated paths
type ManagedFiles struct {
	ID       uint   `gorm:"primaryKey"`
	Filepath string `gorm:"uniqueIndex"`
	Request  string `gorm:"type:TEXT"`
}
