package testutils

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/xmtp/xmtpd/pkg/abis"
	"github.com/xmtp/xmtpd/pkg/utils"
)

// Build an abi encoded MessageSent event struct
func BuildMessageSentEvent(
	groupID [32]byte,
	message []byte,
	sequenceID uint64,
) ([]byte, error) {
	abi, err := abis.GroupMessagesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return abi.Events["MessageSent"].Inputs.Pack(groupID, message, sequenceID)
}

// Build a log message for a MessageSent event
func BuildMessageSentLog(
	t *testing.T,
	groupID [32]byte,
	message []byte,
	sequenceID uint64,
) types.Log {
	eventData, err := BuildMessageSentEvent(groupID, message, sequenceID)
	require.NoError(t, err)

	abi, err := abis.GroupMessagesMetaData.GetAbi()
	require.NoError(t, err)

	topic, err := utils.GetEventTopic(abi, "MessageSent")
	require.NoError(t, err)

	return types.Log{
		Topics: []common.Hash{topic},
		Data:   eventData,
	}
}

/*
*
Deploy a contract and return the contract's address. Will return a different address for each run, making it suitable for testing
*
*/
func deployContract(t *testing.T, filePath string, contractName string) string {
	packageRoot := rootPath(t)
	cmd := exec.Command("./dev/contracts/deploy-ephemeral", filePath, contractName)
	cmd.Dir = packageRoot

	var out bytes.Buffer
	var errors bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errors

	err := cmd.Run()
	require.NoError(t, err, "Failed to execute deploy-ephemeral script", errors.String())

	output := out.String()
	t.Logf("deploy-ephemeral output: %s", output)

	var parsed contractInfo
	require.NoError(t, json.Unmarshal([]byte(output), &parsed))

	return parsed.DeployedTo
}

func DeployNodesContract(t *testing.T) string {
	return deployContract(t, "./src/Nodes.sol", "Nodes")
}

func DeployGroupMessagesContract(t *testing.T) string {
	return deployContract(t, "./src/GroupMessages.sol", "GroupMessages")
}

func DeployIdentityUpdatesContract(t *testing.T) string {
	return deployContract(t, "./src/IdentityUpdates.sol", "IdentityUpdates")
}
