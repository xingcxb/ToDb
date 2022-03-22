package main

import (
	"ToDb/communication"
	"context"
	"testing"
)

func TestPath(t *testing.T) {
	communication.LoadingHistory(context.Background())
}
