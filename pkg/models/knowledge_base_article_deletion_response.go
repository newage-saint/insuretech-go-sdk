package models

// KnowledgeBaseArticleDeletionResponse represents a knowledge_base_article_deletion_response
type KnowledgeBaseArticleDeletionResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
