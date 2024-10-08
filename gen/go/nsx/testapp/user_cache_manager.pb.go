// Code generated by protoc-gen-go-cache-manager. DO NOT EDIT.

package testapp

import (
	"context"
	"fmt"

	"github.com/NSXBet/protoc-gen-go-cache-manager/pkg/gocachemanager"
)

// UserCacheManager for every operation related to this service:
// UserCache is the service that will be used to cache user details.
type UserCacheManager struct {
	userCacheManager_UserDetails *gocachemanager.CacheManager[*UserDetailsRequest, *UserDetailsResponse]
}

// NewUserCacheManager is the constructor method for this service:
// UserCache is the service that will be used to cache user details.
// Required Update Method(s):
//   - updateUserDetailsFn is a function that loads the data from the storage and you
//     can pass any dependencies required to resolve it when calling the Get and Refresh methods
//     and these will be available as the third argument of the method.
func NewUserCacheManager(
	updateUserDetailsFn func(context.Context, *UserDetailsRequest, map[string]any) (*UserDetailsResponse, error),
	options ...gocachemanager.CacheOption,
) (*UserCacheManager, error) {
	userCacheManager_UserDetails, err := gocachemanager.NewCacheManager(
		"usercache::userdetails",
		func() *UserDetailsResponse { return &UserDetailsResponse{} },
		updateUserDetailsFn,
		options...,
	)
	if err != nil {
		return nil, fmt.Errorf("creating cache manager %s: %w", "UserDetails", err)
	}

	return &UserCacheManager{
		userCacheManager_UserDetails: userCacheManager_UserDetails,
	}, nil
}

// GetUserDetails returns the user details for the given user_id from the cache.
// This method is a test of a multi-line comment.
// It should not break other lines.
func (cm *UserCacheManager) GetUserDetails(
	ctx context.Context,
	input *UserDetailsRequest,
	dependencies ...map[string]any,
) (*UserDetailsResponse, error) {
	return cm.userCacheManager_UserDetails.Get(ctx, input, dependencies...)
}

// Eagerly Refresh the cache for the method that:
// UserDetails returns the user details for the given user_id from the cache.
// This method is a test of a multi-line comment.
// It should not break other lines.
func (cm *UserCacheManager) RefreshUserDetails(
	ctx context.Context,
	input *UserDetailsRequest,
	dependencies ...map[string]any,
) (*UserDetailsResponse, error) {
	return cm.userCacheManager_UserDetails.Refresh(
		ctx,
		input,
		dependencies...,
	)
}

// Eagerly Replace the cache for the method that:
// UserDetails returns the user details for the given user_id from the cache.
// This method is a test of a multi-line comment.
// It should not break other lines.
func (cm *UserCacheManager) ReplaceUserDetails(
	ctx context.Context,
	input *UserDetailsRequest,
	newValue *UserDetailsResponse,
	dependencies ...map[string]any,
) (*UserDetailsResponse, error) {
	return cm.userCacheManager_UserDetails.Replace(
		ctx,
		input,
		newValue,
		dependencies...,
	)
}

// Eagerly Delete the cache for the method that:
// UserDetails returns the user details for the given user_id from the cache.
// This method is a test of a multi-line comment.
// It should not break other lines.
func (cm *UserCacheManager) DeleteUserDetails(
	ctx context.Context,
	input *UserDetailsRequest,
	dependencies ...map[string]any,
) error {
	return cm.userCacheManager_UserDetails.Delete(
		ctx,
		input,
		dependencies...,
	)
}

type TournamentCacheManager struct {
	tournamentCacheManager_MainTournaments *gocachemanager.CacheManager[*MainTournamentsRequest, *MainTournamentsResponse]
}

// Required Update Method(s):
// - updateMainTournamentsFn is a function that loads the data from the storage and you
//    can pass any dependencies required to resolve it when calling the Get and Refresh methods
//    and these will be available as the third argument of the method.
func NewTournamentCacheManager(
	updateMainTournamentsFn func(context.Context, *MainTournamentsRequest, map[string]any) (*MainTournamentsResponse, error),
	options ...gocachemanager.CacheOption,
) (*TournamentCacheManager, error) {
	tournamentCacheManager_MainTournaments, err := gocachemanager.NewCacheManager(
		"tournamentcache::maintournaments",
		func() *MainTournamentsResponse { return &MainTournamentsResponse{} },
		updateMainTournamentsFn,
		options...,
	)
	if err != nil {
		return nil, fmt.Errorf("creating cache manager %s: %w", "MainTournaments", err)
	}

	return &TournamentCacheManager{
		tournamentCacheManager_MainTournaments: tournamentCacheManager_MainTournaments,
	}, nil
}

func (cm *TournamentCacheManager) GetMainTournaments(
	ctx context.Context,
	input *MainTournamentsRequest,
	dependencies ...map[string]any,
) (*MainTournamentsResponse, error) {
	return cm.tournamentCacheManager_MainTournaments.Get(ctx, input, dependencies...)
}

func (cm *TournamentCacheManager) RefreshMainTournaments(
	ctx context.Context,
	input *MainTournamentsRequest,
	dependencies ...map[string]any,
) (*MainTournamentsResponse, error) {
	return cm.tournamentCacheManager_MainTournaments.Refresh(
		ctx,
		input,
		dependencies...,
	)
}

func (cm *TournamentCacheManager) ReplaceMainTournaments(
	ctx context.Context,
	input *MainTournamentsRequest,
	newValue *MainTournamentsResponse,
	dependencies ...map[string]any,
) (*MainTournamentsResponse, error) {
	return cm.tournamentCacheManager_MainTournaments.Replace(
		ctx,
		input,
		newValue,
		dependencies...,
	)
}

func (cm *TournamentCacheManager) DeleteMainTournaments(
	ctx context.Context,
	input *MainTournamentsRequest,
	dependencies ...map[string]any,
) error {
	return cm.tournamentCacheManager_MainTournaments.Delete(
		ctx,
		input,
		dependencies...,
	)
}
