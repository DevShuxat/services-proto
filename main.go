package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	grpcserver "github.com/DevShuxat/eater-service/src/application/grpc"
	appsvc "github.com/DevShuxat/eater-service/src/application/services"
	eatersvc "github.com/DevShuxat/eater-service/src/domain/eater/services"
	"github.com/DevShuxat/eater-service/src/infrastructure/config"
	"github.com/DevShuxat/eater-service/src/infrastructure/jwt"
	eaterrepo "github.com/DevShuxat/eater-service/src/infrastructure/repositories/eater"
	"github.com/DevShuxat/eater-service/src/infrastructure/sms"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config,err := config.Load()
	if err != nil{
		panic(err)
	}

	logger, err := config.NewLogger()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?connect_timeout=%d&sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresDatabase,
		60,
	)
	db,err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	
	smsClient := sms.NewClient(config.SmsProvideApiKey)
	tokenSvc := jwt.NewService(config.JWTSecret, config.JWTExpiresInSec)
	eaterRepo := eaterrepo.NewEaterRepository(db)
	
	

	eaterSvc := eatersvc.NewEaterService(eaterRepo,smsClient,logger)

	eaterApp := appsvc.NewEaterApplicationService(eaterSvc,tokenSvc)
	
	root := gin.Default()

	root.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowCredentials: true,
	}))

	ctx,cancel := context.WithCancel(context.Background())
	g,ctx := errgroup.WithContext(ctx)

	osSignals := make(chan os.Signal,1)

	signal.Notify(osSignals,os.Interrupt,syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(osSignals)

	var httpServer *http.Server

	g.Go(func() error{
		httpServer = &http.Server{
			Addr: config.HttpPort,
			Handler: root,
		}

		logger.Debug("main: started http server", zap.String("port",config.HttpPort))
		if err := httpServer.ListenAndServe();err != http.ErrServerClosed{
			return err
		}
		return nil
	})



	var grpcServer *grpc.Server

	g.Go(func () error {
		server := grpcserver.NewServer(
			eaterApp,
		)
		grpcServer = grpc.NewServer()
		pb.RegisterEaterServiceServer(grpcServer,server)

		lis, err := net.Listen("tcp", config.GrpcPort)
		if err != nil {
			logger.Fatal("main: net.Listen", zap.Error(err))
		}

		defer lis.Close()
	
		logger.Debug("main: started grps server", zap.String("port",config.GrpcPort))

		return grpcServer.Serve(lis)
	})


	select {
	case <-osSignals:
		logger.Info("main: received os signal, shutting down")
		break
	case <-ctx.Done():
		break
	}

	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(),5*time.Second)
	defer shutdownCancel()

	if httpServer != nil {
		httpServer.Shutdown(shutdownCtx)
	}

	if grpcServer != nil {
		grpcServer.GracefulStop()
	}

	if err := g.Wait(); err != nil {
		logger.Error("main: server returned an error",zap.Error(err))
	}
}