package main_oo

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
)


func getSave(stock []string)int{

     stks := "ex_ch="

     for _, v := range stock {
      stks = stks + v+".tw|"
    }
   
   url := "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?"+stks

   /* ============================================================== */

   // 根据URL获取资源
      res, err := http.Get(url)
 
      if err != nil {
          fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
          os.Exit(1)
      }
 
      // 读取资源数据 body: []byte
      body, err := ioutil.ReadAll(res.Body)
 
      // 关闭资源流
      res.Body.Close()
 
      if err != nil {
          fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
          os.Exit(1)
      }
 
      // 控制台打印内容 以下两种方法等同
      fmt.Printf("%s", body)
      //fmt.Printf(string(body))
 
      // 写入文件
      // ioutil.WriteFile("site.txt", body, 0644)
      
      return 1 
}

func main() {

    //url := "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=otc_6180.tw|tse_0050.tw|tse_2330.tw"

    var stk1 = []string{"otc_6180","tse_0050","tse_2330","tse_3057","tse_2303","otc_6160"}
    getSave(stk1)
    
}


