package models

// TableOptions represents a table_options
type TableOptions struct {
	SoftDelete        bool               `json:"soft_delete,omitempty"`
	EnableRls         bool               `json:"enable_rls,omitempty"`
	PartitionStrategy *PartitionStrategy `json:"partition_strategy,omitempty"`
	MigrationOrder    int                `json:"migration_order,omitempty"`
	IsTable           bool               `json:"is_table,omitempty"`
	AuditFields       bool               `json:"audit_fields,omitempty"`
	PartitionColumn   string             `json:"partition_column,omitempty"`
	TableName         string             `json:"table_name,omitempty"`
	SchemaName        string             `json:"schema_name,omitempty"`
	Comment           string             `json:"comment,omitempty"`
}
