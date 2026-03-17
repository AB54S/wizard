/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"

	"github.com/ab54s/wizard/cmd"
)

func main() {
    if err := cmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
