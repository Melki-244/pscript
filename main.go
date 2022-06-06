package main

import (
	"fmt"
	"os"

	"github.com/bitfield/script"
)

func main() {
  argumentsOptions()
}
type commands = []string
func argumentsOptions()  {
  switch{
  case os.Args[1] == "sysrestart":
   sysrestart()
  case os.Args[1] == "teste":
   teste()
  default:
   fmt.Printf("Argumento Não Conhecido")
  }
}
func errorDescricao(o int, e error)  {
  if e != nil {
    fmt.Println("Ocorreu Um Erro:", e) 
    fmt.Println("Sáida:", o)
  } 
}
func execCommandsArgs(arg string, commands []string) {
  for _, command := range commands{
    out, err := script.Echo(arg).ExecForEach(command).Stdout()
    if err != nil {
      break
    }
    errorDescricao(out, err)
  }
}
func execCommands(commands []string)  {
  for _, command := range commands {
    out, err := script.Exec(command).Stdout()
    if err != nil {
      break
    }
    errorDescricao(out, err)
  }
}
func sysrestart()  {
  commands := []string{
    "sudo systemctl stop {{.}}",
    "sudo systemctl start {{.}}",
    "sudo systemctl status {{.}}",
  }
  execCommandsArgs(os.Args[2], commands)
}
func teste()  {
  /* Pegando Variáves do Ambiente */
  //home, _ := os.LookupEnv("HOME")
  //user, _ :=  os.LookupEnv("USER")
}
