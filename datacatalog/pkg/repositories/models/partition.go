package models

// Main Use cases:
// 1. Filter artifacts by partition key/val in a dataset from UI [x]
// 2. Get the artifact that has the partitions (x,y,z + tag_name) = latest [x]
type Partition struct {
	BaseModel
	DatasetUUID string `gorm:"primary_key"`
	KeyName     string `gorm:"primary_key"`
	KeyValue    string `gorm:"primary_key"`
	ArtifactID  string `gorm:"index"` // for JOINs with the Tag/Labels table when querying artifacts
}
