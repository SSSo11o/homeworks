//               uslovnie operatori (if, else)

package main
import "fmt"
func main (){
	num := 15
	if num > 9 && num < 100 {
		fmt.Println("Двухзначные")
	}
	n := 6
	if n % 2 == 1 {
		fmt.Println("Nechotnoe")
	} else {
		fmt.Println("chetnoe")
	}

	m := 785
	if m >= 0 && m < 10 { 
		fmt.Println("odnoznachnae")
	} else if m >= 10 && m < 100 {
		fmt.Println("dvukh")
	} else if m >= 100 && m < 1000 {
		fmt.Println("trekhznach")
	} else  {
		fmt.Println("chislo bolshoe")
	}

	
	a := 9
	b := 6
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

	// n1 := 5
	// n2 := 7
	// n3 := 8

	// if n1 <= n2 && n2 <= n3 {
	// 	fmt.Println(n3 + n2)
	//   } else if n1 <= n3 && n3 <= n2 {
	// 	fmt.Println(n2 + n3)
	//   } else if n2 <= n1 && n1 <= n3 {
	// 	fmt.Println(n3 + n1)
	//   } else if n2 <= n3 && n3 <= n1 {
	// 	fmt.Println(n1 + n3)
	//   } else if n3 <= n1 && n1 <= n2 {
	// 	fmt.Println(n2 + n1)
	//   } else {
	// 	fmt.Println(n1 + n2)
	//   }
	

//                        switch
	// fallthrough -- eto oznachaet shto uslovie ne proveraetsa 

	nn := -1
	switch nn {
	case 1 :
		fmt.Println("positive")
	case -1 :
		fmt.Println("negativ")
	}

	n1 := 5
	n2 := 7
	n3 := 8

	if n1 <= n2 && n2 <= n3 {
		fmt.Println(n3 + n2)
	  } else if n1 <= n3 && n3 <= n2 {
		fmt.Println(n2 + n3)
	  } else if n2 <= n1 && n1 <= n3 {
		fmt.Println(n3 + n1)
	  } else if n2 <= n3 && n3 <= n1 {
		fmt.Println(n1 + n3)
	  } else if n3 <= n1 && n1 <= n2 {
		fmt.Println(n2 + n1)
	  } else {
		fmt.Println(n1 + n2)
	  }
 //                    Scan 
	  var s int
	  s = 5
	  fmt.Scan(&s)
	  fmt.Println(s)
	
	  nm := 1
	  switch nm {
	  case 1 :
		  fmt.Println("Январь 31 дней")
	  case 2 :
		  fmt.Println("Феврадя 28 дней")
		case 3 : 
		fmt.Println("март 31 дней")
	  case 4:
		fmt.Println("апрель 30 дней")
	  case 5:
		fmt.Println("май 31 дней")
	  case 6:
		fmt.Println("июнь 30 дней")
	  case 7:
		fmt.Println("июль 31 дней")
	case 8:
		fmt.Println("августь 31 дней")
	case 9:
		fmt.Println("сентябрь 30 дней")
	case 10:
		fmt.Println("октябрь 31 дней")
	case 11:
		fmt.Println("ноябрь 30 дней")
	case 12:
		fmt.Println("декабрь 31 дней")
	  } 

	



}
