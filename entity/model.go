package entity

type ModelsResponse struct {
	Object string   `json:"object"`
	Data   []Models `json:"data"`
}

type Models struct {
	ID         string             `json:"id"`
	Object     string             `json:"object"`
	OwnedBy    string             `json:"owned_by"`
	Permission []ModelsPermission `json:"permission"`
	Root       string             `json:"root"`
	Parent     any                `json:"parent"`
}

type ModelsPermission struct {
	ID                 string `json:"id"`
	Object             string `json:"object"`
	Created            int64  `json:"created"`
	AllowCreateEngine  bool   `json:"allow_create_engine"`
	AllowSampling      bool   `json:"allow_sampling"`
	AllowLogprobs      bool   `json:"allow_logprobs"`
	AllowSearchIndices bool   `json:"allow_search_indices"`
	AllowView          bool   `json:"allow_view"`
	AllowFineTuning    bool   `json:"allow_fine_tuning"`
	Organization       string `json:"organization"`
	Group              any    `json:"group"`
	IsBlocking         bool   `json:"is_blocking"`
}
