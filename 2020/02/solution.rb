#!/usr/bin/env ruby

DATA = File.read(File.join(__dir__, "data"))

class PasswordValidator
  LINE_REGEX = /(?<first_number>\d+)-(?<second_number>\d+)\s+(?<character>\w):\s+(?<password>\S+)/

  def self.from_line(line)
    match = line.match(LINE_REGEX)
    raise "Line did not match expected pattern: line: #{line}, pattern: #{LINE_REGEX}" unless match

    rule = SledRentalRuleValidator.new(match[:character], match[:first_number].to_i, match[:second_number].to_i)
    rule = TobagganRentalRuleValidator.new(match[:character], match[:first_number].to_i, match[:second_number].to_i)

    new(match[:password], rule)
  end

  attr_reader :password, :rule_validator
  def initialize(password, rule_validator)
    @password = password
    @rule_validator = rule_validator
  end

  def valid?
    rule_validator.valid?(password)
  end
end

class SledRentalRuleValidator
  attr_reader :character, :min, :max
  def initialize(character, min, max)
    @character = character
    @min = min
    @max = max
  end

  def valid?(password)
    count = password.count(character)
    count <= max && count >= min
  end
end

class TobagganRentalRuleValidator
  attr_reader :character, :indexes
  def initialize(character, *indexes)
    @character = character
    @indexes = indexes
  end

  def valid?(password)
    count = indexes.inject(0) { |count, i| password[i - 1] == character ? count + 1 : count }
    count == 1
  end
end

total_valid = DATA.lines.select { |line| PasswordValidator.from_line(line).valid? }.count
puts total_valid
