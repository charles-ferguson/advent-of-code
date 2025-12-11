#!/usr/bin/env ruby
#
# frozen_string_literal: true

require 'logger'

FILE_PATH = ARGV[0] || 'sample_input.txt'
LOGGER = Logger.new($stdout)
LOGGER.level = Logger::WARN

# Module containing ID validation rules
module IdRules
  def self.valid_for_solution_1?(id)
    [
      NoLeadingZeros,
      NotReapetedTwice
    ].all? { |rule| rule.valid?(id) }
  end

  def self.valid_for_solution_2?(id)
    [
      NoLeadingZeros,
      NotReapetedTwice,
      NoRepeatedTimes
    ].all? { |rule| rule.valid?(id) }
  end

  # Rejects IDs that start with a leading zero
  class NoLeadingZeros
    def self.valid?(id)
      !id.to_s.start_with?('0')
    end
  end

  # Rejects IDs that are made by repeating the same sequence twice
  class NotReapetedTwice
    def self.valid?(id)
      id.to_s !~ /^(\d+)\1$/
    end
  end

  # Rejects IDs that are composed of repeated sequences of the same digits
  class NoRepeatedTimes
    def self.valid?(id)
      id.to_s !~ /^(\d+)\1+$/
    end
  end
end

def solution_1?(file_handle)
  sum_of_invalid = 0

  until file_handle.eof?
    ids = file_handle.gets(',').chomp(',')
    first_id, second_id = ids.split('-')
    range = ((first_id.to_i)..(second_id.to_i)).to_a
    range.each do |id|
      unless IdRules.valid_for_solution_1?(id)
        sum_of_invalid += id
      end
    end
  end
  sum_of_invalid
end

def solution_2?(file_handle)
  sum_of_invalid = 0

  until file_handle.eof?
    ids = file_handle.gets(',').chomp(',')
    first_id, second_id = ids.split('-')
    range = ((first_id.to_i)..(second_id.to_i)).to_a
    range.each do |id|
      unless IdRules.valid_for_solution_2?(id)
        LOGGER.debug("Invalid ID for solution 2: #{id}")
        sum_of_invalid += id
      end
    end
  end
  sum_of_invalid
end


File.open(FILE_PATH) do |file|
  puts "Solution 1: #{solution_1?(file)}"
  file.rewind
  puts "Solution 2: #{solution_2?(file)}"
end
