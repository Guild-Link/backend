package hypixel

import (
	"context"
)

type Networth struct {
	Total       float64
	Unsoulbound float64
	Profile     *SkyBlockProfile
}

func (c *Client) GetNetworth(ctx context.Context, username, profileName string) (*Networth, error) {
	profile, err := c.GetSkyBlockProfile(ctx, username, profileName)
	if err != nil {
		return nil, err
	}

	museum, err := c.GetMuseumRaw(ctx, profile.ID)
	if err != nil {
		return nil, err
	}

	result, err := c.compat.Networth(ctx, profile.Raw, museum, profile.Player.ID)
	if err != nil {
		return nil, err
	}

	return &Networth{
		Total:       result.Networth,
		Unsoulbound: result.UnsoulboundNetworth,
		Profile:     profile,
	}, nil
}
