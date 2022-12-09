package main

import (
	_ "go_blog/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"go_blog/internal/cmd"
	_ "go_blog/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
