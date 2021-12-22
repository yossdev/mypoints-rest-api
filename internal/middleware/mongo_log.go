package middleware

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/configs"
	"github.com/yossdev/mypoints-rest-api/infrastuctures/db"
	"log"
	"time"
)

// custom middleware logger to mongodb atlas
type logReq struct {
	ReqId     uint64
	Timestamp time.Time
	RemoteIP  string
	Hostname  string
	Protocol  string
	Method    string
	Path      string
	Duration  string
}

type LogMethod interface {
	LogReqRes(c *fiber.Ctx) error
}

type logMongo struct {
	DB db.MongoDB
}

func NewLogMongo(m db.MongoDB) LogMethod {
	return &logMongo{
		DB: m,
	}
}

func (m *logMongo) LogReqRes(c *fiber.Ctx) error {
	id := c.Context().ID()
	timestamp := c.Context().ConnTime()
	ip := c.IP()
	hostname := c.Context().URI().Host()
	protocol := c.Protocol()
	method := c.Context().Method()
	path := c.Context().Path()
	duration := c.Context().Time()
	diff := duration.Sub(timestamp)

	payload := logReq{
		ReqId:     id,
		Timestamp: timestamp,
		RemoteIP:  ip,
		Hostname:  string(hostname),
		Protocol:  protocol,
		Method:    string(method),
		Path:      string(path),
		Duration:  fmt.Sprintf("%v", diff),
	}

	//save log to mongo db
	go func() {
		session := m.DB.DB().Database(configs.Get().MongodbName).Collection(configs.Get().MongodbCollection)
		_, err := session.InsertOne(context.TODO(), payload)
		if err != nil {
			log.Println("Failed to save logResReq to mongo, with err: ", err)
		} else {
			log.Println("Successfully to save logResReq to mongo")
		}
	}()

	return c.Next()
}
