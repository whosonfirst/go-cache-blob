package blob

import (
	"context"
	"github.com/whosonfirst/go-cache"
	_ "gocloud.dev/blob/memblob"
	"testing"
)

func TestFSCache(t *testing.T) {

	ctx := context.Background()

	c, err := cache.NewCache(ctx, "mem://")

	if err != nil {
		t.Fatalf("Failed to create mem:// cache, %v", err)
	}

	opts := &testCacheOptions{}

	// This is defined in testing.go
	err = testCache(ctx, c, opts)

	if err != nil {
		t.Fatalf("Cache tests failed, %v", err)
	}
}
