## Notification System Design

Some of the functionality take care are

- Bulk Notification
- Preference & Rate Limitor
- Extensible

![Notification](../assets/notification.png)

The design is based on the client's input it could be bulk or simple notification.
The input can be from bulk service which take input which determine the type of bulk message and the type of notification.
The bulk services fetches users from a document database which contains different types of information of the users.
Another input to the notification service is directly from the client.

Once getting the input the notification service validates the message and determine the type and based on it sends it to message queue from where the notification processor service picks it up and does the rate limit check and also determine the other user prefences using the data from sql server which contains user information.

Then the message is send to different topic on the message queue which will be picked up by different handles which communicates with different notification types vendors.

### Notification Service On Cloud

The servers used can be aws ec2 or azure vm which can host the different services and message queue we can use azure service bus topics or aws sns to send specific type of message on specific topic. The dynamo/cosmo db can be used instead of cassandra (Column Db)