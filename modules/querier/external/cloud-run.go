package external

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/api/idtoken"
)

// Set the HTTP Authorization header according to our google application
// credentials:
// https://cloud.google.com/run/docs/authenticating/service-to-service#acquire-token
func cloudRunAuthenticator(ctx context.Context, endpoint string, req *http.Request) error {
	ts, err := idtoken.NewTokenSource(ctx, endpoint)
	if err != nil {
		return fmt.Errorf("cloud-run authenticator failed to create token source: %w", err)
	}
	token, err := ts.Token()
	if err != nil {
		return fmt.Errorf("cloud-run authenticator failed to get token: %w", err)
	}

	token.SetAuthHeader(req)
	return nil
}
