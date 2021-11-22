package main
// by haohao
import (
	"fmt"
	"os"
)

func main() {
  fmt.Println("fuck zhi!!! I'm haohao chen...")
	fmt.Println("Hello, Kubernetes！I'm from Jenkins CI！")
	fmt.Println("BRANCH_NAME:", os.Getenv("branch"))
}
