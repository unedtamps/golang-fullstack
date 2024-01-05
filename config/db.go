package config

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	d "github.com/unedtamps/golang-fullstack/internal/repository"
)

type DBConfig struct {
	DBPass string
	DBUser string
	DBHost string
	DBPort uint16
	DBName string
	DBSSL  string
}

func GetDBEnv() DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	DBPort, _ := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 16)
	return DBConfig{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBPort: uint16(DBPort),
		DBName: os.Getenv("DB_NAME"),
		DBSSL:  os.Getenv("DB_SSL"),
	}
}

func NewConnection() (*pgxpool.Pool, func() string) {
	c := GetDBEnv()
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBPass,
		c.DBName,
		c.DBSSL,
	)
	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect :%v", err.Error())
		os.Exit(1)
	}
	err = dbPool.Ping(context.Background())
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("Success connect to Database")
	closeCon := func() string {
		dbPool.Close()
		return "connection pool closed"
	}
	return dbPool, closeCon
}

func NewRepo() (d.Querier, func() string) {
	dbPool, closeCon := NewConnection()
	db := d.New(dbPool)
	return db, closeCon
}
