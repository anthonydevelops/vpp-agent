// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vppcalls

import (
	"bytes"
	"net"
	"strconv"
	"time"

	ipsecapi "github.com/ligato/vpp-agent/plugins/vpp/binapi/ipsec"
	"github.com/ligato/vpp-agent/plugins/vpp/model/interfaces"
	"github.com/ligato/vpp-agent/plugins/vpp/model/ipsec"
)

// IPSecSaDetails holds security association with VPP metadata
type IPSecSaDetails struct {
	Sa   *ipsec.SecurityAssociations_SA
	Meta *IPSecSaMeta
}

// IPSecSaMeta contains all VPP-specific metadata
type IPSecSaMeta struct {
	SaID           uint32
	Interface      string
	IfIdx          uint32
	CryptoKeyLen   uint8
	IntegKeyLen    uint8
	Salt           uint32
	SeqOutbound    uint64
	LastSeqInbound uint64
	ReplayWindow   uint64
	TotalDataSize  uint64
}

func (handler *ipSecVppHandler) DumpIPSecSA() (saList []*IPSecSaDetails, err error) {
	return handler.DumpIPSecSAWithIndex(^uint32(0)) // Get everything
}

func (handler *ipSecVppHandler) DumpIPSecSAWithIndex(saID uint32) (saList []*IPSecSaDetails, err error) {
	defer func(t time.Time) {
		handler.stopwatch.TimeLog(ipsecapi.IpsecSaDump{}).LogTimeEntry(time.Since(t))
	}(time.Now())

	saDetails, err := handler.dumpSecurityAssociations(saID)
	if err != nil {
		return nil, err
	}

	for _, saData := range saDetails {
		// Skip tunnel interfaces
		if saData.SwIfIndex != ^uint32(0) {
			continue
		}

		// Addresses
		var tunnelSrcAddrStr, tunnelDstAddrStr string
		if uintToBool(saData.IsTunnelIP6) {
			var tunnelSrcAddr, tunnelDstAddr net.IP = saData.TunnelSrcAddr, saData.TunnelDstAddr
			tunnelSrcAddrStr, tunnelDstAddrStr = tunnelSrcAddr.String(), tunnelDstAddr.String()
		} else {
			var tunnelSrcAddr, tunnelDstAddr net.IP = saData.TunnelSrcAddr[:4], saData.TunnelDstAddr[:4]
			tunnelSrcAddrStr, tunnelDstAddrStr = tunnelSrcAddr.String(), tunnelDstAddr.String()
		}

		sa := &ipsec.SecurityAssociations_SA{
			Spi:            saData.Spi,
			Protocol:       getSaProto(saData.Protocol),
			CryptoAlg:      getCryptoAlg(saData.CryptoAlg),
			CryptoKey:      string(bytes.SplitN(saData.CryptoKey, []byte{0x00}, 2)[0]),
			IntegAlg:       getIntegAlg(saData.IntegAlg),
			IntegKey:       string(bytes.SplitN(saData.IntegKey, []byte{0x00}, 2)[0]),
			UseEsn:         uintToBool(saData.UseEsn),
			UseAntiReplay:  uintToBool(saData.UseAntiReplay),
			TunnelSrcAddr:  tunnelSrcAddrStr,
			TunnelDstAddr:  tunnelDstAddrStr,
			EnableUdpEncap: uintToBool(saData.UDPEncap),
		}
		meta := &IPSecSaMeta{
			SaID:           saData.SaID,
			IfIdx:          saData.SwIfIndex,
			CryptoKeyLen:   saData.CryptoKeyLen,
			IntegKeyLen:    saData.IntegKeyLen,
			Salt:           saData.Salt,
			SeqOutbound:    saData.SeqOutbound,
			LastSeqInbound: saData.LastSeqInbound,
			ReplayWindow:   saData.ReplayWindow,
			TotalDataSize:  saData.TotalDataSize,
		}
		saList = append(saList, &IPSecSaDetails{
			Sa:   sa,
			Meta: meta,
		})
	}

	return saList, nil
}

// IPSecTunnelInterfaceDetails hold a list of tunnel interfaces with name/index map as metadata
type IPSecTunnelInterfaceDetails struct {
	Tunnels []*ipsec.TunnelInterfaces_Tunnel
	Meta    *IPSecTunnelMeta
}

// IPSecTunnelMeta contains map of name/index pairs
type IPSecTunnelMeta struct {
	IfNameToIdx map[uint32]string
}

