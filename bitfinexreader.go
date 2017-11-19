
package bitfinexreader

import (
	"github.com/jaaapi/bitfinexreader/awsclient"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

func main() {
	svc := awsclient.CreateClient()
	client := bitfinex.NewClient().Auth("user", "pass")
	info, err := client.Account.Info()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(info)
	}
}