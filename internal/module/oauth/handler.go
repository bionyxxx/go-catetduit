package oauth

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   *Service
	validator *validator.Validate
}

func NewHandler(service *Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) Google(w http.ResponseWriter, r *http.Request) {
	url := h.service.oauth2Config.GoogleConfig.AuthCodeURL(h.service.oauth2Config.StateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != h.service.oauth2Config.StateString {
		http.Redirect(w, r, h.service.oauth2Config.FailedRedirectUrl, http.StatusTemporaryRedirect)
		return
	}

	token, err := h.service.oauth2Config.GoogleConfig.Exchange(r.Context(), r.FormValue("code"))
	if err != nil {
		http.Redirect(w, r, h.service.oauth2Config.FailedRedirectUrl, http.StatusTemporaryRedirect)
		return
	}

	client := h.service.oauth2Config.GoogleConfig.Client(r.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Redirect(w, r, h.service.oauth2Config.FailedRedirectUrl, http.StatusTemporaryRedirect)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Redirect(w, r, h.service.oauth2Config.FailedRedirectUrl, http.StatusTemporaryRedirect)
			return
		}
	}(resp.Body)

	var googleUser GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		http.Redirect(w, r, h.service.oauth2Config.FailedRedirectUrl, http.StatusTemporaryRedirect)
		return
	}

	authResp, err := h.service.Google(&googleUser)

	if err != nil {
		http.Redirect(w, r, h.service.oauth2Config.FailedRedirectUrl, http.StatusTemporaryRedirect)
		return
	}

	// 5. Set Cookie
	// Set cookie dengan domain yang shared
	if h.service.mainConfig.IsProduction {
		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    authResp.AccessToken,
			Path:     "/",
			Domain:   "." + h.service.mainConfig.Domain, // Domain shared untuk semua subdomain
			HttpOnly: true,
			Secure:   true,                  // Wajib true untuk SameSite=None
			SameSite: http.SameSiteNoneMode, // Bukan Lax, tapi None untuk cross-site
			MaxAge: func() int {
				remaining := int(authResp.ExpiresAt - time.Now().Unix())
				if remaining < 0 {
					return 0
				}
				return remaining
			}(),
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    authResp.RefreshToken,
			Path:     "/",
			Domain:   "." + h.service.mainConfig.Domain, // Domain shared untuk semua subdomain
			HttpOnly: true,
			Secure:   true, // Wajib true
			SameSite: http.SameSiteNoneMode,
			MaxAge:   int((time.Duration(h.service.jwtHelper.GetJWTRefreshExpiredInHour()) * time.Hour).Seconds()),
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    authResp.AccessToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,                // Wajib true untuk SameSite=None
			SameSite: http.SameSiteLaxMode, // Bukan Lax, tapi None untuk cross-site
			MaxAge: func() int {
				remaining := int(authResp.ExpiresAt - time.Now().Unix())
				if remaining < 0 {
					return 0
				}
				return remaining
			}(),
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    authResp.RefreshToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   false, // Wajib true
			SameSite: http.SameSiteLaxMode,
			MaxAge:   int((time.Duration(h.service.jwtHelper.GetJWTRefreshExpiredInHour()) * time.Hour).Seconds()),
		})
	}

	// 6. Redirect ke Frontend
	http.Redirect(w, r, h.service.oauth2Config.RedirectUrl, http.StatusTemporaryRedirect)
}
