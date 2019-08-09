
package main

import (  
    "fmt"
    "io/ioutil"
    "path/filepath"
     "os"
)

func main() {


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

      fmt.Println("Reading "+ dirname + "\n")

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
