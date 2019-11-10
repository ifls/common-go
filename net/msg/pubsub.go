package msg

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"cloud.google.com/go/iam"
	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
)

func createTopic(client *pubsub.Client, topicName string) error {
	ctx := context.Background()
	// [START pubsub_create_topic]
	t, err := client.CreateTopic(ctx, topicName)
	if err != nil {
		return err
	}
	fmt.Printf("Topic created: %v\n", t)
	// [END pubsub_create_topic]
	return nil
}

func listTopic(client *pubsub.Client) ([]*pubsub.Topic, error) {
	ctx := context.Background()

	// [START pubsub_list_topics]
	var topics []*pubsub.Topic

	it := client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}

	return topics, nil
	// [END pubsub_list_topics]
}

func listSubscriptions(client *pubsub.Client, topicID string) ([]*pubsub.Subscription, error) {
	ctx := context.Background()

	// [START pubsub_list_topic_subscriptions]
	var subs []*pubsub.Subscription

	it := client.Topic(topicID).Subscriptions(ctx)
	for {
		sub, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}
	// [END pubsub_list_topic_subscriptions]
	return subs, nil
}

func deleteTopic(client *pubsub.Client, topic string) error {
	ctx := context.Background()
	// [START pubsub_delete_topic]
	t := client.Topic(topic)
	if err := t.Delete(ctx); err != nil {
		return err
	}
	fmt.Printf("Deleted topic: %v\n", t)
	// [END pubsub_delete_topic]
	return nil
}

func publishMessage(client *pubsub.Client, topic, msg string) error {
	ctx := context.Background()
	// [START pubsub_publish]
	// [START pubsub_quickstart_publisher]
	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	// [END pubsub_publish]
	// [END pubsub_quickstart_publisher]
	return nil
}

func publicMessageN(client *pubsub.Client, topic string, n int) error {
	ctx := context.Background()
	// [START pubsub_publish_with_error_handling_that_scales]
	var wg sync.WaitGroup
	var totalErrors uint64
	t := client.Topic(topic)

	for i := 0; i < n; i++ {
		result := t.Publish(ctx, &pubsub.Message{
			Data: []byte("Message " + strconv.Itoa(i)),
		})

		wg.Add(1)
		go func(i int, res *pubsub.PublishResult) {
			defer wg.Done()
			// The Get method blocks until a server-generated ID or
			// an error is returned for the published message.
			id, err := res.Get(ctx)
			if err != nil {
				// Error handling code can be added here.
				err := log.Output(1, fmt.Sprintf("Failed to publish: %v", err))
				if err != nil {
					log.Println(err)
				}
				atomic.AddUint64(&totalErrors, 1)
				return
			}
			fmt.Printf("Published message %d; msg ID: %v\n", i, id)
		}(i, result)
	}

	wg.Wait()

	if totalErrors > 0 {
		return errors.New(
			fmt.Sprintf("%d of %d messages did not publish successfully",
				totalErrors, n))
	}
	return nil
	// [END pubsub_publish_with_error_handling_that_scales]
}

func publishWithAttributes(client *pubsub.Client, topic string) error {
	ctx := context.Background()
	// [START pubsub_publish_custom_attributes]
	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello world!"),
		Attributes: map[string]string{
			"origin":   "golang",
			"username": "gcp",
		},
	})
	// Block until the result is returned and a server-generated ID is returned for the published message.
	// 阻塞直到结果返回,并返回一个PublishResult, 可以获得 messageID
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Published message with custom attributes; msg ID: %v\n", id)
	// [END pubsub_publish_custom_attributes]
	return nil
}

func publishWithSettings(client *pubsub.Client, topic string, msg []byte) error {
	ctx := context.Background()
	// [START pubsub_publisher_batch_settings]
	t := client.Topic(topic)
	t.PublishSettings = pubsub.PublishSettings{
		ByteThreshold:  5000,
		CountThreshold: 10,
		DelayThreshold: 100 * time.Millisecond,
	}
	result := t.Publish(ctx, &pubsub.Message{Data: msg})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	// [END pubsub_publisher_batch_settings]
	return nil
}

