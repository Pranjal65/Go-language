package main 


import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	//"os"
	"github.com/go-redis/redis/v8"
	"context"
	"strconv"
)

var ctx = context.Background()
var rdb *redis.Client


var counter=0;

func main(){
     fmt.Println("Heyyy welcome");
	r:=redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,

	})
	rdb =r;
	router := httprouter.New()
//    ping,err:=r.Ping(context.Background()).Result()
//    if(err!=nil){
// 	panic(err)
//    }
//    fmt.Println(ping);

//    err=r.Set(ctx,"name","Pranjal",0).Err()
//    if(err!=nil){
// 	fmt.Println("we are not able to set value");
// 	panic(err)
//    }
   
//    val,err:=r.Get(ctx,"name").Result();
//    if(err!=nil){
// 	  fmt.Println("unable to get value");
// 	  panic(err);
//    }
//    fmt.Printf("value retrived from redis: %v ",val);


	router.GET("/",func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		increment_redis_key(w,r,p);
	})
	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":5000", router))

}

func increment_redis_key(writer http.ResponseWriter,r *http.Request, p httprouter.Params) {
	fmt.Println("Increment");
    val, err := rdb.Get(ctx, "counter").Result()
	
	if err == redis.Nil {
		fmt.Println("no counter set")
		err := rdb.Set(ctx, "counter", 1, 0).Err()
		counter++
		if err != nil {
			panic(err)
		}

    } else if err != nil {
        panic(err)
    } else {
		counter,_ = strconv.Atoi(val)
		counter++
		err := rdb.Set(ctx, "counter", counter, 0).Err()

		if err != nil {
			panic(err)
		}
	}
	
	fmt.Fprint(writer, counter)
	fmt.Println("counter", counter)
}


// package main

// import (
//     "context"
//     "fmt"
//     "log"
//     "github.com/go-redis/redis/v8"
// )

// var ctx = context.Background()

// func main() {
//     // Connect to Redis server with authentication
//     rdb := redis.NewClient(&redis.Options{
//         Addr:     "localhost:6379",
//         Password: "", // Replace <your_password> with your actual Redis password
//         DB:       0,
//     })

//     // Ping the Redis server to test the connection
//     _, err := rdb.Ping(ctx).Result()
//     if err != nil {
//         log.Fatalf("Failed to connect to Redis: %v", err)
//     }

//     fmt.Println("Connected to Redis")

//     // Example usage: set a key-value pair
//     err = rdb.Set(ctx, "example_key", "example_value", 0).Err()
//     if err != nil {
//         log.Fatalf("Failed to set key-value pair: %v", err)
//     }

//     // Example usage: get the value for a key
//     val, err := rdb.Get(ctx, "example_key").Result()
//     if err != nil {
//         log.Fatalf("Failed to get value for key: %v", err)
//     }

//     fmt.Printf("Value for 'example_key': %s\n", val)

//     // Close the Redis client connection when done
//     err = rdb.Close()
//     if err != nil {
//         log.Fatalf("Failed to close Redis connection: %v", err)
//     }

//     fmt.Println("Disconnected from Redis")
// }

