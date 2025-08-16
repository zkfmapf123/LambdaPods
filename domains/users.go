package domains

import (
	"time"

	"github.com/zkfmapf123/lambda-pods/internal"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"unique;not null;size:100" json:"email"`
	Password  string    `gorm:"not null;size:255" json:"-"`
	Role      string    `gorm:"not null;default:readonly;size:20;check;role IN ('readonly', 'developer', 'admin)" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// 유저테이블 이름 반환
func (u User) TableName() string {
	return "users"
}

// 유저 역할 반환
func (u User) GetRole() string {
	return u.Role
}

// 유저 역할 확인 - Readonly
func (u User) IsReadonly() bool {
	return u.Role == "readonly"
}

// 유저 역할 확인 - Developer
func (u User) IsDeveloper() bool {
	return u.Role == "developer"
}

// 유저 역할 확인 - Admin
func (u User) IsAdmin() bool {
	return u.Role == "admin"
}

// 유저 역할 모두 반환 (Readonly, Developer, Admin)
func (u User) GetAllRoles() []string {
	return []string{"readonly", "developer", "admin"}
}

// 유저 역할 확인 - 유효성 검사 (Readonly, Developer, Admin)
func (u User) IsValidRole(role string) bool {
	return internal.MatchStringEquals(role, u.GetAllRoles()...)
}

// 유저 생성
func (u *User) CreateUser(email, password, role string) error {
	hasPassword, err := u.GetCryptPassword(password)
	if err != nil {
		return err
	}

	u.Email = email
	u.Password = hasPassword
	u.Role = role

	if !u.IsValidRole(role) {
		return internal.ErrInvalidUserRole
	}

	return nil
}

// 유저 비밀번호 암호화
func (u User) GetCryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// 유저 비밀번호 비교
func (u User) IsComparePassword(inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPassword))
	return err == nil
}
