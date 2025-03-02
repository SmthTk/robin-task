package model

const (
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"
)

type User struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Username string    `json:"username"`
	Password string    `json:"password,omitempty"`
	Role     string    `json:"role"`
	Tasks    []Task    `json:"tasks,omitempty" gorm:"foreignKey:user_id" `
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:user_id" `
}

func GetAllUserRoles() []string {
	return []string{ROLE_ADMIN, ROLE_USER}
}
