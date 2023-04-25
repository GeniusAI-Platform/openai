package entity

import "github.com/GoFarsi/openai/models"

type FineTuneRequest struct {
	Model models.FineTunes `json:"model,omitempty"`
	// TrainingFile The ID of an uploaded file that contains training data
	TrainingFile string `json:"training_file" validate:"required"`
	// ValidationFile The ID of an uploaded file that contains validation data
	ValidationFile string `json:"validation_file,omitempty"`
	// Epochs The number of epochs to train the model for. An epoch refers to one full cycle through the training dataset
	Epochs int `json:"n_epochs,omitempty"`
	// BatchSize The batch size to use for training. The batch size is the number of training examples used to train a single forward and backward pass
	BatchSize int `json:"batch_size,omitempty"`
	// LearningRateMultiplier The learning rate multiplier to use for training. The fine-tuning learning rate is the original learning rate used for pretraining multiplied by this value
	//By default, the learning rate multiplier is the 0.05, 0.1, or 0.2 depending on final batch_size (larger learning rates tend to perform better with larger batch sizes). We recommend experimenting with values in the range 0.02 to 0.2 to see what produces the best results
	LearningRateMultiplier float32 `json:"learning_rate_multiplier,omitempty"`
	// PromptLossWeight The weight to use for loss on the prompt tokens. This controls how much the model tries to learn to generate the prompt (as compared to the completion which always has a weight of 1.0), and can add a stabilizing effect to training when completions are short
	PromptLossWeight float32 `json:"prompt_loss_weight,omitempty"`
	// ComputeClassificationMetrics The number of classes in a classification task
	ComputeClassificationMetrics bool `json:"compute_classification_metrics,omitempty"`
	// ClassificationClasses The number of classes in a classification task
	ClassificationClasses int `json:"classification_n_classes,omitempty"`
	// ClassificationPositiveClass The positive class in binary classification
	ClassificationPositiveClass string `json:"classification_positive_class,omitempty"`
	// ClassificationBetas If this is provided, we calculate F-beta scores at the specified beta values. The F-beta score is a generalization of F-1 score. This is only used for binary classification.
	//With a beta of 1 (i.e. the F-1 score), precision and recall are given the same weight. A larger beta score puts more weight on recall and less on precision. A smaller beta score puts more weight on precision and less on recall.
	ClassificationBetas []float32 `json:"classification_betas,omitempty"`
	// Suffix A string of up to 40 characters that will be added to your fine-tuned model name.
	Suffix string `json:"suffix,omitempty"`
}

type FineTuneResponse struct {
	ID                string              `json:"id"`
	Object            string              `json:"object"`
	Model             string              `json:"model"`
	CreatedAt         int64               `json:"created_at"`
	FineTuneEventList []FineTuneEvent     `json:"events,omitempty"`
	FineTunedModel    string              `json:"fine_tuned_model"`
	HyperParams       FineTuneHyperParams `json:"hyperparams"`
	OrganizationID    string              `json:"organization_id"`
	ResultFiles       []FileResponse      `json:"result_files"`
	Status            string              `json:"status"`
	ValidationFiles   []FileResponse      `json:"validation_files"`
	TrainingFiles     []FileResponse      `json:"training_files"`
	UpdatedAt         int64               `json:"updated_at"`
}

type FineTuneEvent struct {
	Object    string `json:"object"`
	CreatedAt int64  `json:"created_at"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

type FineTuneHyperParams struct {
	BatchSize              int     `json:"batch_size"`
	LearningRateMultiplier float64 `json:"learning_rate_multiplier"`
	Epochs                 int     `json:"n_epochs"`
	PromptLossWeight       float64 `json:"prompt_loss_weight"`
}

type FineTuneList struct {
	Object string             `json:"object"`
	Data   []FineTuneResponse `json:"data"`
}
type FineTuneEventList struct {
	Object string          `json:"object"`
	Data   []FineTuneEvent `json:"data"`
}

type FineTuneDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}
