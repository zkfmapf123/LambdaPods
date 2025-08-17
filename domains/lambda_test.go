package domains

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// @test 람다 삭제 여부 확인
func Test_ValidateDeleteLambda(t *testing.T) {
	lambda_1 := Lambdas{
		DeletedAt: nil,
	}

	lambda_2 := Lambdas{
		DeletedAt: &time.Time{},
	}

	assert.False(t, lambda_1.IsDelete())
	assert.True(t, lambda_2.IsDelete())
}
