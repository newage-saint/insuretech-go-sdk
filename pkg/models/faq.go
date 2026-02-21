package models

// FAQ represents a faq
type FAQ struct {
	Id           string      `json:"id"`
	QuestionBn   string      `json:"question_bn,omitempty"`
	Answer       string      `json:"answer"`
	AnswerBn     string      `json:"answer_bn,omitempty"`
	DisplayOrder int         `json:"display_order,omitempty"`
	IsPublished  bool        `json:"is_published,omitempty"`
	Category     string      `json:"category"`
	Question     string      `json:"question"`
	ViewCount    int         `json:"view_count,omitempty"`
	AuditInfo    interface{} `json:"audit_info"`
}
