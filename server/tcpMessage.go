package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"messageServe/linker"
	"messageServe/model"
	"messageServe/tcp"
	"net"
	"time"
)

var messageHandle = make(chan *model.Message, 1024)

// var hashMap = make([]map[string]*configData,1024)
// var online =make([]string,1024)
type configData struct {
	conn   net.Conn
	reader *bufio.Reader
	userId string
	ctx    context.Context
	cancel context.CancelFunc
}

func (server) messageHandle() {
	defer linker.TcpConn.Close()
	for {
		cofig := &configData{}
		cofig.ctx, cofig.cancel = context.WithCancel(context.Background())
		var err error
		cofig.conn, err = linker.TcpConn.Accept()
		if err != nil {
			fmt.Println("客户端对接失败:", err)
			continue
		}
		cofig.reader = bufio.NewReader(cofig.conn)
		var b bool
		cofig.userId, _, b = authentication(cofig.reader) //牵手连接校验
		if !b {
			log.Println("非法访问！")
			continue
		}
		//userId := "1234"
		log.Println(cofig.userId + "客户端对接完成！")
		go cofig.listen()
		go cofig.write()
	}
}
func authentication(reader *bufio.Reader) (string, int, bool) {
	megbyte, err := tcp.Decode(reader)
	if err != nil || err == io.EOF {
		log.Println(err)
		return "", 0, false
	}
	auten := &model.Authentication{}
	err = json.Unmarshal(megbyte, auten)
	if err != nil {
		log.Println(err)
		return "", 0, false
	}
	if auten.UserId == "dawang" || auten.UserId == "xiaowang" { //鉴权对比
		return auten.UserId, 1, true
	}
	return "", 0, false
}
func (config *configData) listen() {
	//reader := bufio.NewReader(conn)
	for {
		megbyte, err := tcp.Decode(config.reader)
		if err != nil || err == io.EOF {
			log.Println(config.userId+"监听连接已关闭:", err)
			config.conn.Close()
			config.cancel()
			return
		}
		msg := &model.Message{}
		json.Unmarshal(megbyte, msg)
		msg.UserId = config.userId
		msg.CreateTime = time.Now() //未设置标准格式
		//log.Println(msg) //-------------
		log.Println("监听到" + config.userId + "的消息：")
		//log.Println(msg)
		messageHandle <- msg
	}
}
func (config *configData) write() {
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func() {
	//	select {
	//	case <-config.ctx.Done():
	//		wg.Done()
	//		return
	//	}
	//}()
	//go func() {
	for {
		if config.ctx.Err() != nil {
			log.Println(config.userId + "写入连接已关闭")
			return
		}
		v, _ := redisBRPop(config.ctx, 1*time.Second, config.userId)
		if len(v) == 0 {
			continue
		}
		//待办：关闭客户端一瞬间的消息流失问题，加锁
		by, err := tcp.Encode(v[1])
		_, err = config.conn.Write(by)
		if err != nil {
			log.Println(config.userId+"写入连接已关闭:", err)
			config.cancel()
			//wg.Done()
			return
		}
	}
	//}()
	//wg.Wait()
}
func (server) handel() {
	for v := range messageHandle {
		go handelType(v)
	}
}
func handelType(msg *model.Message) {
	switch msg.MessageType {
	//文字通道
	case 1:
		redisLPush(context.Background(), msg)
	case 2:
	case 3:

	}
}
func redisLPush(ctx context.Context, mes *model.Message) error {
	ll, err := json.Marshal(mes)
	if err != nil {
		log.Println(err)
	}
	return linker.RedisDb.LPush(ctx, mes.Target, ll).Err()
}
func redisBRPop(ctx context.Context, time time.Duration, userId string) ([]string, error) {
	v, err := linker.RedisDb.BRPop(ctx, time, userId).Result()
	return v, err
}
