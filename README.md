# together-ai-go

Unofficial API library for [together.ai](https://together.ai)

## Quickstart

### Create a new client

```go
client := togetherai.NewClient(os.Getenv("TOGETHERAI_API_KEY"))
```

### Infer with a model

```go
resp, err := client.NewInference(InferenceConfig{
    Model:     "togethercomputer/RedPajama-INCITE-7B-Instruct",
    Prompt:    "The capital of France is",
    MaxTokens: 128,
    Stop:      &stopString,
})
if err != nil {
    panic(err)
}

fmt.Println(resp.Output.Choices)
```
