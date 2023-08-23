package main_test

import (
	"os/exec"
	"testing"
)

func TestInstructions(t *testing.T) {
	if err := exec.Command("weaver", "generate", ".").Run(); err != nil {
		t.Fatalf("weaver generate : %v", err)
	}

	out, err := exec.Command("go", "run", ".").Output()
	if err != nil {
		t.Fatalf("go run .: %v", err)
	}
	const want = "[🐖 🐗 🐷 🐽]\n"
	if string(out) != want {
		t.Fatalf("go run .: got %q, want %q", string(out), want)
	}
}
