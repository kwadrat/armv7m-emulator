package core

import (
    "reflect"
    "testing"
)

func TestIdentifyAddRegT1(t *testing.T) {
    cases := []IdentifyCase{
        {instr: FetchedInstr16(0x1800), instr_valid: true},  // adds r0, r0, r0
        {instr: FetchedInstr16(0x19ff), instr_valid: true},  // adds r7, r7, r7
        {instr: FetchedInstr16(0x18d1), instr_valid: true},  // adds r1, r2, r3
        {instr: FetchedInstr16(0x0000), instr_valid: false}, // mov r0, r0
        {instr: FetchedInstr16(0x001f), instr_valid: false}, // mov r7, r3
        {instr: FetchedInstr16(0x2000), instr_valid: false}, // mov r0, #0
        {instr: FetchedInstr16(0x2745), instr_valid: false}, // mov r7, #0x45
        {instr: FetchedInstr16(0x4080), instr_valid: false}, // lsl r0, r0, r0
        {instr: FetchedInstr16(0xffff), instr_valid: false},
    }

    test_identify(t, cases, reflect.TypeOf(AddRegT1{}))
}

func TestDecodeAddReg16T1(t *testing.T) {
    cases := []DecodeCase{
        // adds r0, r0, r0
        {instr: FetchedInstr16(0x1800), decoded: AddRegT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT}},
        // adds r7, r7, r7
        {instr: FetchedInstr16(0x19ff), decoded: AddRegT1{Rd: 7, Rm: 7, Rn: 7, Imm: 0, setflags: NOT_IT}},
        // adds r1, r2, r3
        {instr: FetchedInstr16(0x18d1), decoded: AddRegT1{Rd: 1, Rm: 3, Rn: 2, Imm: 0, setflags: NOT_IT}},
    }

    test_decode(t, cases, AddReg16T1)
}

