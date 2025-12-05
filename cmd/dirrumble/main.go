package main

import (
        "github.com/tools-tensorsecurityacademy/dirrumble/internal/fuzzer"
        "github.com/tools-tensorsecurityacademy/dirrumble/internal/models"
)

func main() {
        opts := models.ParseOptions()
        fuzzer.Run(opts)
}
