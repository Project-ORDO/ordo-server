package response

type UserResponse struct {
	Email     string `gorm:"unique"`
	Name      string
}
