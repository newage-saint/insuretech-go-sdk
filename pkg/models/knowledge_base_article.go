package models

// KnowledgeBaseArticle represents a knowledge_base_article
type KnowledgeBaseArticle struct {
	Id           string     `json:"id"`
	TitleBn      string     `json:"title_bn,omitempty"`
	Slug         string     `json:"slug"`
	ContentBn    string     `json:"content_bn,omitempty"`
	IsPublished  bool       `json:"is_published,omitempty"`
	ViewCount    int        `json:"view_count,omitempty"`
	Title        string     `json:"title"`
	Category     string     `json:"category"`
	Content      string     `json:"content"`
	Tags         []string   `json:"tags,omitempty"`
	HelpfulCount int        `json:"helpful_count,omitempty"`
	AuditInfo    *AuditInfo `json:"audit_info,omitempty"`
}
