package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/theankitbhardwaj/latest-wayback-snapshot-redis/cache"
	"github.com/theankitbhardwaj/latest-wayback-snapshot-redis/waybackapi"
)

type APIServer struct {
	addr  string
	cache cache.RedisClient
}

func NewAPIServer(listenAddr string, cache cache.RedisClient) *APIServer {
	return &APIServer{
		addr:  listenAddr,
		cache: cache,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	router.HandleFunc("GET /snapshot", s.handleGetSnapshot)

	log.Print("Server started")
	http.ListenAndServe(s.addr, router)
}

type Request struct {
	Url       string `json:"url"`
	Timestamp string `json:"timestamp"`
}

func (s *APIServer) handleGetSnapshot(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return
	}
	request := &Request{}
	json.Unmarshal(body, request)

	cacheKey := fmt.Sprintf("%v:%v", request.Url, request.Timestamp)
	cacheVal, err := s.cache.Get(cacheKey)

	var snapshotURL string

	if err == redis.Nil {
		log.Print("Cache miss")
		snapshotURL = waybackapi.GetSnapshotUrl(request.Url, request.Timestamp)
		s.cache.Setex(cacheKey, snapshotURL, 5*time.Second)
	} else if err != nil {
		panic(err)
	} else {
		log.Print("Cache Hit")
		snapshotURL = cacheVal
	}

	w.Write([]byte(snapshotURL))
}
