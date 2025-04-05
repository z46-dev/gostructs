package assert

import (
	"slices"

	"github.com/z46-dev/go-logger"
)

type AssertService struct {
	log *logger.Logger
}

func NewAssertService(log *logger.Logger) *AssertService {
	return &AssertService{log: log}
}

func (a *AssertService) AssertEqual(expected, actual any) {
	if expected == actual {
		a.log.Successf("Passed: expected %v, got %v", expected, actual)
		return
	}

	a.log.Errorf("Fail: expected %v, got %v", expected, actual)
}

func (a *AssertService) AssertNotEqual(expected, actual any) {
	if expected != actual {
		a.log.Successf("Passed: expected %v, got %v", expected, actual)
		return
	}

	a.log.Errorf("Fail: expected %v, got %v", expected, actual)
}

func (a *AssertService) AssertTrue(condition bool) {
	if condition {
		a.log.Success("Passed: condition is true")
		return
	}

	a.log.Error("Fail: condition is false")
}

func (a *AssertService) AssertFalse(condition bool) {
	if !condition {
		a.log.Success("Passed: condition is false")
		return
	}

	a.log.Error("Fail: condition is true")
}

func (a *AssertService) AssertNil(value any) {
	if value == nil {
		a.log.Success("Passed: value is nil")
		return
	}

	a.log.Errorf("Fail: expected nil, got %v", value)
}

func (a *AssertService) AssertNotNil(value any) {
	if value != nil {
		a.log.Success("Passed: value is not nil")
		return
	}

	a.log.Error("Fail: expected not nil, got nil")
}

func (a *AssertService) AssertContains(slice []any, value any) {
	if slices.Contains(slice, value) {
		a.log.Successf("Passed: slice contains %v", value)
		return
	}

	a.log.Errorf("Fail: slice does not contain %v", value)
}

func (a *AssertService) AssertNotContains(slice []any, value any) {
	if !slices.Contains(slice, value) {
		a.log.Successf("Passed: slice does not contain %v", value)
		return
	}

	a.log.Errorf("Fail: slice contains %v", value)
}

func (a *AssertService) AssertEqualSlice(expected, actual []any) {
	if slices.Equal(expected, actual) {
		a.log.Successf("Passed: expected %v, got %v", expected, actual)
		return
	}

	a.log.Errorf("Fail: expected %v, got %v", expected, actual)
}

func (a *AssertService) AssertNotEqualSlice(expected, actual []any) {
	if !slices.Equal(expected, actual) {
		a.log.Successf("Passed: expected %v, got %v", expected, actual)
		return
	}

	a.log.Errorf("Fail: expected %v, got %v", expected, actual)
}

func (a *AssertService) AssertEqualMap(expected, actual map[any]any) {
	if len(expected) != len(actual) {
		a.log.Errorf("Fail: expected %v, got %v", expected, actual)
		return
	}

	for k, v := range expected {
		if actual[k] != v {
			a.log.Errorf("Fail: expected %v, got %v", expected, actual)
			return
		}
	}

	a.log.Successf("Passed: expected %v, got %v", expected, actual)
}

func (a *AssertService) AssertNotEqualMap(expected, actual map[any]any) {
	if len(expected) == len(actual) {
		for k, v := range expected {
			if actual[k] == v {
				a.log.Errorf("Fail: expected %v, got %v", expected, actual)
				return
			}
		}
	}

	a.log.Successf("Passed: expected %v, got %v", expected, actual)
}
