package http

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/javiercbk/jayoak/files"
	"github.com/javiercbk/jayoak/api/sound"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	// imports the postgres sql driver
	_ "github.com/lib/pq"
)

const cookieName = "joss"

// Config contains the configuration for the server
type Config struct {
	Address       string
	FilesFolder   string
	RedisAddress  string
	RedisPassword string
	RedisSecret   string
	DBName        string
	DBHost        string
	DBUser        string
	DBPass        string
}

// Serve http connections
func Serve(cnf Config, logger *log.Logger) error {
	store, err := redis.NewStore(10, "tcp", cnf.RedisAddress, cnf.RedisPassword, []byte(cnf.RedisSecret))
	if err != nil {
		return err
	}
	postgresOpts := fmt.Sprintf("dbname=%s host=%s user=%s password=%s", cnf.DBName, cnf.DBHost, cnf.DBUser, cnf.DBPass)
	db, err := sql.Open("postgres", postgresOpts)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	repository := files.NewRepository(cnf.FilesFolder)
	// set new validator
	binding.Validator = new(defaultValidator)
	router := gin.Default()
	apiRouter := router.Group("/api")
	apiRouter.Use(sessions.Sessions(cookieName, store))
	{
		soundHandlers := sound.NewHandlers(logger, db, repository)
		soundHandlers.Routes(apiRouter)
	}
	srv := newServer(router, cnf.Address)
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be catched, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
	return nil
}

func newServer(handler http.Handler, address string) *http.Server {
	// see https://blog.cloudflare.com/exposing-go-on-the-internet/
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},

		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients
			// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	return &http.Server{
		Addr:         address,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      handler,
	}
}
