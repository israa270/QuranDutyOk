package request




type LoginDTO struct {
	UserName     string `json:"username" binding:"required"`        // username  ,email
	Password  string `json:"password" binding:"required,min=8,max=32"` // password
	Captcha   string `json:"captcha,omitempty"`                        // code
	CaptchaId string `json:"captchaId,omitempty"`                      // codeID
}