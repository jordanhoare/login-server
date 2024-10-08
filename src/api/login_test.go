package api

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tibia-oce/login-server/src/api/models"
	"github.com/tibia-oce/login-server/src/grpc/login_proto_messages"
)

var defaultString = "default"
var defaultNumber = uint32(10)

func Test_buildErrorPayloadFromMessage(t *testing.T) {
	type args struct {
		msg *login_proto_messages.LoginResponse
	}
	tests := []struct {
		name string
		args args
		want models.LoginErrorPayload
	}{{
		"default_error_only_message",
		args{&login_proto_messages.LoginResponse{
			Error: &login_proto_messages.Error{
				Code:    10,
				Message: "Failed",
			},
		}},
		models.LoginErrorPayload{
			ErrorCode:    10,
			ErrorMessage: "Failed",
		},
	}, {
		"error_payload_with_more_info",
		args{&login_proto_messages.LoginResponse{
			Error: &login_proto_messages.Error{
				Code:    10,
				Message: "Failed",
			},
			PlayData: &login_proto_messages.PlayData{
				Characters: []*login_proto_messages.Character{
					{WorldId: 0},
					{WorldId: 2},
				},
			},
		}},
		models.LoginErrorPayload{
			ErrorCode:    10,
			ErrorMessage: "Failed",
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildErrorPayloadFromMessage(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildErrorPayloadFromMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildPayloadFromMessage(t *testing.T) {
	type args struct {
		msg *login_proto_messages.LoginResponse
	}
	tests := []struct {
		name string
		args args
		want models.ResponsePayload
	}{{
		name: "default_payload_from_message",
		args: args{&login_proto_messages.LoginResponse{
			Session: &login_proto_messages.Session{
				IsPremium:    false,
				PremiumUntil: 20,
				SessionKey:   "session",
				LastLogin:    0,
			},
			PlayData: &login_proto_messages.PlayData{
				Characters: []*login_proto_messages.Character{{
					WorldId: defaultNumber,
					Info: &login_proto_messages.CharacterInfo{
						Name:      defaultString,
						Vocation:  defaultString,
						Level:     defaultNumber,
						LastLogin: defaultNumber,
						Sex:       defaultNumber,
					},
					Outfit: &login_proto_messages.CharacterOutfit{
						LookType: defaultNumber,
						LookHead: defaultNumber,
						LookBody: defaultNumber,
						LookLegs: defaultNumber,
						LookFeet: defaultNumber,
						Addons:   defaultNumber,
					},
				}},
				Worlds: []*login_proto_messages.World{{
					Id:                         defaultNumber,
					ExternalPort:               defaultNumber,
					ExternalPortProtected:      defaultNumber,
					ExternalPortUnprotected:    defaultNumber,
					Name:                       defaultString,
					ExternalAddress:            defaultString,
					ExternalAddressProtected:   defaultString,
					ExternalAddressUnprotected: defaultString,
					Location:                   defaultString,
				}},
			},
		}},
		want: models.ResponsePayload{
			PlayData: models.PlayData{
				Characters: []models.CharacterPayload{{
					WorldID: defaultNumber,
					CharacterInfo: models.CharacterInfo{
						Name:             defaultString,
						Level:            defaultNumber,
						Vocation:         defaultString,
						IsMale:           false,
						Tutorial:         false,
						IsMainCharacter:  false,
						IsHidden:         false,
						DailyRewardState: 0,
					},
					Outfit: models.Outfit{
						OutfitID:    defaultNumber,
						HeadColor:   defaultNumber,
						TorsoColor:  defaultNumber,
						LegsColor:   defaultNumber,
						DetailColor: defaultNumber,
						AddonsFlags: defaultNumber,
					},
					TournamentInfo: models.TournamentInfo{
						IsTournamentParticipant:          false,
						RemainingDailyTournamentPlayTime: 0,
					},
				}},
				Worlds: []models.World{{
					ID:                         defaultNumber,
					ExternalAddress:            defaultString,
					ExternalAddressProtected:   defaultString,
					ExternalAddressUnprotected: defaultString,
					ExternalPort:               defaultNumber,
					ExternalPortProtected:      defaultNumber,
					ExternalPortUnprotected:    defaultNumber,
					Location:                   defaultString,
					Name:                       defaultString,
					AntiCheatProtection:        false,
					CurrentTournamentPhase:     0,
					IsTournamentWorld:          false,
					PreviewState:               0,
					PvpType:                    0,
					RestrictedStore:            false,
				}},
			},
			Session: models.Session{
				IsPremium:                     false,
				PremiumUntil:                  20,
				SessionKey:                    "session",
				LastLoginTime:                 0,
				EmailCodeRequest:              false,
				FpsTracking:                   false,
				IsReturner:                    false,
				OptionTracking:                false,
				ReturnerNotification:          false,
				ShowRewardNews:                false,
				Status:                        "active",
				TournamentTicketPurchaseState: 0,
				TournamentCyclePhase:          0,
			},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, buildPayloadFromMessage(tt.args.msg))
		})
	}
}
