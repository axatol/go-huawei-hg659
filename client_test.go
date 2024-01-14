package huaweihg659_test

import (
	"context"
	"fmt"
	"testing"

	huaweihg659 "github.com/axatol/go-huawei-hg659"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodePassword(t *testing.T) {
	username := "admin"
	password := "admin"
	csrfParam := "cAMCImb7xhEn6T8eL9cGjIiB4FsfzK0"
	csrfToken := "cl3pAw9VujxiX392oWWObKxqNqVxaOd"
	expected := "e71937a7bfcf970ad7172a3c0702c0d3f68b7c85c60309804d567354fc431222"
	actual := huaweihg659.EncodePassword(username, password, csrfParam, csrfToken)
	assert.Equal(t, expected, actual)
}

func TestClientLogin(t *testing.T) {
	// t.Skip("This test requires a real device")
	ctx := context.Background()
	client, err := huaweihg659.NewClient("http://192.168.1.1")
	require.NoError(t, err)
	err = client.Login(ctx, "admin", "admin")
	require.NoError(t, err)

	fmt.Printf("session id: %s\n", client.SessionID())
	var data any

	// public
	data, err = client.GetBroadcastedWifiNetworks(ctx)
	fmt.Printf("GetBroadcastedWifiNetworks: %+v, %+v\n", data, err)
	require.NoError(t, err)
	data, err = client.GetConnectedDeviceCount(ctx)
	fmt.Printf("GetConnectedDeviceCount: %+v, %+v\n", data, err)
	require.NoError(t, err)
	data, err = client.GetDeviceInfo(ctx)
	fmt.Printf("GetDeviceInfo: %+v, %+v\n", data, err)
	require.NoError(t, err)
	data, err = client.GetInternetDiagnosis(ctx)
	fmt.Printf("GetInternetDiagnosis: %+v, %+v\n", data, err)
	require.NoError(t, err)

	// maintain
	data, err = client.GetLANStatistics(ctx)
	fmt.Printf("GetLANStatistics: %+v, %+v\n", data, err)
	require.NoError(t, err)
	data, err = client.GetUMTSStatistics(ctx)
	fmt.Printf("GetUMTSStatistics: %+v, %+v\n", data, err)
	require.NoError(t, err)
	data, err = client.GetWifiStatistics(ctx)
	fmt.Printf("GetWifiStatistics: %+v, %+v\n", data, err)
	require.NoError(t, err)
	data, err = client.ListDHCPReservation(ctx)
	fmt.Printf("ListDHCPReservation: %+v, %+v\n", data, err)
	require.NoError(t, err)

	// home network
	data, err = client.ListKnownLANDevices(ctx)
	fmt.Printf("ListKnownLANDevices: %+v, %+v\n", data, err)
	require.NoError(t, err)
}
