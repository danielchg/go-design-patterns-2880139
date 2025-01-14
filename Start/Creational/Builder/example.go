package main

import "fmt"

func main() {
	// TODO: Create a NotificationBuilder and use it to set properties
	var bldr = newNotificationBuilder()

	// TODO: Use the builder to set some properties
	bldr.SetTitle("New Notification")
	bldr.SetSubTitle("This is a subtitle")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.png")
	bldr.SetPriority(5)
	bldr.SetMessage("This is a new notification with some text in it")
	bldr.SetType("alert")

	// TODO: Use the Build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating notification:", err)
	} else {
		fmt.Printf("Notification created: %+v", notif)
	}
}
