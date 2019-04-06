package golangexamples

import "github.com/ehteshamz/greetings"

// helpEnc assists making Encrypt function generic
func helpEnc(sliceToEncrypt [] byte, ceaserCount int, eng [] byte) [] byte {

  for i:=0; i<len(sliceToEncrypt); i++{
    check:= 0
    for j:=0; j< len(eng)-ceaserCount;j++{
       if sliceToEncrypt[i] == eng[j]{
         sliceToEncrypt[i] = (eng[j+ceaserCount])
         check =1
         break
       }
    }

    if check ==0{
      for j:=len(eng)-ceaserCount; j<len(eng) ;j++{
          if sliceToEncrypt[i] == eng[j]{

          sliceToEncrypt[i] = (eng[(ceaserCount- (len(eng)-1-j))-1])
        }
      }
    }

  }
    return sliceToEncrypt

}


// ConcatSlice returns the content of the slice concatenated together
func ConcatSlice( sliceToConcat [] byte) string {

  var str string
  for i:=0; i<len(sliceToConcat)-1; i++{
    str+=string(sliceToConcat[i])+"-"
  }
  str+=string(sliceToConcat[len(sliceToConcat)-1])
  // fmt.Println(str)
  return str
}

// Encrypt changes original slice by encrypting its text in english characters form, using ceaser count
func Encrypt( sliceToEncrypt [] byte, ceaserCount int) [] byte {

  var eng = []byte {'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}
  sliceToEncrypt = helpEnc(sliceToEncrypt, ceaserCount, eng)
  var eng_cap = []byte {'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'}
  sliceToEncrypt = helpEnc(sliceToEncrypt, ceaserCount, eng_cap)
  return  sliceToEncrypt
}

// EZGreetings return greetings to you by using github.com/ehteshamz/greetings PrintGreetings function
func EZGreetings( name string) string {
  return greetings.PrintGreetings(name)

}
