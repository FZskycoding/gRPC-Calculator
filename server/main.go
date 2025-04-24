package main

import (
    "log"  // 用於日誌輸出
    "net"  // 用於網絡監聽
    "google.golang.org/grpc"
    pb "gRPC-Calculator/proto"
	"gRPC-Calculator/server/service"  

)

func main() {
    // 4. 創建 TCP 監聽器，監聽 50051 端口
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("監聽端口失敗: %v", err)
    }

    // 5. 創建 gRPC 服務器
    grpcServer := grpc.NewServer()

    // 6. 註冊計算器服務
    calculatorServer := service.NewCalculatorServer()
		pb.RegisterCalculatorServiceServer(grpcServer, calculatorServer)

    // 7. 啟動服務器
    log.Printf("服務器正在監聽端口 :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("服務器運行失敗: %v", err)
    }
}
