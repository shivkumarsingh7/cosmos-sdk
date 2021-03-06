package types

import (
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
	host "github.com/cosmos/cosmos-sdk/x/ibc/24-host"
)

var (
	_ exported.ChannelI      = (*Channel)(nil)
	_ exported.CounterpartyI = (*Counterparty)(nil)
)

// NewChannel creates a new Channel instance
func NewChannel(
	state State, ordering Order, counterparty Counterparty,
	hops []string, version string,
) Channel {
	return Channel{
		State:          state,
		Ordering:       ordering,
		Counterparty:   counterparty,
		ConnectionHops: hops,
		Version:        version,
	}
}

// GetState implements Channel interface.
func (ch Channel) GetState() int32 {
	return int32(ch.State)
}

// GetOrdering implements Channel interface.
func (ch Channel) GetOrdering() int32 {
	return int32(ch.Ordering)
}

// GetCounterparty implements Channel interface.
func (ch Channel) GetCounterparty() exported.CounterpartyI {
	return ch.Counterparty
}

// GetConnectionHops implements Channel interface.
func (ch Channel) GetConnectionHops() []string {
	return ch.ConnectionHops
}

// GetVersion implements Channel interface.
func (ch Channel) GetVersion() string {
	return ch.Version
}

// ValidateBasic performs a basic validation of the channel fields
func (ch Channel) ValidateBasic() error {
	if ch.State.String() == "" {
		return ErrInvalidChannelState
	}
	if !(ch.Ordering == ORDERED || ch.Ordering == UNORDERED) {
		return sdkerrors.Wrap(ErrInvalidChannelOrdering, ch.Ordering.String())
	}
	if len(ch.ConnectionHops) != 1 {
		return sdkerrors.Wrap(
			ErrTooManyConnectionHops,
			"current IBC version only supports one connection hop",
		)
	}
	if err := host.ConnectionIdentifierValidator(ch.ConnectionHops[0]); err != nil {
		return sdkerrors.Wrap(err, "invalid connection hop ID")
	}
	if strings.TrimSpace(ch.Version) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidVersion, "channel version can't be blank")
	}
	return ch.Counterparty.ValidateBasic()
}

// NewCounterparty returns a new Counterparty instance
func NewCounterparty(portID, channelID string) Counterparty {
	return Counterparty{
		PortID:    portID,
		ChannelID: channelID,
	}
}

// GetPortID implements CounterpartyI interface
func (c Counterparty) GetPortID() string {
	return c.PortID
}

// GetChannelID implements CounterpartyI interface
func (c Counterparty) GetChannelID() string {
	return c.ChannelID
}

// ValidateBasic performs a basic validation check of the identifiers
func (c Counterparty) ValidateBasic() error {
	if err := host.PortIdentifierValidator(c.PortID); err != nil {
		return sdkerrors.Wrap(err, "invalid counterparty port ID")
	}
	if err := host.ChannelIdentifierValidator(c.ChannelID); err != nil {
		return sdkerrors.Wrap(err, "invalid counterparty channel ID")
	}
	return nil
}

// NewIdentifiedChannel creates a new IdentifiedChannel instance
func NewIdentifiedChannel(portID, channelID string, ch Channel) IdentifiedChannel {
	return IdentifiedChannel{
		State:          ch.State,
		Ordering:       ch.Ordering,
		Counterparty:   ch.Counterparty,
		ConnectionHops: ch.ConnectionHops,
		Version:        ch.Version,
		PortID:         portID,
		ChannelID:      channelID,
	}
}

// ValidateBasic performs a basic validation of the identifiers and channel fields.
func (ic IdentifiedChannel) ValidateBasic() error {
	if err := host.ChannelIdentifierValidator(ic.ChannelID); err != nil {
		return sdkerrors.Wrap(err, "invalid channel ID")
	}
	if err := host.PortIdentifierValidator(ic.PortID); err != nil {
		return sdkerrors.Wrap(err, "invalid port ID")
	}
	channel := NewChannel(ic.State, ic.Ordering, ic.Counterparty, ic.ConnectionHops, ic.Version)
	return channel.ValidateBasic()
}