func TestExecuteAddRegT1(t *testing.T) {
    cases := []ExecuteCase{
        // adds r0, r0, r0
        {instr: AddRegT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
            expected: Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: true}}},
        // adds r0, r0, r0
        {instr: AddRegT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: true}},
            expected: Registers{r: GeneralRegs{2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: false}}},
        // adds r1, r2, r3
        {instr: AddRegT1{Rd: 1, Rm: 3, Rn: 2, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
            expected: Registers{r: GeneralRegs{1, 7, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}}},
        // adds r0, r1, r2
        {instr: AddRegT1{Rd: 0, Rm: 2, Rn: 1, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 0x7fffffff, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
            expected: Registers{r: GeneralRegs{0x80000000, 0x7fffffff, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Apsr: Apsr{V: true, N: true}}},
        // adds r0, r1, r1
        {instr: AddRegT1{Rd: 0, Rm: 1, Rn: 1, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 0x80000000, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
            expected: Registers{r: GeneralRegs{0, 0x80000000, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Apsr: Apsr{Z: true, C: true, V: true}}},
    }

    test_execute(t, cases)
}

func TestIdentifyAddRegT2(t *testing.T) {
    cases := []IdentifyCase{
        {instr: FetchedInstr16(0x4400), instr_valid: true},  // add r0, r0
        {instr: FetchedInstr16(0x4487), instr_valid: true},  // add pc, r0
        {instr: FetchedInstr16(0x44f7), instr_valid: true},  // add pc, lr
        {instr: FetchedInstr16(0x4468), instr_valid: false}, // add r0, sp -> add r0, sp, r0 (sp encoding)
        {instr: FetchedInstr16(0x4485), instr_valid: false}, // add sp, r0 -> add sp, r0 (sp encoding)
        {instr: FetchedInstr16(0x1800), instr_valid: false}, // adds r0, r0, r0
        {instr: FetchedInstr16(0x19ff), instr_valid: false}, // adds r7, r7, r7
        {instr: FetchedInstr16(0x18d1), instr_valid: false}, // adds r1, r2, r3
        {instr: FetchedInstr16(0x2745), instr_valid: false}, // mov r7, #0x45
        {instr: FetchedInstr16(0x4080), instr_valid: false}, // lsl r0, r0, r0
        {instr: FetchedInstr16(0xffff), instr_valid: false},
    }

    test_identify(t, cases, reflect.TypeOf(AddRegT2{}))
}

func TestDecodeAddReg16T2(t *testing.T) {
    cases := []DecodeCase{
        // add r0, r0
        {instr: FetchedInstr16(0x4400), decoded: AddRegT2{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NEVER}},
        // add pc, r0
        {instr: FetchedInstr16(0x4487), decoded: AddRegT2{Rd: PC, Rm: 0, Rn: PC, Imm: 0, setflags: NEVER}},
        // add pc, lr
        {instr: FetchedInstr16(0x44f7), decoded: AddRegT2{Rd: PC, Rm: LR, Rn: PC, Imm: 0, setflags: NEVER}},
        // add r0, sp -> add r0, sp, r0 (sp encoding)
        {instr: FetchedInstr16(0x4468), decoded: AddRegSPT1{Rd: 0, Rm: 0, Rn: SP, Imm: 0, setflags: NEVER}},
        // add sp, r0 -> add sp, r0 (sp encoding)
        {instr: FetchedInstr16(0x4485), decoded: AddRegSPT2{Rd: SP, Rm: 0, Rn: SP, Imm: 0, setflags: NEVER}},
    }

    test_decode(t, cases, AddReg16T2)
}

var share_t *testing.T

func TestExecuteAddRegT2(t *testing.T) {
    cases := []ExecuteCase{
        // add r0, r0
        {instr: AddRegT2{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
            expected: Registers{r: GeneralRegs{2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}}},
        // add pc, r0
        {instr: AddRegT2{Rd: PC, Rm: 0, Rn: PC, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{4, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, pc: 1000},
            expected: Registers{r: GeneralRegs{4, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, pc: 1004}},
        // add pc, lr
        {instr: AddRegT2{Rd: PC, Rm: LR, Rn: PC, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, pc: 1000, lr: 2000},
            expected: Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, pc: 3000, lr: 2000}},
    }

    share_t = t

    test_execute(t, cases)
}

func TestIdentifyAddRegSPT1(t *testing.T) {
    cases := []IdentifyCase{
        {instr: FetchedInstr16(0x4468), instr_valid: true},  // add r0, sp, r0
        {instr: FetchedInstr16(0x446f), instr_valid: true},  // add r7, sp, r7
        {instr: FetchedInstr16(0x44ef), instr_valid: true},  // add pc, sp, pc
        {instr: FetchedInstr16(0x44ed), instr_valid: true},  // add sp, sp, sp
        {instr: FetchedInstr16(0x1800), instr_valid: false}, // adds r0, r0, r0
        {instr: FetchedInstr16(0x19ff), instr_valid: false}, // adds r7, r7, r7
        {instr: FetchedInstr16(0x18d1), instr_valid: false}, // adds r1, r2, r3
        {instr: FetchedInstr16(0x0000), instr_valid: false}, // mov r0, r0
        {instr: FetchedInstr16(0x001f), instr_valid: false}, // mov r7, r3
        {instr: FetchedInstr16(0xffff), instr_valid: false},
    }

    test_identify(t, cases, reflect.TypeOf(AddRegSPT1{}))
}

func TestDecodeAddRegSP16T1(t *testing.T) {
    cases := []DecodeCase{
        // add r0, sp, r0
        {instr: FetchedInstr16(0x4468), decoded: AddRegSPT1{Rd: 0, Rm: 0, Rn: SP, Imm: 0, setflags: NEVER}},
        // add r7, sp, r7
        {instr: FetchedInstr16(0x446f), decoded: AddRegSPT1{Rd: 7, Rm: 7, Rn: SP, Imm: 0, setflags: NEVER}},
        // add pc, sp, pc
        {instr: FetchedInstr16(0x44ef), decoded: AddRegSPT1{Rd: PC, Rm: PC, Rn: SP, Imm: 0, setflags: NEVER}},
        // add sp, sp, sp
        {instr: FetchedInstr16(0x44ed), decoded: AddRegSPT1{Rd: SP, Rm: SP, Rn: SP, Imm: 0, setflags: NEVER}},
    }

    test_decode(t, cases, AddRegSP16T1)
}

func TestExecuteAddRegSPT1(t *testing.T) {
    cases := []ExecuteCase{
        // add r0, sp, r0
        {instr: AddRegSPT1{Rd: 0, Rm: 0, Rn: SP, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, sp: SPRegs{1, 0}, Control: Control{Spsel: MSP}},
            expected: Registers{r: GeneralRegs{2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, sp: SPRegs{1, 0}, Control: Control{Spsel: MSP}}},
        // add r0, sp, r0
        {instr: AddRegSPT1{Rd: 0, Rm: 0, Rn: SP, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, sp: SPRegs{0, 1}, Control: Control{Spsel: PSP}},
            expected: Registers{r: GeneralRegs{2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, sp: SPRegs{0, 1}, Control: Control{Spsel: PSP}}},
        // add r7, sp, r7
        {instr: AddRegSPT1{Rd: 7, Rm: 7, Rn: SP, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 0xffffffff, 9, 10, 11, 12, 13}, sp: SPRegs{1, 0}, Control: Control{Spsel: MSP}},
            expected: Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 0, 9, 10, 11, 12, 13}, sp: SPRegs{1, 0}, Control: Control{Spsel: MSP}}},
        // add pc, sp, pc
        {instr: AddRegSPT1{Rd: PC, Rm: PC, Rn: SP, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, pc: 0x80000000, sp: SPRegs{4, 0}, Control: Control{Spsel: MSP}},
            expected: Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, pc: 0x80000004, sp: SPRegs{4, 0}, Control: Control{Spsel: MSP}}},
        // add sp, sp, sp
        {instr: AddRegSPT1{Rd: SP, Rm: SP, Rn: SP, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, sp: SPRegs{4, 0}, Control: Control{Spsel: MSP}},
            expected: Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, sp: SPRegs{8, 0}, Control: Control{Spsel: MSP}}},
    }

    test_execute(t, cases)
}

func TestIdentifyAddRegSPT2(t *testing.T) {
    cases := []IdentifyCase{
        {instr: FetchedInstr16(0x4485), instr_valid: true},  // add sp, r0
        {instr: FetchedInstr16(0x44fd), instr_valid: true},  // add sp, pc
        {instr: FetchedInstr16(0x44ed), instr_valid: false}, // add sp, sp -> add sp, sp, sp
        {instr: FetchedInstr16(0x4468), instr_valid: false}, // add r0, sp, r0
        {instr: FetchedInstr16(0x446f), instr_valid: false}, // add r7, sp, r7
        {instr: FetchedInstr16(0x44ef), instr_valid: false}, // add pc, sp, pc
        {instr: FetchedInstr16(0x1800), instr_valid: false}, // adds r0, r0, r0
        {instr: FetchedInstr16(0x19ff), instr_valid: false}, // adds r7, r7, r7
        {instr: FetchedInstr16(0x18d1), instr_valid: false}, // adds r1, r2, r3
        {instr: FetchedInstr16(0x0000), instr_valid: false}, // mov r0, r0
        {instr: FetchedInstr16(0x001f), instr_valid: false}, // mov r7, r3
        {instr: FetchedInstr16(0xffff), instr_valid: false},
    }

    test_identify(t, cases, reflect.TypeOf(AddRegSPT2{}))
}

func TestDecodeAddRegSP16T2(t *testing.T) {
    cases := []DecodeCase{
        // add sp, r0
        {instr: FetchedInstr16(0x4485), decoded: AddRegSPT2{Rd: SP, Rm: 0, Rn: SP, Imm: 0, setflags: NEVER}},
        // add sp, pc
        {instr: FetchedInstr16(0x44fd), decoded: AddRegSPT2{Rd: SP, Rm: PC, Rn: SP, Imm: 0, setflags: NEVER}},
        // add sp, sp -> add sp, sp, sp
        {instr: FetchedInstr16(0x44ed), decoded: AddRegSPT1{Rd: SP, Rm: SP, Rn: SP, Imm: 0, setflags: NEVER}},
    }

    test_decode(t, cases, AddRegSP16T2)
}

func TestExecuteAddRegSPT2(t *testing.T) {
    cases := []ExecuteCase{
        // add sp, r0
        {instr: AddRegSPT2{Rd: SP, Rm: 0, Rn: SP, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, sp: SPRegs{0, 0}, Control: Control{Spsel: MSP}},
            expected: Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, sp: SPRegs{1, 0}, Control: Control{Spsel: MSP}}},
        // add sp, r0
        {instr: AddRegSPT2{Rd: SP, Rm: 0, Rn: SP, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, sp: SPRegs{0, 1}, Control: Control{Spsel: PSP}},
            expected: Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, sp: SPRegs{0, 2}, Control: Control{Spsel: PSP}}},
        // add sp, pc
        {instr: AddRegSPT2{Rd: SP, Rm: PC, Rn: SP, Imm: 0, setflags: NEVER},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, pc: 0x80000000, sp: SPRegs{4, 0}, Control: Control{Spsel: MSP}},
            expected: Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, pc: 0x80000000, sp: SPRegs{0x80000004, 0}, Control: Control{Spsel: MSP}}},
    }

    test_execute(t, cases)
}

