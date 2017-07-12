package vpn

import (
	"time"

	"github.com/kapmahc/fly/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

type fmSignIn struct {
	Email    string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"min=6,max=32"`
}

func (p *Engine) apiAuth(c *gin.Context) (interface{}, error) {
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	lang := c.MustGet(web.LOCALE).(string)
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return nil, err
	}
	now := time.Now()
	if user.Enable && user.StartUp.Before(now) && user.ShutDown.After(now) {
		return gin.H{}, nil
	}
	return nil, p.I18n.E(lang, "ops.vpn.errors.user-is-not-available")
}

type fmStatus struct {
	Email       string  `form:"common_name" binding:"required,email"`
	TrustedIP   string  `form:"trusted_ip" binding:"required"`
	TrustedPort uint    `form:"trusted_port" binding:"required"`
	RemoteIP    string  `form:"ifconfig_pool_remote_ip" binding:"required"`
	RemotePort  uint    `form:"remote_port_1" binding:"required"`
	Received    float64 `form:"bytes_received" binding:"required"`
	Send        float64 `form:"bytes_sent" binding:"required"`
}

func (p *Engine) apiConnect(c *gin.Context) (interface{}, error) {
	var fm fmStatus
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Create(&Log{
		UserID:      user.ID,
		RemoteIP:    fm.RemoteIP,
		RemotePort:  fm.RemotePort,
		TrustedIP:   fm.TrustedIP,
		TrustedPort: fm.TrustedPort,
		Received:    fm.Received,
		Send:        fm.Send,
		StartUp:     time.Now(),
	}).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&User{}).
		Where("id = ?", user.ID).
		UpdateColumn("online", true).Error; err != nil {
		return nil, err
	}

	return gin.H{}, nil
}

func (p *Engine) apiDisconnect(c *gin.Context) (interface{}, error) {
	var fm fmStatus
	if err := c.Bind(&fm); err != nil {
		return nil, err
	}
	var user User
	if err := p.Db.Where("email = ?", fm.Email).First(&user).Error; err != nil {
		return nil, err
	}
	if err := p.Db.Model(&User{}).
		Where("id = ?", user.ID).
		UpdateColumn("online", false).Error; err != nil {
		return nil, err
	}

	if err := p.Db.
		Model(&Log{}).
		Where(
			"trusted_ip = ? AND trusted_port = ? AND user_id = ? AND shut_down IS NULL",
			fm.TrustedIP,
			fm.TrustedPort,
			user.ID,
		).Update(map[string]interface{}{
		"shut_down": time.Now(),
		"received":  fm.Received,
		"send":      fm.Send,
	}).Error; err != nil {
		return nil, err
	}
	return gin.H{}, nil
}
