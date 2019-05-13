package hello

import (
    "testing"
)

func TestGreeting(t *testing.T) {
    if Japanese() != "こんにちは、世界" {
        t.Errorf("Excepted %s, But %s.", "こんにちは、世界", Japanese())
    }
    if Chinese() != "你好,是世界" {
        t.Errorf("Excepted %s, But %s.", "你好,是世界", Chinese())
    }
    if English() != "Hello, world" {
        t.Errorf("Excepted %s, But %s.", "Hello, world", English())
    }
}
