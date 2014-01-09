package main

import (
    "testing"
)


func TestHaveSubCommands(t *testing.T) {
    if len(subCommands) == 0 {
        t.Errorf("jobs is nil")
    }
}
