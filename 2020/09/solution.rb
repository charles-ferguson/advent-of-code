#!/usr/bin/env ruby

class XmasCracker
  attr_reader :numbers, :cursor, :preamble
  def initialize(data, preamble: 25)
    @numbers = data.lines.map(&:to_i)
    @cursor = preamble
    @preamble = preamble
  end

  def advance
    @cursor += 1
  end

  def all_invalid
    @invalid ||= begin
      all_invalid = []
      while cursor < numbers.size
        all_invalid << numbers[cursor] unless current_valid?
        advance
      end
      all_invalid
    end
  end

  def continuous_summands_for_invalid
    range = ContinuousSummandSearcher.new(all_invalid.first, numbers).find_summand_range
    numbers[range.first..range.last]
  end

  def current_valid?
    search_space = current_preamble
    !search_space.find do |candidate|
      (search_space - [candidate]).include?(numbers[cursor] - candidate)
    end.nil?
  end

  def current_preamble
    numbers[cursor - preamble..cursor - 1]
  end
end

class ContinuousSummandSearcher
  attr_reader :cursor, :numbers, :target
  def initialize(target, numbers)
    @numbers = numbers
    @cursor = [0, 0]
    @target = target
  end

  def find_summand_range
    loop do
      raise "Could not find sum: #{target}" if cursor.last > numbers.size
      current_sum = numbers[cursor.first..cursor.last].reduce(:+)
      puts cursor.inspect if current_sum.nil?
      return cursor if current_sum == target

      if current_sum < target
        @cursor = [cursor.first, cursor.last + 1]
      else
        @cursor = [cursor.first + 1, cursor.first + 1]
      end
    end
  end
end

if $PROGRAM_NAME =~ /solution.rb$/
  data = File.read(File.join(__dir__, 'data'))
  cracker = XmasCracker.new(data)

  puts "Part 1: #{cracker.all_invalid.first}"

  puts sum_numbers.inspect
  puts "Part 2: #{sum_numbers.minmax.reduce(:+)}"
end
