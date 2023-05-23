package main

import (
	"context"
	"fmt"
	"time"

	gcrauthn "github.com/google/go-containerregistry/pkg/authn"
	gcrname "github.com/google/go-containerregistry/pkg/name"
	gcrremote "github.com/google/go-containerregistry/pkg/v1/remote"
)

func main() {
	const repo = "my_repo"
	const gcrToken = "token"

	ctx := context.Background()

	// gcpRepo, err := gcrname.NewRepository(repo)
	// if err != nil {
	// 	panic(err)
	// }

	// tags, err := gcrgoogle.List(gcpRepo)

	// keychain := gcrauthn.NewKeychainFromHelper(&gcrauthn.Bearer{Token: gcrToken})

	// keychain := gcrauthn.NewMultiKeychain(
	// 	// bearerkeychain.New(os.Getenv("GCRCLEANER_TOKEN")),
	// 	gcrauthn.DefaultKeychain,
	// 	gcrgoogle.Keychain,
	// )

	const image = "my_image"
	tag, err := gcrname.NewTag(fmt.Sprintf("%s/%s:latest", repo, image))
	if err != nil {
		panic(err)
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := gcrremote.Delete(tag,
		gcrremote.WithContext(timeoutCtx),
		gcrremote.WithAuth(&gcrauthn.Bearer{Token: gcrToken}),
	); err != nil {
		panic(err)
	}
}
