package msg

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestPub(t *testing.T) {
	ctx := context.Background()

	// 项目名
	proj := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if proj == "" {
		_, err := fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		if err != nil {
			t.Fatal(err)
		}
		os.Exit(1)
	}

	// 连接
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}

	// 获取所有topic
	fmt.Println("Listing all topics from the project:")
	topics, err := listTopic(client)
	if err != nil {
		log.Fatalf("Failed to list topics: %v", err)
	}
	for _, t := range topics {
		fmt.Println(t)
	}

	const topic = "proto"
	// Create a new topic called my-topic.
	// if err := createTopic(client, topic); err != nil {
	// proto.Fatalf("Failed to create a topic: %v", err)
	// }
	subs, err := listSubscriptions(client, topic)
	if err != nil {
		log.Fatalf("Failed to list topics: %v", err)
	}
	for _, t := range subs {
		fmt.Println(t)
	}
	// Publish a text message on the created topic.
	for {
		fmt.Printf("%v\n", time.Now())
		t := time.Now().Format("2006-01-02 15:04:05.000000")
		fmt.Printf("%v\n", t)
		if err := publishMessage(client, topic, "hello world! now is "+t); err != nil {
			log.Fatalf("Failed to publish: %v", err)
		}
		time.Sleep(time.Duration(5) * time.Second)
	}

	// Publish 10 messages with asynchronous error handling.
	// if err := publicMessageN(client, topic, 10); err != nil {
	// proto.Fatalf("Failed to publish: %v", err)
	// }

	// Delete the topic.
	// if err := deleteTopic(client, topic); err != nil {
	// proto.Fatalf("Failed to delete the topic: %v", err)
	// }
}

func TestSub(t *testing.T) {
	ctx := context.Background()
	proj := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if proj == "" {
		_, err := fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		if err != nil {
			t.Fatal(err)
		}
		os.Exit(1)
	}
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}

	// Print all the subscriptions in the project.
	fmt.Println("Listing all subscriptions from the project:")
	subs, err := list(client)
	if err != nil {
		log.Fatal(err)
	}
	for _, sub := range subs {
		fmt.Println(sub)
	}

	// t := createTopicIfNotExists(client)

	// Create a new subscription.
	// if err := create(client, sub, t); err != nil {
	// 	proto.Fatal(err)
	// }

	// Pull messages via the subscription.
	// if err := pullMsgs(client, sub, t); err != nil {
	// proto.Fatal(err)
	// }

	//getPolicy(client, sub)

	// Delete the subscription.
	// if err := delete(client, sub); err != nil {
	// 	proto.Fatal(err)
	// }
}
func TestGooglePubSub(t *testing.T) {
	err := createTopic(nil, "")
	if err != nil {
		t.Fatal(err)
	}
	_ = deleteTopic(nil, "")
	_ = publicMessageN(nil, "", 3)
	_ = publishWithAttributes(nil, "")
	_ = publishWithSettings(nil, "", []byte(""))
	_ = publishSingleGoroutine(nil, "", []byte(""))
	_, _ = getPubPolicy(nil, "")
	_ = addPubUsers(nil, "")
	_, _ = testPubPermissions(nil, "")
	_ = pullMsgs(nil, "", nil)
	_ = pullMsgsSettings(nil, "")
	_ = create(nil, "", nil)
	_ = createWithEndpoint(nil, "", nil, "")
	_ = updateEndpoint(nil, "", "")
	_ = delete2(nil, "")
	exists := createTopicIfNotExists(nil)
	log.Println(exists)
	_, _ = getSubPolicy(nil, "")
	_ = addSubUsers(nil, "")
	_, _ = testSubPermissions(nil, "")

	_ = pullMsgsError(nil, "")
}
