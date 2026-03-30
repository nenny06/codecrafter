func hexbin(s string) int {
	var output int
	word := strings.Fields(s)
	for i := 1; i < len(word); i++ {
		prev := word[i-1]
		// var hex int
		// var bin int
		if len(word) == 0 {
			break
		}
		if word[i] == "hex" || word[i] == "bin" && i > 0 {
			if word[i] == "hex" {
				hex, err := strconv.ParseInt(prev, 16, 64)
				if err != nil {
					fmt.Println("Error: incorrect base")
				}
				prev = strconv.FormatInt(hex, 10)
				result, _ := strconv.Atoi(prev)
				output = result
			} else if word[i] == "bin" {
				bin, err := strconv.ParseInt(prev, 2, 64)
				if err != nil {
					fmt.Println("Error: incorrect base")
				}
				prev = strconv.FormatInt(bin, 10)
				den, _ := strconv.Atoi(prev)
				output = den
			}
		}
		//fmt.Printf()
	}
	// word = append(word[:i], word[i+1:]...)
	// }
 	return output
 }

func decToBinHex(s string)string {
  var den string
  word := strings.Fields(s)
  for i := 1; i < len(word); i++{
    prev := word[i - 1]
    if word[i] == "dec" && i > 0{
      dec, err := strconv.ParseInt(prev, 10, 64)
      if err != nil{
        fmt.Println("Error: Not a base")
        }
      hex := strings.ToUpper(strconv.FormatInt(dec, 16))
      bin := strconv.FormatInt(dec, 2)
      den = fmt.Sprintf("%s\n", hex, bin)
      }
    }
  return den
  }
      
