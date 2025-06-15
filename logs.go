/*
 * Copyright (c) 2022 Avi Misra
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan_v2

import "strconv"

// GetLogs gets logs that match "topic" emitted by the specified "address" between the "fromBlock" and "toBlock"
func (c *Client) GetLogs(fromBlock, toBlock int, address string, topicsList []string, topicsOperatorList []string, page int, offset int) (logs []Log, err error) {
	param := M{
		"fromBlock": fromBlock,
		"toBlock":   toBlock,
	}

	if address != "" {
		param["address"] = address
	}

	for i, topics := range topicsList {
		param["topic"+strconv.Itoa(i)] = topics
	}

	for i, topicsOperator := range topicsOperatorList {
		if topicsOperator == "" {
			continue
		}
		switch i {
		case 0:
			param["topic0_1_opr"] = topicsOperator
		case 1:
			param["topic1_2_opr"] = topicsOperator
		case 2:
			param["topic2_3_opr"] = topicsOperator
		case 3:
			param["topic0_2_opr"] = topicsOperator
		case 4:
			param["topic0_3_opr"] = topicsOperator
		case 5:
			param["topic1_3_opr"] = topicsOperator
		}
	}

	err = c.call("logs", "getLogs", param, &logs)
	return
}
