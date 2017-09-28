package greetlib
//           key type -\/  \/- value type
var greetings = map[string]string{
  "English": "Hello",
  "Spanish": "Hola",
}

//                        (string, error) { // if you want to return an error
func GreetIn(lang string) string {
  if greeting, ok := greetings[lang]; ok {
    return greeting
  }
  return greetings["English"]
}

/*
  don't build, install ` $ go install `
  This will put the file in the pkg directory: ~/go/pkg/darwin_amd64/gotutorials/greetlib.a
*/
