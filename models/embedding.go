package models

type Embedding uint8

const (
	// TEXT_EMBEDDING_ADA_002 is a designed to replace the previous 16 first-generation embedding models at a fraction of the cost.
	TEXT_EMBEDDING_ADA_002 Embedding = iota + 1
	TEXT_SEARCH_ADA_DOC_001
)

func (e Embedding) String() string {
	switch e {
	case TEXT_EMBEDDING_ADA_002:
		return "text-embedding-ada-002"
	case TEXT_SEARCH_ADA_DOC_001:
		return "text-search-ada-doc-001"
	default:
		return ""
	}
}
