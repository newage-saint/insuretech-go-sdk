package models

// ColumnOptions represents a column_options
type ColumnOptions struct {
	AutoIncrement     bool          `json:"auto_increment,omitempty"`
	DefaultValue      string        `json:"default_value,omitempty"`
	CheckConstraint   string        `json:"check_constraint,omitempty"`
	Comment           string        `json:"comment,omitempty"`
	ExcludeFromInsert bool          `json:"exclude_from_insert,omitempty"`
	IsJson            bool          `json:"is_json,omitempty"`
	ColumnName        string        `json:"column_name,omitempty"`
	NotNull           bool          `json:"not_null,omitempty"`
	Unique            bool          `json:"unique,omitempty"`
	PrimaryKey        bool          `json:"primary_key,omitempty"`
	ExcludeFromUpdate bool          `json:"exclude_from_update,omitempty"`
	SqlType           string        `json:"sql_type,omitempty"`
	ForeignKey        *ForeignKey   `json:"foreign_key,omitempty"`
	Encrypted         bool          `json:"encrypted,omitempty"`
	Index             *IndexOptions `json:"index,omitempty"`
}
