package core

type Node interface {
	InsertNode()
	FindNode()
	DeleteNode()
	UpdateNode()

	// Init(address string, port int64, account string, passwd string)
	// Create() error
	// Update() error
	// Delete() error
	// Select(query interface{}, args ...interface{}) *gorm.DB
	// Where(query interface{}, args ...interface{}) *gorm.DB
	// First(out interface{}, where ...interface{}) *gorm.DB
	// Find(out interface{}, where ...interface{}) *gorm.DB
	// Connect()
}
