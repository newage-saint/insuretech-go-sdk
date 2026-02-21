package models

// KnowledgeBaseArticle represents a knowledge_base_article
type KnowledgeBaseArticle struct {
	Id           string      `json:"id"`
	Title        string      `json:"title"`
	Tags         []string    `json:"tags,omitempty"`
	ViewCount    int         `json:"view_count,omitempty"`
	HelpfulCount int         `json:"helpful_count,omitempty"`
	AuditInfo    interface{} `json:"audit_info"`
	TitleBn      string      `json:"title_bn,omitempty"`
	Slug         string      `json:"slug"`
	Category     string      `json:"category"`
	Content      string      `json:"content"`
	ContentBn    string      `json:"content_bn,omitempty"`
	IsPublished  bool        `json:"is_published,omitempty"`
}
