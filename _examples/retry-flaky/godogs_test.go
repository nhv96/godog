package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func Test_RetryFlaky(t *testing.T) {
	suite := godog.TestSuite{
		Name:                "retry flaky tests",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/retry.feature"},
			MaxRetry: 5,
		},
	}

	passed := 0

	expected := 5

	for i := 0; i < expected; i++ {
		if suite.Run() == 0 {
			passed++
		}
	}

	assert.Equal(t, expected, passed, fmt.Sprintf("expected the scenario to pass %d times, got %d", expected, passed))
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Step(`^a step randomly failed`, func(ctx context.Context) (context.Context, error) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		if r.Intn(2) != 1 {
			return ctx, errors.New("unexpected error")
		}
		return ctx, nil

	})

	sc.Step(`^godog retry the scenario`, func(ctx context.Context) (context.Context, error) {
		return ctx, nil
	})

	sc.Step(`^the scenario must pass`, func(ctx context.Context) (context.Context, error) {
		return ctx, nil
	})

	// sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	// 	fmt.Println(">> before scenario")
	// 	return ctx, nil
	// })

	// sc.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	// 	fmt.Println(">> after scenario")
	// 	return ctx, nil
	// })

	// sc.StepContext().Before(func(ctx context.Context, st *godog.Step) (context.Context, error) {
	// 	fmt.Println(">>>> before step")
	// 	return ctx, nil
	// })

	// sc.StepContext().After(func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
	// 	fmt.Println(">>>> after step")
	// 	return ctx, nil
	// })
}
