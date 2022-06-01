package main

import (
	"fmt"
	"os"

	"github.com/bitfield/script"
)

func main() {
  argumentsOptions()
}
func argumentsOptions()  {
  switch{
  case os.Args[1] == "systemctl":
   instalaSystemctl()
  case os.Args[1] == "teste" : 
   teste()
  case os.Args[1] == "restart" :
   sysrestart(os.Args[2])
  default:
   fmt.Printf("Argumento Não Conhecido")
  }
}
func errorDescribe(o int, e error)  {
  if e != nil {
    fmt.Println("Ocorreu Um Erro:", e) 
  } 
  fmt.Println("Sáida:", o)
}
func sysrestart(arg string)  {
  command := "sudo pacman -S" + arg
  out, err := script.Exec(command).Stdout() 
  errorDescribe(out,err)
}
func instalaSystemctl()  {
  yes := script.Echo("y")
  commands := []string{
    "sudo pacman -Syu",
    "sudo pacman -S python2",
    "sudo mv /usr/bin/systemctl /usr/bin/systemctl.old",
    "sudo curl https://raw.githubusercontent.com/gdraheim/docker-systemctl-replacement/master/files/docker/systemctl.py -o /usr/bin/systemctl",
    "sudo chown root /usr/bin/systemctl",
    "sudo chgrp root /usr/bin/systemctl",
    "sudo chmod +x /usr/bin/systemctl",
  }
  for i, command := range commands {
    out, err := yes.Exec(command).Stdout()
    if err != nil {
      fmt.Println("O Comando", i, "Apresentou um Erro", err, out) 
      break
    }
  }
}
func teste()  {
}
