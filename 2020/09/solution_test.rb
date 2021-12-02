#!/usr/bin/env ruby

require 'pry'
require 'minitest/autorun'
require './2020/09/solution'

class Day09Test < Minitest::Test
  TEST_DATA = <<~EODATA
    35
    20
    15
    25
    47
    40
    62
    55
    65
    95
    102
    117
    150
    182
    127
    219
    299
    277
    309
    576
  EODATA

  def test_xmas_cracker_can_select_the_preamble_correctly
    cracker = XmasCracker.new(TEST_DATA, preamble: 5)
    assert_equal [35, 20, 15, 25, 47], cracker.current_preamble
  end

  def test_xmax_creacker_preamble_moves_as_we_advannce
    cracker = XmasCracker.new(TEST_DATA, preamble: 5)
    cracker.advance
    assert_equal [20, 15, 25, 47, 40], cracker.current_preamble
  end

  def test_xmas_cracker_knows_when_current_is_valid
    cracker = XmasCracker.new(<<~EODATA, preamble: 2)
      1
      2
      3
    EODATA

    assert cracker.current_valid?
  end

  def test_xmas_cracker_knows_when_current_is_invalid
    cracker = XmasCracker.new(<<~EODATA, preamble: 2)
      3
      2
      3
    EODATA

    assert !cracker.current_valid?
  end

  def test_xmas_cracker_can_find_invalid_numbers
    cracker = XmasCracker.new(TEST_DATA, preamble: 5)
    assert_equal 127, cracker.all_invalid.first
  end

  def test_xmas_cracker_can_find_continuous_number_summing_to_invalid
    cracker = XmasCracker.new(TEST_DATA, preamble: 5)
    assert_equal [15, 25, 47, 40], cracker.continuous_summands_for_invalid
  end
end
