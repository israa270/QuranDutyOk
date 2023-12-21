package response

// SysCaptchaResponse captcha response
type SysCaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
	LoginStatus   string `json:"loginStatus,omitempty"`
}
