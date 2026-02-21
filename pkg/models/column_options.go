package models

// ColumnOptions represents a column_options
type ColumnOptions struct {
	Comment           string        `json:"comment,omitempty"`
	IsJson            bool          `json:"is_json,omitempty"`
	NotNull           bool          `json:"not_null,omitempty"`
	Unique            bool          `json:"unique,omitempty"`
	AutoIncrement     bool          `json:"auto_increment,omitempty"`
	DefaultValue      string        `json:"default_value,omitempty"`
	ForeignKey        *ForeignKey   `json:"foreign_key,omitempty"`
	ExcludeFromInsert bool          `json:"exclude_from_insert,omitempty"`
	ColumnName        string        `json:"column_name,omitempty"`
	SqlType           string        `json:"sql_type,omitempty"`
	CheckConstraint   string        `json:"check_constraint,omitempty"`
	ExcludeFromUpdate bool          `json:"exclude_from_update,omitempty"`
	Encrypted         bool          `json:"encrypted,omitempty"`
	PrimaryKey        bool          `json:"primary_key,omitempty"`
	Index             *IndexOptions `json:"index,omitempty"`
}