func (handler *ipSecVppHandler) DumpIPSecTunnelInterfaces() (tun *IPSecTunnelInterfaceDetails, err error) {
	defer func(t time.Time) {
		handler.stopwatch.TimeLog(ipsecapi.IpsecSaDump{}).LogTimeEntry(time.Since(t))
	}(time.Now())

	var tunnels []*ipsec.TunnelInterfaces_Tunnel
	meta := &IPSecTunnelMeta{
		IfNameToIdx: make(map[uint32]string),
	}
	saDetails, err := handler.dumpSecurityAssociations(^uint32(0))
	if err != nil {
		return nil, err
	}

	for _, saData := range saDetails {
		// Skip non-tunnel security associations
		if saData.SwIfIndex == ^uint32(0) {
			continue
		}

		// Interface
		var ifName string
		var ifData *interfaces.Interfaces_Interface
		if saData.SwIfIndex != ^uint32(1) {
			var found bool
			ifName, ifData, found = handler.ifIndexes.LookupName(saData.SwIfIndex)
			if !found {
				handler.log.Warnf("IPSec SA dump: interface name not found for %d", saData.SwIfIndex)
				continue
			}
			if ifData == nil {
				handler.log.Warnf("IPSec SA dump: interface %s has no metadata", ifName)
				continue
			}
		}

		// Addresses
		var tunnelSrcAddrStr, tunnelDstAddrStr string
		if uintToBool(saData.IsTunnelIP6) {
			var tunnelSrcAddr, tunnelDstAddr net.IP = saData.TunnelSrcAddr, saData.TunnelDstAddr
			tunnelSrcAddrStr, tunnelDstAddrStr = tunnelSrcAddr.String(), tunnelDstAddr.String()
		} else {
			var tunnelSrcAddr, tunnelDstAddr net.IP = saData.TunnelSrcAddr[:4], saData.TunnelDstAddr[:4]
			tunnelSrcAddrStr, tunnelDstAddrStr = tunnelSrcAddr.String(), tunnelDstAddr.String()
		}

		// Prepare tunnel interface data
		tunnel := &ipsec.TunnelInterfaces_Tunnel{
			Name:        ifName,
			Esn:         uintToBool(saData.UseEsn),
			AntiReplay:  uintToBool(saData.UseAntiReplay),
			LocalIp:     tunnelSrcAddrStr,
			RemoteIp:    tunnelDstAddrStr,
			LocalSpi:    saData.Spi,
			RemoteSpi:   saData.Spi,
			CryptoAlg:   getCryptoAlg(saData.CryptoAlg),
			IntegAlg:    getIntegAlg(saData.IntegAlg),
			Enabled:     ifData.Enabled,
			IpAddresses: ifData.IpAddresses,
			Vrf:         ifData.Vrf,
		}
		tunnels = append(tunnels, tunnel)

		// Put metadata entry
		meta.IfNameToIdx[saData.SwIfIndex] = ifName
	}

	return &IPSecTunnelInterfaceDetails{
		Tunnels: tunnels,
		Meta:    meta,
	}, nil
}

// IPSecSpdDetails represents IPSec policy databases with particular metadata
type IPSecSpdDetails struct {
	Spd  *ipsec.SecurityPolicyDatabases_SPD
	Meta *IPSecSpdMeta
}

// IPSecSpdMeta is map where key is a generated security association name, and value is an SpdMeta object
type IPSecSpdMeta struct {
	SpdMeta map[string]*SpdMeta // SA-generated name is a key
}

// SpdMeta hold VPP-specific data related to SPD
type SpdMeta struct {
	SaID    uint32
	SpdID   uint32
	Policy  uint8
	Bytes   uint64
	Packets uint64
}

