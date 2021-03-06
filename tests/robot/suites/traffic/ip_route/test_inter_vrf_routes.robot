*** Settings ***

Library     OperatingSystem
Library     String
#Library     RequestsLibrary

Resource     ../../../variables/${VARIABLES}_variables.robot
Resource    ../../../libraries/all_libs.robot
Resource    ../../../libraries/pretty_keywords.robot

Force Tags        traffic     IPv4
Suite Setup       Run Keywords    Discard old results     Test Setup
Suite Teardown    Test Teardown

*** Variables ***
${VARIABLES}=          common
${ENV}=                common
${WAIT_TIMEOUT}=       20s
${SYNC_SLEEP}=         2s
${FINAL_SLEEP}=        1s
${IP_1}=               10.1.1.1
${IP_2}=               10.1.1.2
${IP_3}=               10.1.2.1
${IP_4}=               10.1.2.2
${NET1}=               10.1.1.0
${NET2}=               10.1.2.0

*** Test Cases ***
# Non default VRF table 2 used in Agent VPP Node agent_vpp_2
Start Two Agents And Then Configure With Default And Non Default VRF
    Add Agent VPP Node    agent_vpp_1
    Add Agent VPP Node    agent_vpp_2

    Create Master memif0 on agent_vpp_1 with IP ${IP_1}, MAC 02:f1:be:90:00:00, key 1 and m0.sock socket
    Create Slave memif0 on agent_vpp_2 with IP ${IP_2}, MAC 02:f1:be:90:00:02, key 1 and m0.sock socket

    Create Master memif1 on agent_vpp_1 with VRF 2, IP ${IP_3}, MAC 02:f1:be:90:02:00, key 1 and m1.sock socket
    Create Slave memif1 on agent_vpp_2 with VRF 2, IP ${IP_4}, MAC 02:f1:be:90:02:02, key 1 and m1.sock socket

    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    List of interfaces On agent_vpp_1 Should Contain Interface memif1/1
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    List of interfaces On agent_vpp_2 Should Contain Interface memif1/1
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    List of interfaces On agent_vpp_1 Should Contain Interface memif2/1
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    List of interfaces On agent_vpp_2 Should Contain Interface memif2/1
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    IP Fib Table 0 On agent_vpp_1 Should Contain Route With IP ${IP_1}/32
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    IP Fib Table 2 On agent_vpp_1 Should Contain Route With IP ${IP_3}/32
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    IP Fib Table 0 On agent_vpp_2 Should Contain Route With IP ${IP_2}/32
    Wait Until Keyword Succeeds   ${WAIT_TIMEOUT}   ${SYNC_SLEEP}    IP Fib Table 2 On agent_vpp_2 Should Contain Route With IP ${IP_4}/32


Check Normal Ping Inside VRF
    # try ping
    Ping From agent_vpp_1 To ${IP_2}
    Ping From agent_vpp_2 To ${IP_1}

Ping From Source Agent 1
    Ping On agent_vpp_1 With IP ${IP_2}, Source memif1/1
    Ping On agent_vpp_1 With IP ${IP_4}, Source memif2/1

Ping From Source Agent 2
#    ${int}=    Get Interface Internal Name    agent_vpp_2    memif0
#    Ping On agent_vpp_2 With IP 10.1.1.1, Source ${int}
    Ping On agent_vpp_2 With IP ${IP_1}, Source memif1/1
    Ping On agent_vpp_2 With IP ${IP_3}, Source memif2/1

Ping Should Fail
    #no route from vrf 0 to vrf 2, then shoul fail
    Command: Ping On agent_vpp_1 With IP ${IP_4}, Source memif1/1 should fail
    Command: Ping On agent_vpp_1 With IP ${IP_3}, Source memif1/1 should fail
    Command: Ping On agent_vpp_2 With IP ${IP_4}, Source memif1/1 should fail
    Command: Ping On agent_vpp_2 With IP ${IP_3}, Source memif1/1 should fail

Create Route For Inter Vrf Routing
    Create Route On agent_vpp_1 With IP ${NET2}/24 With Next Hop ${IP_2} And Vrf Id 0
    Create Route On agent_vpp_1 With IP ${NET1}/24 With Next Hop ${IP_4} And Vrf Id 2
    Create Route On agent_vpp_2 With IP ${NET2}/24 With Next Hop VRF 2 From Vrf Id 0 And Type 1
    Create Route On agent_vpp_2 With IP ${NET1}/24 With Next Hop VRF 0 From Vrf Id 2 And Type 1

Config Done
    No Operation

Check Inter VRF Routing
    Show IP Fib On agent_vpp_1
    IP Fib Table 0 On agent_vpp_1 Should Contain Route With IP ${NET2}/24
    IP Fib Table 0 On agent_vpp_1 Should Contain Vrf ipv4 via ${IP_2} memif1/1
    Show IP Fib On agent_vpp_2
    IP Fib Table 2 On agent_vpp_2 Should Contain Route With IP ${NET1}/24
    IP Fib Table 2 On agent_vpp_2 Should Contain Vrf unicast lookup in ipv4-VRF:
    IP Fib Table 0 On agent_vpp_2 Should Contain Route With IP ${NET2}/24
    IP Fib Table 0 On agent_vpp_2 Should Contain Vrf unicast lookup in ipv4-VRF:

Check Route With Ping
    Ping On agent_vpp_1 With IP ${IP_4}, Source memif1/1
    Ping On agent_vpp_1 With IP ${IP_4}, Source memif2/1
    Ping On agent_vpp_1 With IP ${IP_3}, Source memif1/1

Final Sleep For Manual Checking
    Sleep   ${FINAL_SLEEP}

*** Keywords ***
List of interfaces On ${node} Should Contain Interface ${int}
    ${out}=   vpp_term: Show Interfaces    ${node}
    Should Match Regexp        ${out}  ${int}

IP Fib Table ${table_id} On ${node} Should Contain Vrf ${inter_vrf_string}
    ${out}=    vpp_term: Show IP Fib Table    ${node}    ${table_id}
    Should Contain  ${out}  ${inter_vrf_string}
