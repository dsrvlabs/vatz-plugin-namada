package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"
	"os/exec"
	"strings"
)

const (
	// Default values.
	defaultAddr     = "127.0.0.1"
	defaultPort     = 10001
	defaultBaseDir  = "/mnt/namada"
	defaultChainId  = "shielded-expedition.88f17d1d14"
	defaultNodeInfo = "http://0.0.0.0:26657"
	pluginName      = "node-governance-alarm"
)

type Proposal struct {
	ProposalID int    `json:"proposal_id"`
	Type       string `json:"type"`
	Author     string `json:"author"`
	StartEpoch int    `json:"start_epoch"`
	EndEpoch   int    `json:"end_epoch"`
	GraceEpoch int    `json:"grace_epoch"`
}

var (
	addr         string
	port         int
	chainId      string
	baseDir      string
	nodeInfo     string
	lastProposal Proposal
)

func init() {
	flag.StringVar(&addr, "addr", defaultAddr, "IP Address(e.g. 0.0.0.0, 127.0.0.1)")
	flag.IntVar(&port, "port", defaultPort, "Port number")
	flag.StringVar(&chainId, "chainId", defaultChainId, "Need to Chain ID")
	flag.StringVar(&baseDir, "baseDir", defaultBaseDir, "Need to Base Dir")
	flag.StringVar(&nodeInfo, "nodeInfo", defaultNodeInfo, "Need to know Node Info")
	flag.Parse()

	initialProposalsStr, err := GetLatestProposalList()
	if err != nil {
		panic(err)
	}
	initialProposals, err := ParseProposals(initialProposalsStr)
	lastProposal = initialProposals[len(initialProposals)-1]
}

func GetLatestProposalList() (string, error) {
	cmd := exec.Command("namadac", "--base-dir", baseDir, "query-proposal", "--chain-id", chainId, "--node", nodeInfo)
	// Create a buffer to capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	// Return the output as a string
	return out.String(), nil
}

func ParseProposals(input string) ([]Proposal, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var (
		proposals []Proposal
		proposal  Proposal
	)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Proposal Id:") {
			if proposal.ProposalID != 0 {
				proposals = append(proposals, proposal)
				proposal = Proposal{} // reset the proposal struct for the next one
			}
			fmt.Sscanf(line, "Proposal Id: %d", &proposal.ProposalID)
		} else if strings.HasPrefix(line, "Type:") {
			fmt.Sscanf(line, "Type: %s", &proposal.Type)
		} else if strings.HasPrefix(line, "Author:") {
			fmt.Sscanf(line, "Author: %s", &proposal.Author)
		} else if strings.HasPrefix(line, "Start Epoch:") {
			fmt.Sscanf(line, "Start Epoch: %d", &proposal.StartEpoch)
		} else if strings.HasPrefix(line, "End Epoch:") {
			fmt.Sscanf(line, "End Epoch: %d", &proposal.EndEpoch)
		} else if strings.HasPrefix(line, "Grace Epoch:") {
			fmt.Sscanf(line, "Grace Epoch: %d", &proposal.GraceEpoch)
		}
	}
	// Append the last proposal if it exists
	if proposal.ProposalID != 0 {
		proposals = append(proposals, proposal)
	}

	return proposals, scanner.Err()
}

func main() {
	p := sdk.NewPlugin(pluginName)
	p.Register(pluginFeature)

	ctx := context.Background()
	if err := p.Start(ctx, addr, port); err != nil {
		fmt.Println("exit")
	}
}

func pluginFeature(info, option map[string]*structpb.Value) (sdk.CallResponse, error) {
	severity := pluginpb.SEVERITY_INFO
	state := pluginpb.STATE_NONE
	latestProposalAnnounced := false
	var msg string

	nextProposalsStr, err := GetLatestProposalList()
	if err != nil {
		panic(err)
	}

	nextProposals, err := ParseProposals(nextProposalsStr)
	currentLastProposal := nextProposals[len(nextProposals)-1]
	if currentLastProposal.ProposalID > lastProposal.ProposalID {
		lastProposal = currentLastProposal
		latestProposalAnnounced = true
	}

	if latestProposalAnnounced {
		log.Info().Str("module", "plugin").Msgf("INFO : new proposal is released - proposalId: %d", lastProposal.ProposalID)
		severity = pluginpb.SEVERITY_CRITICAL
		state = pluginpb.STATE_SUCCESS
		msg = fmt.Sprintf("New Release Proposal Info - Proposal ID: %d\nType: %s\nAuthor: %s\nStart Epoch: %d\nEnd Epoch: %d\nGrace Epoch: %d",
			lastProposal.ProposalID, lastProposal.Type, lastProposal.Author, lastProposal.StartEpoch, lastProposal.EndEpoch, lastProposal.GraceEpoch)
	} else {
		log.Debug().Str("module", "plugin").Msg("DEBUG : tmp != proposalId")
		severity = pluginpb.SEVERITY_INFO
		state = pluginpb.STATE_SUCCESS
		msg = fmt.Sprintf("Last Proposal ID: %d\nType: %s\nAuthor: %s\nStart Epoch: %d\nEnd Epoch: %d\nGrace Epoch: %d",
			lastProposal.ProposalID, lastProposal.Type, lastProposal.Author, lastProposal.StartEpoch, lastProposal.EndEpoch, lastProposal.GraceEpoch)
	}
	log.Info().Str("module", "plugin").Msg(msg)
	ret := sdk.CallResponse{
		FuncName: info["execute_method"].GetStringValue(),
		Message:  msg,
		Severity: severity,
		State:    state,
	}
	return ret, nil
}
