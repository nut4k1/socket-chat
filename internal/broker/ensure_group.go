package broker

import (
	"context"
	"log"
	"strings"
)

func EnsureGroup(stream string, group string, ctx context.Context) error {
	err := client.XGroupCreateMkStream(
		ctx,
		stream,
		group,
		"0-0",
	).Err()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		log.Fatal("client XGroupCreateMkStream error:", err)
		return err
	}

	log.Printf("EnsureGroup stream = %s and group = %s is ready\n", stream, group)
	return nil
}
