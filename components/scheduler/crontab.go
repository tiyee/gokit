package scheduler

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Job 表示一个任务
type Job struct {
	ID       string
	Name     string
	Func     func()
	NextRun  time.Time
	Interval time.Duration // 0表示只执行一次
}

// Scheduler 调度器
type Scheduler struct {
	jobs    *Heap[*Job]
	mu      sync.RWMutex
	timer   *time.Timer
	ctx     context.Context
	cancel  context.CancelFunc
	running bool
	jobMap  map[string]*Job // 用于快速查找任务
}

// New 创建新的调度器
func New() *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())

	// 创建以NextRun时间为比较条件的最小堆
	jobHeap := NewHeap[*Job](func(i, j *Job) bool {
		return i.NextRun.Before(j.NextRun)
	})

	return &Scheduler{
		jobs:   jobHeap,
		ctx:    ctx,
		cancel: cancel,
		jobMap: make(map[string]*Job),
	}
}

// Start 启动调度器
func (s *Scheduler) Start() {
	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return
	}
	s.running = true
	s.mu.Unlock()

	go s.run()
}

// Stop 停止调度器
func (s *Scheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return
	}

	s.running = false
	s.cancel()
	if s.timer != nil {
		s.timer.Stop()
	}
}

// RunAt 在指定时间执行任务
func (s *Scheduler) RunAt(id, name string, t time.Time, fn func()) error {
	job := &Job{
		ID:      id,
		Name:    name,
		Func:    fn,
		NextRun: t,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查是否已存在相同ID的任务
	if _, exists := s.jobMap[id]; exists {
		return ErrJobExists
	}

	s.jobs.Push(job)
	s.jobMap[id] = job
	s.updateTimer()

	return nil
}

// Every 按固定间隔循环执行任务
func (s *Scheduler) Every(id, name string, interval time.Duration, fn func()) error {
	// 最低一分钟
	if interval <= time.Second {
		return ErrInvalidInterval
	}

	job := &Job{
		ID:       id,
		Name:     name,
		Func:     fn,
		NextRun:  time.Now().Add(interval),
		Interval: interval,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查是否已存在相同ID的任务
	if _, exists := s.jobMap[id]; exists {
		return ErrJobExists
	}

	s.jobs.Push(job)
	s.jobMap[id] = job
	s.updateTimer()

	return nil
}

// Remove 移除指定任务
func (s *Scheduler) Remove(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.jobMap[id]
	if !exists {
		return ErrJobNotFound
	}

	// 在堆中找到该任务并移除
	index := s.jobs.FindIndex(func(j *Job) bool {
		return j.ID == id
	})

	if index >= 0 {
		s.jobs.Remove(index)
	}

	delete(s.jobMap, id)
	s.updateTimer()

	return nil
}

// Jobs 获取所有任务信息
func (s *Scheduler) Jobs() []JobInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]JobInfo, 0, len(s.jobMap))
	for _, job := range s.jobMap {
		result = append(result, JobInfo{
			ID:       job.ID,
			Name:     job.Name,
			NextRun:  job.NextRun,
			Interval: job.Interval,
		})
	}

	return result
}

// run 调度器主循环
func (s *Scheduler) run() {
	s.mu.Lock()
	s.updateTimer()
	s.mu.Unlock()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-s.timer.C:
			s.executeReadyJobs()
		}
	}
}

// executeReadyJobs 执行到期的任务
func (s *Scheduler) executeReadyJobs() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()

	for s.jobs.Len() > 0 {
		job := s.jobs.Top()

		// 如果最早的任务还没到执行时间，则退出
		if job.NextRun.After(now) {
			break
		}

		// 移除任务
		s.jobs.Pop()

		// 异步执行任务
		go func(j *Job) {
			defer func() {
				if r := recover(); r != nil {
					// 可以在这里添加日志记录
					fmt.Printf("任务 %s 执行出错: %v\n", j.Name, r)
				}
			}()
			j.Func()
		}(job)

		// 如果是循环任务，重新调度
		if job.Interval > 0 {
			job.NextRun = now.Add(job.Interval)
			s.jobs.Push(job)
		} else {
			// 一次性任务，从映射中删除
			delete(s.jobMap, job.ID)
		}
	}

	s.updateTimer()
}

// updateTimer 更新定时器
func (s *Scheduler) updateTimer() {
	if s.timer != nil {
		s.timer.Stop()
	}

	if s.jobs.Len() == 0 {
		// 没有任务时，设置一个很长的定时器
		s.timer = time.NewTimer(24 * time.Hour)
		return
	}

	nextJob := s.jobs.Top()
	duration := time.Until(nextJob.NextRun)
	if duration < 0 {
		duration = 0
	}

	s.timer = time.NewTimer(duration)
}

// JobInfo 任务信息
type JobInfo struct {
	ID       string
	Name     string
	NextRun  time.Time
	Interval time.Duration
}

// 错误定义
var (
	ErrJobExists       = fmt.Errorf("job with this ID already exists")
	ErrJobNotFound     = fmt.Errorf("job not found")
	ErrInvalidInterval = fmt.Errorf("interval must be positive")
)
