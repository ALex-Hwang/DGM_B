package token

import (
    jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(expireDuration time.Duration) (string, error) {
    expire := time.Now().Add(expireDuration)
    // 将 uid，用户角色， 过期时间作为数据写入 token 中
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, util.LoginClaims{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expire.Unix(),
        },
    })
    // SecretKey 用于对用户数据进行签名，不能暴露
    return token.SignedString([]byte(util.SecretKey))
}

func (a *AuthFilter) Filter(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        tokenStr := req.Header.Get("Authorization")
        token, err := jwt.ParseWithClaims(tokenStr, &util.LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte(util.SecretKey), nil 
        }   
        if err != nil {
            httputil.Error(rw, errors.ErrUnauthorized)
            return
        }   

        if claims, ok := token.Claims.(*util.LoginClaims); ok && token.Valid {
            log.Infof("uid: %s, role: %v", claims.Uid, claims.Role)
        } else {
            httputil.Error(rw, errors.ErrUnauthorized)
            return
        }   
        next.ServeHTTP(rw, req)
    }   
}

func CheckToken(tokenStr string) (bool) {
	token, err := jwt.ParseWithClaims(tokenStr, &util.LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.SecretKey), nil 
	}   
	if err != nil {
		return false
	}   

	if claims, ok := token.Claims.(*util.LoginClaims); ok && token.Valid {
		return true
	} else {
		return false
	}   
}