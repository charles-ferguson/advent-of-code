#!/usr/bin/env ruby

require 'minitest/autorun'
require 'pry'
require_relative 'solution'

class Day07Test < Minitest::Test
  TEST_DATA = <<~EODATA
    light red bags contain 1 bright white bag, 2 muted yellow bags.
    dark orange bags contain 3 bright white bags, 4 muted yellow bags.
    bright white bags contain 1 shiny gold bag.
    muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
    shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
    dark olive bags contain 3 faded blue bags, 4 dotted black bags.
    vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
    faded blue bags contain no other bags.
    dotted black bags contain no other bags.
  EODATA

  def test_gets_all_the_bags
    parser = BagParser.new(TEST_DATA)
    assert_equal 9, parser.bag_collection.bags.length
  end

  def test_bags_know_their_description
    parser = BagParser.new(TEST_DATA.lines[0])
    assert_equal "light red", parser.bag_collection.bags.first.description
  end

  def test_bags_that_can_not_contain_other_bags_dont
    test_data = "dotted black bags contain no other bags.\n"
    bag = BagParser.new(test_data).bag_collection.bags.last

    assert_equal 0, bag.container_allowances.size
  end

  def test_bags_that_contain_one_other_type_of_bag_do
    test_data = <<~EODATA
      bright white bags contain 1 shiny gold bag.
      shiny gold bags contain no other bags.
    EODATA
    bags = BagParser.new(test_data).bag_collection.bags

    assert bags.first.can_contain?(bags.last)
  end

  def test_bags_can_contain_bags_within_bags_do
    test_data = <<~EODATA
      dark orange bags contain 1 bright white bag.
      bright white bags contain 1 shiny gold bag.
      shiny gold bags contain no other bags.
    EODATA
    bags = BagParser.new(test_data).bag_collection.bags

    assert bags.first.can_contain?(bags.last)
  end

  def test_bags_can_contain_multiple_types_of_bags
    test_data = 'light red bags contain 1 bright white bag, 2 muted yellow bags.\n'
    bag_collection = BagParser.new(test_data).bag_collection

    assert bag_collection.find_or_create('light red').can_contain?(
      bag_collection.find_or_create('muted yellow')
    )
  end

  def test_bags_containing_no_bag_know_they_contain_zero
    test_data = "shiny gold bags contain no other bags.\n"
    bag_collection = BagParser.new(test_data).bag_collection

    assert_equal 0, bag_collection.find_or_create('shiny gold').number_contained_bags
  end

  def test_bags_containing_bags_contain_the_bags_in_those_bags
    test_data = <<~EODATA
      dark orange bags contain 1 bright white bag.
      bright white bags contain 3 shiny gold bag.
      shiny gold bags contain no other bags.
    EODATA
    bag_collection = BagParser.new(test_data).bag_collection

    assert_equal 4, bag_collection.find_or_create('dark orange').number_contained_bags
  end

end