func TestIdentifyAddImmT1(t *testing.T) {
    cases := []IdentifyCase{
        {instr: FetchedInstr16(0x1c00), instr_valid: true},  // adds r0, r0, #0
        {instr: FetchedInstr16(0x1dff), instr_valid: true},  // adds r7, r7, #7
        {instr: FetchedInstr16(0x1cd1), instr_valid: true},  // adds r1, r2, #3
        {instr: FetchedInstr16(0x1800), instr_valid: false}, // adds r0, r0, r0
        {instr: FetchedInstr16(0x19ff), instr_valid: false}, // adds r7, r7, r7
        {instr: FetchedInstr16(0x18d1), instr_valid: false}, // adds r1, r2, r3
        {instr: FetchedInstr16(0x0000), instr_valid: false}, // mov r0, r0
        {instr: FetchedInstr16(0x2745), instr_valid: false}, // mov r7, #0x45
        {instr: FetchedInstr16(0x4080), instr_valid: false}, // lsl r0, r0, r0
        {instr: FetchedInstr16(0xffff), instr_valid: false},
    }

    test_identify(t, cases, reflect.TypeOf(AddImmT1{}))
}

func TestDecodeAddImm16T1(t *testing.T) {
    cases := []DecodeCase{
        // adds r0, r0, #0
        {instr: FetchedInstr16(0x1c00), decoded: AddImmT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT}},
        // adds r7, r7, #7
        {instr: FetchedInstr16(0x1dff), decoded: AddImmT1{Rd: 7, Rm: 0, Rn: 7, Imm: 7, setflags: NOT_IT}},
        // adds r1, r2, #3
        {instr: FetchedInstr16(0x1cd1), decoded: AddImmT1{Rd: 1, Rm: 0, Rn: 2, Imm: 3, setflags: NOT_IT}},
    }

    test_decode(t, cases, AddImm16T1)
}

