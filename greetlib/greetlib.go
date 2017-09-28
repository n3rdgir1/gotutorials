package greetlib

var greetings = map[string]string{
  "English": "Hello",
  "Spanish": "Hola" }

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