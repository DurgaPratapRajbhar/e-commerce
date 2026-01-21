package service

import (
	"fmt"
)

type EmailService interface {
	SendVerificationEmail(to, token string) error
	SendPasswordResetEmail(to, token string) error
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type SMTPService struct {
	config SMTPConfig
}

func NewSMTPService(config SMTPConfig) *SMTPService {
	return &SMTPService{config: config}
}

func (s *SMTPService) SendVerificationEmail(to, token string) error {
	fmt.Printf("Sending verification email to %s with token %s\n", to, token)
	return nil
}

func (s *SMTPService) SendPasswordResetEmail(to, token string) error {
	fmt.Printf("Sending password reset email to %s with token %s\n", to, token)
	return nil
}