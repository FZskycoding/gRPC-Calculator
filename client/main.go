package main

import (
    "log"  // 用於日誌輸出
    "context"  // 用於上下文管理
		"fmt"
		"bufio"
		"strconv"
		"os"
		"strings"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "github.com/FZskycoding/gRPC-Calculator/proto"
)

func main() {
    // 連接到服務器
    conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("無法連接到服務器: %v", err)
    }
    defer conn.Close()

    // 創建客戶端
    client := pb.NewCalculatorServiceClient(conn)

		// 讀取用戶輸入
    reader := bufio.NewReader(os.Stdin)
		// 讀取第一個數字
		fmt.Print("請輸入第一個數字: ")
		firstInput, _ := reader.ReadString('\n')
		firstNum, err := strconv.ParseFloat(strings.TrimSpace(firstInput), 64)
    if err != nil {
        log.Fatalf("無效的數字: %v", err)
    }
		// 讀取第二個數字
    fmt.Print("請輸入第二個數字: ")
    secondInput, _ := reader.ReadString('\n')
    secondNum, err := strconv.ParseFloat(strings.TrimSpace(secondInput), 64)
    if err != nil {
        log.Fatalf("無效的數字: %v", err)
    }

		// 讀取運算符
    fmt.Print("請輸入運算符 (+、-、*、/) :")
    op, _ := reader.ReadString('\n')
    op = strings.TrimSpace(op)

    //準備請求數據
    req := &pb.CalculateRequest{
        FirstNumber:  firstNum,    
        SecondNumber: secondNum,    
        Operation:    op,    
    }

    //發送請求並獲取響應
    resp, err := client.Calculate(context.Background(), req)
    if err != nil {
        log.Fatalf("調用計算服務失敗: %v", err)
    }

    //處理響應
    if resp.Success {
        log.Printf("計算結果: %.2f\n", resp.Result)
    } else {
        log.Printf("計算失敗: %s\n", resp.ErrorMessage)
    }
}
