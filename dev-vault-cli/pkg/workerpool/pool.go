package workerpool

import (
	"dev-vault-cli/models"
	"sync"
)

type Job struct {
	Snippet models.Snippet
}

type Result struct {
	SnippetID string
	Tags      []string
	Err       error
}

type WorkerPool struct {
	numWorkers int
	jobs       chan Job
	results    chan Result
	wg         sync.WaitGroup
	mu         sync.Mutex
	index      map[string][]string
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Job, 100),
		results:    make(chan Result, 100),
		index:      map[string][]string{},
	}
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for job := range wp.jobs {
		temp := Result{
			SnippetID: job.Snippet.Id,
			Tags:      job.Snippet.Tags,
		}
		wp.results <- temp
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

func (wp *WorkerPool) Submit(snippet models.Snippet) {
	wp.jobs <- Job{Snippet: snippet}
}

func (wp *WorkerPool) Close() {
	close(wp.jobs)
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.results)
}

func (wp *WorkerPool) BuildIndex() map[string][]string {
	for result := range wp.results {
		if result.Err != nil {
			continue
		}

		for _, tag := range result.Tags {
			wp.index[tag] = append(wp.index[tag], result.SnippetID)
		}
	}

	return wp.index
}

func IndexSnippets(snippets []models.Snippet, numWorkers int) map[string][]string {
	pool := NewWorkerPool(numWorkers)
	pool.Start()

	for _, snippet := range snippets {
		pool.Submit(snippet)
	}

	pool.Close()

	go pool.Wait()

	return pool.BuildIndex()
}
