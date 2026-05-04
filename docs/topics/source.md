---
title: "6502.org Source Code Library"
source: "https://6502.org/source/"
tags:
  - 6502
  - source
  - code
---

# Source Code Library

Here you will find a collection of useful programs and routines (mostly written in assembly language) that can be
adapted to any 6502-based computer.

## Games

### MicroChess

by Peter Jennings, modified by Daryl Rictor

MicroChess is a very famous game written by Peter Jennings in 1976. It was originally written for the MOS Technology
KIM-1 but is easily ported and was used on other systems like the Rockwell AIM-65. Daryl Rictor has updated the program
for use on projects with a 6551-based serial port. Both the original and the enhanced version are provided.

## Floating Point Math

### Floating Point Routines for the 6502

by Roy Rankin and Steve Wozniak

Originally published in the August 1976 issue of Dr. Dobb's Journal, these floating point routines allow 6502 users to
perform most of the more popular and desired floating point and transcendental functions, namely: Natural Log, Common
Log, Addition, Subtraction, Multiplication, Division, and conversions between floating and fixed point numbers.

### Errata for Rankin's 6502 Floating Point Routines

by Roy Rankin

In the November/December issue of Dr. Dobb's Journal Roy Rankin published three error corrections to the Floating Point
Routines presented above.

### Floating Point Implementation in the Apple II

by Steve Wozniak

An almost identical set of the above routines appeared in the original manual for the Apple II (the Red Book, January
1978). Documentation for these routines appeared in another book, the Wozpak II, in November 1979.

### IEEE 754 quadruple precision floating point for 65C816

by Marco Granati

The main goal is build a "virtual numeric co-processor" with some internal pseudo-registers stack-oriented, interfaced
by 'cop' instruction. About 16k binary, 10k lines of source. See announcement.

### Calc65 floating-point arithmetic and transcendental function package

by Charles R Bond

"The math routines use Cordic algorithms and similar strategies involving pseudo-multiplication and pseudo-division. The
square root routine employs non-restoring pseudo-division. The BCD number format occupies 8 bytes: note that there are
12 mantissa digits and 3 exponent digits."

## General Purpose

### Calculation of Clock Frequency

by Leo Nechaev

This program calculates the clock frequency of the 6502-based system it is running on and reports it back via RS-232
through a 6551 ACIA.

### Clearing a Section of Memory

from 6502 Software Gourmet Guide & Cookbook, Simplified by Mickael Pointier

Clear up to 256 bytes anywhere in memory using this routine.

### Day of the Week

by Paul Guertin

Computes the day of the week from (almost) any given date within the years 1900 and 2155.

### Enhanced Day of the Week

by W.J. Brier

This is an alternative to Paul Guertin's program. It is longer but has the advantage of supporting a wider range of
dates.

### Practical Memory Move Routines

by Bruce Clark

Here are some reasonably fast general-purpose routines for moving blocks of memory. You simply specify the address to
move from, the address to move to, and the size of the block.

### Address Line Test

by Tony Gonzalez

The typical RAM test consists of reading and writing $AA and $55 to each memory location. This works well for finding
bad RAM locations. However, it does not specifically find bad address lines. This simple method does just that.

### Efficient Nybble Swap

by Garth Wilson

Efficient nybble swap on 6502, taking only 8 bytes and 12 clock cycles, and no variables, no stack usage, no look-up
table, no X or Y usage. It uses only the accumulator and status register.

## Integer Math

### Large Multiply & Divide

from 6502 Software Design, Expanded by Greg

Two routines for multiplying and dividing larger integers. The first is a 32-bit multiply routine which computes a
64-bit product and the second is a 64-bit divide routine which gives a 32-bit quotient.

### Fast Multiply by 10

by Leo Nechaev

Multiplying a number by ten is fast and easy with this routine.

### Fast, Table-Driven Multiplication

by Martin Arndt

This routine uses tables of squares to quickly multiply two 8-bit numbers. The algorithm is very useful when you want to
multiply multiple bytes by the same factor, and could easily be expanded for 16-bit numbers.

### Division (32-bit)

by Garth Wilson

Routines for unsigned division of a 32-bit dividend by a 16-bit divisor, resulting with a 16-bit quotient and a 16-bit
remainder. Presented for the 65C02 (with changes for use on 6502) and 65C816, with notes of interest to Forth users.

