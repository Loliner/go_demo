package main

import "fmt"

// ============================================================
// 练习：error 处理模式
// ============================================================
//
// 目标：实现一个简易用户注册流程，综合运用本章知识
//
// 1. 定义 sentinel error：
//    var ErrDuplicateEmail = errors.New("duplicate email")
//
// 2. 定义自定义 error 类型 FieldError：
//    { Field string, Message string }，实现 error interface
//
// 3. 实现 validateEmail(email string) error：
//    - email 为空 → 返回 &FieldError{Field: "email", Message: "required"}
//    - email 不含 "@" → 返回 &FieldError{Field: "email", Message: "invalid format"}
//    - 否则返回 nil
//    提示：用 strings.Contains(email, "@") 检查格式
//
// 4. 实现 registerUser(email string) error：
//    - 先调用 validateEmail，如果有错就 wrap 后返回：
//      fmt.Errorf("registerUser: %w", err)
//    - 模拟已注册：如果 email == "taken@example.com"，返回包装后的 ErrDuplicateEmail
//    - 否则打印 "registered: <email>" 并返回 nil
//
// 5. 在 exercise() 里测试以下场景，用 errors.Is / errors.As 判断错误类型：
//    - registerUser("")          → FieldError
//    - registerUser("bad-email") → FieldError
//    - registerUser("taken@example.com") → ErrDuplicateEmail
//    - registerUser("new@example.com")   → 成功

// TODO: 在这里写你的代码

func exercise() {
	fmt.Println("=== Exercise: error 处理模式 ===")
}
