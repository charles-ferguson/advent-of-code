#!/usr/bin/env ruby

require 'set'
DATA = File.read(File.join(__dir__, 'data'))

class Parser
  attr_reader :data
  attr_accessor :position
  def initialize(data)
    @data = data.lines
    @position = 0
  end

  def all_groups
    groups = []
    loop do
      break unless next_group?

      groups << next_group
      self.position += 1
    end

    groups
  end

  def next_group
    group = []

    until data.fetch(position).match(/^\s*$/)
      group << data.fetch(position).chomp
      self.position += 1

      break if position >= data.size
    end

    group
  end

  def next_group?
    position <= data.size
  end
end

class Group
  attr_reader :yeses
  def initialize(yeses)
    @yeses = yeses
  end

  def part_1_score
    unique_yeses.size
  end

  def part_2_score
    unique_yeses.select do |yes|
      yeses.all? { |person_yeses| person_yeses.include?(yes) }
    end.count
  end

  def unique_yeses
    Set.new(yeses.flatten.join.chars)
  end
end


groups = Parser.new(DATA).all_groups.map { |data| Group.new(data) }
puts "Part 1: #{groups.map(&:part_1_score).reduce(:+)}"
puts "Part 1: #{groups.map(&:part_2_score).reduce(:+)}"
