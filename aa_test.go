package main

import (
	"ToDb/communication"
	"context"
	"testing"
)

func TestTt(t *testing.T) {
	aa := `{"alias":"11","hostURL":"11","port":"11","username":"11","password":"11","savePassword":true}`
	communication.Ok(context.Background(), aa)
}

func TestReadFile(t *testing.T) {
	//communication.LoadingBaseHistoryInfo()
	communication.LoadingHistoryInfo("这是redis2")
}
