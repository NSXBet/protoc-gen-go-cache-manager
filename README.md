<h1 align="center"><img src="./docs/logo-gcm.png" width="480" align="center" /></h1>

go-cache-manager is an extensible way of safely, concurrently, scalably and observably managing cached data.

## Features

- **Smart**: go-cache-manager uses the singleflight package to ensure that only one goroutine is fetching the data at a
  time. This is useful when you have multiple goroutines trying to fetch the same data at the same time. This avoids the
  issue of the [thundering herd](https://en.wikipedia.org/wiki/Thundering_herd_problem), that can easily DDoS your
  database.
- **Type-Safe**: go-cache-manager uses protobuf definitions to generate the cache management code. This ensures that the
  cache keys are always correct and that the stored cache data is always used correctly.
- **Binary**: since go-cache-manager stores the marshalled version of the protobuf definitions, it will use less space
  than your usual JSON storing cache.
- **Compatible**: go-cache-manager uses protobuf, and with that, the stored data is already backwards and forwards
  compatible, provided you follow the protobuf rules.
- **Layered**: go-cache-manager provides layers of cache. By default you'll be using an in-memory cache, as well as a Redis cache.
- **Configurable**: go-cache-manager allows options to be passed to the cache manager. This allows you to configure the
  cache manager to your needs with options such as Prometheus prefix, redis endpoint and skipping the in-memory-cache
  layer. For all the available options check the Cache Manager Options section below.

## Usage

First we'll install it with:

```shell
go get github.com/NSXBet/go-cache-manager
```

For this you'll need to configure protobuf files in your repository. The recommended way is to create a `proto` directory in the root of your repository and put all your protobuf configuration files there. 

Inside the proto folder you can create a namespace for your protobuf files. For example, if you have a service called
`user` you can create a `user` directory inside the `proto` directory and put all your protobuf files there.

You can also use a parent namespace like the name of your company, for example, if your company is called `acme` you can
create a `acme` directory inside the `proto` directory and put the `user` directory inside the `acme` directory.

For the contents of the proto files, you can refer to the section below `Creating the protobuf definitions for the
cache`.

### Configuring buf

After creating the protobuf files you need to configure the `buf.yaml` file in the root of your repository. Here is an example of a `buf.yaml` file:

```yaml
version: v1
breaking:
  use:
    - FILE
lint:
  allow_comment_ignores: true
  rpc_allow_google_protobuf_empty_requests: true
  use:
    - DEFAULT
  except:
    - PACKAGE_DIRECTORY_MATCH
```

Then you need to configure the `buf.gen.yaml` file in the root of your repository. This is where you'll configure go-cache-manager. Here is an example of a `buf.gen.yaml` file:

```yaml
version: v1
managed:
  enabled: true
  go_package_prefix:
    default: github.com/acme/user-svc/gen/go  # Here you can put whatever package path you want
plugins:
  - plugin: buf.build/protocolbuffers/go  # this plugin will be used to generate regular protobuf
    out: ../gen/go
    opt: paths=source_relative
  - plugin: github.com/NSXbet/go-cache-manager  # this plugin will be used to generate go-cache-manager code
    out: ../gen/go
    opt:
      - paths=source_relative
```

Then you can run the following command to generate the code:

```shell
cd ./protos && buf dep update && buf generate
```

### Creating the protobuf definitions for the cache

`go-cache-manager` uses your protobuf definitions to generate cache management code. You need to create a protobuf file with the cache configuration. 

Imagine that we have a service called `user` and we want to cache the user data by the user's id.

Here is an example of that protobuf file:

```protobuf
syntax = "proto3";

package nsx.testapp;

option go_package = "github.com/NSXBet/acme/gen/go/acme/usersvc";

// UserDetailsRequest contains the parameters that will be used to vary the cache with.
// In this case, the user_id will be used to vary the cache. Any other parameters
// you add to this message will be used to vary the cache.
// This message should be named after the cache method it is used in, so a cache method called
// UserDetails will have a UserDetailsRequest for varying the cache items.
// WARNING: Refrain from using complex types here like other messages, repeated fields, etc.
//          This should contain simple scalar types.
message UserDetailsRequest {
  string user_id = 1;
}

// UserDetailsResponse contains the response that will be cached.
// This is the actual data that will be cached. It can contain any protobuf fields you want like:
// - scalar types
// - messages
// - repeated fields
// - oneof fields
// - timestamp
// - etc.
// This message should be named after the cache method it is used in, so a cache method called
// UserDetails will have a UserDetailsResponse for caching the response.
message UserDetailsResponse {
  User user = 1;
}

message User {
  string user_id = 1;
  string name = 2;
  string email = 3;
}

// UserCache is the service that will be used to cache user details.
// This service should contain the cache methods you want to use.
// Each method should have a corresponding request and response message.
// The service MUST have a name ending in `Cache`. That's how go-cache-manager
// knows that it must generate code for this service.
service UserCache {
  // UserDetails returns the user details for the given user_id from the cache.
  // This will generate a `.GetUserDetails` method to either return cached data or refresh and return,
  // and a `.RefreshUserDetails` method to refresh the cache, independently of the cache state,
  // while also returning the data.
  rpc UserDetails(UserDetailsRequest) returns (UserDetailsResponse) {}

  // You can add as many cache methods as you want here.
  // Remember that they will all share the same cache configuration since a CacheManager is created
  // for each service definition and will contain all the cache methods defined in the service as `rpcs`.
}

// You can also add any number of different cache services here. They will all be generated in the same package.
// go-cache-manager is smart enough to generate the cache manager for each service separately.
// The service name will be used as a prefix for all cache entries.
```

### Using the generated cache managers

After generating the code you can use the generated cache managers to cache your data. Here is an example of how you can use the generated cache manager:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/NSXBet/acme/gen/go/acme/usersvc"
)

