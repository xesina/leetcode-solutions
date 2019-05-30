package main
import ( 
    "fmt"
    "strconv"
    "bufio"
    "os"
)


func twoSumStdin() {
    s := bufio.NewScanner(os.Stdin)
    s.Scan()
    sum, _ := strconv.Atoi(s.Text())
    s.Scan()
    l, _ := strconv.Atoi(s.Text())

    m := make(map[int]int)
    for i:=0; i<l; i++ {
        s.Scan()
        v, _ := strconv.Atoi(s.Text())
        
        diff := sum - v

        if _, ok := m[diff]; ok {
           fmt.Println(1)
           return
       }
       
       m[v] = 1
    }
    
    fmt.Println(0)
}