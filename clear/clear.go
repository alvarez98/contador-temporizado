package clear 

import (
    "os" 
    "os/exec" 
    "runtime" 
) 

var clear map[string]func() // Create a map for storing clear funcs 

func init() { 
	clear = make(map[string]func()) 
	
    clear["linux"] = func() { 
     cmd := exec.Command("clear") 
     cmd.Stdout = os.Stdout 
     cmd.Run() 
	} 
	
    clear["windows"] = func() { 
     cmd := exec.Command("cmd", "/c", "cls")
     cmd.Stdout = os.Stdout 
     cmd.Run() 
    } 
} 

func CallClear() { 
    value, ok := clear[runtime.GOOS] // Runtime.GOOS -> linux, windows, darwin etc. 
    if ok { 
     value() 
    } else { 
     panic("Su plataforma no esta soportada para limpiar la pantalla: (") 
    } 
} 