func publishSingleGoroutine(client *pubsub.Client, topic string, msg []byte) error {
	ctx := context.Background()
	// [START pubsub_publisher_concurrency_control]
	t := client.Topic(topic)
	t.PublishSettings = pubsub.PublishSettings{
		NumGoroutines: 1,
	}
	result := t.Publish(ctx, &pubsub.Message{Data: msg})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	// [END pubsub_publisher_concurrency_control]
	return nil
}

func getPubPolicy(c *pubsub.Client, topicName string) (*iam.Policy, error) {
	ctx := context.Background()

	// [START pubsub_get_topic_policy]
	policy, err := c.Topic(topicName).IAM().Policy(ctx)
	if err != nil {
		return nil, err
	}
	for _, role := range policy.Roles() {
		log.Print(policy.Members(role))
	}
	// [END pubsub_get_topic_policy]
	return policy, nil
}

func addPubUsers(c *pubsub.Client, topicName string) error {
	ctx := context.Background()

	// [START pubsub_set_topic_policy]
	topic := c.Topic(topicName)
	policy, err := topic.IAM().Policy(ctx)
	if err != nil {
		return err
	}
	// Other valid prefixes are "serviceAccount:", "proto:"
	// See the documentation for more values.
	policy.Add(iam.AllUsers, iam.Viewer)
	policy.Add("group:cloud-logs@google.com", iam.Editor)
	if err := topic.IAM().SetPolicy(ctx, policy); err != nil {
		log.Fatalf("SetPolicy: %v", err)
	}
	// NOTE: It may be necessary to retry this operation if IAM policies are
	// being modified concurrently. SetPolicy will return an error if the policy
	// was modified since it was retrieved.
	// [END pubsub_set_topic_policy]
	return nil
}

func testPubPermissions(c *pubsub.Client, topicName string) ([]string, error) {
	ctx := context.Background()

	// [START pubsub_test_topic_permissions]
	topic := c.Topic(topicName)
	perms, err := topic.IAM().TestPermissions(ctx, []string{
		"pubsub.topics.publish",
		"pubsub.topics.update",
	})
	if err != nil {
		return nil, err
	}
	for _, perm := range perms {
		log.Printf("Allowed: %v", perm)
	}
	// [END pubsub_test_topic_permissions]
	return perms, nil
}

const (
	PubsubClientVersion = "0.0.1"
)

//列出项目所有订阅
func list(client *pubsub.Client) ([]*pubsub.Subscription, error) {
	ctx := context.Background()
	// [START pubsub_list_subscriptions]
	var subs []*pubsub.Subscription
	it := client.Subscriptions(ctx)
	for {
		s, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	// [END pubsub_list_subscriptions]
	return subs, nil
}

func pullMsgs(client *pubsub.Client, subName string, topic *pubsub.Topic) error {
	ctx := context.Background()
	log.Println(topic)
	// 异步拉取 n条消息
	maxCount := 99999999999
	var mu sync.Mutex
	received := 0
	sub := client.Subscription(subName)
	cctx, cancel := context.WithCancel(ctx)
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		fmt.Printf(PubsubClientVersion+"->Got message: %q\n", string(msg.Data))
		mu.Lock()
		defer mu.Unlock()
		received++
		if received >= maxCount {
			cancel()
		}
	})
	return err
}

func pullMsgsError(client *pubsub.Client, subName string) error {
	ctx := context.Background()
	// [START pubsub_subscriber_error_listener]
	// If the service returns a non-retryable error, Receive returns that error after
	// all of the outstanding calls to the handler have returned.
	err := client.Subscription(subName).Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Got error message: %q\n", string(msg.Data))
		msg.Ack()
	})
	return err
}

func pullMsgsSettings(client *pubsub.Client, subName string) error {
	ctx := context.Background()
	// [START pubsub_subscriber_flow_settings]
	sub := client.Subscription(subName)
	sub.ReceiveSettings.MaxOutstandingMessages = 10
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
	})
	return err
}

