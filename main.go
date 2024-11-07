package main

import "github.com/01-edu/z01"

func main(){
  for i:='a' ; i <= 'z' ;i++{
    
  z01.PrintRune(i)
  }
    z01.PrintRune('\n')
}


package main

import "github.com/01-edu/z01"

func main(){
  for i:='1' ; i<= '9'  ;i++{
    
  z01.PrintRune(i)
  }
    z01.PrintRune('\n')
}

package piscine 

import "github.com/01-edu/z01" 

func IsNegative(nb int) {
  if nb < 0{
    z01.PrintRune('T')
  }else{
      z01.PrintRune('F') 
  }
    z01.PrintRune('\n')
}


