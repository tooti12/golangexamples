package golangexamples                       //importing packages.
import	"github.com/ehteshamz/greetings"


func ConcatSlice(sliceToConcat []byte) string {      //function named as concatslice to concatinate the string.
	var temp string
	for  i:=0;i<len(sliceToConcat);i++ {

		temp += string(sliceToConcat[i])
		if i<len(sliceToConcat)-1{ 
			temp+=string ('-')
		}
	}

	return temp
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i := 0; i < len(sliceToEncrypt); i++ {
		sliceToEncrypt[i] = sliceToEncrypt[i] + byte(ceaserCount-48)
	}
}


func EZGreetings(name string) string{
	return greetings.PrintGreetings(name)
}
