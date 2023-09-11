package togetherai

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInference(t *testing.T) {
	client := NewClient(os.Getenv("TOGETHERAI_API_KEY"))

	stopStrings := []string{"*"}

	respBody, err := client.NewInference(InferenceConfig{
		Model:     "togethercomputer/RedPajama-INCITE-7B-Instruct",
		Prompt:    "The capital of France is",
		MaxTokens: 128,
		Stop:      &stopStrings,
	})
	assert.NoError(t, err)
	assert.NotNil(t, respBody)
	log.Println(respBody)

	assert.Equal(t, "finished", respBody.Status)
	assert.Equal(t, "language-model-inference", respBody.Output.ResultType)
}
