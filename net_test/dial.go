package main
import (
	"net"
	"io/ioutil"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	bs, _ := ioutil.ReadAll(conn)

	fmt.Print(string(bs))
}
