package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func SendVerificationEmail(to, token string) error {
	apiKey := os.Getenv("BREVO_API_KEY")
	fromEmail := os.Getenv("EMAIL_FROM")

	link := fmt.Sprintf("http://localhost:8080/api/v1/auth/verify?token=%s", token)

	// HTML email template
	htmlTemplate := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Verify Your Email - ORDO</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            line-height: 1.6;
            color: #333;
            background: linear-gradient(135deg, #f5f7fa 0%%, #c3cfe2 100%%);
            min-height: 100vh;
            padding: 20px;
        }
        
        .email-container {
            max-width: 600px;
            margin: 0 auto;
            background: white;
            border-radius: 16px;
            box-shadow: 0 20px 40px rgba(0, 146, 244, 0.1);
            overflow: hidden;
            border: 1px solid rgba(0, 146, 244, 0.1);
        }
        
        .header {
            background: linear-gradient(135deg, #0092f4 0%%, #0074cc 100%%);
            padding: 40px 30px;
            text-align: center;
            color: white;
            position: relative;
            overflow: hidden;
        }
        
        .header::before {
            content: '';
            position: absolute;
            top: -50%%;
            left: -50%%;
            width: 200%%;
            height: 200%%;
            background: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><defs><pattern id="grain" width="100" height="100" patternUnits="userSpaceOnUse"><circle cx="25" cy="25" r="1" fill="white" opacity="0.1"/><circle cx="75" cy="75" r="1" fill="white" opacity="0.1"/><circle cx="50" cy="10" r="0.5" fill="white" opacity="0.15"/><circle cx="20" cy="80" r="0.5" fill="white" opacity="0.15"/></pattern></defs><rect width="100" height="100" fill="url(%%23grain)"/></svg>');
            animation: float 20s ease-in-out infinite;
        }
        
        @keyframes float {
            0%%, 100%% { transform: translateY(0px) rotate(0deg); }
            50%% { transform: translateY(-10px) rotate(1deg); }
        }
        
        .logo {
            font-size: 2.5rem;
            font-weight: 800;
            letter-spacing: 2px;
            margin-bottom: 10px;
            position: relative;
            z-index: 1;
        }
        
        .header-subtitle {
            font-size: 1.1rem;
            opacity: 0.9;
            font-weight: 300;
            position: relative;
            z-index: 1;
        }
        
        .content {
            padding: 50px 40px;
            text-align: center;
        }
        
        .welcome-text {
            font-size: 1.8rem;
            font-weight: 600;
            color: #2c3e50;
            margin-bottom: 20px;
        }
        
        .description {
            font-size: 1.1rem;
            color: #666;
            margin-bottom: 40px;
            line-height: 1.7;
        }
        
        .verify-button {
            display: inline-block;
            background: linear-gradient(135deg, #0092f4 0%%, #0074cc 100%%);
            color: white;
            text-decoration: none;
            padding: 18px 40px;
            border-radius: 50px;
            font-weight: 600;
            font-size: 1.1rem;
            letter-spacing: 0.5px;
            transition: all 0.3s ease;
            box-shadow: 0 8px 25px rgba(0, 146, 244, 0.3);
            position: relative;
            overflow: hidden;
        }
        
        .verify-button::before {
            content: '';
            position: absolute;
            top: 0;
            left: -100%%;
            width: 100%%;
            height: 100%%;
            background: linear-gradient(90deg, transparent, rgba(255,255,255,0.2), transparent);
            transition: left 0.5s;
        }
        
        .verify-button:hover::before {
            left: 100%%;
        }
        
        .verify-button:hover {
            transform: translateY(-2px);
            box-shadow: 0 12px 35px rgba(0, 146, 244, 0.4);
        }
        
        .security-note {
            margin-top: 40px;
            padding: 25px;
            background: #f8f9ff;
            border-radius: 12px;
            border-left: 4px solid #0092f4;
        }
        
        .security-title {
            font-weight: 600;
            color: #0092f4;
            margin-bottom: 8px;
            font-size: 1rem;
        }
        
        .security-text {
            color: #666;
            font-size: 0.9rem;
            line-height: 1.6;
        }
        
        .footer {
            background: #f8f9fa;
            padding: 30px 40px;
            text-align: center;
            color: #888;
            font-size: 0.9rem;
            border-top: 1px solid #eee;
        }
        
        .footer-links {
            margin-top: 15px;
        }
        
        .footer-links a {
            color: #0092f4;
            text-decoration: none;
            margin: 0 10px;
            font-weight: 500;
        }
        
        .footer-links a:hover {
            text-decoration: underline;
        }
        
        /* Mobile Responsiveness */
        @media (max-width: 600px) {
            body {
                padding: 10px;
            }
            
            .email-container {
                border-radius: 12px;
            }
            
            .header {
                padding: 30px 20px;
            }
            
            .logo {
                font-size: 2rem;
            }
            
            .header-subtitle {
                font-size: 1rem;
            }
            
            .content {
                padding: 40px 25px;
            }
            
            .welcome-text {
                font-size: 1.5rem;
            }
            
            .description {
                font-size: 1rem;
            }
            
            .verify-button {
                padding: 16px 35px;
                font-size: 1rem;
            }
            
            .security-note {
                padding: 20px;
                margin-top: 30px;
            }
            
            .footer {
                padding: 25px 20px;
            }
        }
        
        @media (max-width: 400px) {
            .verify-button {
                padding: 14px 30px;
                font-size: 0.95rem;
            }
            
            .welcome-text {
                font-size: 1.3rem;
            }
        }
    </style>
</head>
<body>
    <div class="email-container">
        <div class="header">
            <div class="logo">ORDO</div>
            <div class="header-subtitle">Welcome to the future</div>
        </div>
        
        <div class="content">
            <h1 class="welcome-text">Welcome to ORDO!</h1>
            <p class="description">
                Thank you for joining us! To complete your registration and secure your account, 
                please verify your email address by clicking the button below.
            </p>
            
            <a href="%s" class="verify-button">Verify Email Address</a>
            
            <div class="security-note">
                <div class="security-title">🔒 Security Notice</div>
                <div class="security-text">
                    This verification link will expire in 24 hours for your security. 
                    If you didn't create an account with ORDO, please ignore this email.
                </div>
            </div>
        </div>
        
        <div class="footer">
            <p>&copy; 2024 ORDO. All rights reserved.</p>
            <div class="footer-links">
                <a href="#">Privacy Policy</a>
                <a href="#">Terms of Service</a>
                <a href="#">Contact Support</a>
            </div>
        </div>
    </div>
</body>
</html>`

	// Create request body
	data := map[string]interface{}{
		"sender": map[string]string{
			"name":  "ORDO",
			"email": fromEmail,
		},
		"to": []map[string]string{
			{"email": to},
		},
		"subject": "Verify Your Email - ORDO",
		"htmlContent": fmt.Sprintf(htmlTemplate, link),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to send email: status code %d", resp.StatusCode)
	}

	return nil
}