### Large Look-up Tables for 16-Bit Fixed-Point Math

by Garth Wilson

Uses large look-up tables for very fast 16-bit fixed-point math, including trig and log functions. Also includes an
introduction to fixed-point/scaled-integer math.

### BCD to Binary Conversion (32-bit)

from 6502 Software Design, Expanded by Greg

Two routines for converting between larger binary and BCD numbers. The first routine converts a 32-bit packed BCD number
to its binary equivalent and the second performs the reverse operation.

### Hexadecimal to Decimal Conversion

by Garth Wilson

Routines for converting 8- and 16-bit hexadecimal numbers to their decimal (BCD) equivalents.

### More Hexadecimal to Decimal Conversion

by Andrew Jacobs

Alternate versions of Garth's conversion routines, slower but eliminating the lookup tables.

### CRC Calculations

by Paul Guertin

Three routines for computing 32-, 16-, and 8-bit CRC (Cyclic Redundancy Check) error detecting codes.

### More CRC Calculations

by Greg Cook

Alternate versions of Paul's CRC calculation routines that execute in constant cycle time and eliminate the lookup
tables.

### Permutation Generator

by Paul Guertin

Generating permutations is usually done with a recursive procedure, but here is a cute iterative routine that does it
simply and efficiently.

### Square Calculator

by Lee Davison

Calculates the 16-bit unsigned integer square of a signed 16-bit integer.

### Root Calculator

by Lee Davison

Calculates the 8-bit root and 9-bit remainder of a 16-bit unsigned integer.

### Linear Congruential Pseudo-Random Number Generator Routines

by Bruce Clark

A set of routines for generating pseudo-random numbers, using a linear congruential type generator.

## I/O Techniques

### Counting Push-Button Switch Closures, With Debouncing

from 6502 Software Design

Counts closures on a push-button switch until a second push-button is pushed. Uses 6522 VIA Timer #1 to generate a
10-millisecond delay for the debounce routine.

### One-Shot Timer Examples for the 6522

from 6502 Applications

Two small programs demonstrate how to use Timer 1 and 2 in one-shot mode and also a short explanation of how to generate
a pulse.

### Identifying 6522 Internal Interrupt Sources

from 6502 Applications

A simple program to identify which of the seven possible sources caused a 6522 interrupt.

### Interrupt-Serviced 256-Byte Data Buffer

by Lee Davison

Demonstrates how to set up a data buffer suitable for many operations such as sending data to a printer or serial
device.

### Print Immediate (PRIMM)

from Commodore 128 KERNAL ROM

By cleverly manipulating the stack this useful routine allows you to inline data to be printed with your assembly code,
almost like a BASIC PRINT statement.

### Intel HEX File Downloader

by Ross Archer

This program allows any 6502 system with a 6551-based serial port to painlessly download programs from a host computer
by sending them in the Intel HEX file format.

### X-Modem File Transfer Protocol

by Daryl Rictor

This implementation of the X-Modem file transfer protocol will work on any 6502-based computer with a 6551 serial port.
Using X-Modem allows PC's to easily exchange files with your project using a terminal program.

## Monitors, Assemblers, and Interpreters

### 2KSA [2KSA.PDF: 4.34MB]

by Robert Ford Denison

The 2K Symbolic Assembler is a two-pass assembler for 6502 systems which assembles to less than two kilobytes. This PDF
includes all 65 pages of the original documentation which includes commented source code and hex dumps for the program.
It is a really nice assembler considering its size and it so well documented that making changes to it are easy.

### FIG FORTH 6502 [FORTH65.ZIP: 32KB]

from the FORTH Interest Group

FORTH Interest Group (FIG) FORTH kernel for the 6502. It is written for the Rockwell System-65 but is easily adaptable
to your own systems. Note that this has the "UM/MOD" division bug. An explanation of the bug and a correction can be
found in Garth Wilson's 32-bit Division entry, which can also be found here in the Source Code Library.

### 32-bit DO LOOP & related words in 6502 ITC Forth [text source-code file]

by Garth Wilson

All in 32-bit (including the increment value for 2+LOOP): 32-bit equivalents for DO, ?DO, LOOP, +LOOP, I, BOUNDS, LEAVE,
?LEAVE, UNLOOP, and the internals, plus explanation. These augment, not replace, the normal 16-bit ones, putting "2" in
front of the names to tell the double-precision ones apart from the normal single-precision ones.

