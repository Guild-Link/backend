package hypixel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	sc "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/guild-link/backend/pkg/mojang"
)

type SkyBlockProfile struct {
	ID       string
	Name     string
	GameMode string
	Selected bool
	Player   *mojang.Profile
	Raw      json.RawMessage

	Data              *sc.Member
	Banking           *sc.Banking
	CommunityUpgrades *sc.CommunityUpgrades
}

func (c *Client) GetProfilesRaw(ctx context.Context, username string) (*mojang.Profile, []json.RawMessage, error) {
	profile, err := c.mojang.GetProfile(ctx, username)
	if err != nil {
		return nil, nil, err
	}

	body, err := c.get(ctx, "/skyblock/profiles"+"?uuid="+url.QueryEscape(profile.ID))
	if err != nil {
		return nil, nil, err
	}

	var data struct {
		Profiles []json.RawMessage `json:"profiles"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, nil, fmt.Errorf("decode Hypixel profiles response: %w", err)
	}

	return profile, data.Profiles, nil
}

func (c *Client) GetSkyBlockProfiles(ctx context.Context, username string) ([]SkyBlockProfile, error) {
	player, rawProfiles, err := c.GetProfilesRaw(ctx, username)
	if err != nil {
		return nil, err
	}

	data := make([]sc.Profile, len(rawProfiles))
	for i := range rawProfiles {
		if err := json.Unmarshal(rawProfiles[i], &data[i]); err != nil {
			return nil, fmt.Errorf("decode Hypixel profile: %w", err)
		}
	}

	profiles := make([]SkyBlockProfile, 0, len(data))
	for i := range data {
		var memberData *sc.Member
		p := &data[i]

		member, ok := p.Members[player.ID]
		if ok {
			memberData = &member
		}

		profiles = append(profiles, SkyBlockProfile{
			ID:                p.ProfileID,
			Name:              p.CuteName,
			Player:            player,
			Raw:               rawProfiles[i],
			GameMode:          p.GameMode,
			Selected:          p.Selected,
			Data:              memberData,
			Banking:           p.Banking,
			CommunityUpgrades: p.CommunityUpgrades,
		})
	}

	return profiles, nil
}

func (c *Client) GetSkyBlockProfile(ctx context.Context, username, profileName string) (*SkyBlockProfile, error) {
	profiles, err := c.GetSkyBlockProfiles(ctx, username)
	if err != nil {
		return nil, err
	}

	if profileName == "" {
		for _, profile := range profiles {
			if !profile.Selected {
				continue
			}

			if profile.Data == nil {
				return nil, fmt.Errorf("member %s not found in selected profile", profile.Player.ID)
			}

			return &profile, nil
		}

		return nil, fmt.Errorf("selected profile not found")
	}

	for _, profile := range profiles {
		if strings.EqualFold(profile.Name, profileName) {
			if profile.Data == nil {
				return nil, fmt.Errorf("member %s not found in profile %s", profile.Player.ID, profileName)
			}
			return &profile, nil
		}
	}

	return nil, fmt.Errorf("profile %s not found", profileName)
}
