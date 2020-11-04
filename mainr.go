package main
import(
  	"log"
	"time"
	"fmt"
        "github.com/streadway/amqp"
        )
func main(){
        conn,_:= amqp.Dial("amqp://172.20.0.3:5672")
        ch1,_:= conn.Channel()
	cha:= "Hello"
        q1,_:= ch1.QueueDeclare(cha,false,false,false,false,nil)
	msgs,_:= ch1.Consume(q1.Name,"",true,false,false,false,nil)
	fmt.Println("Listening")
	for d:= range msgs {
		time.Sleep(500*time.Millisecond)
		log.Printf("Received: %s",d.Body)
		}


	}
