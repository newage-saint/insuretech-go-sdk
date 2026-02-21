package models

// FAQ represents a faq
type FAQ struct {
	Question     string      `json:"question"`
	DisplayOrder int         `json:"display_order,omitempty"`
	AuditInfo    interface{} `json:"audit_info"`
	Category     string      `json:"category"`
	QuestionBn   string      `json:"question_bn,omitempty"`
	Answer       string      `json:"answer"`
	AnswerBn     string      `json:"answer_bn,omitempty"`
	IsPublished  bool        `json:"is_published,omitempty"`
	ViewCount    int         `json:"view_count,omitempty"`
	Id           string      `json:"id"`
}
