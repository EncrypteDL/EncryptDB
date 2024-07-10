package driver

import (
	"math"
	"math/rand"
	"time"
)

// BackoffStrategy is an interface for implementing a delay before retrying the provided function with a new transaction.
type BackoffStrategy interface {
	// Get the time to delay before retrying, using an exponential function on the retry attempt, and jitter.
	Delay(retryAttempt int) time.Duration
}

// RetryPolicy defines the policy to use to for retrying the provided function in the case of a non-fatal error.
type RetryPolicy struct {
	// The maximum amount of times to retry.
	MaxRetryLimit int
	// The strategy to use for delaying before the retry attempt.
	Backoff BackoffStrategy
}

// ExponentialBackoffStrategy exponentially increases the delay per retry attempt given a base and a cap.
//
// This is the default strategy implementation.
type ExponentialBackoffStrategy struct {
	// The time in milliseconds to use as the exponent base for the delay calculation.
	SleepBase time.Duration
	// The maximum delay time in milliseconds.
	SleepCap time.Duration
}

// Delay gets the time to delay before retrying, using an exponential function on the retry attempt, and jitter.
func (s ExponentialBackoffStrategy) Delay(retryAttempt int) time.Duration {
	rand.Seed(time.Now().UTC().UnixNano())
	jitter := rand.Float64()*0.5 + 0.5

	return time.Duration(jitter*math.Min(float64(s.SleepCap.Milliseconds()), float64(s.SleepBase.Milliseconds())*math.Pow(2, float64(retryAttempt)))) * time.Millisecond
}
