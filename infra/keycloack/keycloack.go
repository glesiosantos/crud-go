package keycloack

import (
	"context"
	"fmt"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
)

type Keycloak struct {
	Verifier *oidc.IDTokenVerifier
}

func NewKeycloak(ctx context.Context) (*Keycloak, error) {
	keycloakURL := os.Getenv("KEYCLOAK_URL")
	realm := os.Getenv("KEYCLOAK_REALM")
	clientID := os.Getenv("KEYCLOAK_CLIENT_ID")

	if keycloakURL == "" {
		return nil, fmt.Errorf("KEYCLOAK_URL não configurada")
	}

	if realm == "" {
		return nil, fmt.Errorf("KEYCLOAK_REALM não configurado")
	}

	if clientID == "" {
		return nil, fmt.Errorf("KEYCLOAK_CLIENT_ID não configurado")
	}

	issuer := fmt.Sprintf(
		"%s/realms/%s",
		keycloakURL,
		realm,
	)

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao Keycloak: %w", err)
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: clientID,
	})

	return &Keycloak{
		Verifier: verifier,
	}, nil
}