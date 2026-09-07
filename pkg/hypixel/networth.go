package hypixel

import (
	"context"
	"fmt"
)

type Networth struct {
	Total       float64
	Unsoulbound float64
	Profile     *SkyBlockProfile
}

func (c *Client) GetNetworth(ctx context.Context, username, profileName string) (*Networth, error) {
	player, rawProfile, err := c.GetRawProfile(ctx, username, profileName)
	if err != nil {
		return nil, err
	}

	profile, err := parseRawProfile(player, rawProfile)
	if err != nil {
		return nil, err
	}

	if profile.Data == nil {
		return nil, fmt.Errorf("member %s not found in profile %s", player.ID, profile.Name)
	}

	museum, err := c.getRawMuseum(ctx, profile.ID)
	if err != nil {
		return nil, err
	}

	result, err := c.compat.Networth(ctx, rawProfile, museum, profile.Mojang.ID)
	if err != nil {
		return nil, err
	}

	return &Networth{
		Total:       result.Networth,
		Unsoulbound: result.UnsoulboundNetworth,
		Profile:     profile,
	}, nil
}
