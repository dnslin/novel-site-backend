package middleware

import (
	"novel-site-backend/pkg/log"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	CtxClientIPKey = "client_ip"
)

func ClientIPMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIP := getClientIP(ctx)

		// 将 IP 存储到上下文中
		ctx.Set(CtxClientIPKey, clientIP)

		// 添加到日志字段中
		logger.WithValue(ctx, zap.String("client_ip", clientIP))

		ctx.Next()
	}
}

// getClientIP 获取客户端真实IP
func getClientIP(ctx *gin.Context) string {
	// 从 X-Forwarded-For 获取
	xForwardedFor := ctx.GetHeader("X-Forwarded-For")
	if xForwardedFor != "" {
		// X-Forwarded-For 可能包含多个 IP，第一个是真实客户端 IP
		ips := strings.Split(xForwardedFor, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// 从 X-Real-IP 获取
	xRealIP := ctx.GetHeader("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	// 从 RemoteAddr 获取
	return ctx.ClientIP()
}

// GetClientIP 从上下文获取客户端IP
func GetClientIP(ctx *gin.Context) string {
	if ip, exists := ctx.Get(CtxClientIPKey); exists {
		return ip.(string)
	}
	return ""
}