// 创建订阅
func create(client *pubsub.Client, subName string, topic *pubsub.Topic) error {
	ctx := context.Background()
	// [START pubsub_create_pull_subscription]
	sub, err := client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created subscription: %v\n", sub)
	// [END pubsub_create_pull_subscription]
	return nil
}

func createWithEndpoint(client *pubsub.Client, subName string, topic *pubsub.Topic, endpoint string) error {
	ctx := context.Background()
	// [START pubsub_create_push_subscription]

	// For example, endpoint is "https://my-test-project.appspot.com/push". 消息应该被push到的url
	sub, err := client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10 * time.Second,
		PushConfig:  pubsub.PushConfig{Endpoint: endpoint},
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created subscription: %v\n", sub)
	// [END pubsub_create_push_subscription]
	return nil
}

func updateEndpoint(client *pubsub.Client, subName string, endpoint string) error {
	ctx := context.Background()
	// [START pubsub_update_push_configuration]

	// For example, endpoint is "https://my-test-project.appspot.com/push".
	subConfig, err := client.Subscription(subName).Update(ctx, pubsub.SubscriptionConfigToUpdate{
		PushConfig: &pubsub.PushConfig{Endpoint: endpoint},
	})
	if err != nil {
		return err
	}
	fmt.Printf("Updated subscription config: %#v", subConfig)
	// [END pubsub_update_push_configuration]
	return nil
}

// 删除订阅
func delete2(client *pubsub.Client, subName string) error {
	ctx := context.Background()
	// [START pubsub_delete_subscription]
	sub := client.Subscription(subName)
	if err := sub.Delete(ctx); err != nil {
		return err
	}
	fmt.Println("Subscription deleted.")
	// [END pubsub_delete_subscription]
	return nil
}

func createTopicIfNotExists(c *pubsub.Client) *pubsub.Topic {
	ctx := context.Background()

	const topic = "my-topic"
	// 获取 a topic to subscribe to.
	t := c.Topic(topic)
	ok, err := t.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return t
	}

	//不存在, 则创建
	t, err = c.CreateTopic(ctx, topic)
	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}
	return t
}

func getSubPolicy(c *pubsub.Client, subName string) (*iam.Policy, error) {
	ctx := context.Background()

	// [START pubsub_get_subscription_policy]
	policy, err := c.Subscription(subName).IAM().Policy(ctx)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	for _, role := range policy.Roles() {
		log.Printf("%q: %q", role, policy.Members(role))
	}
	fmt.Printf("policy : %v \n", *policy)
	// [END pubsub_get_subscription_policy]
	return policy, nil
}

func addSubUsers(c *pubsub.Client, subName string) error {
	ctx := context.Background()

	// [START pubsub_set_subscription_policy]
	sub := c.Subscription(subName)
	policy, err := sub.IAM().Policy(ctx)
	if err != nil {
		return err
	}
	// Other valid prefixes are "serviceAccount:", "proto:"
	// See the documentation for more values.
	policy.Add(iam.AllUsers, iam.Viewer)
	policy.Add("group:cloud-logs@google.com", iam.Editor)
	if err := sub.IAM().SetPolicy(ctx, policy); err != nil {
		return err
	}
	// NOTE: It may be necessary to retry this operation if IAM policies are
	// being modified concurrently. SetPolicy will return an error if the policy
	// was modified since it was retrieved.
	// [END pubsub_set_subscription_policy]
	return nil
}

func testSubPermissions(c *pubsub.Client, subName string) ([]string, error) {
	ctx := context.Background()

	// [START pubsub_test_subscription_permissions]
	sub := c.Subscription(subName)
	perms, err := sub.IAM().TestPermissions(ctx, []string{
		"pubsub.subscriptions.consume",
		"pubsub.subscriptions.update",
	})
	if err != nil {
		return nil, err
	}
	for _, perm := range perms {
		log.Printf("Allowed: %v", perm)
	}
	// [END pubsub_test_subscription_permissions]
	return perms, nil
}
