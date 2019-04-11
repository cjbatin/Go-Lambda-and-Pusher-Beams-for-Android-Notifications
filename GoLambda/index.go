package main
import (
        "fmt"
        "context"
				"github.com/pusher/push-notifications-go"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        Title string `json:"title"`
				Message string `json:"message"`
}

const (
  instanceId = "0fb4a756-8780-4425-b393-3c7855470d99"
  secretKey  = "7A57DBC7C435D04039C301492C05F75F9AD65835D5DCF41F50FC901F68AA93C2"
)

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	beamsClient, _ := pushnotifications.New(instanceId, secretKey)

	publishRequest := map[string]interface{}{
      "fcm": map[string]interface{}{
        "notification": map[string]interface{}{
          "title": event.Title,
            "body":  event.Message,
          },
        },
	}

	pubId, err := beamsClient.PublishToInterests([]string{"hello"}, publishRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Publish Id:", pubId)
	}
	return fmt.Sprintf("Completed"), nil
}

func main() {
  lambda.Start(HandleRequest)
}
