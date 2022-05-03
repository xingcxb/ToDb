// package os
// @Author: easymbol
// @Date: 2022/5/3
// @LastEditors: easymbol
// @LastEditTime: 2022/5/3 22:23
// @FilePath:
// @Description: TODO

// Copyright (c) 2022 by easymbol, All Rights Reserved.

package os

import (
	"context"
	"testing"
)

func TestCreateFile(t *testing.T) {
	File().CreateFile(context.Background(), "test.txt", true)
}
