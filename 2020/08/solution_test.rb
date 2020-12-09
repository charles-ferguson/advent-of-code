#!/usr/bin/env ruby

require 'minitest/autorun'
require 'pry'
require_relative 'solution'

class Day08Test < Minitest::Test
  TEST_DATA = <<~EODATA
    nop +0
    acc +1
    jmp +4
    acc +3
    jmp -3
    acc -99
    acc +1
    jmp -4
    acc +6
  EODATA

  def test_reads_all_the_data
    test_data = <<~EODATA
      nop +0
      acc +1
      jmp +4
    EODATA
    program = Program.new(test_data)

    assert_equal 3, program.instructions.size
  end

  def test_can_parse_nop_instructions
    test_data = 'nop +0'
    instruction = Instructions.for(test_data)

    assert_kind_of Instructions::NopInstruction, instruction
  end

  def test_can_parse_acc_instructions
    test_data = 'acc +5'
    instruction = Instructions.for(test_data)

    assert_kind_of Instructions::AccInstruction, instruction
    assert_equal 5, instruction.operand
  end

  def test_can_parse_jmp_instructions
    test_data = 'jmp -3'
    instruction = Instructions.for(test_data)

    assert_kind_of Instructions::JmpInstruction, instruction
    assert_equal -3, instruction.operand
  end

  def test_knows_the_value_of_accumular_when_hits_loop
    program = Program.new(TEST_DATA)
    program.run

    assert_equal 5, program.accumulator
  end
end
