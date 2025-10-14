package main

import (
	"bookingapp/internal/infrastructure/db"
	"bookingapp/internal/infrastructure/db/models"
	mysqlrepo "bookingapp/internal/infrastructure/repository/mysql"
	userrepo "bookingapp/internal/infrastructure/repository/mysql/user"
	httpi "bookingapp/internal/interface/http"
	"bookingapp/internal/usecase"
	"log"
	"net/http"
	"os"
	"strconv"

	"gorm.io/gorm"
)

func main() {
	// ---- 環境変数から接続情報 ----
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnvInt("DB_PORT", 3306)
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASS", "password")
	name := getEnv("DB_NAME", "booking")

	gdb, err := db.Open(db.Config{
		User: user, Pass: pass, Host: host, Port: port, Name: name,
	})
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	if err := db.Ping(gdb); err != nil {
		log.Fatalf("ping db: %v", err)
	}
	if err := db.Migrate(gdb); err != nil {
		log.Fatalf("migrate: %v", err)
	}
	if err := seedIfEmpty(gdb); err != nil {
		log.Fatalf("seed: %v", err)
	}

	planRepo := mysqlrepo.NewPlanRepo(gdb)
	resvRepo := mysqlrepo.NewReservationRepo(gdb)
	userRepo := userrepo.NewUserRepo(gdb)

	reservationUC := &usecase.ReservationUsecase{Plans: planRepo, Resv: resvRepo, Users: userRepo}
	userUC := &usecase.UserUsecase{Users: userRepo}

	reservationHandler := &httpi.ReservationHandler{UC: reservationUC}
	userHandler := &httpi.UserHandler{UC: userUC}

	mux := http.NewServeMux()

	// 予約登録、予約一覧、予約取得、プラン検索、ユーザ登録API
	mux.HandleFunc("POST /reservations", reservationHandler.Create)
	mux.HandleFunc("GET /reservations", reservationHandler.List)
	mux.HandleFunc("GET /reservations/", reservationHandler.Get)
	mux.HandleFunc("GET /plans", reservationHandler.SearchPlans)
	mux.HandleFunc("POST /register", userHandler.Register)

	// ユーザ情報取得APIを追加
	mux.HandleFunc("GET /users/", userHandler.GetUser)

	addr := ":8080"
	log.Printf("listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func seedIfEmpty(gdb *gorm.DB) error {
	var count int64
	if err := gdb.Model(&models.PlanModel{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	seed := []models.PlanModel{
		{ID: 100, Name: "富士プレミアム", Keyword: "富士 山 静岡", Price: 12000},
		{ID: 175, Name: "サウスベーシック", Keyword: "サウス 南", Price: 8000},
		{ID: 200, Name: "北の宿", Keyword: "北海道 北", Price: 10000},
	}
	return gdb.Create(&seed).Error
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
func getEnvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}
