/*
This code displays some shit
*/

package main
import(
	"fmt"
	 consul"github.com/hashicorp/consul/api"
)

func main(){
	addr:="172.20.0.2:8500"
	config := consul.DefaultConfig()
	config.Address=addr
	c,_:=consul.NewClient(config)
	fmt.Println(c)
}
