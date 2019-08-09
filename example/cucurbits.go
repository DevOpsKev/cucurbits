
package main

import (  
    "fmt"
    "io/ioutil"
    "path/filepath"
     "os"
)

func main() {

    //Bubblehead font generated online at http://patorjk.com/software/taag/#p=display&f=Bulbhead


    fmt.Println(`
      ___  __  __  ___  ____  ____  ____  ____  ____  ___ 
     / __)(  )(  )/ __)( ___)(  _ \(  _ \(_  _)(_  _)/ __)
    ( (__  )(__)(( (__  )__)  )   / ) _ < _)(_   )(  \__ \
     \___)(______)\___)(____)(_)\_)(____/(____) (__) (___/                                    
    `)


    fmt.Println(`
     ____  ____    __    ____  ____  _  _  ___    ____  ____  __    ____  ____  ____  ___       
    (  _ \( ___)  /__\  (  _ \(_  _)( \( )/ __)  (  _ \( ___)(  )  (_  _)( ___)( ___)/ __)      
     )   / )__)  /(__)\  )(_) )_)(_  )  (( (_-.   ) _ < )__)  )(__  _)(_  )__)  )__) \__ \      
    (_)\_)(____)(__)(__)(____/(____)(_)\_)\___/  (____/(____)(____)(____)(____)(__)  (___/()()()


    `)


    dirname := "." + string(filepath.Separator)

      d, err := os.Open(dirname)
      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }
      defer d.Close()

      files, err := d.Readdir(-1)
      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }

      for _, file := range files {
          
          if file.Mode().IsRegular() {
              if filepath.Ext(file.Name()) == ".belief" {
                
                fmt.Println(file.Name())
              
              	data, err := ioutil.ReadFile(file.Name())
    
			    if err != nil {
			        fmt.Println("File reading error", err)
			        return
			    }

    			fmt.Println(string(data))

              }
          }


      }
 
}
