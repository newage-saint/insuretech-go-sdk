package models

// KnowledgeBaseArticle represents a knowledge_base_article
type KnowledgeBaseArticle struct {
	HelpfulCount int         `json:"helpful_count,omitempty"`
	AuditInfo    interface{} `json:"audit_info"`
	TitleBn      string      `json:"title_bn,omitempty"`
	Slug         string      `json:"slug"`
	Category     string      `json:"category"`
	Content      string      `json:"content"`
	Id           string      `json:"id"`
	Title        string      `json:"title"`
	ContentBn    string      `json:"content_bn,omitempty"`
	Tags         []string    `json:"tags,omitempty"`
	IsPublished  bool        `json:"is_published,omitempty"`
	ViewCount    int         `json:"view_count,omitempty"`
}
