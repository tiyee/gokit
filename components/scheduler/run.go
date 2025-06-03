package scheduler

import (
	"fmt"
	"time"
)

func main() {
	// 创建调度器
	s := New()

	// 启动调度器
	s.Start()
	defer s.Stop()

	// 添加一个定点执行的任务
	err := s.RunAt("task1", "一次性任务", time.Now().Add(2*time.Second), func() {
		fmt.Println("定点任务执行:", time.Now().Format("15:04:05"))
	})
	if err != nil {
		fmt.Println("添加任务失败:", err)
	}

	// 添加一个循环执行的任务
	err = s.Every("task2", "循环任务", 3*time.Second, func() {
		fmt.Println("循环任务执行:", time.Now().Format("15:04:05"))
	})
	if err != nil {
		fmt.Println("添加任务失败:", err)
	}

	// 等待一段时间观察执行结果
	time.Sleep(15 * time.Second)

	// 移除循环任务
	err = s.Remove("task2")
	if err != nil {
		fmt.Println("移除任务失败:", err)
	}

	// 查看当前所有任务
	jobs := s.Jobs()
	fmt.Printf("当前任务数量: %d\n", len(jobs))
	for _, job := range jobs {
		fmt.Printf("任务: %s, 下次执行: %s\n", job.Name, job.NextRun.Format("15:04:05"))
	}

	time.Sleep(5 * time.Second)
}
