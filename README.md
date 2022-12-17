# Coding
Struct to Json = ตัวเเปลง Type Struct to Json 
   - กำหนด type Struct เป็น Json
        type Exmaple struct {
            ID int `json:"id"`
            Name string `json:"name"`
            Age int `json:"age"`
        }
   - ใช้คำสั่ง json.Marshal(struct name) ซึ่งจะ Return ออกมา 2 ตัว คือ Result And Error
        b, err := json.Marshal(t)
   - เช็ด log เเต่ละ type ได้จากคำสั่ง
        fmt.Printf("type : %T \n", b) // type : []uint8
        fmt.Printf("byte : %v \n", b) // byte : [123 34 105 100 …]
        fmt.Printf("string: %s \n", b) // string: {"id":1,"name":"AnuchitO","age":18}
        fmt.Println(err) // nil
          
 