---
title: "6502 Instructions - Obelisk by Andrew Jacobs"
source: "https://6502.org/users/obelisk/6502/instructions.html"
word_count: 928
---

## The Instruction Set

The 6502 has a relatively basic set of instructions, many having similar functions (e.g. memory access, arithmetic,
etc.). The following sections list the complete set of 56 instructions in functional groups.

### Load/Store Operations

These instructions transfer a single byte between memory and one of the registers. Load operations set the
negative ([[registers#Negative Flag|N]]) and zero ([[registers#Zero Flag|Z]]) flags depending on the value of
transferred. Store operations do not affect the flag settings.

| Instruction                                        | Operation         | Flags                                                            |
|----------------------------------------------------|-------------------|------------------------------------------------------------------|
| [LDA](reference.md#LDA%20-%20Load%20Accumulator)   | Load Accumulator  | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [LDX](reference.md#LDX%20-%20Load%20X%20Register)  | Load X Register   | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [LDY](reference.md#LDY%20-%20Load%20Y%20Register)  | Load Y Register   | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [STA](reference.md#STA%20-%20Store%20Accumulator)  | Store Accumulator |                                                                  |
| [STX](reference.md#STX%20-%20Store%20X%20Register) | Store X Register  |                                                                  |
| [STY](reference.md#STY%20-%20Store%20Y%20Register) | Store Y Register  |                                                                  |

### Register Transfers

The contents of the X and Y registers can be moved to or from the accumulator, setting the
negative ([[registers#Negative Flag|N]]) and zero ([[registers#Zero Flag|Z]]) flags as appropriate.

| Instruction                                                   | Operation                 | Flags                                                            |
|---------------------------------------------------------------|---------------------------|------------------------------------------------------------------|
| [TAX](reference.md#TAX%20-%20Transfer%20Accumulator%20to%20X) | Transfer accumulator to X | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [TAY](reference.md#TAY%20-%20Transfer%20Accumulator%20to%20Y) | Transfer accumulator to Y | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [TXA](reference.md#TXA%20-%20Transfer%20X%20to%20Accumulator) | Transfer X to accumulator | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [TYA](reference.md#TYA%20-%20Transfer%20Y%20to%20Accumulator) | Transfer Y to accumulator | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |

### Stack Operations

The 6502 microprocessor supports a 256 byte stack fixed between memory locations $0100 and $01FF. A special 8-bit
register, S, is used to keep track of the next free byte of stack space. Pushing a byte on to the stack causes the value
to be stored at the current free location (e.g. $0100,S) and then the stack pointer is post decremented. Pull operations
reverse this procedure.

The stack register can only be accessed by transferring its value to or from the X register. Its value is automatically
modified by push/pull instructions, subroutine calls and returns, interrupts and returns from interrupts.

| Instruction                                                       | Operation                        | Flags                                                            |
|-------------------------------------------------------------------|----------------------------------|------------------------------------------------------------------|
| [TSX](reference.md#TSX%20-%20Transfer%20Stack%20Pointer%20to%20X) | Transfer stack pointer to X      | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [TXS](reference.md#TXS%20-%20Transfer%20X%20to%20Stack%20Pointer) | Transfer X to stack pointer      |                                                                  |
| [PHA](reference.md#PHA%20-%20Push%20Accumulator)                  | Push accumulator on stack        |                                                                  |
| [PHP](reference.md#PHP%20-%20Push%20Processor%20Status)           | Push processor status on stack   |                                                                  |
| [PLA](reference.md#PLA%20-%20Pull%20Accumulator)                  | Pull accumulator from stack      | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [PLP](reference.md#PLP%20-%20Pull%20Processor%20Status)           | Pull processor status from stack | All                                                              |

### Logical

The following instructions perform logical operations on the contents of the accumulator and another value held in
memory. The BIT instruction performs a logical AND to test the presence of bits in the memory value to set the flags but
does not keep the result.

| Instruction                                            | Operation            | Flags                                                                                               |
|--------------------------------------------------------|----------------------|-----------------------------------------------------------------------------------------------------|
| [AND](reference.md#AND%20-%20Logical%20AND)            | Logical AND          | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag)                                    |
| [EOR](reference.md#EOR%20-%20Exclusive%20OR)           | Exclusive OR         | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag)                                    |
| [ORA](reference.md#ORA%20-%20Logical%20Inclusive%20OR) | Logical Inclusive OR | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag)                                    |
| [BIT](reference.md#BIT%20-%20Bit%20Test)               | Bit Test             | [N](registers.md#Negative%20Flag), [V](registers.md#Overflow%20Flag), [Z](registers.md#Zero%20Flag) |

### Arithmetic

The arithmetic operations perform addition and subtraction on the contents of the accumulator. The compare operations
allow the comparison of the accumulator and X or Y with memory values.

| Instruction                                           | Operation           | Flags                                                                                                                               |
|-------------------------------------------------------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| [ADC](reference.md#ADC%20-%20Add%20with%20Carry)      | Add with Carry      | [N](registers.md#Negative%20Flag), [V](registers.md#Overflow%20Flag), [Z](registers.md#Zero%20Flag), [C](registers.md#Carry%20Flag) |
| [SBC](reference.md#SBC%20-%20Subtract%20with%20Carry) | Subtract with Carry | [N](registers.md#Negative%20Flag), [V](registers.md#Overflow%20Flag), [Z](registers.md#Zero%20Flag), [C](registers.md#Carry%20Flag) |
| [CMP](reference.md#CMP%20-%20Compare)                 | Compare accumulator | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag), [C](registers.md#Carry%20Flag)                                    |
| [CPX](reference.md#CPX%20-%20Compare%20X%20Register)  | Compare X register  | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag), [C](registers.md#Carry%20Flag)                                    |
| [CPY](reference.md#CPY%20-%20Compare%20Y%20Register)  | Compare Y register  | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag), [C](registers.md#Carry%20Flag)                                    |

### Increments & Decrements

Increment or decrement a memory location or one of the X or Y registers by one setting the
negative ([[registers#Negative Flag|N]]) and zero ([[registers#Zero Flag|Z]]) flags as appropriate,

| Instruction                                            | Operation                   | Flags                                                            |
|--------------------------------------------------------|-----------------------------|------------------------------------------------------------------|
| [INC](reference.md#INC%20-%20Increment%20Memory)       | Increment a memory location | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [INX](reference.md#INX%20-%20Increment%20X%20Register) | Increment the X register    | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [INY](reference.md#INY%20-%20Increment%20Y%20Register) | Increment the Y register    | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [DEC](reference.md#DEC%20-%20Decrement%20Memory)       | Decrement a memory location | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [DEX](reference.md#DEX%20-%20Decrement%20X%20Register) | Decrement the X register    | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |
| [DEY](reference.md#DEY%20-%20Decrement%20Y%20Register) | Decrement the Y register    | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag) |

### Shifts

Shift instructions cause the bits within either a memory location or the accumulator to be shifted by one bit position.
The rotate instructions use the contents if the carry flag ([[registers#Carry Flag|C]]) to fill the vacant position
generated by the shift and to catch the overflowing bit. The arithmetic and logical shifts shift in an appropriate 0 or
1 bit as appropriate but catch the overflow bit in the carry flag ([[registers#Carry Flag|C]]).

| Instruction                                             | Operation             | Flags                                                                                            |
|---------------------------------------------------------|-----------------------|--------------------------------------------------------------------------------------------------|
| [ASL](reference.md#ASL%20-%20Arithmetic%20Shift%20Left) | Arithmetic Shift Left | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag), [C](registers.md#Carry%20Flag) |
| [LSR](reference.md#LSR%20-%20Logical%20Shift%20Right)   | Logical Shift Right   | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag), [C](registers.md#Carry%20Flag) |
| [ROL](reference.md#ROL%20-%20Rotate%20Left)             | Rotate Left           | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag), [C](registers.md#Carry%20Flag) |
| [ROR](reference.md#ROR%20-%20Rotate%20Right)            | Rotate Right          | [N](registers.md#Negative%20Flag), [Z](registers.md#Zero%20Flag), [C](registers.md#Carry%20Flag) |

### Jumps & Calls

The following instructions modify the program counter causing a break to normal sequential execution.
The [[reference#JSR - Jump to Subroutine|JSR]] instruction pushes the old [[registers#Program Counter|PC]] onto the
stack before changing it to the new location allowing a subsequent [[reference#RTS - Return from Subroutine|RTS]] to
return execution to the instruction after the call.

| Instruction                                              | Operation                | Flags |
|----------------------------------------------------------|--------------------------|-------|
| [JMP](reference.md#JMP%20-%20Jump)                       | Jump to another location |       |
| [JSR](reference.md#JSR%20-%20Jump%20to%20Subroutine)     | Jump to a subroutine     |       |
| [RTS](reference.md#RTS%20-%20Return%20from%20Subroutine) | Return from subroutine   |       |

### Branches

Branch instructions break the normal sequential flow of execution by changing the program counter if a specified
condition is met. All the conditions are based on examining a single bit within the processor status.

| Instruction                                                  | Operation                     | Flags |
|--------------------------------------------------------------|-------------------------------|-------|
| [BCC](reference.md#BCC%20-%20Branch%20if%20Carry%20Clear)    | Branch if carry flag clear    |       |
| [BCS](reference.md#BCS%20-%20Branch%20if%20Carry%20Set)      | Branch if carry flag set      |       |
| [BEQ](reference.md#BEQ%20-%20Branch%20if%20Equal)            | Branch if zero flag set       |       |
| [BMI](reference.md#BMI%20-%20Branch%20if%20Minus)            | Branch if negative flag set   |       |
| [BNE](reference.md#BNE%20-%20Branch%20if%20Not%20Equal)      | Branch if zero flag clear     |       |
| [BPL](reference.md#BPL%20-%20Branch%20if%20Positive)         | Branch if negative flag clear |       |
| [BVC](reference.md#BVC%20-%20Branch%20if%20Overflow%20Clear) | Branch if overflow flag clear |       |
| [BVS](reference.md#BVS%20-%20Branch%20if%20Overflow%20Set)   | Branch if overflow flag set   |       |

Branch instructions use relative address to identify the target instruction if they are executed. As relative addresses
are stored using a signed 8 bit byte the target instruction must be within 126 bytes before the branch or 128 bytes
after the branch.

### Status Flag Changes

The following instructions change the values of specific status flags.

| Instruction                                               | Operation                    | Flags                                 |
|-----------------------------------------------------------|------------------------------|---------------------------------------|
| [CLC](reference.md#CLC%20-%20Clear%20Carry%20Flag)        | Clear carry flag             | [C](registers.md#Carry%20Flag)        |
| [CLD](reference.md#CLD%20-%20Clear%20Decimal%20Mode)      | Clear decimal mode flag      | [D](registers.md#Decimal%20Mode)      |
| [CLI](reference.md#CLI%20-%20Clear%20Interrupt%20Disable) | Clear interrupt disable flag | [I](registers.md#Interrupt%20Disable) |
| [CLV](reference.md#CLV%20-%20Clear%20Overflow%20Flag)     | Clear overflow flag          | [V](registers.md#Overflow%20Flag)     |
| [SEC](reference.md#SEC%20-%20Set%20Carry%20Flag)          | Set carry flag               | [C](registers.md#Carry%20Flag)        |
| [SED](reference.md#SED%20-%20Set%20Decimal%20Flag)        | Set decimal mode flag        | [D](registers.md#Decimal%20Mode)      |
| [SEI](reference.md#SEI%20-%20Set%20Interrupt%20Disable)   | Set interrupt disable flag   | [I](registers.md#Interrupt%20Disable) |

### System Functions

The remaining instructions perform useful but rarely used functions.

| Instruction                                             | Operation             | Flags                             |
|---------------------------------------------------------|-----------------------|-----------------------------------|
| [BRK](reference.md#BRK%20-%20Force%20Interrupt)         | Force an interrupt    | [B](registers.md#Break%20Command) |
| [NOP](reference.md#NOP%20-%20No%20Operation)            | No Operation          |                                   |
| [RTI](reference.md#RTI%20-%20Return%20from%20Interrupt) | Return from Interrupt | All                               |

---

This page was last updated on 2nd January 2002
