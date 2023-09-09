package togetherai

import (
	"encoding/json"
	"fmt"
)

type InferenceConfig struct {
	Model             string   `json:"model"`
	Prompt            string   `json:"prompt"`
	MaxTokens         int32    `json:"max_tokens"`
	Stop              *string  `json:"stop"`
	Temperature       *float32 `json:"temperature"`
	TopP              *float32 `json:"top_p"`
	TopK              *int32   `json:"top_k"`
	RepetitionPenalty *float32 `json:"repetition_penalty"`
	LogProbs          *int32   `json:"logprobs"`
}

type inferenceRequestBody struct {
	InferenceConfig
	StreamTokens bool `json:"stream_tokens"`
}

type InferenceResponseBody struct {
	Status     string   `json:"status"`
	Prompt     []string `json:"prompt"`
	Model      string   `json:"model"`
	ModelOwner string   `json:"model_owner"`
	Tags       struct {
	} `json:"tags"`
	NumReturns int `json:"num_returns"`
	Args       struct {
		Model             string   `json:"model"`
		Prompt            string   `json:"prompt"`
		MaxTokens         int      `json:"max_tokens"`
		Stop              []string `json:"stop"`
		Temperature       float64  `json:"temperature"`
		TopP              float64  `json:"top_p"`
		TopK              int      `json:"top_k"`
		RepetitionPenalty int      `json:"repetition_penalty"`
	} `json:"args"`
	Subjobs []interface{} `json:"subjobs"`
	Output  struct {
		Choices []struct {
			FinishReason string `json:"finish_reason"`
			Index        int    `json:"index"`
			Text         string `json:"text"`
		} `json:"choices"`
		RawComputeTime float64 `json:"raw_compute_time"`
		ResultType     string  `json:"result_type"`
	} `json:"output"`
}

// NewInference creates a new inference using the provided configuration.
//
// config: The configuration for the inference.
// respBody: The response body containing the inference result.
// err: An error if the inference creation fails.
func (c *Client) NewInference(config InferenceConfig) (respBody *InferenceResponseBody, err error) {
	resp, err := c.getRestyClient().
		R().
		SetBody(inferenceRequestBody{
			InferenceConfig: config,
			StreamTokens:    false,
		}).
		Post("/inference")
	if err != nil {
		return
	}

	if resp.StatusCode() != 200 {
		err = fmt.Errorf("received status code %d with body %s", resp.StatusCode(), resp.Body())
		return
	}

	respBody = &InferenceResponseBody{}
	err = json.Unmarshal(resp.Body(), respBody)
	return
}
