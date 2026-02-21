package models

// ColumnOptions represents a column_options
type ColumnOptions struct {
	Unique            bool          `json:"unique,omitempty"`
	Index             *IndexOptions `json:"index,omitempty"`
	DefaultValue      string        `json:"default_value,omitempty"`
	Encrypted         bool          `json:"encrypted,omitempty"`
	IsJson            bool          `json:"is_json,omitempty"`
	SqlType           string        `json:"sql_type,omitempty"`
	AutoIncrement     bool          `json:"auto_increment,omitempty"`
	ExcludeFromUpdate bool          `json:"exclude_from_update,omitempty"`
	PrimaryKey        bool          `json:"primary_key,omitempty"`
	CheckConstraint   string        `json:"check_constraint,omitempty"`
	ForeignKey        *ForeignKey   `json:"foreign_key,omitempty"`
	Comment           string        `json:"comment,omitempty"`
	ExcludeFromInsert bool          `json:"exclude_from_insert,omitempty"`
	ColumnName        string        `json:"column_name,omitempty"`
	NotNull           bool          `json:"not_null,omitempty"`
}
