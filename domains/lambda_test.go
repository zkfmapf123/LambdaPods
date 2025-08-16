package domains

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateDeleteLambda(t *testing.T) {
	lambda_1 := Lambda{
		DeletedAt: nil,
	}

	lambda_2 := Lambda{
		DeletedAt: &time.Time{},
	}

	assert.False(t, lambda_1.IsDelete())
	assert.True(t, lambda_2.IsDelete())
}
