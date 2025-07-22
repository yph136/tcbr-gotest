package cloudbase

import "fmt"

// 参考腾讯云开发 OPENAPI 文档 https://docs.cloudbase.net/http-api/auth/auth-sign-in

// SignInReq
type SignInReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignInRes
type SignInRes struct {
	TokenType    string   `json:"token_type"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	Sub          string   `json:"sub"`
	Groups       []string `json:"groups"`
}

// SignIn
func (c *Client) SignIn() (string, error) {
	req := &SignInReq{
		Username: c.Username,
		Password: c.Password,
	}

	res := &SignInRes{}

	url := fmt.Sprintf("https://%s.%s/auth/v1/signin", c.EnvId, domain)
	err := httpPost(url, nil, req, res)
	if err != nil {
		return "", err
	}
	if res.AccessToken != "" {
		return fmt.Sprintf("%s %s", res.TokenType, res.AccessToken), nil
	}
	return "", fmt.Errorf("Get AccessToken Failed")
}
