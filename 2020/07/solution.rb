#!/usr/bin/env ruby

require 'set'
class BagParser
  BAG_REGEX = /^(?<description>[\S\s]*?) bags/
  CONTAIN_REGEX = /(\d+) ([\S\s]*?) bags*[,.]/
  attr_reader :data, :bags
  def initialize(data, collection = BagCollection.new)
    @data = data
    @bags = collection
  end

  def bag_collection
    data.lines.map(&:chomp).map { |line| line_to_bag(line) }
    bags
  end

  def line_to_bag(line)
    bag_data, containing_data = line.split(' contain ')
    bag_match = BAG_REGEX.match(bag_data)
    bag = bags.find_or_create(bag_match['description'])
    return bag if containing_data == 'no other bags.'

    containing_data.scan(CONTAIN_REGEX).each do |match_group|
      count = match_group.first.to_i
      other_bag = bags.find_or_create(match_group.last)
      bag.add_container_rule(count, other_bag)
    end

    bag
  end
end

class BagCollection
  attr_reader :bags
  def initialize
    @bags = []
  end

  def add_bag(bag)
    bags << []
  end

  def find_or_create(description)
    bag = bags.find { |a_bag| a_bag.description == description }
    return bag if bag

    bag = Bag.new(description)
    bags << bag

    bag
  end
end

class Bag
  attr_reader :description, :container_allowances
  def initialize(description)
    @description = description
    @container_allowances = []
  end

  def add_container_rule(number, other_bag)
    container_allowances << BagContainerRule.new(number, other_bag)
  end

  def can_contain?(other_bag)
    !!container_allowances.find do |container_allowance|
      inner_bag =container_allowance.bag
      inner_bag == other_bag || inner_bag.can_contain?(other_bag)
    end
  end

  def number_contained_bags
    container_allowances.inject(0) do |sum, contained_bag|
      sum + contained_bag.number * (1 + contained_bag.bag.number_contained_bags)
    end
  end

  def ==(other)
    other.respond_to?(:description) && other.description == description
  end

  def eql?(other)
    self == other
  end
end

BagContainerRule = Struct.new(:number, :bag)

if $PROGRAM_NAME =~ /solution.rb$/
  data = File.read(File.join(__dir__, 'data'))
  bag_collection = BagParser.new(data).bag_collection

  objective = bag_collection.find_or_create('shiny gold')
  part1_solution = bag_collection.bags.select { |bag| bag.can_contain?(objective) }.count

  puts "Part 1: #{part1_solution}"
  puts "Part 2: #{objective.number_contained_bags}"

end
