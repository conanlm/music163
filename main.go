package main

import "music163/api"

// func main() {
// 	// api.Getmusic163()
// 	// api.Getsong()

// 	// wg := sync.WaitGroup{}
// 	// wg.Add(5)
// 	// for i := 0; i < 5; i++ {
// 	// 	go func(i int) {
// 	// 		fmt.Println("i", i)
// 	// 		wg.Done()
// 	// 	}(i)
// 	// }
// 	// wg.Wait()

// 	t := time.Now()

// 	tasks := []Task{

// 		{Id: 0, f: func() { time.Sleep(2 * time.Second); fmt.Println(0) }},
// 		{Id: 1, f: func() { time.Sleep(time.Second); fmt.Println(1) }},
// 		{Id: 2, f: func() { fmt.Println(2) }},
// 	}
// 	pool := NewWorkerPool(tasks, 2)
// 	pool.Start()

// 	tasks = pool.Results()
// 	fmt.Printf("all tasks finished, timeElapsed: %f s\n", time.Now().Sub(t).Seconds())
// 	for _, task := range tasks {
// 		fmt.Printf("result of task %d is %v\n", task.Id, task.Err)
// 	}

// }

// type Task struct {
// 	Id  int
// 	Err error
// 	f   func()
// }

// func (task *Task) Do() {
// 	return
// }

// type WorkerPool struct {
// 	PoolSize    int
// 	tasksSize   int
// 	tasksChan   chan Task
// 	resultsChan chan Task
// 	Results     func() []Task
// }

// func NewWorkerPool(tasks []Task, size int) *WorkerPool {
// 	tasksChan := make(chan Task, len(tasks))
// 	resultsChan := make(chan Task, len(tasks))
// 	for _, task := range tasks {
// 		tasksChan <- task
// 	}
// 	close(tasksChan)
// 	pool := &WorkerPool{PoolSize: size, tasksSize: len(tasks), tasksChan: tasksChan, resultsChan: resultsChan}
// 	pool.Results = pool.results
// 	return pool
// }

// func (pool *WorkerPool) Start() {
// 	for i := 0; i < pool.PoolSize; i++ {
// 		go pool.worker()
// 	}
// }

// func (pool *WorkerPool) worker() {
// 	for task := range pool.tasksChan {
// 		task.Err = errors.New("error")
// 		pool.resultsChan <- task
// 	}
// }

// func (pool *WorkerPool) results() []Task {
// 	tasks := make([]Task, pool.tasksSize)
// 	for i := 0; i < pool.tasksSize; i++ {
// 		tasks[i] = <-pool.resultsChan
// 	}
// 	return tasks
// }

// var sum int32

// func myFunc(i interface{}) {
// 	n := i.(int32)
// 	atomic.AddInt32(&sum, n)
// 	fmt.Printf("run with %d\n", n)
// }

// func demoFunc() {
// 	time.Sleep(10 * time.Millisecond)
// 	fmt.Println("Hello World!")
// }

func main() {

	api.Mygoroutinepool()

}
