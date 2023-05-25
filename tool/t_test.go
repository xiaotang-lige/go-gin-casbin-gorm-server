package tool

import (
	"log"
	"messageServe/model"
	"testing"
)

func Test_casbin(t *testing.T) {

}

const name = "eyJbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjMiLCJ1c2VyTmFtZSI6IjIzIiwic3RhdHVzSWQiOjF9.G99yF3opGfUkT573Eg_AIFJ2s8mDUeWyFIW4e0uLbQo"

func Test_token(t *testing.T) {
	//Api.Token.Create()
	//token := &model.Token{}
	//tx, err := Api.Token.Verify(name, token.Key())
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(tx.Valid)
	//log.Println(tx.Header)
	//log.Println(tx.Claims)
	//log.Println(tx.Method)
	config := &model.UserConfig{
		UserId:   "123133",
		UserName: "111222",
		Password: "123",
		Status:   2,
	}
	token, err := Api.Token.Create(config)
	if err != nil {
		log.Println("创建token错误！", err)
	}

	m, err := Api.Token.UnPack(token)
	if err != nil {
		log.Println("获取token模型失败！", err)
	}
	log.Println(m.UserId)
	log.Println(m)
}
