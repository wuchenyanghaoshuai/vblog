package context

import (
	"context"
	"testing"
	"time"
)

func TestCurl(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := Curl(ctx, "https://httpbin.org/delay/1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPaymente(t *testing.T) {
	ctx := context.Background()
	err := Payment(ctx, "xxxx")
	if err != nil {
		t.Fatal(err)
	}
}
