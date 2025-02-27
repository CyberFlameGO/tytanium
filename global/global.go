package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/vysiondev/tytanium/api"
)

// Configuration stores the current configuration for the server and should not be modified as it is not thread safe.
var Configuration *api.Configuration

// RedisClient holds the Redis client used to communicate with Redis databases.
var RedisClient *redis.Client

const (
	// Version is the current version of the server.
	Version = "1.3.1"
)
