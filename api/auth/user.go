package auth

// type createUserRequest struct {
// 	Username string `json:"username" binding:"required,alphanum"`
// 	Password string `json:"password" binding:"required,min=6"`
// 	FullName string `json:"full_name" binding:"required"`
// 	Email    string `json:"email" binding:"required,email"`
// }

// type userResponse struct {
//     Username          string    `json:"username"`
//     FullName          string    `json:"full_name"`
//     Email             string    `json:"email"`
//     PasswordChangedAt time.Time `json:"password_changed_at"`
//     CreatedAt         time.Time `json:"created_at"`
// }

//TODO: password changes at

// type loginUserRequest struct {
// 	Name     string `json:"name"  binding:"required,alphanum"`
// 	Password string `json:"password" binding:"required,min=6"`
// }
// type loginUserResponse struct {
// 	AccessToken string                 `json:"access_token"`
// 	User        api.CreateUserResponse `json:"user"`
// }

// func NewLogin() {
// 	return &Login{
// 		s: &api.Server{},
// 	}
// }

// func LoginUser(ctx *gin.Context, s ) {
// 	var req loginUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// }
