package main
import(
	 "fmt"
	 "github.com/tooti12/golangexamples"
	)
func main() {
	slc := make([]byte,2)
	slc[0]=48
	slc[1]=49
	fmt.Println(golangexamples.ConcatSlice(slc))

	golangexamples.Encrypt(slc,2)

	fmt.Println(slc	)

}
