```go
words.Words("?! I'm what I am")
    --> []string{"i'm", "what", "i", "am"}
```

```go
words.RemoveWords("Hello, my dear friend!", []string{"my"}, "!")
    --> "Hello, ! dear friend!"
```