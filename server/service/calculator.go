package service

import (
    // 1. 導入必要的包
    "context"  // 用於上下文
    "fmt"  // 用於格式化字符串

    // 2. 導入生成的 proto 包
    pb "gRPC-Calculator/proto"
)

// 3. 定義計算器服務器結構體
type CalculatorServer struct {
    // 嵌入未實現的服務接口
    pb.UnimplementedCalculatorServiceServer
}

// 4. 創建新的計算器服務實例
func NewCalculatorServer() *CalculatorServer {
    return &CalculatorServer{}
}

// 5. 實現 Calculate 方法
func (s *CalculatorServer) Calculate(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
    // 獲取請求中的數字和操作符
    firstNum := req.FirstNumber
    secondNum := req.SecondNumber
    operation := req.Operation

    // 定義結果變量
    var result float64
    var success bool = true
    var errMessage string

    // 根據操作類型進行計算
    switch operation {
    case "+":
        result = firstNum + secondNum
    case "-":
        result = firstNum - secondNum
    case "*":
        result = firstNum * secondNum
    case "/":
        // 處理除法，需要檢查除數是否為零
        if secondNum == 0 {
            return &pb.CalculateResponse{
                Success: false,
                ErrorMessage: "除數不能為零",
            }, nil
        }
        result = firstNum / secondNum
    default:
        return &pb.CalculateResponse{
            Success: false,
            ErrorMessage: fmt.Sprintf("不支持的操作：%s", operation),
        }, nil
    }

    // 返回計算結果
    return &pb.CalculateResponse{
        Result: result,
        Success: success,
        ErrorMessage: errMessage,
    }, nil
}