func TestExecuteAddImmT1(t *testing.T) {
    cases := []ExecuteCase{
        // adds r0, r0, #0
        {instr: AddImmT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
            expected: Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: true}}},
        // adds r0, r0, #1
        {instr: AddImmT1{Rd: 0, Rm: 0, Rn: 0, Imm: 1, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: true}},
            expected: Registers{r: GeneralRegs{2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: false}}},
        // adds r1, r2, #3
        {instr: AddImmT1{Rd: 1, Rm: 0, Rn: 2, Imm: 3, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
            expected: Registers{r: GeneralRegs{1, 6, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}}},
        // adds r0, r1, #1
        {instr: AddImmT1{Rd: 0, Rm: 0, Rn: 1, Imm: 1, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 0x7fffffff, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
            expected: Registers{r: GeneralRegs{0x80000000, 0x7fffffff, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Apsr: Apsr{V: true, N: true}}},
        // adds r0, r1, #1
        {instr: AddImmT1{Rd: 0, Rm: 0, Rn: 1, Imm: 1, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 0xffffffff, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
            expected: Registers{r: GeneralRegs{0, 0xffffffff, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Apsr: Apsr{Z: true, C: true}}},
    }

    test_execute(t, cases)
}

func TestIdentifyAddImmT2(t *testing.T) {
    cases := []IdentifyCase{
        {instr: FetchedInstr16(0x3000), instr_valid: true},  // adds r0, #0
        {instr: FetchedInstr16(0x37ff), instr_valid: true},  // adds r7, #255
        {instr: FetchedInstr16(0x3280), instr_valid: true},  // adds r2, #128
        {instr: FetchedInstr16(0x1c00), instr_valid: false}, // adds r0, r0, #0
        {instr: FetchedInstr16(0x1dff), instr_valid: false}, // adds r7, r7, #7
        {instr: FetchedInstr16(0x1cd1), instr_valid: false}, // adds r1, r2, #3
        {instr: FetchedInstr16(0x1800), instr_valid: false}, // adds r0, r0, r0
        {instr: FetchedInstr16(0x19ff), instr_valid: false}, // adds r7, r7, r7
        {instr: FetchedInstr16(0x0000), instr_valid: false}, // mov r0, r0
        {instr: FetchedInstr16(0x2745), instr_valid: false}, // mov r7, #0x45
        {instr: FetchedInstr16(0x4080), instr_valid: false}, // lsl r0, r0, r0
        {instr: FetchedInstr16(0xffff), instr_valid: false},
    }

    test_identify(t, cases, reflect.TypeOf(AddImmT2{}))
}

func TestDecodeAddImm16T2(t *testing.T) {
    cases := []DecodeCase{
        // adds r0, #0
        {instr: FetchedInstr16(0x3000), decoded: AddImmT2{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT}},
        // adds r7, #255
        {instr: FetchedInstr16(0x37ff), decoded: AddImmT2{Rd: 7, Rm: 0, Rn: 7, Imm: 255, setflags: NOT_IT}},
        // adds r2, #128
        {instr: FetchedInstr16(0x3280), decoded: AddImmT2{Rd: 2, Rm: 0, Rn: 2, Imm: 128, setflags: NOT_IT}},
    }

    test_decode(t, cases, AddImm16T2)
}

func TestExecuteAddImmT2(t *testing.T) {
    cases := []ExecuteCase{
        // adds r0, #0
        {instr: AddImmT2{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
            expected: Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: true}}},
        // adds r0, #1
        {instr: AddImmT2{Rd: 0, Rm: 0, Rn: 0, Imm: 1, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: true}},
            expected: Registers{r: GeneralRegs{2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: false}}},
        // adds r2, #255
        {instr: AddImmT2{Rd: 2, Rm: 0, Rn: 2, Imm: 255, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
            expected: Registers{r: GeneralRegs{1, 2, 258, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}}},
        // adds r0, #1
        {instr: AddImmT2{Rd: 0, Rm: 0, Rn: 0, Imm: 1, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0x7fffffff, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
            expected: Registers{r: GeneralRegs{0x80000000, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Apsr: Apsr{V: true, N: true}}},
        // adds r1, #1
        {instr: AddImmT2{Rd: 1, Rm: 0, Rn: 1, Imm: 1, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 0xffffffff, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
            expected: Registers{r: GeneralRegs{0, 0, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Apsr: Apsr{Z: true, C: true}}},
    }

    test_execute(t, cases)
}

func TestIdentifySubRegT1(t *testing.T) {
    cases := []IdentifyCase{
        {instr: FetchedInstr16(0x1a00), instr_valid: true},  // subs r0, r0, r0
        {instr: FetchedInstr16(0x1bff), instr_valid: true},  // subs r7, r7, r7
        {instr: FetchedInstr16(0x1ad1), instr_valid: true},  // subs r1, r2, r3
        {instr: FetchedInstr16(0x19ff), instr_valid: false}, // adds r7, r7, r7
        {instr: FetchedInstr16(0x18d1), instr_valid: false}, // adds r1, r2, r3
        {instr: FetchedInstr16(0x0000), instr_valid: false}, // mov r0, r0
        {instr: FetchedInstr16(0x2000), instr_valid: false}, // mov r0, #0
        {instr: FetchedInstr16(0x2745), instr_valid: false}, // mov r7, #0x45
        {instr: FetchedInstr16(0x4080), instr_valid: false}, // lsl r0, r0, r0
        {instr: FetchedInstr16(0xffff), instr_valid: false},
    }

    test_identify(t, cases, reflect.TypeOf(SubRegT1{}))
}

func TestDecodeSubReg16T1(t *testing.T) {
    cases := []DecodeCase{
        // subs r0, r0, r0
        {instr: FetchedInstr16(0x1a00), decoded: SubRegT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT}},
        // subs r7, r7, r7
        {instr: FetchedInstr16(0x1bff), decoded: SubRegT1{Rd: 7, Rm: 7, Rn: 7, Imm: 0, setflags: NOT_IT}},
        // subs r1, r2, r3
        {instr: FetchedInstr16(0x1ad1), decoded: SubRegT1{Rd: 1, Rm: 3, Rn: 2, Imm: 0, setflags: NOT_IT}},
    }

    test_decode(t, cases, SubReg16T1)
}

func TestExecuteSubRegT1(t *testing.T) {
    cases := []ExecuteCase{
        // subs r0, r0, r0
        {instr: SubRegT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
            expected: Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: true, C: true}}},
        // subs r0, r0, r1
        {instr: SubRegT1{Rd: 0, Rm: 1, Rn: 0, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: true}},
            expected: Registers{r: GeneralRegs{0xfffffffe, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{N: true}}},
        // subs r1, r2, r3
        {instr: SubRegT1{Rd: 1, Rm: 3, Rn: 2, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
            expected: Registers{r: GeneralRegs{1, 0xffffffff, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{N: true}}},
        // subs r0, r1, r2
        {instr: SubRegT1{Rd: 0, Rm: 2, Rn: 1, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 0x7fffffff, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
            expected: Registers{r: GeneralRegs{0x7ffffffe, 0x7fffffff, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Apsr: Apsr{C: true}}},
        // subs r0, r1, r1
        {instr: SubRegT1{Rd: 0, Rm: 1, Rn: 1, Imm: 0, setflags: NOT_IT},
            regs:     Registers{r: GeneralRegs{0, 0x80000000, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
            expected: Registers{r: GeneralRegs{0, 0x80000000, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Apsr: Apsr{Z: true, C: true}}},
    }

    test_execute(t, cases)
}
