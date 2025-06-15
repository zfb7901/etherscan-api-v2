package etherscan_v2

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_GetLogs(t *testing.T) {
	expectedLogs := []Log{
		Log{
			Address:         "0x33990122638b9132ca29c723bdf037f1a891a70c",
			Topics:          []string{"0xf63780e752c6a54a94fc52715dbc5518a3b4c3c2833d301a204226548a2a8545", "0x72657075746174696f6e00000000000000000000000000000000000000000000", "0x000000000000000000000000d9b2f59f3b5c7b3c67047d2f03c3e8052470be92"},
			Data:            "0x",
			BlockNumber:     "0x5c958",
			BlockHash:       "0xe32a9cac27f823b18454e8d69437d2af41a1b81179c6af2601f1040a72ad444b",
			TransactionHash: "0x0b03498648ae2da924f961dda00dc6bb0a8df15519262b7e012b7d67f4bb7e83",
			LogIndex:        "0x",
		},
	}

	actualLogs, err := api.GetLogs(54772722, 54772722, "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd", []string{"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}, []string{}, 1, 100)

	noError(t, err, "api.GetLogs")

	equal := cmp.Equal(expectedLogs, actualLogs)

	if !equal {
		t.Errorf("api.GetLogs not working\n: %s\n", cmp.Diff(expectedLogs, actualLogs))
	}
}
