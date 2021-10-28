package redisclient

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"parsons.com/fds/goserver/globalhelpers"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "100.70.1.51:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var ctx = context.Background()

var requestsMade = make(map[string][]int)

func RPCRequest(client string, id int, method string, params interface{}) {
	// Get the nodeid
	servicesString := "services.handlerid." + method
	val, err := rdb.Get(ctx, servicesString).Result()
	if err != nil {
		panic(err)
	} else {

		// Construct the Redis RPC information and push it
		rpcInString := "rpcin." + val
		rpcParams := map[string]interface{}{"client": client, "id": id, "method": method, "params": params}
		// Marshal params to convert them from Go map back to json
		marshalledParams, _ := json.Marshal(rpcParams)
		rdb.RPush(ctx, rpcInString, string(marshalledParams))
		_, err = rdb.Expire(ctx, rpcInString, 180*time.Second).Result()
		if err != nil {
			panic(err)
		}
		// Check if this client has already been keyed in requestsMade

		// Add the id to the slice of ids mapped to the client
		requestsMade[client] = append(requestsMade[client], id)
	}
}

func RPCResponse(client string, id int) string {
	var response string
	// First check to see if the client exists as a key in requestsMade
	if val, ok := requestsMade[client]; ok {
		idFound, idx := globalhelpers.Contains(val, id)
		if idFound {
			rpcOutString := "rpcout." + client
			// Create a loop that calls LPop every half second for 180 sec until response length > 0
			for i := 0; i < 360; i++ {
				response = rdb.LPop(ctx, rpcOutString).Val()
				if len(response) > 0 {
					globalhelpers.Remove(val, idx)
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
		} else {
			response = "Not a past request"
		}
	}

	return response
}
