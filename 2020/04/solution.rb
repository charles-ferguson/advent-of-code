#!/usr/bin/env ruby
require 'pry'

DATA = File.read(File.join(__dir__, 'data'))

class Passport
  REQUIRED_FIELDS = %w[byr iyr eyr hgt hcl ecl pid].freeze
  DATA_VALID_REGEXES = {
    'byr' => /^(19[2-9]\d|200[0-2])$/,
    'iyr' => /^(201\d|2020)$/,
    'eyr' => /^(202\d|2030)$/,
    'hgt' => /^(1[5-8]\dcm|19[0-3]cm|59in|6\din|7[0-6]in)$/,
    'hcl' => /^#[0-9a-f]{6,6}$/,
    'ecl' => /^(amb|blu|brn|gry|grn|hzl|oth)$/,
    'pid' => /^\d{9,9}$/
  }

  attr_reader :fields
  def initialize(fields)
    @fields = fields
  end

  def valid_fields?
    REQUIRED_FIELDS.all? { |required_field| fields.keys.include?(required_field) }
  end

  def valid_data?
    valid_fields? && DATA_VALID_REGEXES.all? { |field, regex| fields[field] =~ regex }
  end

  def print_invalid_fields
    return unless valid_fields?

    DATA_VALID_REGEXES.each do |field, regex|
      puts "field: #{field}, invalid for #{fields[field]} validated by #{regex}" unless fields[field] =~ regex
    end
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
