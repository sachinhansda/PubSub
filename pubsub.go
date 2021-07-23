package main

import "fmt"

func CreateTopic(topicID int) {

}

func DeleteTopic(TopicID int) {

}

func AddSubscription(topicID,SubscriptionID int) {

}

func DeleteSubscription(SubscriptionID int) {

}

func Publish(topicID int, message string) {

}

func Subscribe(SubscriptionID int, SubscriberFunc) {

}

func UnSubscribe(SubscriptionID int) {

}

func Ack(SubscriptionID, MessageID int) {

}

func main() {
	CreateTopic(topicID)
  DeleteTopic(TopicID)
  AddSubscription(topicID,SubscriptionID)
  DeleteSubscription(SubscriptionID)
  Publish(topicID, message)
  Subscribe(SubscriptionID, SubscriberFunc)
  UnSubscribe(SubscriptionID)
  Ack(SubscriptionID, MessageID)
}
