package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/mohrezfadaei/marzban-node-go/internal/config"
	"github.com/mohrezfadaei/marzban-node-go/internal/logger"
	"github.com/mohrezfadaei/marzban-node-go/internal/xray"
)

type Service struct {
	core *xray.XRayCore
}

func NewService(core *xray.XRayCore) *Service {
	return &Service{core: core}
}

func (s *Service) base(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"connected":    s.core.Connected(),
		"started":      s.core.Started(),
		"core_version": s.core.GetVersion(),
	})
}

func (s *Service) connect(c *gin.Context) {
	sessionID := uuid.New().String()
	clientIP := c.ClientIP()

	if s.core.Connected() {
		logger.Warn.Printf("New connection from %s, Core control access was taken away from previous client.", clientIP)
		if s.core.Started() {
			if err := s.core.Stop(); err != nil {
				logger.Error.Println(err)
			}
		}
	}

	s.core.Connect(sessionID, clientIP)
	logger.Info.Printf("%s connected, Session ID = %s.", clientIP, sessionID)
	c.JSON(http.StatusOK, gin.H{
		"session_id": sessionID,
	})
}

func (s *Service) disconnect(c *gin.Context) {
	if s.core.Connected() {
		logger.Info.Printf("%s disconnected, Session ID = %s.", s.core.ClientIP(), s.core.SessionID())
		s.core.Disconnect()
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (s *Service) start(c *gin.Context) {
	var req struct {
		SessionID string `json:"session_id"`
		Config    string `json:"config"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !s.core.MatchSessionID(req.SessionID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Session ID mismatch"})
		return
	}

	config, err := xray.NewConfig(req.Config, s.core.ClientIP())
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if err := s.core.Start(config); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (s *Service) stop(c *gin.Context) {
	var req struct {
		SessionID string `json:"session_id"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !s.core.MatchSessionID(req.SessionID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Session ID mismatch"})
		return
	}

	if err := s.core.Stop(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (s *Service) restart(c *gin.Context) {
	var req struct {
		SessionID string `json:"session_id"`
		Config    string `json:"config"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !s.core.MatchSessionID(req.SessionID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Session ID mismatch"})
		return
	}

	config, err := xray.NewConfig(req.Config, s.core.ClientIP())
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if err := s.core.Restart(config); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func RunRestServer(conf *config.Config, core *xray.XRayCore) error {
	r := gin.Default()

	service := NewService(core)
	r.POST("/", service.base)
	r.POST("/connect", service.connect)
	r.POST("/disconnect", service.disconnect)
	r.POST("/start", service.start)
	r.POST("/stop", service.stop)
	r.POST("/restart", service.restart)

	return r.RunTLS(fmt.Sprintf(":%d", conf.ServicePort), conf.SslCertFile, conf.SslKeyFile)
}
