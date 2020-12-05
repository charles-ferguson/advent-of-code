#!/usr/bin/env ruby
require 'pry'

DATA = File.read(File.join(__dir__, 'data'))

module DataValidators
  NumericRangeRule = ::Struct.new(:min, :max)
  class NumericRangeRule
    def valid?(data)
      (min..max).include?(Float(data))
    rescue ArgumentError
      false
    end
  end

  RegexRule = ::Struct.new(:pattern)
  class RegexRule
    def valid?(data)
      data =~ pattern
    end
  end

  class FixesValuesRule
    attr_reader :allowed_values
    def initialize(allowed_values)
      @allowed_values = allowed_values
    end

    def valid?(data)
      allowed_values.include?(data)
    end
  end

  NumericRangeFromCaptureRule = Struct.new(:pattern, :min, :max)
  class NumericRangeFromCaptureRule
    def valid?(data)
      match = pattern.match(data)

      NumericRangeRule.new(min, max).valid?(match.captures.first) if match
    end
  end

  class OrRule
    attr_reader :rule1, :rule2
    def initialize(rule1, rule2)
      @rule1 = rule1
      @rule2 = rule2
    end

    def valid?(data)
      rule1.valid?(data) || rule2.valid?(data)
    end
  end
end

class Passport
  REQUIRED_FIELDS = %w[byr iyr eyr hgt hcl ecl pid].freeze
  DATA_VALID_REGEXES = {
    'byr' => DataValidators::NumericRangeRule.new(1920, 2002),
    'iyr' => DataValidators::NumericRangeRule.new(2010, 2020),
    'eyr' => DataValidators::NumericRangeRule.new(2020, 2030),
    'hgt' => DataValidators::OrRule.new(
      DataValidators::NumericRangeFromCaptureRule.new(/^(\d+)cm$/, 150, 193),
      DataValidators::NumericRangeFromCaptureRule.new(/^(\d+)in$/, 59, 76)
    ),
    'hcl' => DataValidators::RegexRule.new(/^#[0-9a-f]{6,6}$/),
    'ecl' => DataValidators::FixesValuesRule.new(%w[amb blu brn gry grn hzl oth]),
    'pid' => DataValidators::RegexRule.new(/^\d{9,9}$/)
  }

  attr_reader :fields
  def initialize(fields)
    @fields = fields
  end

  def valid_fields?
    REQUIRED_FIELDS.all? { |required_field| fields.keys.include?(required_field) }
  end

  def valid_data?
    valid_fields? && DATA_VALID_REGEXES.all? { |field, validator| validator.valid?(fields[field]) }
  end
end

class Parser
  REGEX = /(?<key>\w+):(?<value>[^\s]+)/.freeze

  attr_reader :data, :position
  def initialize(data)
    @data = data.lines
    @position = 0
  end

  def parse
    all_records.map { |record| record_to_hash(record) }
  end

  private

  def all_records
    records = []
    loop do
      break unless next_record?

      records << next_full_record
    end

    records
  end

  def next_full_record
    record = []

    until data.fetch(position).match(/^\s*$/)
      record << data.fetch(position).chomp
      @position += 1

      break if position >= data.size
    end

    @position += 1 # advance to start of next record

    record.join(' ')
  end

  def next_record?
    position <= data.size
  end

  def record_to_hash(record_string)
    record_string.scan(REGEX).each_with_object({}) do |kv, hsh|
      hsh[kv.first] = kv.last
    end
  end
end

passports = Parser.new(DATA).parse.map { |record| Passport.new(record) }
puts "Part 1: #{passports.select(&:valid_fields?).count } valid of #{passports.count}"
puts "Part 2: #{passports.select(&:valid_data?).count } valid of #{passports.count}"
