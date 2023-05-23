package test

import (
	"testing"
)

func TestIsOk(t *testing.T) {
	if !isOk() {
		t.Error("发生错误了")
	}
}
