package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"messageServe/model"
	"strconv"
)

type Token struct {
	Menber *menber
}
type menber struct {
}

func (*Token) key() []byte {
	return []byte("长夜漫漫！！")
}

func (*Token) Pack() {

}
func (t *Token) GetClaimsModel(c *gin.Context) *model.Token {
	claims, _ := c.Get("claims")
	return claims.(*model.Token)
}

func (t *Token) Verify(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &model.Token{}, func(token *jwt.Token) (interface{}, error) {
		return t.key(), nil
	})
}

func (t *Token) Create(userConfig *model.UserConfig) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Token{
		UserId:         userConfig.UserId,
		UserName:       userConfig.UserName,
		StatusId:       strconv.Itoa(userConfig.Status),
		StandardClaims: jwt.StandardClaims{},
	})
	return claims.SignedString(t.key())
}
func (t *Token) Issue(userConfig *model.UserConfig) string {
	token, err := t.Create(userConfig)
	if err != nil {
		return ""
	}
	return token
}
func TokenHande() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			Api.Response.Write(c, &model.Response{State: 400, Err: "非法访问！"})
			c.Abort()
			return
		}
		t, err := Api.Token.Verify(token)
		if err != nil {
			Api.Response.Write(c, &model.Response{State: 400, Err: "token验证失败！"})
			c.Abort()
			return
		}
		if !t.Valid {
			Api.Response.Write(c, &model.Response{State: 400, Err: "token验证失败！"})
			c.Abort()
			return
		}
		claims := t.Claims
		c.Set("claims", claims)
		c.Next()
	}
}
