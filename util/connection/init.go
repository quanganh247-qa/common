package connection

import (
	"context"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/quanganh247-qa/common/db/sqlc"
	"github.com/quanganh247-qa/common/service/mail"
	"github.com/quanganh247-qa/common/service/nats"
	"github.com/quanganh247-qa/common/service/redis"
	"github.com/quanganh247-qa/common/service/token"
	"github.com/quanganh247-qa/common/service/worker"
	"github.com/quanganh247-qa/common/util"
)

type Connection struct {
	Close      func()
	NatsClient *nats.Client
	Store      db.Store
}

func Init(config util.Config) (*Connection, error) {
	// Initialize JWT token maker
	_, err := token.NewJWTMaker(config.SymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can't create token maker: %w", err)
	}

	// Initialize database connection pool
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %w", err)
	}
	store := db.NewStore(connPool)

	_ = asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}
	err = redis.InitRedis(config.RedisAddress)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to redis: %w", err)
	}

	DB := db.InitStore(connPool)
	go runTaskProcessor(&config, asynq.RedisClientOpt{Addr: config.RedisAddress}, DB)

	// nats
	nc, err := nats.NewNATsClient(config.NATs)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to nats: %w", err)
	}

	conn := &Connection{
		NatsClient: nc,
		Store:      store,
		Close: func() {
			connPool.Close()
			nc.Close()
		},
	}
	return conn, nil
}

func runTaskProcessor(config *util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	// Kiểm tra mailer
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	if mailer == nil {
		log.Fatal("Failed to create mailer")
	}

	// Khởi tạo task processor
	taskProcessor := worker.NewRedisTaskProccessor(redisOpt, store, mailer)
	if taskProcessor == nil {
		log.Fatal("Failed to create task processor")
	}

	// Bắt đầu task processor
	err := taskProcessor.Start()
	if err != nil {
		log.Fatalf("Failed to start task processor: %v", err)
	}
}
