package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

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
//    - registerUser("new@example.com")   → 成

// TODO: 在这里写你的代码
var ErrDuplicateEmail = errors.New("duplicate email")

type FieldError struct {
	Field   string
	Message string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("FieldError: Field(%s) is %s", e.Field, e.Message)
}

func validateEmail(email string) error {
	if email == "" {
		return &FieldError{Field: "email", Message: "required"}
	}
	if !strings.Contains(email, "@") {
		return &FieldError{Field: "email", Message: "invalid format"}
	}
	return nil
}

func registerUser(email string) error {
	err := validateEmail(email)
	if err != nil {
		return fmt.Errorf("registerUser: %w", err)
	}
	takenEmails := []string{"taken@example.com", "taken2@example.com"}
	if slices.Contains(takenEmails, email) {
		return fmt.Errorf("registerUser: %w", ErrDuplicateEmail)
	}
	fmt.Println("registered: ", email)
	return nil
}

func errorCheck(err error) {
	var validErr *FieldError
	if errors.As(err, &validErr) {
		fmt.Printf("FieldError Field(%s) is %s\n", validErr.Field, validErr.Message)
	}
	if errors.Is(err, ErrDuplicateEmail) {
		fmt.Println("ErrDuplicateEmail")
	}
}

func exercise() {
	fmt.Println("=== Exercise: error 处理模式 ===")

	errorCheck(registerUser(""))                  // FieldError
	errorCheck(registerUser("bad-email"))         // FieldError
	errorCheck(registerUser("taken@example.com")) // ErrDuplicateEmail
	errorCheck(registerUser("new@example.com"))   // Success
}
