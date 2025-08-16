package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zkfmapf123/lambda-pods/internal"
)

// @test 유저 생성 테스트
func Test_CreateUser(t *testing.T) {
	user := User{}
	err := user.CreateUser("test@test.com", "1234", "readonly")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, user.Email, "test@test.com")
	assert.NotEqual(t, user.Password, "1234")
	assert.True(t, user.IsReadonly())
	assert.False(t, user.IsDeveloper())
	assert.False(t, user.IsAdmin())
}

// @test 유저 생성 테스트 - 유효하지 않은 역할
func Test_CreateUserInvalidRole(t *testing.T) {
	user := User{}
	err := user.CreateUser("test@test.com", "1234", "invalid_role")
	if err != nil {
		assert.Equal(t, err, internal.ErrInvalidUserRole)
	}
}

// @test 유저 비밀번호 암호화 테스트
func Test_ComparePassword(t *testing.T) {
	password := "1234"
	user := User{}
	err := user.CreateUser("test@test.com", password, "readonly")
	if err != nil {
		t.Fatal(err)
	}

	isMatch := user.IsComparePassword(password)
	assert.True(t, isMatch)
}
