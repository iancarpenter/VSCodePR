// go-hello-world/main_test.go
package main

import (
    "bytes"
    "os"
    "testing"
)

func TestMainOutput(t *testing.T) {
    // Save original stdout
    oldStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Call main
    main()

    // Restore stdout and read output
    w.Close()
    var buf bytes.Buffer
    _, err := buf.ReadFrom(r)
    os.Stdout = oldStdout
    if err != nil {
        t.Fatalf("Failed to read stdout: %v", err)
    }

    got := buf.String()
    want := "Hello, World!\n"
    if got != want {
        t.Errorf("Expected %q, got %q", want, got)
    }
}