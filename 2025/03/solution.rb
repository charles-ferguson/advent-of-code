#!/usr/bin/env ruby
# frozen_string_literal: true

require 'pry'
require 'logger'

INPUT_PATH = ARGV[0] || 'sample_input.txt'
LOGGER = Logger.new($stdout)
LOGGER.level = ARGV[1] == "-d" ? Logger::DEBUG : Logger::WARN

class BatteryBank
  attr_reader :joltages
  def initialize(joltages)
    @joltages = joltages.chomp.split("").map(&:to_i)
  end

  def max_two_battery_joltage
    max_joltage = 0
    @joltages.each_with_index do |joltage, index|
      @joltages[index + 1..-1].each do |next_joltage|
        total_joltage = joltage * 10 + next_joltage
        max_joltage = total_joltage if total_joltage > max_joltage
      end
    end
    max_joltage
  end

  def max_twelve_battery_joltage
    max_joltages = []

    search_bounds = [0, @joltages.length - 12]
    while max_joltages.length < 12
      LOGGER.debug(@joltages.join)
      LOGGER.debug("#{' ' * search_bounds[0]}^#{' ' * (search_bounds[1] + 1 - search_bounds[0])}^")
      search_space = @joltages[search_bounds[0]..search_bounds[1]]
      max_joltages << search_space.max
      search_bounds[0] += search_space.index(max_joltages.last) + 1
      search_bounds[1] += 1
    end

    max_joltages.join.to_i
  end
end

def solution_1(file_handle)
  sum = 0
  file_handle.each_line do |line|
    battery_bank = BatteryBank.new(line)
    joltage = battery_bank.max_two_battery_joltage
    sum += joltage
    LOGGER.debug("Max joltage: #{joltage}, Current sum: #{sum}")
  end
  puts "Solution 1: #{sum}"
end

def solution_2(file_handle)
  sum = 0
  file_handle.each_line do |line|
    battery_bank = BatteryBank.new(line)
    joltage = battery_bank.max_twelve_battery_joltage
    sum += joltage
    LOGGER.debug("Max joltage (12 batteries): #{joltage}, line: #{line}, Current sum: #{sum}")
  end
  puts "Solution 2: #{sum}"
end

File.open(INPUT_PATH) do |file|
  solution_1(file)
  file.rewind
  solution_2(file)
end