func (handler *ipSecVppHandler) DumpIPSecSPD() (spdList []*IPSecSpdDetails, err error) {
	defer func(t time.Time) {
		handler.stopwatch.TimeLog(ipsecapi.IpsecSpdDump{}).LogTimeEntry(time.Since(t))
	}(time.Now())

	metadata := &IPSecSpdMeta{
		SpdMeta: make(map[string]*SpdMeta),
	}

	// TODO IPSec SPD dump request requires SPD ID, otherwise it returns nothing. There is currently no way
	// to dump all SPDs available on the VPP, so let's dump at least the ones configurator knows about.
	for _, spdName := range handler.spdIndexes.GetMapping().ListNames() {
		spdIdx, _, found := handler.spdIndexes.LookupIdx(spdName)
		if !found {
			// Shouldn't happen, call the police or something
			continue
		}
		spd := &ipsec.SecurityPolicyDatabases_SPD{}

		// Prepare VPP binapi request
		req := &ipsecapi.IpsecSpdDump{
			SpdID: spdIdx,
			SaID:  0xffffffff,
		}
		requestCtx := handler.callsChannel.SendMultiRequest(req)

		// Policy association index, used to generate SA name
		var paIdx int

		for {
			spdDetails := &ipsecapi.IpsecSpdDetails{}
			stop, err := requestCtx.ReceiveReply(spdDetails)
			if stop {
				break
			}
			if err != nil {
				return nil, err
			}

			// Security association name, to distinguish metadata
			saGenName := "sa-generated-" + strconv.Itoa(paIdx)
			paIdx++

			// Addresses
			var remoteStartAddrStr, remoteStopAddrStr, localStartAddrStr, localStopAddrStr string
			if uintToBool(spdDetails.IsIpv6) {
				var remoteStartAddr, remoteStopAddr net.IP = spdDetails.RemoteStartAddr, spdDetails.RemoteStopAddr
				remoteStartAddrStr, remoteStopAddrStr = remoteStartAddr.String(), remoteStopAddr.String()
				var localStartAddr, localStopAddr net.IP = spdDetails.LocalStartAddr, spdDetails.LocalStopAddr
				localStartAddrStr, localStopAddrStr = localStartAddr.String(), localStopAddr.String()
			} else {
				var remoteStartAddr, remoteStopAddr net.IP = spdDetails.RemoteStartAddr[:4], spdDetails.RemoteStopAddr[:4]
				remoteStartAddrStr, remoteStopAddrStr = remoteStartAddr.String(), remoteStopAddr.String()
				var localStartAddr, localStopAddr net.IP = spdDetails.LocalStartAddr[:4], spdDetails.LocalStopAddr[:4]
				localStartAddrStr, localStopAddrStr = localStartAddr.String(), localStopAddr.String()
			}

			// Prepare policy entry and put to the SPD
			policyEntry := &ipsec.SecurityPolicyDatabases_SPD_PolicyEntry{
				Sa:              saGenName,
				Priority:        spdDetails.Priority,
				IsOutbound:      uintToBool(spdDetails.IsOutbound),
				RemoteAddrStart: remoteStartAddrStr,
				RemoteAddrStop:  remoteStopAddrStr,
				LocalAddrStart:  localStartAddrStr,
				LocalAddrStop:   localStopAddrStr,
				Protocol:        uint32(spdDetails.Protocol),
				RemotePortStart: uint32(spdDetails.RemoteStartPort),
				RemotePortStop:  uint32(spdDetails.RemoteStopPort),
				LocalPortStart:  uint32(spdDetails.LocalStartPort),
				LocalPortStop:   uint32(spdDetails.LocalStopPort),
			}
			spd.PolicyEntries = append(spd.PolicyEntries, policyEntry)

			// Prepare meta and put to the metadata map
			meta := &SpdMeta{
				SpdID:   spdDetails.SpdID,
				SaID:    spdDetails.SaID,
				Policy:  spdDetails.Policy,
				Bytes:   spdDetails.Bytes,
				Packets: spdDetails.Packets,
			}
			metadata.SpdMeta[saGenName] = meta
		}
		// Store STD in list
		spdList = append(spdList, &IPSecSpdDetails{
			Spd:  spd,
			Meta: metadata,
		})
	}

	return spdList, nil
}

// Get all security association (used also for tunnel interfaces) in binary api format
func (handler *ipSecVppHandler) dumpSecurityAssociations(saID uint32) (saList []*ipsecapi.IpsecSaDetails, err error) {
	defer func(t time.Time) {
		handler.stopwatch.TimeLog(ipsecapi.IpsecSaDump{}).LogTimeEntry(time.Since(t))
	}(time.Now())

	req := &ipsecapi.IpsecSaDump{
		SaID: saID,
	}
	requestCtx := handler.callsChannel.SendMultiRequest(req)

	for {
		saDetails := &ipsecapi.IpsecSaDetails{}
		stop, err := requestCtx.ReceiveReply(saDetails)
		if stop {
			break
		}
		if err != nil {
			return nil, err
		}

		saList = append(saList, saDetails)
	}

	return saList, nil

}

func getCryptoAlg(alg uint8) ipsec.CryptoAlgorithm {
	switch alg {
	case 0:
		return ipsec.CryptoAlgorithm_NONE_CRYPTO
	case 1:
		return ipsec.CryptoAlgorithm_AES_CBC_128
	case 2:
		return ipsec.CryptoAlgorithm_AES_CBC_192
	case 3:
		return ipsec.CryptoAlgorithm_AES_CBC_256
	default:
		return ipsec.CryptoAlgorithm_NONE_CRYPTO // As default
	}
}

func getIntegAlg(alg uint8) ipsec.IntegAlgorithm {
	switch alg {
	case 0:
		return ipsec.IntegAlgorithm_NONE_INTEG
	case 1:
		return ipsec.IntegAlgorithm_MD5_96
	case 2:
		return ipsec.IntegAlgorithm_SHA1_96
	case 3:
		return ipsec.IntegAlgorithm_SHA_256_96
	case 4:
		return ipsec.IntegAlgorithm_SHA_256_128
	case 5:
		return ipsec.IntegAlgorithm_SHA_384_192
	case 6:
		return ipsec.IntegAlgorithm_SHA_512_256
	default:
		return ipsec.IntegAlgorithm_NONE_INTEG // As default
	}
}

func getSaProto(protocol uint8) ipsec.SecurityAssociations_SA_IPSecProtocol {
	if protocol == 0 {
		return ipsec.SecurityAssociations_SA_AH
	}
	return ipsec.SecurityAssociations_SA_ESP
}

func uintToBool(input uint8) bool {
	if input == 1 {
		return true
	}
	return false
}
