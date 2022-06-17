package service

import (
	v1 "arod-im/api/logic/v1"
	"context"
	"errors"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// var _ jwt.Claims = (*Claims)(nil)

// type Claims struct {
// 	Uid string `json:"uid"`
// 	jwt.StandardClaims
// }

// GenerateToken 生成token
func GenerateToken(uid string) (string, error) {
	// file, _ := exec.LookPath(os.Args[0])
	// path, _ := filepath.Abs(file)
	privateKeyByte, err := ioutil.ReadFile("../../internal/conf/private.key")
	if err != nil {
		return "", err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		return "", err
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"uid": uid,
		"iat": time.Now().Unix(),                    // Token颁发时间
		"nbf": time.Now().Unix(),                    // Token生效时间
		"exp": time.Now().Add(time.Hour * 6).Unix(), // Token过期时间，目前是6小时
		"iss": "arod-im",                            // 颁发者
		"sub": "AuthToken",                          // 主题
		// "role": uid,                                 // 角色（附加）
	})

	token, err := tokenClaims.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ParseToken 解析token
func ParseToken(token string) (string, error) {
	publicKeyByte, _ := ioutil.ReadFile("../../internal/conf/public.key")
	publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)

	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("验证token加密类型错误")
		}
		return publicKey, nil
	})

	if tokenClaims != nil {
		if Claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
			return Claims["sub"].(string), nil
		}
	}
	return "", err
}

func (s *MessageService) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReplay, error) {
	uid := req.Uid
	token, err := GenerateToken(uid)
	if err != nil {
		return &v1.LoginReplay{
			ActionStatus: "FAIL",
			ErrorInfo:    err.Error(),
			ErrorCode:    90001,
		}, err
	}
	return &v1.LoginReplay{
		ActionStatus: "OK",
		AccessToken:  token,
	}, nil
}
