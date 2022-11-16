package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreatorPoolMint = "creator_pool_mint"

var _ sdk.Msg = &MsgCreatorPoolMint{}

func NewMsgCreatorPoolMint(creator string) *MsgCreatorPoolMint {
	return &MsgCreatorPoolMint{
		Creator: creator,
	}
}

func (msg *MsgCreatorPoolMint) Route() string {
	return RouterKey
}

func (msg *MsgCreatorPoolMint) Type() string {
	return TypeMsgCreatorPoolMint
}

func (msg *MsgCreatorPoolMint) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatorPoolMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatorPoolMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
