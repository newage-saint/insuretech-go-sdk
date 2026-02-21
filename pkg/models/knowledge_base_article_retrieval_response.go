package models

// KnowledgeBaseArticleRetrievalResponse represents a knowledge_base_article_retrieval_response
type KnowledgeBaseArticleRetrievalResponse struct {
	Error   *Error                `json:"error,omitempty"`
	Article *KnowledgeBaseArticle `json:"article,omitempty"`
}
