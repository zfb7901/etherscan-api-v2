/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan_v2

import (
	"fmt"
	"testing"
)

func TestClient_rpcProxy(t *testing.T) {
	params := map[string]interface{}{
		"tag":     fmt.Sprintf("0x%x", 9138018),
		"boolean": true,
	}
	response, err := api.RpcProxy("eth_getBlockByNumber", params)
	if err != nil {
		fmt.Println(response, err, "api.RpcProxy")
	}
}
