# gRPC 計算器服務

這是一個基於 gRPC 的遠程計算器服務實現，展示了現代微服務架構中的關鍵技術應用。

## 技術棧

### 1. Protocol Buffers (protobuf)
- 使用 Protocol Buffers v3 定義服務接口
- 實現了強類型的訊息結構
- 自動生成客戶端和服務端代碼
- 二進制序列化，高效的數據傳輸

### 2. gRPC
- 實現客戶端-服務器通信
- 使用 HTTP/2 作為傳輸協議
- 支持雙向流通信
- 內建負載均衡和安全機制

### 3. Go 語言特性應用
- goroutines 併發處理
- channels 用於數據流控制
- error handling 錯誤處理
- defer 資源管理
- interface 接口實現

### 4. 系統設計模式
- 客戶端-服務器架構
- 遠程過程調用 (RPC)
- 接口驅動設計
- 模塊化設計



