package models

// TableOptions represents a table_options
type TableOptions struct {
	TableName         string             `json:"table_name,omitempty"`
	SchemaName        string             `json:"schema_name,omitempty"`
	MigrationOrder    int                `json:"migration_order,omitempty"`
	IsTable           bool               `json:"is_table,omitempty"`
	Comment           string             `json:"comment,omitempty"`
	SoftDelete        bool               `json:"soft_delete,omitempty"`
	PartitionStrategy *PartitionStrategy `json:"partition_strategy,omitempty"`
	PartitionColumn   string             `json:"partition_column,omitempty"`
	AuditFields       bool               `json:"audit_fields,omitempty"`
	EnableRls         bool               `json:"enable_rls,omitempty"`
}
