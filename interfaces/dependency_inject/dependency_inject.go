package dependency_inject_iface

/* Dependency injection is a technique in which an object receives other objects that it depends on.
These other objects are called dependencies.

In the example below, the SendNotification function is dependent on the Notifier interface.
The EmailNotifier struct implements the Notifier interface.
The SendNotification function can be called with any object that implements the Notifier interface.
This allows for easy testing and swapping of implementations.
*/
func Main() {
	notifier := EmailNotifier{
		Email: "bob@bob.bob",
	}
	SendNotification(notifier, "Hello!")
}

type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Notify(message string) error {
	return nil
}

func SendNotification(n Notifier, message string) error {
	return n.Notify(message)
}
