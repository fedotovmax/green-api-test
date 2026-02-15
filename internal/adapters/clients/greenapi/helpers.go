package greenapi

import "fmt"

func buildURL(base, instanceID, tokenApi string, method Methods) string {
	return fmt.Sprintf("%s/waInstance%s/%s/%s", base, instanceID, method, tokenApi)

}
