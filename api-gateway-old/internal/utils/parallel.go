package utils

import (
	"context"
	"sync"
	"time"
)

// Parallel provides utilities for executing operations in parallel
type Parallel struct{}

// NewParallel creates a new parallel utility
func NewParallel() *Parallel {
	return &Parallel{}
}

// Execute executes multiple functions in parallel and returns results and errors
func (p *Parallel) Execute(ctx context.Context, funcs []func() (interface{}, error)) ([]interface{}, []error) {
	var wg sync.WaitGroup
	results := make([]interface{}, len(funcs))
	errors := make([]error, len(funcs))

	for i, fn := range funcs {
		wg.Add(1)
		go func(index int, f func() (interface{}, error)) {
			defer wg.Done()
			
			// Check if context is cancelled
			select {
			case <-ctx.Done():
				results[index] = nil
				errors[index] = ctx.Err()
				return
			default:
			}
			
			result, err := f()
			results[index] = result
			errors[index] = err
		}(i, fn)
	}

	wg.Wait()
	return results, errors
}

// ExecuteWithTimeoutSeconds executes multiple functions in parallel with a timeout in seconds
func (p *Parallel) ExecuteWithTimeoutSeconds(ctx context.Context, timeoutSeconds int, funcs []func() (interface{}, error)) ([]interface{}, []error) {
	if timeoutSeconds <= 0 {
		return p.Execute(ctx, funcs)
	}
	
	// Convert seconds to time.Duration (seconds * time.Second)
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
	defer cancel() // Clean up the context
	
	var wg sync.WaitGroup
	results := make([]interface{}, len(funcs))
	errors := make([]error, len(funcs))

	for i, fn := range funcs {
		wg.Add(1)
		go func(index int, f func() (interface{}, error)) {
			defer wg.Done()
			
			// Check if context is cancelled
			select {
			case <-ctxWithTimeout.Done():
				results[index] = nil
				errors[index] = ctxWithTimeout.Err()
				return
			default:
			}
			
			result, err := f()
			results[index] = result
			errors[index] = err
		}(i, fn)
	}

	wg.Wait()
	return results, errors
}