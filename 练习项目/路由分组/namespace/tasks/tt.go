package tasks

import (
	"context"
	"fmt"
	"log"

	"github.com/beego/beego/v2/server/task"
)

func CrondTask() {
	// 秒 分 时 日 月 周
	tk1 := task.NewTask("tk1", "0 12 * * * *", func(ctx context.Context) error { fmt.Println("tk1"); return nil })
	err := tk1.Run()
	if err != nil {
		log.Fatal(err)
	}

	task.AddTask("tk1", tk1)
	task.StartTask()
	defer task.StopTask()
}
