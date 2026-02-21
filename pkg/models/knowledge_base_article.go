package models

// KnowledgeBaseArticle represents a knowledge_base_article
type KnowledgeBaseArticle struct {
	Title        string      `json:"title"`
	TitleBn      string      `json:"title_bn,omitempty"`
	Slug         string      `json:"slug"`
	Tags         []string    `json:"tags,omitempty"`
	HelpfulCount int         `json:"helpful_count,omitempty"`
	AuditInfo    interface{} `json:"audit_info"`
	Id           string      `json:"id"`
	Category     string      `json:"category"`
	Content      string      `json:"content"`
	ContentBn    string      `json:"content_bn,omitempty"`
	IsPublished  bool        `json:"is_published,omitempty"`
	ViewCount    int         `json:"view_count,omitempty"`
}
