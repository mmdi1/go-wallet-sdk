// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine_v2

import (
	"errors"

	"github.com/mmdi1/go-wallet-sdk/coins/solana/base"
)

// InitializeCandyMachine is the `initializeCandyMachine` instruction.
type InitializeCandyMachine struct {
	Data *CandyMachineData

	// [0] = [WRITE] candyMachine
	//
	// [1] = [] wallet
	//
	// [2] = [] authority
	//
	// [3] = [SIGNER] payer
	//
	// [4] = [] systemProgram
	//
	// [5] = [] rent
	base.AccountMetaSlice `bin:"-"`
}

// NewInitializeCandyMachineInstructionBuilder creates a new `InitializeCandyMachine` instruction builder.
func NewInitializeCandyMachineInstructionBuilder() *InitializeCandyMachine {
	nd := &InitializeCandyMachine{
		AccountMetaSlice: make(base.AccountMetaSlice, 6),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *InitializeCandyMachine) SetData(data CandyMachineData) *InitializeCandyMachine {
	inst.Data = &data
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *InitializeCandyMachine) SetCandyMachineAccount(candyMachine base.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[0] = base.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *InitializeCandyMachine) GetCandyMachineAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetWalletAccount sets the "wallet" account.
func (inst *InitializeCandyMachine) SetWalletAccount(wallet base.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[1] = base.Meta(wallet)
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *InitializeCandyMachine) GetWalletAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *InitializeCandyMachine) SetAuthorityAccount(authority base.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[2] = base.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *InitializeCandyMachine) GetAuthorityAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *InitializeCandyMachine) SetPayerAccount(payer base.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[3] = base.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *InitializeCandyMachine) GetPayerAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeCandyMachine) SetSystemProgramAccount(systemProgram base.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[4] = base.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeCandyMachine) GetSystemProgramAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetRentAccount sets the "rent" account.
func (inst *InitializeCandyMachine) SetRentAccount(rent base.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[5] = base.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitializeCandyMachine) GetRentAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst InitializeCandyMachine) Build() *Instruction {
	return &Instruction{BaseVariant: base.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeCandyMachine,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeCandyMachine) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeCandyMachine) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (obj InitializeCandyMachine) MarshalWithEncoder(encoder *base.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeCandyMachineInstruction declares a new InitializeCandyMachine instruction with the provided parameters and accounts.
func NewInitializeCandyMachineInstruction(
	// Parameters:
	data CandyMachineData,
	// Accounts:
	candyMachine base.PublicKey,
	wallet base.PublicKey,
	authority base.PublicKey,
	payer base.PublicKey,
	systemProgram base.PublicKey,
	rent base.PublicKey) *InitializeCandyMachine {
	return NewInitializeCandyMachineInstructionBuilder().
		SetData(data).
		SetCandyMachineAccount(candyMachine).
		SetWalletAccount(wallet).
		SetAuthorityAccount(authority).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