### Porting Sweet 16

by Carsten Strotmann

Sweet 16 is a metaprocessor or "pseudo microprocessor" implemented in 6502 assembly language. Originally written by
Steve Wozniak and used in the Apple II, Sweet 16 can also be ported to other 6502-based systems to provide useful 16-bit
functionality. This article includes the source code for Sweet 16, along with a brief history, programming instructions,
and notes to help port it.

### VTL02 (VTL-2 for the 6502)

by Mike Barry

VTL-2 is a very tiny language (the interpreter is under 1k.) Originally written in 1976 by Frank McCoy and Gary Shannon
for the 8800 and 6800, a port is now available for the 6502. There's a discussion on the forum with a series of updates.

### Supermon 816

by BigDumbDinosaur

Supermon 816 is a full-featured machine-language monitor that is adaptable to most computers that are powered by the
Western Design Center's W65C816S 16-bit microprocessor if operating in native mode.

## Operating System Kernels

### Mini Multitasking Kernel

by Joachim Deboy

This mini-kernel uses IRQ interrupts from a free running 6502 timer for task switching. It supports four concurrent
tasks, each receiving fixed time slices in a round-robin method.

## Sorting Algorithms

### Bubble Sort (for 8-bit Elements)

from 6502 Software Design

A short description of the bubble sort technique followed by a subroutine that arranges the 8-bit elements of a list in
ascending order.

### Bubble Sort (for 16-bit Elements)

from 6502 Software Design

The sorting routine above expanded to sort a list containing 16-bit elements.

### Combination Sort

by Daryl Rictor

The Combination Sort is a modified bubble sort (iterative) that takes much less time to complete.

### Optimal Sort (for 8-bit elements)

by Mats Rosengren

The Optimal Sort is another improvement on the bubble sort but uses a different approach than the Combination Sort.

### Optimal Sort (for 16-bit elements)

by Mats Rosengren

An expansion on the 8-bit Optimal Sort above, this version allows for sorting of large numbers (over 255) of 16-bit
elements.

### Radix 256 Counting Sort

by Dwight Elvey

A radix 256 counting sort to demonstrate that other sorts can be significantly faster than QuickSort. This
implementation does a signed 16-bit integer sort of 1K of data.

### Shell Sort (for 16-bit elements)

by Fredrik Ramsberg

The Shell Sort is a highly-efficient sorting algorithm. This implementation can sort a large number (>32,000) of 16-bit
elements. Sorting a completely unsorted array of 10,000 elements takes just over thirty seconds on a 1 Mhz machine.

## String Manipulations

### Pattern Matcher

by Paul Guertin

Matches a string against a pattern and returns with the carry bit set if they match, or clear if they don't. Supports
wildcard characters "?" and "*".

### Simple Comparisons

by Jonathyn Bet'nct

Compare two strings in memory stored in Pascal string format (length byte followed by string). Case-sensitive and
-insensitive versions are provided.

### Convert ASCII Number String to 32-bit Binary

by BigDumbDinosaur

This 6502 assembly language program converts a null-terminated ASCII number string into a 32-bit unsigned binary value
in little-endian format. It can accept a number in binary, octal, decimal or hexadecimal format.

### Convert 32-bit Binary to ASCII Number String

by BigDumbDinosaur

This program converts 32-bit binary back to an ASCII string. It is a companion to "Convert ASCII Number String to 32-bit
Binary" above.

### W65C816 String Manipulation Library [816STRINGLIB.ZIP: 648KB]

by BigDumbDinosaur

This W65C816S string manipulation library is a collection of assembly language subroutines that perform common character
string operations, such as copying and concatenating.

## Categories

- Games
- Floating Point Math
- General Purpose
- Integer Math
- I/O Techniques
- Monitors, Assemblers, Interpreters
- Operating System Kernels
- Sorting Algorithms
- String Manipulations

## Related Sites

The Fridge -- 6502 assembly language routines for the Commodore 64 and other computers.

Apple Assembly Line -- archive of an old Apple II publication with a lot of good general 6502 content.

Codebase 64 -- wiki with general 6502 code examples and material specific to the C64 and other Commodore machines.

## Contribute

Have source code that should be included? Send it to us in any format.

6502.org is an ongoing project by Mike Naberezny and contributors.

About · News · GitHub
