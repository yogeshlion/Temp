package main
import(
	"fmt"
	"github.com/streadway/amqp"
	)
func main(){
	conn,err:= amqp.Dial("amqp://172.20.0.3:5672")
	fmt.Println(err)
	ch1,_:= conn.Channel()
	cha:="Hello"
	body:="Data"
	q1,_:= ch1.QueueDeclare(cha,false,false,false,false,nil)
	er:= ch1.Publish("",q1.Name,false,false,amqp.Publishing{ContentType:"text/plain",Body:[]byte(body)})
	if er!=nil{
		fmt.Println(er)
	}
	fmt.Println("Sent")
}
