---
title: "6502 Reference - Obelisk by Andrew Jacobs"
source: "https://6502.org/users/obelisk/6502/reference.md"
word_count: 5618
---

## Instruction Reference

Click on any of following links to go straight to the information for that instruction.

| [ADC](#ADC%20-%20Add%20with%20Carry)         | [AND](#AND%20-%20Logical%20AND)               | [ASL](#ASL%20-%20Arithmetic%20Shift%20Left) | [BCC](#BCC%20-%20Branch%20if%20Carry%20Clear) | [BCS](#BCS%20-%20Branch%20if%20Carry%20Set) | [BEQ](#BEQ%20-%20Branch%20if%20Equal)    | [BIT](#BIT%20-%20Bit%20Test)               | [BMI](#BMI%20-%20Branch%20if%20Minus)      | [BNE](#BNE%20-%20Branch%20if%20Not%20Equal)       | [BPL](#BPL%20-%20Branch%20if%20Positive)          | [BRK](#BRK%20-%20Force%20Interrupt)                   | [BVC](#BVC%20-%20Branch%20if%20Overflow%20Clear)  | [BVS](#BVS%20-%20Branch%20if%20Overflow%20Set)        | [CLC](#CLC%20-%20Clear%20Carry%20Flag)            |
|----------------------------------------------|-----------------------------------------------|---------------------------------------------|-----------------------------------------------|---------------------------------------------|------------------------------------------|--------------------------------------------|--------------------------------------------|---------------------------------------------------|---------------------------------------------------|-------------------------------------------------------|---------------------------------------------------|-------------------------------------------------------|---------------------------------------------------|
| [CLD](#CLD%20-%20Clear%20Decimal%20Mode)     | [CLI](#CLI%20-%20Clear%20Interrupt%20Disable) | [CLV](#CLV%20-%20Clear%20Overflow%20Flag)   | [CMP](#CMP%20-%20Compare)                     | [CPX](#CPX%20-%20Compare%20X%20Register)    | [CPY](#CPY%20-%20Compare%20Y%20Register) | [DEC](#DEC%20-%20Decrement%20Memory)       | [DEX](#DEX%20-%20Decrement%20X%20Register) | [DEY](#DEY%20-%20Decrement%20Y%20Register)        | [EOR](#EOR%20-%20Exclusive%20OR)                  | [INC](#INC%20-%20Increment%20Memory)                  | [INX](#INX%20-%20Increment%20X%20Register)        | [INY](#INY%20-%20Increment%20Y%20Register)            | [JMP](#JMP%20-%20Jump)                            |
| [JSR](#JSR%20-%20Jump%20to%20Subroutine)     | [LDA](#LDA%20-%20Load%20Accumulator)          | [LDX](#LDX%20-%20Load%20X%20Register)       | [LDY](#LDY%20-%20Load%20Y%20Register)         | [LSR](#LSR%20-%20Logical%20Shift%20Right)   | [NOP](#NOP%20-%20No%20Operation)         | [ORA](#ORA%20-%20Logical%20Inclusive%20OR) | [PHA](#PHA%20-%20Push%20Accumulator)       | [PHP](#PHP%20-%20Push%20Processor%20Status)       | [PLA](#PLA%20-%20Pull%20Accumulator)              | [PLP](#PLP%20-%20Pull%20Processor%20Status)           | [ROL](#ROL%20-%20Rotate%20Left)                   | [ROR](#ROR%20-%20Rotate%20Right)                      | [RTI](#RTI%20-%20Return%20from%20Interrupt)       |
| [RTS](#RTS%20-%20Return%20from%20Subroutine) | [SBC](#SBC%20-%20Subtract%20with%20Carry)     | [SEC](#SEC%20-%20Set%20Carry%20Flag)        | [SED](#SED%20-%20Set%20Decimal%20Flag)        | [SEI](#SEI%20-%20Set%20Interrupt%20Disable) | [STA](#STA%20-%20Store%20Accumulator)    | [STX](#STX%20-%20Store%20X%20Register)     | [STY](#STY%20-%20Store%20Y%20Register)     | [TAX](#TAX%20-%20Transfer%20Accumulator%20to%20X) | [TAY](#TAY%20-%20Transfer%20Accumulator%20to%20Y) | [TSX](#TSX%20-%20Transfer%20Stack%20Pointer%20to%20X) | [TXA](#TXA%20-%20Transfer%20X%20to%20Accumulator) | [TXS](#TXS%20-%20Transfer%20X%20to%20Stack%20Pointer) | [TYA](#TYA%20-%20Transfer%20Y%20to%20Accumulator) |

### ADC - Add with Carry

A,Z,C,N = A+M+C

This instruction adds the contents of a memory location to the accumulator together with the carry bit. If overflow
occurs the carry bit is set, this enables multiple byte addition to be performed.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set if overflow in bit 7     |
|-----------------------------------------|---------------------------------------------------------|------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0                 |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                 |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                 |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                 |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Set if sign bit is incorrect |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 set             |

| **Addressing Mode**                              | **Opcode** | **Bytes** | **Cycles**             |
|--------------------------------------------------|------------|-----------|------------------------|
| [Immediate](addressing.md#Immediate)             | $69        | 2         | 2                      |
| [Zero Page](addressing.md#Zero%20Page)           | $65        | 2         | 3                      |
| [Zero Page,X](addressing.md#Zero%20Page,X)       | $75        | 2         | 4                      |
| [Absolute](addressing.md#Absolute)               | $6D        | 3         | 4                      |
| [Absolute,X](addressing.md#Absolute,X)           | $7D        | 3         | 4 (+1 if page crossed) |
| [Absolute,Y](addressing.md#Absolute,Y)           | $79        | 3         | 4 (+1 if page crossed) |
| [(Indirect,X)](addressing.md#Indexed%20Indirect) | $61        | 2         | 6                      |
| [(Indirect),Y](addressing.md#Indirect%20Indexed) | $71        | 2         | 5 (+1 if page crossed) |

See also: [SBC](#SBC%20-%20Subtract%20with%20Carry)

### AND - Logical AND

A,Z,N = A&M

A logical AND is performed, bit by bit, on the accumulator contents using the contents of a byte of memory.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected     |
|-----------------------------------------|---------------------------------------------------------|------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0     |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected     |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected     |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected     |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected     |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 set |

| **Addressing Mode**                              | **Opcode** | **Bytes** | **Cycles**             |
|--------------------------------------------------|------------|-----------|------------------------|
| [Immediate](addressing.md#Immediate)             | $29        | 2         | 2                      |
| [Zero Page](addressing.md#Zero%20Page)           | $25        | 2         | 3                      |
| [Zero Page,X](addressing.md#Zero%20Page,X)       | $35        | 2         | 4                      |
| [Absolute](addressing.md#Absolute)               | $2D        | 3         | 4                      |
| [Absolute,X](addressing.md#Absolute,X)           | $3D        | 3         | 4 (+1 if page crossed) |
| [Absolute,Y](addressing.md#Absolute,Y)           | $39        | 3         | 4 (+1 if page crossed) |
| [(Indirect,X)](addressing.md#Indexed%20Indirect) | $21        | 2         | 6                      |
| [(Indirect),Y](addressing.md#Indirect%20Indexed) | $31        | 2         | 5 (+1 if page crossed) |

See also: [EOR](#EOR%20-%20Exclusive%20OR), [ORA](#ORA%20-%20Logical%20Inclusive%20OR)

### ASL - Arithmetic Shift Left

A,Z,C,N = M\*2 or M,Z,C,N = M\*2

This operation shifts all the bits of the accumulator or memory contents one bit left. Bit 0 is set to 0 and bit 7 is
placed in the carry flag. The effect of this operation is to multiply the memory contents by 2 (ignoring 2's complement
considerations), setting the carry if the result will not fit in 8 bits.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set to contents of old bit 7      |
|-----------------------------------------|---------------------------------------------------------|-----------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0                      |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                      |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                      |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                      |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected                      |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of the result is set |

| **Addressing Mode**                        | **Opcode** | **Bytes** | **Cycles** |
|--------------------------------------------|------------|-----------|------------|
| [Accumulator](addressing.md#Accumulator)   | $0A        | 1         | 2          |
| [Zero Page](addressing.md#Zero%20Page)     | $06        | 2         | 5          |
| [Zero Page,X](addressing.md#Zero%20Page,X) | $16        | 2         | 6          |
| [Absolute](addressing.md#Absolute)         | $0E        | 3         | 6          |
| [Absolute,X](addressing.md#Absolute,X)     | $1E        | 3         | 7          |

See also: [LSR](#LSR%20-%20Logical%20Shift%20Right), [ROL](#ROL%20-%20Rotate%20Left), [ROR](#ROR%20-%20Rotate%20Right)

### BCC - Branch if Carry Clear

If the carry flag is clear then add the relative displacement to the program counter to cause a branch to a new
location.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                | **Opcode** | **Bytes** | **Cycles**                                      |
|------------------------------------|------------|-----------|-------------------------------------------------|
| [Relative](addressing.md#Relative) | $90        | 2         | 2 (+1 if branch succeeds   +2 if to a new page) |

See also: [BCS](#BCS%20-%20Branch%20if%20Carry%20Set)

### BCS - Branch if Carry Set

If the carry flag is set then add the relative displacement to the program counter to cause a branch to a new location.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                | **Opcode** | **Bytes** | **Cycles**                                      |
|------------------------------------|------------|-----------|-------------------------------------------------|
| [Relative](addressing.md#Relative) | $B0        | 2         | 2 (+1 if branch succeeds   +2 if to a new page) |

See also: [BCC](#BCC%20-%20Branch%20if%20Carry%20Clear)

### BEQ - Branch if Equal

If the zero flag is set then add the relative displacement to the program counter to cause a branch to a new location.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                | **Opcode** | **Bytes** | **Cycles**                                      |
|------------------------------------|------------|-----------|-------------------------------------------------|
| [Relative](addressing.md#Relative) | $F0        | 2         | 2 (+1 if branch succeeds   +2 if to a new page) |

See also: [BNE](#BNE%20-%20Branch%20if%20Not%20Equal)

### BIT - Bit Test

A & M, N = M7, V = M6

This instructions is used to test if one or more bits are set in a target memory location. The mask pattern in A is
ANDed with the value in memory to set or clear the zero flag, but the result is not kept. Bits 7 and 6 of the value from
memory are copied into the N and V flags.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected                         |
|-----------------------------------------|---------------------------------------------------------|--------------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if the result if the AND is zero |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                         |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                         |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                         |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Set to bit 6 of the memory value     |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set to bit 7 of the memory value     |

| **Addressing Mode**                    | **Opcode** | **Bytes** | **Cycles** |
|----------------------------------------|------------|-----------|------------|
| [Zero Page](addressing.md#Zero%20Page) | $24        | 2         | 3          |
| [Absolute](addressing.md#Absolute)     | $2C        | 3         | 4          |

### BMI - Branch if Minus

If the negative flag is set then add the relative displacement to the program counter to cause a branch to a new
location.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                | **Opcode** | **Bytes** | **Cycles**                                      |
|------------------------------------|------------|-----------|-------------------------------------------------|
| [Relative](addressing.md#Relative) | $30        | 2         | 2 (+1 if branch succeeds   +2 if to a new page) |

See also: [BPL](#BPL%20-%20Branch%20if%20Positive)

### BNE - Branch if Not Equal

If the zero flag is clear then add the relative displacement to the program counter to cause a branch to a new location.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                | **Opcode** | **Bytes** | **Cycles**                                      |
|------------------------------------|------------|-----------|-------------------------------------------------|
| [Relative](addressing.md#Relative) | $D0        | 2         | 2 (+1 if branch succeeds   +2 if to a new page) |

See also: [BEQ](#BEQ%20-%20Branch%20if%20Equal)

### BPL - Branch if Positive

If the negative flag is clear then add the relative displacement to the program counter to cause a branch to a new
location.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                | **Opcode** | **Bytes** | **Cycles**                                      |
|------------------------------------|------------|-----------|-------------------------------------------------|
| [Relative](addressing.md#Relative) | $10        | 2         | 2 (+1 if branch succeeds   +2 if to a new page) |

See also: [BMI](#BMI%20-%20Branch%20if%20Minus)

### BRK - Force Interrupt

The BRK instruction forces the generation of an interrupt request. The program counter and processor status are pushed
on the stack then the IRQ interrupt vector at $FFFE/F is loaded into the PC and the break flag in the status set to one.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Set to 1     |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $00        | 1         | 7          |

The interpretation of a BRK depends on the operating system. On the BBC Microcomputer it is used by language ROMs to
signal run time errors but it could be used for other purposes (e.g. calling operating system functions, etc.).

### BVC - Branch if Overflow Clear

If the overflow flag is clear then add the relative displacement to the program counter to cause a branch to a new
location.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                | **Opcode** | **Bytes** | **Cycles**                                      |
|------------------------------------|------------|-----------|-------------------------------------------------|
| [Relative](addressing.md#Relative) | $50        | 2         | 2 (+1 if branch succeeds   +2 if to a new page) |

See also: [BVS](#BVS%20-%20Branch%20if%20Overflow%20Set)

### BVS - Branch if Overflow Set

If the overflow flag is set then add the relative displacement to the program counter to cause a branch to a new
location.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                | **Opcode** | **Bytes** | **Cycles**                                      |
|------------------------------------|------------|-----------|-------------------------------------------------|
| [Relative](addressing.md#Relative) | $70        | 2         | 2 (+1 if branch succeeds   +2 if to a new page) |

See also: [BVC](#BVC%20-%20Branch%20if%20Overflow%20Clear)

### CLC - Clear Carry Flag

C = 0

Set the carry flag to zero.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set to 0     |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $18        | 1         | 2          |

See also: [SEC](#SEC%20-%20Set%20Carry%20Flag)

### CLD - Clear Decimal Mode

D = 0

Sets the decimal mode flag to zero.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Set to 0     |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $D8        | 1         | 2          |

**NB**:  
The state of the decimal flag is uncertain when the CPU is powered up and it is not reset when an interrupt is
generated. In both cases you should include an explicit CLD to ensure that the flag is cleared before performing
addition or subtraction.

See also: [SED](#SED%20-%20Set%20Decimal%20Flag)

### CLI - Clear Interrupt Disable

I = 0

Clears the interrupt disable flag allowing normal interrupt requests to be serviced.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Set to 0     |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $58        | 1         | 2          |

See also: [SEI](#SEI%20-%20Set%20Interrupt%20Disable)

### CLV - Clear Overflow Flag

V = 0

Clears the overflow flag.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Set to 0     |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $B8        | 1         | 2          |

### CMP - Compare

Z,C,N = A-M

This instruction compares the contents of the accumulator with another memory held value and sets the zero and carry
flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set if A \>\= M                   |
|-----------------------------------------|---------------------------------------------------------|-----------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = M                      |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                      |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                      |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                      |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected                      |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of the result is set |

| **Addressing Mode**                              | **Opcode** | **Bytes** | **Cycles**             |
|--------------------------------------------------|------------|-----------|------------------------|
| [Immediate](addressing.md#Immediate)             | $C9        | 2         | 2                      |
| [Zero Page](addressing.md#Zero%20Page)           | $C5        | 2         | 3                      |
| [Zero Page,X](addressing.md#Zero%20Page,X)       | $D5        | 2         | 4                      |
| [Absolute](addressing.md#Absolute)               | $CD        | 3         | 4                      |
| [Absolute,X](addressing.md#Absolute,X)           | $DD        | 3         | 4 (+1 if page crossed) |
| [Absolute,Y](addressing.md#Absolute,Y)           | $D9        | 3         | 4 (+1 if page crossed) |
| [(Indirect,X)](addressing.md#Indexed%20Indirect) | $C1        | 2         | 6                      |
| [(Indirect),Y](addressing.md#Indirect%20Indexed) | $D1        | 2         | 5 (+1 if page crossed) |

See also: [CPX](#CPX%20-%20Compare%20X%20Register), [CPY](#CPY%20-%20Compare%20Y%20Register)

### CPX - Compare X Register

Z,C,N = X-M

This instruction compares the contents of the X register with another memory held value and sets the zero and carry
flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set if X \>\= M                   |
|-----------------------------------------|---------------------------------------------------------|-----------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if X = M                      |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                      |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                      |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                      |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected                      |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of the result is set |

| **Addressing Mode**                    | **Opcode** | **Bytes** | **Cycles** |
|----------------------------------------|------------|-----------|------------|
| [Immediate](addressing.md#Immediate)   | $E0        | 2         | 2          |
| [Zero Page](addressing.md#Zero%20Page) | $E4        | 2         | 3          |
| [Absolute](addressing.md#Absolute)     | $EC        | 3         | 4          |

See also: [CMP](#CMP%20-%20Compare), [CPY](#CPY%20-%20Compare%20Y%20Register)

### CPY - Compare Y Register

Z,C,N = Y-M

This instruction compares the contents of the Y register with another memory held value and sets the zero and carry
flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set if Y \>\= M                   |
|-----------------------------------------|---------------------------------------------------------|-----------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if Y = M                      |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                      |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                      |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                      |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected                      |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of the result is set |

| **Addressing Mode**                    | **Opcode** | **Bytes** | **Cycles** |
|----------------------------------------|------------|-----------|------------|
| [Immediate](addressing.md#Immediate)   | $C0        | 2         | 2          |
| [Zero Page](addressing.md#Zero%20Page) | $C4        | 2         | 3          |
| [Absolute](addressing.md#Absolute)     | $CC        | 3         | 4          |

See also: [CMP](#CMP%20-%20Compare), [CPX](#CPX%20-%20Compare%20X%20Register)

### DEC - Decrement Memory

M,Z,N = M-1

Subtracts one from the value held at a specified memory location setting the zero and negative flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected                      |
|-----------------------------------------|---------------------------------------------------------|-----------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if result is zero             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                      |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                      |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                      |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected                      |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of the result is set |

| **Addressing Mode**                        | **Opcode** | **Bytes** | **Cycles** |
|--------------------------------------------|------------|-----------|------------|
| [Zero Page](addressing.md#Zero%20Page)     | $C6        | 2         | 5          |
| [Zero Page,X](addressing.md#Zero%20Page,X) | $D6        | 2         | 6          |
| [Absolute](addressing.md#Absolute)         | $CE        | 3         | 6          |
| [Absolute,X](addressing.md#Absolute,X)     | $DE        | 3         | 7          |

See also: [DEX](#DEX%20-%20Decrement%20X%20Register), [DEY](#DEY%20-%20Decrement%20Y%20Register)

### DEX - Decrement X Register

X,Z,N = X-1

Subtracts one from the X register setting the zero and negative flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if X is zero         |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of X is set |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $CA        | 1         | 2          |

See also: [DEC](#DEC%20-%20Decrement%20Memory), [DEY](#DEY%20-%20Decrement%20Y%20Register)

### DEY - Decrement Y Register

Y,Z,N = Y-1

Subtracts one from the Y register setting the zero and negative flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if Y is zero         |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of Y is set |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $88        | 1         | 2          |

See also: [DEC](#DEC%20-%20Decrement%20Memory), [DEX](#DEX%20-%20Decrement%20X%20Register)

### EOR - Exclusive OR

A,Z,N = A^M

An exclusive OR is performed, bit by bit, on the accumulator contents using the contents of a byte of memory.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected     |
|-----------------------------------------|---------------------------------------------------------|------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0     |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected     |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected     |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected     |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected     |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 set |

| **Addressing Mode**                              | **Opcode** | **Bytes** | **Cycles**             |
|--------------------------------------------------|------------|-----------|------------------------|
| [Immediate](addressing.md#Immediate)             | $49        | 2         | 2                      |
| [Zero Page](addressing.md#Zero%20Page)           | $45        | 2         | 3                      |
| [Zero Page,X](addressing.md#Zero%20Page,X)       | $55        | 2         | 4                      |
| [Absolute](addressing.md#Absolute)               | $4D        | 3         | 4                      |
| [Absolute,X](addressing.md#Absolute,X)           | $5D        | 3         | 4 (+1 if page crossed) |
| [Absolute,Y](addressing.md#Absolute,Y)           | $59        | 3         | 4 (+1 if page crossed) |
| [(Indirect,X)](addressing.md#Indexed%20Indirect) | $41        | 2         | 6                      |
| [(Indirect),Y](addressing.md#Indirect%20Indexed) | $51        | 2         | 5 (+1 if page crossed) |

See also: [AND](#AND%20-%20Logical%20AND), [ORA](#ORA%20-%20Logical%20Inclusive%20OR)

### INC - Increment Memory

M,Z,N = M+1

Adds one to the value held at a specified memory location setting the zero and negative flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected                      |
|-----------------------------------------|---------------------------------------------------------|-----------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if result is zero             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                      |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                      |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                      |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected                      |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of the result is set |

| **Addressing Mode**                        | **Opcode** | **Bytes** | **Cycles** |
|--------------------------------------------|------------|-----------|------------|
| [Zero Page](addressing.md#Zero%20Page)     | $E6        | 2         | 5          |
| [Zero Page,X](addressing.md#Zero%20Page,X) | $F6        | 2         | 6          |
| [Absolute](addressing.md#Absolute)         | $EE        | 3         | 6          |
| [Absolute,X](addressing.md#Absolute,X)     | $FE        | 3         | 7          |

See also: [INX](#INX%20-%20Increment%20X%20Register), [INY](#INY%20-%20Increment%20Y%20Register)

### INX - Increment X Register

X,Z,N = X+1

Adds one to the X register setting the zero and negative flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if X is zero         |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of X is set |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $E8        | 1         | 2          |

See also: [INC](#INC%20-%20Increment%20Memory), [INY](#INY%20-%20Increment%20Y%20Register)

### INY - Increment Y Register

Y,Z,N = Y+1

Adds one to the Y register setting the zero and negative flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if Y is zero         |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of Y is set |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $C8        | 1         | 2          |

See also: [INC](#INC%20-%20Increment%20Memory), [INX](#INX%20-%20Increment%20X%20Register)

### JMP - Jump

Sets the program counter to the address specified by the operand.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                | **Opcode** | **Bytes** | **Cycles** |
|------------------------------------|------------|-----------|------------|
| [Absolute](addressing.md#Absolute) | $4C        | 3         | 3          |
| [Indirect](addressing.md#Indirect) | $6C        | 3         | 5          |

NB:  
An original 6502 has does not correctly fetch the target address if the indirect vector falls on a page boundary (
e.g. $xxFF where xx is any value from $00 to $FF). In this case fetches the LSB from $xxFF as expected but takes the MSB
from $xx00. This is fixed in some later chips like the 65SC02 so for compatibility always ensure the indirect vector is
not at the end of the page.

### JSR - Jump to Subroutine

The JSR instruction pushes the address (minus one) of the return point on to the stack and then sets the program counter
to the target memory address.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                | **Opcode** | **Bytes** | **Cycles** |
|------------------------------------|------------|-----------|------------|
| [Absolute](addressing.md#Absolute) | $20        | 3         | 6          |

See also: [RTS](#RTS%20-%20Return%20from%20Subroutine)

### LDA - Load Accumulator

A,Z,N = M

Loads a byte of memory into the accumulator setting the zero and negative flags as appropriate.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of A is set |

| **Addressing Mode**                              | **Opcode** | **Bytes** | **Cycles**             |
|--------------------------------------------------|------------|-----------|------------------------|
| [Immediate](addressing.md#Immediate)             | $A9        | 2         | 2                      |
| [Zero Page](addressing.md#Zero%20Page)           | $A5        | 2         | 3                      |
| [Zero Page,X](addressing.md#Zero%20Page,X)       | $B5        | 2         | 4                      |
| [Absolute](addressing.md#Absolute)               | $AD        | 3         | 4                      |
| [Absolute,X](addressing.md#Absolute,X)           | $BD        | 3         | 4 (+1 if page crossed) |
| [Absolute,Y](addressing.md#Absolute,Y)           | $B9        | 3         | 4 (+1 if page crossed) |
| [(Indirect,X)](addressing.md#Indexed%20Indirect) | $A1        | 2         | 6                      |
| [(Indirect),Y](addressing.md#Indirect%20Indexed) | $B1        | 2         | 5 (+1 if page crossed) |

See also: [LDX](#LDX%20-%20Load%20X%20Register), [LDY](#LDY%20-%20Load%20Y%20Register)

### LDX - Load X Register

X,Z,N = M

Loads a byte of memory into the X register setting the zero and negative flags as appropriate.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if X = 0             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of X is set |

| **Addressing Mode**                        | **Opcode** | **Bytes** | **Cycles**             |
|--------------------------------------------|------------|-----------|------------------------|
| [Immediate](addressing.md#Immediate)       | $A2        | 2         | 2                      |
| [Zero Page](addressing.md#Zero%20Page)     | $A6        | 2         | 3                      |
| [Zero Page,Y](addressing.md#Zero%20Page,Y) | $B6        | 2         | 4                      |
| [Absolute](addressing.md#Absolute)         | $AE        | 3         | 4                      |
| [Absolute,Y](addressing.md#Absolute,Y)     | $BE        | 3         | 4 (+1 if page crossed) |

See also: [LDA](#LDA%20-%20Load%20Accumulator), [LDY](#LDY%20-%20Load%20Y%20Register)

### LDY - Load Y Register

Y,Z,N = M

Loads a byte of memory into the Y register setting the zero and negative flags as appropriate.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if Y = 0             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of Y is set |

| **Addressing Mode**                        | **Opcode** | **Bytes** | **Cycles**             |
|--------------------------------------------|------------|-----------|------------------------|
| [Immediate](addressing.md#Immediate)       | $A0        | 2         | 2                      |
| [Zero Page](addressing.md#Zero%20Page)     | $A4        | 2         | 3                      |
| [Zero Page,X](addressing.md#Zero%20Page,X) | $B4        | 2         | 4                      |
| [Absolute](addressing.md#Absolute)         | $AC        | 3         | 4                      |
| [Absolute,X](addressing.md#Absolute,X)     | $BC        | 3         | 4 (+1 if page crossed) |

See also: [LDA](#LDA%20-%20Load%20Accumulator), [LDX](#LDX%20-%20Load%20X%20Register)

### LSR - Logical Shift Right

A,C,Z,N = A/2 or M,C,Z,N = M/2

Each of the bits in A or M is shift one place to the right. The bit that was in bit 0 is shifted into the carry flag.
Bit 7 is set to zero.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set to contents of old bit 0      |
|-----------------------------------------|---------------------------------------------------------|-----------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if result = 0                 |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                      |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                      |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                      |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected                      |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of the result is set |

| **Addressing Mode**                        | **Opcode** | **Bytes** | **Cycles** |
|--------------------------------------------|------------|-----------|------------|
| [Accumulator](addressing.md#Accumulator)   | $4A        | 1         | 2          |
| [Zero Page](addressing.md#Zero%20Page)     | $46        | 2         | 5          |
| [Zero Page,X](addressing.md#Zero%20Page,X) | $56        | 2         | 6          |
| [Absolute](addressing.md#Absolute)         | $4E        | 3         | 6          |
| [Absolute,X](addressing.md#Absolute,X)     | $5E        | 3         | 7          |

See also: [ASL](#ASL%20-%20Arithmetic%20Shift%20Left), [ROL](#ROL%20-%20Rotate%20Left), [ROR](#ROR%20-%20Rotate%20Right)

### NOP - No Operation

The NOP instruction causes no changes to the processor other than the normal incrementing of the program counter to the
next instruction.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $EA        | 1         | 2          |

### ORA - Logical Inclusive OR

A,Z,N = A|M

An inclusive OR is performed, bit by bit, on the accumulator contents using the contents of a byte of memory.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected     |
|-----------------------------------------|---------------------------------------------------------|------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0     |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected     |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected     |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected     |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected     |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 set |

| **Addressing Mode**                              | **Opcode** | **Bytes** | **Cycles**             |
|--------------------------------------------------|------------|-----------|------------------------|
| [Immediate](addressing.md#Immediate)             | $09        | 2         | 2                      |
| [Zero Page](addressing.md#Zero%20Page)           | $05        | 2         | 3                      |
| [Zero Page,X](addressing.md#Zero%20Page,X)       | $15        | 2         | 4                      |
| [Absolute](addressing.md#Absolute)               | $0D        | 3         | 4                      |
| [Absolute,X](addressing.md#Absolute,X)           | $1D        | 3         | 4 (+1 if page crossed) |
| [Absolute,Y](addressing.md#Absolute,Y)           | $19        | 3         | 4 (+1 if page crossed) |
| [(Indirect,X)](addressing.md#Indexed%20Indirect) | $01        | 2         | 6                      |
| [(Indirect),Y](addressing.md#Indirect%20Indexed) | $11        | 2         | 5 (+1 if page crossed) |

See also: [AND](#AND%20-%20Logical%20AND), [EOR](#EOR%20-%20Exclusive%20OR)

### PHA - Push Accumulator

Pushes a copy of the accumulator on to the stack.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $48        | 1         | 3          |

See also: [PLA](#PLA%20-%20Pull%20Accumulator)

### PHP - Push Processor Status

Pushes a copy of the status flags on to the stack.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $08        | 1         | 3          |

See also: [PLP](#PLP%20-%20Pull%20Processor%20Status)

### PLA - Pull Accumulator

Pulls an 8 bit value from the stack and into the accumulator. The zero and negative flags are set as appropriate.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of A is set |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $68        | 1         | 4          |

See also: [PHA](#PHA%20-%20Push%20Accumulator)

### PLP - Pull Processor Status

Pulls an 8 bit value from the stack and into the processor flags. The flags will take on new states as determined by the
value pulled.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set from stack |
|-----------------------------------------|---------------------------------------------------------|----------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set from stack |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Set from stack |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Set from stack |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Set from stack |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Set from stack |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set from stack |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $28        | 1         | 4          |

See also: [PHP](#PHP%20-%20Push%20Processor%20Status)

### ROL - Rotate Left

Move each of the bits in either A or M one place to the left. Bit 0 is filled with the current value of the carry flag
whilst the old bit 7 becomes the new carry flag value.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set to contents of old bit 7      |
|-----------------------------------------|---------------------------------------------------------|-----------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0                      |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                      |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                      |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                      |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected                      |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of the result is set |

| **Addressing Mode**                        | **Opcode** | **Bytes** | **Cycles** |
|--------------------------------------------|------------|-----------|------------|
| [Accumulator](addressing.md#Accumulator)   | $2A        | 1         | 2          |
| [Zero Page](addressing.md#Zero%20Page)     | $26        | 2         | 5          |
| [Zero Page,X](addressing.md#Zero%20Page,X) | $36        | 2         | 6          |
| [Absolute](addressing.md#Absolute)         | $2E        | 3         | 6          |
| [Absolute,X](addressing.md#Absolute,X)     | $3E        | 3         | 7          |

See
also: [ASL](#ASL%20-%20Arithmetic%20Shift%20Left), [LSR](#LSR%20-%20Logical%20Shift%20Right), [ROR](#ROR%20-%20Rotate%20Right)

### ROR - Rotate Right

Move each of the bits in either A or M one place to the right. Bit 7 is filled with the current value of the carry flag
whilst the old bit 0 becomes the new carry flag value.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set to contents of old bit 0      |
|-----------------------------------------|---------------------------------------------------------|-----------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0                      |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                      |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                      |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                      |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected                      |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of the result is set |

| **Addressing Mode**                        | **Opcode** | **Bytes** | **Cycles** |
|--------------------------------------------|------------|-----------|------------|
| [Accumulator](addressing.md#Accumulator)   | $6A        | 1         | 2          |
| [Zero Page](addressing.md#Zero%20Page)     | $66        | 2         | 5          |
| [Zero Page,X](addressing.md#Zero%20Page,X) | $76        | 2         | 6          |
| [Absolute](addressing.md#Absolute)         | $6E        | 3         | 6          |
| [Absolute,X](addressing.md#Absolute,X)     | $7E        | 3         | 7          |

See
also [ASL](#ASL%20-%20Arithmetic%20Shift%20Left), [LSR](#LSR%20-%20Logical%20Shift%20Right), [ROL](#ROL%20-%20Rotate%20Left)

### RTI - Return from Interrupt

The RTI instruction is used at the end of an interrupt processing routine. It pulls the processor flags from the stack
followed by the program counter.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set from stack |
|-----------------------------------------|---------------------------------------------------------|----------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set from stack |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Set from stack |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Set from stack |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Set from stack |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Set from stack |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set from stack |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $40        | 1         | 6          |

### RTS - Return from Subroutine

The RTS instruction is used at the end of a subroutine to return to the calling routine. It pulls the program counter (
minus one) from the stack.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $60        | 1         | 6          |

See also: [JSR](#JSR%20-%20Jump%20to%20Subroutine)

### SBC - Subtract with Carry

A,Z,C,N = A-M-(1-C)

This instruction subtracts the contents of a memory location to the accumulator together with the not of the carry bit.
If overflow occurs the carry bit is clear, this enables multiple byte subtraction to be performed.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Clear if overflow in bit 7   |
|-----------------------------------------|---------------------------------------------------------|------------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0                 |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected                 |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected                 |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected                 |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Set if sign bit is incorrect |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 set             |

| **Addressing Mode**                              | **Opcode** | **Bytes** | **Cycles**             |
|--------------------------------------------------|------------|-----------|------------------------|
| [Immediate](addressing.md#Immediate)             | $E9        | 2         | 2                      |
| [Zero Page](addressing.md#Zero%20Page)           | $E5        | 2         | 3                      |
| [Zero Page,X](addressing.md#Zero%20Page,X)       | $F5        | 2         | 4                      |
| [Absolute](addressing.md#Absolute)               | $ED        | 3         | 4                      |
| [Absolute,X](addressing.md#Absolute,X)           | $FD        | 3         | 4 (+1 if page crossed) |
| [Absolute,Y](addressing.md#Absolute,Y)           | $F9        | 3         | 4 (+1 if page crossed) |
| [(Indirect,X)](addressing.md#Indexed%20Indirect) | $E1        | 2         | 6                      |
| [(Indirect),Y](addressing.md#Indirect%20Indexed) | $F1        | 2         | 5 (+1 if page crossed) |

See also: [ADC](#ADC%20-%20Add%20with%20Carry)

### SEC - Set Carry Flag

C = 1

Set the carry flag to one.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Set to 1     |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $38        | 1         | 2          |

See also: [CLC](#CLC%20-%20Clear%20Carry%20Flag)

### SED - Set Decimal Flag

D = 1

Set the decimal mode flag to one.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Set to 1     |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $F8        | 1         | 2          |

See also: [CLD](#CLD%20-%20Clear%20Decimal%20Mode)

### SEI - Set Interrupt Disable

I = 1

Set the interrupt disable flag to one.

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Set to 1     |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $78        | 1         | 2          |

See also: [CLI](#CLI%20-%20Clear%20Interrupt%20Disable)

### STA - Store Accumulator

M = A

Stores the contents of the accumulator into memory.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                              | **Opcode** | **Bytes** | **Cycles** |
|--------------------------------------------------|------------|-----------|------------|
| [Zero Page](addressing.md#Zero%20Page)           | $85        | 2         | 3          |
| [Zero Page,X](addressing.md#Zero%20Page,X)       | $95        | 2         | 4          |
| [Absolute](addressing.md#Absolute)               | $8D        | 3         | 4          |
| [Absolute,X](addressing.md#Absolute,X)           | $9D        | 3         | 5          |
| [Absolute,Y](addressing.md#Absolute,Y)           | $99        | 3         | 5          |
| [(Indirect,X)](addressing.md#Indexed%20Indirect) | $81        | 2         | 6          |
| [(Indirect),Y](addressing.md#Indirect%20Indexed) | $91        | 2         | 6          |

See also: [STX](#STX%20-%20Store%20X%20Register), [STY](#STY%20-%20Store%20Y%20Register)

### STX - Store X Register

M = X

Stores the contents of the X register into memory.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                        | **Opcode** | **Bytes** | **Cycles** |
|--------------------------------------------|------------|-----------|------------|
| [Zero Page](addressing.md#Zero%20Page)     | $86        | 2         | 3          |
| [Zero Page,Y](addressing.md#Zero%20Page,Y) | $96        | 2         | 4          |
| [Absolute](addressing.md#Absolute)         | $8E        | 3         | 4          |

See also: [STA](#STA%20-%20Store%20Accumulator), [STY](#STY%20-%20Store%20Y%20Register)

### STY - Store Y Register

M = Y

Stores the contents of the Y register into memory.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**                        | **Opcode** | **Bytes** | **Cycles** |
|--------------------------------------------|------------|-----------|------------|
| [Zero Page](addressing.md#Zero%20Page)     | $84        | 2         | 3          |
| [Zero Page,X](addressing.md#Zero%20Page,X) | $94        | 2         | 4          |
| [Absolute](addressing.md#Absolute)         | $8C        | 3         | 4          |

See also: [STA](#STA%20-%20Store%20Accumulator), [STX](#STX%20-%20Store%20X%20Register)

### TAX - Transfer Accumulator to X

X = A

Copies the current contents of the accumulator into the X register and sets the zero and negative flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if X = 0             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of X is set |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $AA        | 1         | 2          |

See also: [TXA](#TXA%20-%20Transfer%20X%20to%20Accumulator)

### TAY - Transfer Accumulator to Y

Y = A

Copies the current contents of the accumulator into the Y register and sets the zero and negative flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if Y = 0             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of Y is set |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $A8        | 1         | 2          |

See also: [TYA](#TAY%20-%20Transfer%20Accumulator%20to%20Y)

### TSX - Transfer Stack Pointer to X

X = S

Copies the current contents of the stack register into the X register and sets the zero and negative flags as
appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if X = 0             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of X is set |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $BA        | 1         | 2          |

See also: [TXS](#TXS%20-%20Transfer%20X%20to%20Stack%20Pointer)

### TXA - Transfer X to Accumulator

A = X

Copies the current contents of the X register into the accumulator and sets the zero and negative flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of A is set |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $8A        | 1         | 2          |

See also: [TAX](#TAX%20-%20Transfer%20Accumulator%20to%20X)

### TXS - Transfer X to Stack Pointer

S = X

Copies the current contents of the X register into the stack register.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected |
|-----------------------------------------|---------------------------------------------------------|--------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Not affected |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Not affected |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $9A        | 1         | 2          |

See also: [TSX](#TSX%20-%20Transfer%20Stack%20Pointer%20to%20X)

### TYA - Transfer Y to Accumulator

A = Y

Copies the current contents of the Y register into the accumulator and sets the zero and negative flags as appropriate.

Processor Status after use:

| [C](registers.md#Carry%20Flag)          | [Carry Flag](registers.md#Carry%20Flag)                 | Not affected             |
|-----------------------------------------|---------------------------------------------------------|--------------------------|
| [Z](registers.md#Zero%20Flag)           | [Zero Flag](registers.md#Zero%20Flag)                   | Set if A = 0             |
| [I](registers.md#Interrupt%20Disable)   | [Interrupt Disable](registers.md#Interrupt%20Disable)   | Not affected             |
| [D](registers.md#Decimal%20Mode%20Flag) | [Decimal Mode Flag](registers.md#Decimal%20Mode%20Flag) | Not affected             |
| [B](registers.md#Break%20Command)       | [Break Command](registers.md#Break%20Command)           | Not affected             |
| [V](registers.md#Overflow%20Flag)       | [Overflow Flag](registers.md#Overflow%20Flag)           | Not affected             |
| [N](registers.md#Negative%20Flag)       | [Negative Flag](registers.md#Negative%20Flag)           | Set if bit 7 of A is set |

| **Addressing Mode**               | **Opcode** | **Bytes** | **Cycles** |
|-----------------------------------|------------|-----------|------------|
| [Implied](addressing.md#Implicit) | $98        | 1         | 2          |

See also: [TAY](#TAY%20-%20Transfer%20Accumulator%20to%20Y)

---

This page was last updated on 17th February, 2008