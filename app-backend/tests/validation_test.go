package test

import (
	"testing"
	"willdo/internal/models"
	"willdo/internal/validator"

	"github.com/stretchr/testify/assert"
)

func TestEventInvalidDateReturnsErr(t *testing.T) {
	p := models.Event{
		Description: "abc",
		DueDate:     "this is not a valid date",
	}

	v := validator.New()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestValidEventDoesNotReturnErr(t *testing.T) {
	p := models.Event{
		Description: "abc",
		DueDate:     "2000-01-01T00:00",
	}

	v := validator.New()
	err := v.Validate(p)
	assert.Len(t, err, 0)
}
