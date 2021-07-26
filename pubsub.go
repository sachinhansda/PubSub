package main

import "fmt"

//function to create a topic
func CreateTopic(topicID int) {

}

//function to delete a topic
func DeleteTopic(TopicID int) {

}

//function creates and adds subscription with id SubscriptionID to topic with id topicID
func AddSubscription(topicID,SubscriptionID int) {

}

//function to delete a subscription with id SubscriptionID
func DeleteSubscription(SubscriptionID int) {

}

//function to publish message on the topic with id topicID
func Publish(topicID int, message string) {

}

//function to subscribe a subscription with id SubscriptionID
func Subscribe(SubscriptionID int, SubscriberFunc) {

}

//function to unsubscribe a subscription with id SubscriptionID
func UnSubscribe(SubscriptionID int) {

}

//function to acknowledge that the message has been received and processed
func Ack(SubscriptionID, MessageID int) {

}

//main function
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
