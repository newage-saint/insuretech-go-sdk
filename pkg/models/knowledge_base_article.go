package models

// KnowledgeBaseArticle represents a knowledge_base_article
type KnowledgeBaseArticle struct {
	Tags         []string    `json:"tags,omitempty"`
	IsPublished  bool        `json:"is_published,omitempty"`
	ViewCount    int         `json:"view_count,omitempty"`
	HelpfulCount int         `json:"helpful_count,omitempty"`
	Id           string      `json:"id"`
	Title        string      `json:"title"`
	TitleBn      string      `json:"title_bn,omitempty"`
	Slug         string      `json:"slug"`
	ContentBn    string      `json:"content_bn,omitempty"`
	AuditInfo    interface{} `json:"audit_info"`
	Category     string      `json:"category"`
	Content      string      `json:"content"`
}
