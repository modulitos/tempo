package app

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/grafana/tempo/modules/distributor"
	"github.com/grafana/tempo/modules/distributor/forwarder"
	"github.com/grafana/tempo/modules/overrides/userconfigurableapi"
)

func Test_overridesValidator(t *testing.T) {
	testCases := []struct {
		name   string
		cfg    Config
		limits userconfigurableapi.UserConfigurableLimits
		expErr string
	}{
		{
			name: "forwarders - valid",
			cfg: Config{
				Distributor: distributor.Config{
					Forwarders: []forwarder.Config{
						{Name: "forwarder-1"},
						{Name: "forwarder-2"},
						{Name: "forwarder-3"},
					},
				},
			},
			limits: userconfigurableapi.UserConfigurableLimits{
				Forwarders: &[]string{"forwarder-1", "forwarder-3"},
			},
		},
		{
			name: "forwarders - invalid",
			cfg: Config{
				Distributor: distributor.Config{
					Forwarders: []forwarder.Config{
						{Name: "forwarder-1"},
						{Name: "forwarder-2"},
					},
				},
			},
			limits: userconfigurableapi.UserConfigurableLimits{
				Forwarders: &[]string{"forwarder-1", "some-forwarder"},
			},
			expErr: "forwarder \"some-forwarder\" is not a known forwarder, contact your system administrator",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := NewOverridesValidator(&tc.cfg)

			err := validator.Validate(&tc.limits)
			if tc.expErr != "" {
				assert.EqualError(t, err, tc.expErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
