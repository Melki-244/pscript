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
  case os.Args[1] == "install":
   switch{
   case os.Args[2] == "pgadmin4":
    pgAdminInstall()
   case os.Args[2] == "systemctl":
    systemctlInstall()
   }
  case os.Args[1] == "sysrestart":
   sysrestart()
  case os.Args[1] == "teste":
   teste()
   teste()
  case os.Args[1] == "start":
   switch{
   case os.Args[2] == "pgadmin4":
    pgAdminStart()
   }
  default:
   fmt.Printf("Argumento Não Conhecido")
  }
}
func errorDescribe(o int, e error)  {
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
    errorDescribe(out, err)
  }
}
func execCommands(commands []string)  {
  for _, command := range commands {
    out, err := script.Exec(command).Stdout()
    if err != nil {
      break
    }
    errorDescribe(out, err)
  }
}
func systemctlInstall()  {
  commands := []string{
    "sudo pacman -Syuu --noconfirm",
    "sudo pacman -S python2 --noconfirm",
    "sudo mv /usr/bin/systemctl /usr/bin/systemctl.old",
    "sudo curl https://raw.githubusercontent.com/gdraheim/docker-systemctl-replacement/master/files/docker/systemctl.py -o /usr/bin/systemctl",
    "sudo chown root /usr/bin/systemctl",
    "sudo chgrp root /usr/bin/systemctl",
    "sudo chmod +x /usr/bin/systemctl",
  } 
  execCommands(commands)
}
func sysrestart()  {
  commands := []string{
    "sudo systemctl stop {{.}}",
    "sudo systemctl start {{.}}",
    "sudo systemctl status {{.}}",
  }
  execCommandsArgs(os.Args[2], commands)
}
func pgAdminInstall()  {
  commands := []string{
    "sudo mkdir /var/lib/pgadmin",
    "sudo mkdir /var/log/pgadmin",
    "sudo chown {{.}} /var/lib/pgadmin",
    "sudo chown {{.}} /var/log/pgadmin",
    "sudo pacman -Syyuu --noconfirm",
    "pip install --upgrade pip",
    "python -m venv ~/.pgadmin4",
    "source ~/.pgadmin4/bin/activate",
    "pip install pgadmin4",
  } 
  user, _ :=  os.LookupEnv("USER")
  execCommandsArgs(user,commands)
}
func pgAdminStart()  {
  commands := []string{
    "source ~/.pgadmin4/bin/activate",
    "pgadmin4",
  } 
  execCommands(commands)
}
func teste()  {

}