func main() {
	manager, err := usersvc.NewUserCacheManager(
		func(ctx context.Context, input *usersvc.UserDetailsRequest) (*usersvc.UserDetailsResponse, error) {
            // This is where you would query a database, a service, or anything else where you
            // need to get the source of truth data from.
            // This is simply an example of what you need to return as a result of this method.
			return &usersvc.UserDetailsResponse{
				User: &usersvc.User{
					UserId: input.UserId,
					Name:   "Test User",
					Email:  "test@user.com",
				},
			}, nil
		},
        // Here you can add as many different settings as you'd like.
        // In this example we will configure the cache to connect to redis,
        // and to use a prometheus prefix of `acme_user_svc`.
		gocachemanager.WithRedisConnection("localhost:6379"), 
        gocachemanager.WithPrometheusPrefix("acme_user_svc"),
	)

    // After this we can use the manager to get the user details.
    // The first time we call this method it will call the function we passed to the manager
    // and cache the result. The next time we call this method it will return the cached result.
    // If we want to refresh the cache we can call the `RefreshUserDetails` method.
    // This will call the function we passed to the manager and cache the result.
    // This is useful when you want to refresh the cache independently of the cache state.
    // 
    // The call below will generate a redis entry with the key `usercache::userdetails::CgEx` where the
    // last part is the base64 encoded version of the proto.Marshal of the UserDetailsRequest, and will vary with
    // the `UserId` passed to the `GetUserDetails` method.
    userDetails, err := manager.GetUserDetails(context.Background(), &usersvc.UserDetailsRequest{
        UserId: "123",
    })
    if err != nil {
        log.Fatalf("error getting user details: %v", err)
    }

    fmt.Printf("User details: %+v\n", userDetails)
}
```

## Cache Manager Options

### WithRedisConnection

This option allows you to configure the cache manager to use a Redis cache. This option takes a single string parameter
which is the Redis endpoint. Here is an example of how you can use this option:

```go
manager, err := usersvc.NewUserCacheManager(
    func(ctx context.Context, input *usersvc.UserDetailsRequest) (*usersvc.UserDetailsResponse, error) {
        return &usersvc.UserDetailsResponse{
            User: &usersvc.User{
                UserId: input.UserId,
                Name:   "Test User",
                Email:  "
            },
        }, nil
    },
    gocachemanager.WithRedisConnection("localhost:6379"),
)
```

### WithPrometheusPrefix

This option allows you to configure the cache manager to use a Prometheus prefix. This option takes a single string
parameter which is the Prometheus prefix that will be used for all published metrics. Here is an example of how you can use this option:

```go
manager, err := usersvc.NewUserCacheManager(
    func(ctx context.Context, input *usersvc.UserDetailsRequest) (*usersvc.UserDetailsResponse, error) {
        return &usersvc.UserDetailsResponse{
            User: &usersvc.User{
                UserId: input.UserId,
                Name:   "Test User",
                Email:  "
            },
        }, nil
    },
    gocachemanager.WithPrometheusPrefix("acme_user_svc"),
)
```

### WithSkipInMemoryCache

This option allows you to configure the cache manager to skip the in-memory cache. This option takes no parameters. Here is an example of how you can use this option:

```go
manager, err := usersvc.NewUserCacheManager(
    func(ctx context.Context, input *usersvc.UserDetailsRequest) (*usersvc.UserDetailsResponse, error) {
        return &usersvc.UserDetailsResponse{
            User: &usersvc.User{
                UserId: input.UserId,
                Name:   "Test User",
                Email:  "
            },
        }, nil
    },
    gocachemanager.WithSkipInMemoryCache(),
)
```

### WithInMemoryCacheSize

This option allows you to configure the cache manager to use a specific size for the in-memory cache. This option takes a single int64 parameter which is the size of the in-memory cache. Here is an example of how you can use this option:

```go
manager, err := usersvc.NewUserCacheManager(
    func(ctx context.Context, input *usersvc.UserDetailsRequest) (*usersvc.UserDetailsResponse, error) {
        return &usersvc.UserDetailsResponse{
            User: &usersvc.User{
                UserId: input.UserId,
                Name:   "Test User",
                Email:  "
            },
        }, nil
    },
    gocachemanager.WithInMemoryCacheSize(512_000_000), // 512MB
)
```
