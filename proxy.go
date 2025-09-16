/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan_v2

// RpcProxy AccountBalance gets ether balance for a single address
func (c *Client) RpcProxy(action string, param map[string]interface{}) ([]byte, error) {
	var result []byte
	err := c.callRpc("proxy", action, param, &result)
	return result, err
}
