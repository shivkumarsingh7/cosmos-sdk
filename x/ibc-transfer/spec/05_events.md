<!--
order: 5
-->

# Events

## MsgTransfer

| Type         | Attribute Key | Attribute Value |
|--------------|---------------|-----------------|
| ibc_transfer | sender        | {sender}        |
| ibc_transfer | receiver      | {receiver}      |
| message      | action        | transfer        |
| message      | module        | transfer        |

## OnRecvPacket callback

| Type                  | Attribute Key | Attribute Value |
|-----------------------|---------------|-----------------|
| fungible_token_packet | module        | transfer        |
| fungible_token_packet | receiver      | {receiver}      |
| fungible_token_packet | value         | {amount}        |

## OnAcknowledgePacket callback

| Type                  | Attribute Key | Attribute Value |
|-----------------------|---------------|-----------------|
| fungible_token_packet | module        | transfer        |
| fungible_token_packet | receiver      | {receiver}      |
| fungible_token_packet | value         | {amount}        |
| fungible_token_packet | success       | {ackSuccess}    |

## OnTimeoutPacket callback

| Type                  | Attribute Key   | Attribute Value |
|-----------------------|-----------------|-----------------|
| fungible_token_packet | module          | transfer        |
| fungible_token_packet | refund_receiver | {receiver}      |
| fungible_token_packet | value           | {amount}        |






