package models

// FAQ represents a faq
type FAQ struct {
	Id           string     `json:"id"`
	Category     string     `json:"category"`
	QuestionBn   string     `json:"question_bn,omitempty"`
	DisplayOrder int        `json:"display_order,omitempty"`
	IsPublished  bool       `json:"is_published,omitempty"`
	ViewCount    int        `json:"view_count,omitempty"`
	Question     string     `json:"question"`
	Answer       string     `json:"answer"`
	AnswerBn     string     `json:"answer_bn,omitempty"`
	AuditInfo    *AuditInfo `json:"audit_info,omitempty"`
}
